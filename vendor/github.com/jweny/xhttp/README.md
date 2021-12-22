# xhttp
## Intro

http client for scanner

应用于扫描器场景下的http基础库。

1. client

   - 精准的http client配置：目前支持支持19项
   - 多client共享cookie
   - 跳转策略
   - 失败重试
   - 代理
   - tls
   - limiter：qps限制
2. request

   - context
   - trace
   - getbody：获取请求body
   - getRaw：获取请求报文
3. response

   - getLatency：发起请求到收到响应的整个持续时间，可用于判断时间延时场景，如盲注

   - getbody：获取响应body
   - getRaw：获取响应报文
4. requestMiddleware：请求发起之前，对请求的修饰
   - context
   - method 限制策略
   - 启用 trace 
   - 根据配置修改header
   - 根据配置修改cookie
5. responseMiddleware：响应获取后，对响应的处理
   - 读body
   - 响应长度限制策略
6. debug模式：debug模式下将打印请求和响应完整信息
7. 完整的 testhttp server

## Install

```
go get github.com/jweny/xhttp
```

## Demo

```go
//获取默认配置
options := NewDefaultClientOptions()
// 如果要继承cookie，传入cookie jar；否则填nil。
cookieJar, _ := cookiejar.New(&cookiejar.Options{PublicSuffixList: publicsuffix.List})

// 创建client
client, err := NewClient(options, nil)

// 构造请求
hr, _ := http.NewRequest("GET", "<TARGET URL>" , nil)
req := &Request{RawRequest: hr,}

// 发起请求
ctx := context.Background()
resp, err := client.Do(ctx, req)
```

## Todo

- errorHook
- others...

## Ref

- https://github.com/go-resty/resty
- https://github.com/hashicorp/go-retryablehttp
