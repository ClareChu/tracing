
### Jaeger组件

#### Agent

Agent是一个网络守护进程，监听通过UDP发送过来的Span，它会将其批量发送给collector。按照设计，Agent要被部署到所有主机上，作为基础设施。Agent将collector和客户端之间的路由与发现机制抽象了出来。

#### Collector

Collector从Jaeger Agent接收Trace，并通过一个处理管道对其进行处理。目前的管道会校验Trace、建立索引、执行转换并最终进行存储。存储是一个可插入的组件，现在支持Cassandra和elasticsearch。

#### Query

Query服务会从存储中检索Trace并通过UI界面进行展现，该UI界面通过React技术实现，其页面UI如下图所示，展现了一条Trace的详细信息。

#### 存储

jaeger采集到的数据必须存储到某个存储引擎，目前支持Cassandra和elasticsearch

 

在开启 span 记录一个过程时，还可以通过 api 进行 tag，logs等操作 ，并能在 UI 看到相应设置的键z值
```go
span.SetTag("value", helloStr)
span.LogFields(
    log.String("event", "sayhello"),
    log.String("value", helloStr),
)
//span.LogKV("event", "sayhello") // 单一设置
```

tag 和 logs 在opentarcing中提到一些推荐命名：语义惯例

使用 tag 是用于描述 span 中的特性，是对整个过程而言，而 log 是用于记录 span 这个过程中的一个时间，因为记录 log 时会携带一个发生的时间戳，是有先后之分的。

###### baggage
相比 tag，log 限制在 span 中， baggage 同样提供保存键值对设置，但是  baggage 数据有效是全 trace 的，所以使用的时候避免设置不必要的值，导致传递开销。
```go
// set
span.SetBaggageItem("greeting", greeting)
// get
greeting := span.BaggageItem("greeting")
```


###### 使用上下文传递 span
当我们提到调用链，一般涉及多个函数，多个进程甚至多个机器上运行的过程，用 tracer 开启 root span 后，需要向其他过程传递以保持他们之间的关联性，我们通过上下文来存储 span 并传递。

```go
// 存储到 context 中
ctx := context.Background()
ctx = opentracing.ContextWithSpan(ctx, span)
//....

// 其他过程获取并开始子 span
span, ctx := opentracing.StartSpanFromContext(ctx, "newspan")
defer span.Finish()
// StartSpanFromContext 会将新span保存到ctx中更新
```

或者

###### 先取出 parent span，然后在以 childof 开启span，需要手动写入新 span 到 ctx中。

```go
//获取上一级 span
parent := opentracing.SpanFromContext(ctx) 
span1 := opentracing.StartSpan("from-sayhello-1", opentracing.ChildOf(span2.Context()))
...
span1.Finish()
ctx = opentracing.ContextWithSpan(ctx, span2) //更新ctx

span2 := opentracing.StartSpan("from-sayhello-2", opentracing.ChildOf(span2.Context()))
...
span2.Finish()
ctx = opentracing.ContextWithSpan(ctx, span2) //更新ctx

```

##### tracing  grpc 调用
由于 grpc 调用和服务端都声明了 UnaryInterceptor 和 StreamInterceptor 两回调函数，因此只需要重写这两个函数，在函数中调用 opentracing 的借口进行链路追踪，并初始化客户端或者服务端时候注册进去就可以。

相应的函数已经有现成的包 `grpc-opentracing`

使用如下：
```go
var tracer opentracing.Tracer = ...
//client
conn, err := grpc.Dial(
    address,
    ... // other options
    grpc.WithUnaryInterceptor(
        otgrpc.OpenTracingClientInterceptor(tracer)),
    grpc.WithStreamInterceptor(
        otgrpc.OpenTracingStreamClientInterceptor(tracer)))


// server
s := grpc.NewServer(
    ... // other options
    grpc.UnaryInterceptor(
        otgrpc.OpenTracingServerInterceptor(tracer)),
    grpc.StreamInterceptor(
        otgrpc.OpenTracingStreamServerInterceptor(tracer)))
```

Serializing to the wire
```go
func makeSomeRequest(ctx context.Context) ... {
        if span := opentracing.SpanFromContext(ctx); span != nil {
            httpClient := &http.Client{}
            httpReq, _ := http.NewRequest("GET", "http://myservice/", nil)

            // Transmit the span's TraceContext as HTTP headers on our
            // outbound request.
            opentracing.GlobalTracer().Inject(
                span.Context(),
                opentracing.HTTPHeaders,
                opentracing.HTTPHeadersCarrier(httpReq.Header))

            resp, err := httpClient.Do(httpReq)
            ...
        }
        ...
    }
```

Deserializing from the wire

```go
http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
        var serverSpan opentracing.Span
        appSpecificOperationName := ...
        wireContext, err := opentracing.GlobalTracer().Extract(
            opentracing.HTTPHeaders,
            opentracing.HTTPHeadersCarrier(req.Header))
        if err != nil {
            // Optionally record something about err here
        }

        // Create the span referring to the RPC client if available.
        // If wireContext == nil, a root span will be created.
        serverSpan = opentracing.StartSpan(
            appSpecificOperationName,
            ext.RPCServerOption(wireContext))

        defer serverSpan.Finish()

        ctx := opentracing.ContextWithSpan(context.Background(), serverSpan)
        ...
    }
```