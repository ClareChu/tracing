package service

import (
	"context"
	"fmt"
	"github.com/opentracing-contrib/go-stdlib/nethttp"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"github.com/opentracing/opentracing-go/log"
	"io/ioutil"
	"net/http"
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

func runClient(ctx context.Context) {
	// nethttp.Transport from go-stdlib will do the tracing
	c := &http.Client{Transport: &nethttp.Transport{}}
	tracer := opentracing.SpanFromContext(ctx).Tracer()

	// create a top-level span to represent full work of the client

	req, err := http.NewRequest(
		"GET",
		fmt.Sprintf("http://localhost:%s/", "8081"),
		nil,
	)

	req = req.WithContext(ctx)
	// wrap the request in nethttp.TraceRequest
	req, ht := nethttp.TraceRequest(tracer, req)
	defer ht.Finish()
	res, err := c.Do(req)
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return
	}
	fmt.Printf("Received result: %s\n", string(body))
}

func onError(span opentracing.Span, err error) {
	// handle errors by recording them in the span
	span.SetTag(string(ext.Error), true)
	span.LogKV(log.Error(err))
}
