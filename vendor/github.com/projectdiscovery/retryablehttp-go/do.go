package retryablehttp

import (
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"sync/atomic"
	"time"
)

// PassthroughErrorHandler is an ErrorHandler that directly passes through the
// values from the net/http library for the final request. The body is not
// closed.
func PassthroughErrorHandler(resp *http.Response, err error, _ int) (*http.Response, error) {
	return resp, err
}

// Do wraps calling an HTTP method with retries.
func (c *Client) Do(req *Request) (*http.Response, error) {
	var resp *http.Response
	var err error

	// Create a main context that will be used as the main timeout
	mainCtx, cancel := context.WithTimeout(context.Background(), c.options.Timeout)
	defer cancel()

	for i := 0; ; i++ {
		// Always rewind the request body when non-nil.
		if req.body != nil {
			body, err := req.body()
			if err != nil {
				c.closeIdleConnections()
				return resp, err
			}
			if c, ok := body.(io.ReadCloser); ok {
				req.Body = c
			} else {
				req.Body = ioutil.NopCloser(body)
			}
		}

		if c.RequestLogHook != nil {
			c.RequestLogHook(req.Request, i)
		}

		// Attempt the request
		resp, err = c.HTTPClient.Do(req.Request)

		// Check if we should continue with retries.
		checkOK, checkErr := c.CheckRetry(req.Context(), resp, err)

		// if err is equal to missing minor protocol version retry with http/2
		if err != nil && strings.Contains(err.Error(), "net/http: HTTP/1.x transport connection broken: malformed HTTP version \"HTTP/2\"") {
			resp, err = c.HTTPClient2.Do(req.Request)
			checkOK, checkErr = c.CheckRetry(req.Context(), resp, err)
		}

		if err != nil {
			// Increment the failure counter as the request failed
			req.Metrics.Failures++
		} else {
			// Call this here to maintain the behavior of logging all requests,
			// even if CheckRetry signals to stop.
			if c.ResponseLogHook != nil {
				// Call the response logger function if provided.
				c.ResponseLogHook(resp)
			}
		}

		// Now decide if we should continue.
		if !checkOK {
			if checkErr != nil {
				err = checkErr
			}
			c.closeIdleConnections()
			return resp, err
		}

		// We do this before drainBody beause there's no need for the I/O if
		// we're breaking out
		remain := c.options.RetryMax - i
		if remain <= 0 {
			break
		}

		// Increment the retries counter as we are going to do one more retry
		req.Metrics.Retries++

		// We're going to retry, consume any response to reuse the connection.
		if err == nil && resp != nil {
			c.drainBody(req, resp)
		}

		// Wait for the time specified by backoff then retry.
		// If the context is cancelled however, return.
		wait := c.Backoff(c.options.RetryWaitMin, c.options.RetryWaitMax, i, resp)

		// Exit if the main context or the request context is done
		// Otherwise, wait for the duration and try again.
		select {
		case <-mainCtx.Done():
			break
		case <-req.Context().Done():
			c.closeIdleConnections()
			return nil, req.Context().Err()
		case <-time.After(wait):
		}
	}

	if c.ErrorHandler != nil {
		c.closeIdleConnections()
		return c.ErrorHandler(resp, err, c.options.RetryMax+1)
	}

	// By default, we close the response body and return an error without
	// returning the response
	if resp != nil {
		resp.Body.Close()
	}
	c.closeIdleConnections()
	return nil, fmt.Errorf("%s %s giving up after %d attempts: %w", req.Method, req.URL, c.options.RetryMax+1, err)
}

// Try to read the response body so we can reuse this connection.
func (c *Client) drainBody(req *Request, resp *http.Response) {
	_, err := io.Copy(ioutil.Discard, io.LimitReader(resp.Body, c.options.RespReadLimit))
	if err != nil {
		req.Metrics.DrainErrors++
	}
	resp.Body.Close()
}

const closeConnectionsCounter = 100

func (c *Client) closeIdleConnections() {
	if c.options.KillIdleConn {
		requestCounter := atomic.LoadUint32(&c.requestCounter)
		if requestCounter < closeConnectionsCounter {
			atomic.AddUint32(&c.requestCounter, 1)
		} else {
			atomic.StoreUint32(&c.requestCounter, 0)
			c.HTTPClient.CloseIdleConnections()
		}
	}
}
