+++
title = "How to Create a Simple Go web server "
description = "So, for my dynamic DNS project, I needed to verify my public IP. Naturally, I couldn't just use one of the million existing services. No, I had to create my own web service to do it. This guide will take you on a thrilling adventure of setting up a simple Go web server. We’ll cover the exhilarating process of handling HTTP requests and deploying the service, because who doesn't love reinventing the wheel? By the end, you’ll be rolling your eyes and wondering why you didn’t just use a pre-made solution in the first place."
date = 2024-08-01
url = "/blog/9998-gowebserver"

[author]
name = "mugund10"
email = "bjmugundhan@gmail.com"
+++


# From Confusion to Clarity

*   I never really understood how the web works and never had much interest in it. It always seemed like something too complex for me. But then, I needed to work on a project that required web knowledge. So, I started learning about HTTP and web protocols. That’s when I had an Eureka moment—I discovered that every web request includes the sender's IP address. Excited by this revelation, I made a simple [webserver](https://mugund10.openwaves.in/ip) in Go. It reads incoming requests, grabs the IP address, and sends it back to the response. It’s been quite a journey, from knowing nothing about web tech to making something that works!

![](https://media3.giphy.com/media/v1.Y2lkPTc5MGI3NjExOHJkNGlzZW4zYzFhbGo2a210aHYzdHB4b3pld3Bzb2UzaTZvaXhocSZlcD12MV9pbnRlcm5hbF9naWZfYnlfaWQmY3Q9Zw/IwTWTsUzmIicM/giphy.webp)


# Whats a Webserver ?

*   A web server is a program that delivers web content over HTTP by handling requests from clients. It works by listening on a specific port for incoming connections. When a client, like a web browser, wants to access a website, it first sends a request to a DNS server to find the website’s address. This request then travels through the internet, passing through routers and reaching the web server.
![](https://media4.giphy.com/media/v1.Y2lkPTc5MGI3NjExa20zcGE2Z3Q1cDBkN2FkbjM0dnYzdG84aHFldnV0b2k3Y3Z3NGgzeiZlcD12MV9pbnRlcm5hbF9naWZfYnlfaWQmY3Q9Zw/d2BCGOX8usZK0zdxw5/giphy.webp)

*   So, how does the server handle this? It uses its operating system and network hardware to manage connections and process requests. Once the server receives the request, it generates a response, which may include the web page or other content. This response is then sent back through the same network route to the client's browser, allowing the user to view the requested content.

![](https://media0.giphy.com/media/v1.Y2lkPTc5MGI3NjExNnJ2ZXhwZmYycDc4dzM4Y21zejNyMm15bjQyZTg5NmptYmRtMWI3bSZlcD12MV9pbnRlcm5hbF9naWZfYnlfaWQmY3Q9Zw/AqV8uSb8ptxyo7Wyog/giphy.webp)

#   Lets GO Now

*   Go has a built-in package for HTTP, which makes setting up a web server straight forward and easy. Here’s a simple example of how you can create a basic web server in Go:

    ```GO
    package main

    import (
        "fmt"
        "net/http"
    )

    func handler(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Your IP address is: %s", r.RemoteAddr)
    }

    func main() {
        http.HandleFunc("/", handler)
        fmt.Println("Starting server on :8080...")
        http.ListenAndServe(":8080", nil)
    }

    ```

- `http.ListenAndServe(":8080", nil)` starts the server and listens on port 8080. The second argument, `nil`, indicates that we are using the default `ServeMux` for routing.
- Since we didn’t create a custom multiplexer (mux), the default mux is used. This is why we passed `nil` as the second argument. The `http.HandleFunc` function maps the root path ("/") to the `handler` function, which will be called for all GET requests to this path.
- The `handler` function writes the client’s IP address to the response, allowing us to see the incoming IP address for each request.

![](https://media4.giphy.com/media/v1.Y2lkPTc5MGI3NjExc2lpZ2s0cHplM2syYTg1cmg3ajA2ZnNoaDNrYmNmNXB3YXIzdHBiOCZlcD12MV9pbnRlcm5hbF9naWZfYnlfaWQmY3Q9Zw/U634xW7LKU0sU/giphy.webp)



here is the sample full response and i mapped it with the same struct fields of [http.Request](https://pkg.go.dev/net/http#Request) for a good understanding
```
{GET / HTTP/1.1 1 1 map[Accept:[text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/png,image/svg+xml,*/*;q=0.8] Accept-Encoding:[gzip, deflate, br, zstd] Accept-Language:[en-US,en;q=0.5] Connection:[keep-alive] Cookie:[userToken=9fzz3o7786gm76adfwlth; wh_theme=light; cookie-consent=true] Priority:[u=0, i] Sec-Fetch-Dest:[document] Sec-Fetch-Mode:[navigate] Sec-Fetch-Site:[none] Sec-Fetch-User:[?1] Upgrade-Insecure-Requests:[1] User-Agent:[Mozilla/5.0 (X11; Linux x86_64; rv:128.0) Gecko/20100101 Firefox/128.0]] {} <nil> 0 [] false localhost:8102 map[] map[] <nil> map[] 127.0.0.1:36302 / <nil> <nil> <nil> 0xc00008c0a0 0xc00012a180 [] map[]}
``` 


```GO
type Request struct {
    Method     string            // "GET"
    RequestURI string            // "/"
    Proto      string            // "HTTP/1.1"
    ProtoMajor int               // 1
    ProtoMinor int               // 1
    Header     Header            // map[Accept:[text/html,...] ...]
    Body       io.ReadCloser     // <nil>
    GetBody    func() (io.ReadCloser, error) // <nil>
    ContentLength int64          // 0
    TransferEncoding []string   // []
    Close      bool              // false
    Host       string            // "localhost:8102"
    Form       url.Values        // map[]
    PostForm   url.Values        // map[]
    MultipartForm *multipart.Form // <nil>
    Trailer    Header            // map[]
    RemoteAddr string            // "127.0.0.1:36302"
    RequestURI string            // "/"
    TLS        *tls.ConnectionState // <nil>
    Cancel     <-chan struct{}   // <nil>
    Response   *http.Response    // <nil>
    ctx        context.Context   // 0xc00008c0a0
    // Additional fields for request context and other details
}
```

#   summary

-   In this guide, we built a basic Go web server that responds with your IP address. We explored how web servers work and utilized Go's net/http package for a straight forward setup We didn’t cover everything, though:

    1.  HTTP Methods: We only used GET. There are others like POST and PUT.
    2.  Error Handling: It’s important to handle errors when starting the server.
    -   For more info, check out the [Go documentation](https://go.dev/doc/).

-   Feel free to visit mugund10.openwaves.in/ip to see the implementation in action. This example offers a hands-on look at handling web requests and responses. Whether you're new to web development or experimenting with new concepts, it's a valuable starting point.

- okay Thats it Bye

![](https://media0.giphy.com/media/v1.Y2lkPTc5MGI3NjExem55MWtqamJ2dzFxczN4bTl5NzY3ZHM0ODVzZGpiejl0OHA5d2loMyZlcD12MV9pbnRlcm5hbF9naWZfYnlfaWQmY3Q9Zw/w89ak63KNl0nJl80ig/giphy.webp)