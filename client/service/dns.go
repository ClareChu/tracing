package service

import (
	"context"
	"fmt"
	"github.com/ClareChu/gorequest"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"github.com/opentracing/opentracing-go/log"
	"net/http/httptrace"
)

func NewClientTrace(span opentracing.Span) *httptrace.ClientTrace {
	trace := &clientTrace{span: span}
	return &httptrace.ClientTrace{
		DNSStart: trace.dnsStart,
		DNSDone:  trace.dnsDone,
	}
}

// clientTrace holds a reference to the Span and
// provides methods used as ClientTrace callbacks
type clientTrace struct {
	span opentracing.Span
}

func (h *clientTrace) dnsStart(info httptrace.DNSStartInfo) {
	h.span.LogKV(
		log.String("event", "DNS start"),
		log.Object("host", info.Host),
	)
}

func (h *clientTrace) dnsDone(httptrace.DNSDoneInfo) {
	h.span.LogKV(log.String("event", "DNS done"))
}

func RunClient(ctx context.Context) {
	url := "http://localhost:8081/dns/start"
	resp, _, _ := gorequest.New().Get(url).SetSpanContext(ctx).End()
	if resp != nil {
		fmt.Println(resp.StatusCode)
	}

}

func onError(span opentracing.Span, err error) {
	// handle errors by recording them in the span
	span.SetTag(string(ext.Error), true)
	span.LogKV(log.Error(err))
}
