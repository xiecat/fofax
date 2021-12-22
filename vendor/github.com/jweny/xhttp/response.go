package xhttp

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"time"
)

//‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾
// Response struct
//_______________________________________________________________________

type Response struct {
	Request     *Request
	RawResponse *http.Response
	Body        []byte

	raw        []byte
	receivedAt time.Time
}

//‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾
// Response methods get
//_______________________________________________________________________

func (r *Response) GetHeaders() http.Header {
	return r.RawResponse.Header
}

func (r *Response) GetContentType() string {
	return r.RawResponse.Header.Get("content-type")
}

func (r *Response) GetUrl() *url.URL {
	return r.RawResponse.Request.URL
}

func (r *Response) GetLatency() (time.Duration, error) {
	var latency time.Duration
	if r.Request.clientTrace != nil {
		latency = r.Request.getTraceInfo().TotalTime
	}
	latency = r.receivedAt.Sub(r.Request.sendAt)
	if latency == 0 {
		return 0, fmt.Errorf("response latency %d, maybe not enable trace", latency)
	}
	return latency, nil
}

// GetStatus method returns the HTTP status string for the executed request.
func (r *Response) GetStatus() int {
	return r.RawResponse.StatusCode
}

func (r *Response) GetBody() []byte {
	return r.Body
}

func (r *Response) GetRaw() ([]byte, error) {
	// dump 响应头
	respHeaderRaw, err := httputil.DumpResponse(r.RawResponse, false)
	if err != nil {
		return nil, err
	}
	// 拼接 Body
	r.raw = append(respHeaderRaw, r.Body...)
	return r.raw, nil
}

//// Error method returns the error object if it has one
//func (r *Response) Error() interface{} {
//	return r.Request.Error
//}

// getReceivedAt method returns when response got received from server for the request.
func (r *Response) getReceivedAt() time.Time {
	return r.receivedAt
}

//‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾
// Response methods set
//_______________________________________________________________________

func (r *Response) setReceivedAt() {
	r.receivedAt = time.Now()
	if r.Request.clientTrace != nil {
		r.Request.clientTrace.endTime = r.receivedAt
	}
}
