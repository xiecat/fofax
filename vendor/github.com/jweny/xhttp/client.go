package xhttp

import (
	"context"
	"errors"
	"fmt"
	"github.com/jweny/xhttp/xtls"
	"golang.org/x/net/http2"
	"golang.org/x/net/publicsuffix"
	"net"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"time"
)

type (
	// requestMiddleware run before request
	requestMiddleware func(*Request, *Client) error
	// responseMiddleware run after receive response
	responseMiddleware func(*Response, *Client) error
	// errorHook after retry deal error
	errorHook func(*Request, error)
)


func (o *HTTPClientOptions) verify() error {
	if o == nil || o.Limiter == nil {
		return errors.New("client options or limiter cannot nil")
	}
	return nil
}

// Client struct
type Client struct {
	HTTPClient    *http.Client
	ClientOptions *HTTPClientOptions
	// if debug == true will run responseLogger middleware
	Debug         bool
	// todo Error handle exp:Error interface{}

	// Middleware
	defaultBeforeRequest []requestMiddleware
	extraBeforeRequest   []requestMiddleware
	afterResponse        []responseMiddleware
	errorHooks           []errorHook

	// handle
	closeConnection		 bool
}

// NewClient 不跟随跳转
func NewClient(options *HTTPClientOptions, jar *cookiejar.Jar) (*Client, error) {
	err := options.verify()
	if err != nil {
		return nil, err
	}
	hc, err := createHttpClient(false, options, jar)
	if err != nil {
		return nil, err
	}

	client := createClient(hc, options)
	return client, nil
}

// NewRedirectClient 跟随跳转
func NewRedirectClient(options *HTTPClientOptions, jar *cookiejar.Jar) (*Client, error) {
	err := options.verify()
	if err != nil {
		return nil, err
	}
	hc, err := createHttpClient(true, options, jar)
	if err != nil {
		return nil, err
	}

	client := createClient(hc, options)
	return client, nil
}

// NewWithHTTPClient 根据 http client 创建
func NewWithHTTPClient(hc *http.Client, options *HTTPClientOptions) (*Client, error) {
	err := options.verify()
	if err != nil {
		return nil, err
	}
	return createClient(hc, options), nil
}

// Do request
func (c *Client) Do(ctx context.Context, req *Request) (*Response, error) {
	var (
		resp                 *http.Response
		shouldRetry          bool
		err, doErr, checkErr error
	)
	err = c.verify()
	if err != nil {
		return nil, err
	}

	req.SetContext(ctx)
	req.attempt = 0

	err = c.ClientOptions.Limiter.Wait(req.GetContext())
	if err != nil {
		return nil, err
	}

	// user diy requestMiddleware
	for _, f := range c.extraBeforeRequest {
		if err = f(req, c); err != nil {
			return nil, err
		}
	}

	// default diy requestMiddleware
	for _, f := range c.defaultBeforeRequest {
		if err = f(req, c); err != nil {
			return nil, err
		}
	}

	// do request with retry
	for i := 0; ; i++ {
		req.attempt++
		// 尝试请求
		req.setSendAt()
		resp, doErr = c.HTTPClient.Do(req.RawRequest)
		// need retry
		shouldRetry, checkErr = defaultRetryPolicy(req.GetContext(), resp, doErr)
		if !shouldRetry {
			break
		}
		remain := c.ClientOptions.FailRetries - i
		if remain <= 0 {
			break
		}
		// waitTime
		waitTime := defaultBackoff(defaultRetryWaitMin, defaultRetryWaitMax, i, resp)
		select {
		case <-time.After(waitTime):
		case <-req.GetContext().Done():
			return nil, req.GetContext().Err()
		}
	}

	if doErr == nil && checkErr == nil && !shouldRetry {
		// request success
		response := &Response{
			Request:     req,
			RawResponse: resp,
		}
		response.setReceivedAt()

		//responseMiddleware
		for _, f := range c.afterResponse {
			if err = f(response, c); err != nil {
				return nil, err
			}
		}
		return response, nil
	} else {
		// request fail
		finalErr := doErr
		if checkErr != nil {
			err = checkErr
		}
		return nil, fmt.Errorf("%s %s giving up after %d attempt(s): %w",
			req.RawRequest.Method, req.RawRequest.URL, req.attempt, finalErr)
	}
}

func (c *Client) SetCloseConnection(close bool) *Client {
	c.closeConnection = close
	return c
}

func (c *Client) verify() error {
	if c == nil {
		return errors.New("current client is nil")
	}
	return nil
}

func createClient(hc *http.Client, options *HTTPClientOptions) *Client {
	c := &Client{
		HTTPClient:    hc,
		ClientOptions: options,
		Debug:         options.Debug,
	}

	c.extraBeforeRequest = []requestMiddleware{}
	c.defaultBeforeRequest = []requestMiddleware{
		verifyRequestMethod,
		createHTTPRequest,
	}
	c.afterResponse = []responseMiddleware{
		readResponseBody,
		verifyResponseBodyLength,
		responseLogger,
	}
	return c
}

func createHttpClient(followRedirects bool, options *HTTPClientOptions, jar *cookiejar.Jar) (*http.Client, error) {

	tlsClientConfig, err := xtls.NewTLSClientConfig(options.TlsOptions)
	if err != nil {
		return nil, err
	}
	transport := &http.Transport{
		DialContext: (&net.Dialer{
			Timeout: time.Duration(options.DialTimeout) * time.Second,
		}).DialContext,
		MaxConnsPerHost:       options.MaxConnsPerHost,
		ResponseHeaderTimeout: time.Duration(options.ReadTimeout) * time.Second,
		IdleConnTimeout:       time.Duration(options.IdleConnTimeout) * time.Second,
		TLSHandshakeTimeout:   time.Duration(options.TLSHandshakeTimeout) * time.Second,
		MaxIdleConns:          options.MaxIdleConns,
		TLSClientConfig:       tlsClientConfig,
		DisableKeepAlives:     options.DisableKeepAlives,
	}
	if options.EnableHTTP2 {
		err := http2.ConfigureTransport(transport)
		if err != nil {
			return nil, err
		}
	}

	if options.Proxy != "" {
		proxy, err := url.Parse(options.Proxy)
		if err != nil {
			return nil, err
		}
		transport.Proxy = http.ProxyURL(proxy)
	}
	// 如果没有传cookiejar 就生成个默认的
	if jar == nil {
		cookieJar, err := cookiejar.New(&cookiejar.Options{PublicSuffixList: publicsuffix.List})
		if err != nil {
			return nil, err
		}
		jar = cookieJar
	}

	return &http.Client{
		Jar:           jar,
		Transport:     transport,
		CheckRedirect: makeCheckRedirectFunc(followRedirects, options.MaxRedirect),
	}, nil
}

type checkRedirectFunc func(req *http.Request, via []*http.Request) error

func makeCheckRedirectFunc(followRedirects bool, maxRedirects int) checkRedirectFunc {
	return func(req *http.Request, via []*http.Request) error {
		if !followRedirects {
			return http.ErrUseLastResponse
		}
		if len(via) >= maxRedirects {
			return http.ErrUseLastResponse
		}
		return nil
	}
}
