package controller

import (
	"context"
	"fmt"
	"github.com/opentracing/opentracing-go"
	"runtime"
	"strconv"
	"strings"
	"time"
)

type Dns struct {
}

type Response struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (d *Dns) Start(ctx context.Context) (rep *Response) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "startService")
	span.LogKV("message", "server dns start...")
	defer span.Finish()
	return &Response{
		Code:    0,
		Message: "start",
	}
}

func (d *Dns) Done(ctx context.Context) (rep *Response) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "startService")
	defer span.Finish()
	time.Sleep(30 * time.Second)
	span.LogKV("message", "server dns Done...")
	return &Response{
		Code:    0,
		Message: "done",
	}
}


func GoID() int {
	var buf [64]byte
	n := runtime.Stack(buf[:], false)
	idField := strings.Fields(strings.TrimPrefix(string(buf[:n]), "goroutine "))[0]
	id, err := strconv.Atoi(idField)
	if err != nil {
		panic(fmt.Sprintf("cannot get goroutine id: %v", err))
	}
	return id
}