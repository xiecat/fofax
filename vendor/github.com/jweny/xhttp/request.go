package xhttp

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
	"time"
)

//‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾
// Request struct
//_______________________________________________________________________

type Request struct {
	RawRequest *http.Request
	Error      interface{}
	Body       []byte

	attempt     int
	ctx         context.Context
	raw         []byte
	trace       bool
	sendAt      time.Time
	clientTrace *clientTrace
}

//‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾
// Request methods get
//_______________________________________________________________________

// GetContext get
func (r *Request) GetContext() context.Context {
	if r.ctx == nil {
		return context.Background()
	}
	return r.ctx
}

// GetAttempt get
func (r *Request) GetAttempt() int {
	return r.attempt
}

func (r *Request) GetUrl() *url.URL {
	return r.RawRequest.URL
}

func (r *Request) GetMethod() string {
	return r.RawRequest.Method
}

func (r *Request) GetHeaders() http.Header {
	return r.RawRequest.Header
}

func (r *Request) GetContentType() string {
	return r.RawRequest.Header.Get("content-type")

}

func (r *Request) GetBody() ([]byte, error) {
	if r.Body != nil {
		return r.Body, nil
	}
	if r.RawRequest.Body == nil {
		return nil, nil
	}
	bodyBytes, err := ioutil.ReadAll(r.RawRequest.Body)
	if err != nil {
		return nil, err
	}
	// 修复 read 后的偏移量
	r.RawRequest.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
	r.Body = bodyBytes
	return r.Body, nil
}

func (r *Request) GetRaw() ([]byte, error) {
	if r.raw != nil {
		return r.raw, nil
	}
	// Dump请求头
	reqHeaderRaw, err := httputil.DumpRequest(r.RawRequest, false)
	if err != nil {
		return nil, err
	}
	// 拼接 Body
	r.raw = append(reqHeaderRaw, r.Body...)
	return r.raw, nil
}

func (r *Request) getTraceInfo() TraceInfo {
	ct := r.clientTrace
	if ct == nil {
		return TraceInfo{}
	}
	ti := TraceInfo{
		DNSLookup:     ct.dnsDone.Sub(ct.dnsStart),
		TLSHandshake:  ct.tlsHandshakeDone.Sub(ct.tlsHandshakeStart),
		ServerTime:    ct.gotFirstResponseByte.Sub(ct.gotConn),
		IsConnReused:  ct.gotConnInfo.Reused,
		IsConnWasIdle: ct.gotConnInfo.WasIdle,
		ConnIdleTime:  ct.gotConnInfo.IdleTime,
		//RequestAttempt: r.Attempt,
	}

	// Calculate the total time accordingly, when connection is reused
	if ct.gotConnInfo.Reused {
		ti.TotalTime = ct.endTime.Sub(ct.getConn)
	} else {
		ti.TotalTime = ct.endTime.Sub(ct.dnsStart)
	}

	// Only calculate on successful connections
	if !ct.connectDone.IsZero() {
		ti.TCPConnTime = ct.connectDone.Sub(ct.dnsDone)
	}

	// Only calculate on successful connections
	if !ct.gotConn.IsZero() {
		ti.ConnTime = ct.gotConn.Sub(ct.getConn)
	}

	// Only calculate on successful connections
	if !ct.gotFirstResponseByte.IsZero() {
		ti.ResponseTime = ct.endTime.Sub(ct.gotFirstResponseByte)
	}

	// Capture remote address info when connection is non-nil
	if ct.gotConnInfo.Conn != nil {
		ti.RemoteAddr = ct.gotConnInfo.Conn.RemoteAddr()
	}

	return ti
}

//‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾
// Request methods set
//_______________________________________________________________________

func (r *Request) EnableTrace() *Request {
	r.trace = true
	return r
}

func (r *Request) setSendAt() *Request {
	r.sendAt = time.Now()
	return r
}

// SetContext set
func (r *Request) SetContext(ctx context.Context) *Request {
	r.ctx = ctx
	return r
}

// SetHeader set a single header field and its value in the current request
func (r *Request) SetHeader(key, value string) *Request {
	r.RawRequest.Header.Set(key, value)
	return r
}

// SetHeaders set multiple headers field
func (r *Request) SetHeaders(headers map[string]string) *Request {
	for h, v := range headers {
		r.SetHeader(h, v)
	}
	return r
}

// SetHeaderMultiValues sets multiple headers fields and its values is list of strings
// For Example: To set `Accept` as `text/html, application/xhtml+xml, application/xml;q=0.9, image/webp, */*;q=0.8`
func (r *Request) SetHeaderMultiValues(headers map[string][]string) *Request {
	for key, values := range headers {
		r.SetHeader(key, strings.Join(values, ", "))
	}
	return r
}

func (r *Request) SetCookie(hc *http.Cookie) *Request {
	r.RawRequest.AddCookie(hc)
	return r
}
