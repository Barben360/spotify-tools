package utils

import (
	"context"
	"net/http"

	"github.com/Barben360/spotify-tools/app/services/logger"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type httpClientTransport struct {
	options *TransportOptions
}

type TransportOptions struct {
	reqTransformers []func(*http.Request)
	logLevel        zapcore.Level
	log             *logger.Logger
}

type TransportOption func(*TransportOptions)

func WithReqTransformer(reqTransformer func(*http.Request)) TransportOption {
	return func(o *TransportOptions) {
		o.reqTransformers = append(o.reqTransformers, reqTransformer)
	}
}

func WithLog(logger *logger.Logger, level zapcore.Level) TransportOption {
	return func(o *TransportOptions) {
		o.log = logger
		o.logLevel = level
	}
}

func NewTransport(ctx context.Context, opt ...TransportOption) http.RoundTripper {
	options := &TransportOptions{}
	for _, o := range opt {
		o(options)
	}
	return &httpClientTransport{
		options: options,
	}
}

func (t *httpClientTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	for _, reqTransformer := range t.options.reqTransformers {
		reqTransformer(req)
	}
	// Logging at the end to have the final URL
	if t.options.log != nil {
		t.options.log.Logger.Log(t.options.logLevel, "http request", zap.String("method", req.Method), zap.String("url", req.URL.String()), zap.Any("headers", req.Header))
	}
	resp, err := http.DefaultTransport.RoundTrip(req)
	if err != nil {
		return resp, err
	}
	if t.options.log != nil {
		t.options.log.Logger.Log(t.options.logLevel, "http response", zap.String("status", resp.Status), zap.String("url", req.URL.String()), zap.Int("status", resp.StatusCode))
	}
	return resp, err
}
