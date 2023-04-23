package models

import "net/http"

// ResponseWriterWrapper wraps the natural `http.ResponseWriter` and simply exposes the status code we set
//in the handlers.  This way we can intercept it in `app.ProxyHandler` and log appropriately for each request.
type ResponseWriterWrapper struct {
    http.ResponseWriter
    StatusCode int
}

func NewResponseWriterWrapper(w http.ResponseWriter) *ResponseWriterWrapper {
    return &ResponseWriterWrapper{w, http.StatusOK}
}

func (lrw *ResponseWriterWrapper) WriteHeader(code int) {
    lrw.StatusCode = code
    lrw.ResponseWriter.WriteHeader(code)
}
