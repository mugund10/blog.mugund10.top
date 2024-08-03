+++
title = "How to Create a Simple Go web server "
description = "So, for my dynamic DNS project, I needed to verify my public IP. Naturally, I couldn't just use one of the million existing services. No, I had to create my own web service to do it. This guide will take you on a thrilling adventure of setting up a simple Go web server. We’ll cover the exhilarating process of handling HTTP requests and deploying the service, because who doesn't love reinventing the wheel? By the end, you’ll be rolling your eyes and wondering why you didn’t just use a pre-made solution in the first place."
date = 2024-03-03
url = "/blog/aug1"

[author]
name = "mugund10"
email = "bjmugundhan@gmail.com"
+++


# From Confusion to Clarity

*   I never really understood how the web works and never had much interest in it. It always seemed like something too complex for me. But then, I needed to work on a project that required web knowledge. So, I started learning about HTTP and web protocols. That’s when I had an Eureka moment—I discovered that every web request includes the sender's IP address. Excited by this revelation, I made a simple [handler](https://mugund10.openwaves.in/ip) in Go. It reads incoming requests, grabs the IP address, and sends it back in the response. It’s been quite a journey, from knowing nothing about web tech to making something that works!

![](https://media3.giphy.com/media/v1.Y2lkPTc5MGI3NjExOHJkNGlzZW4zYzFhbGo2a210aHYzdHB4b3pld3Bzb2UzaTZvaXhocSZlcD12MV9pbnRlcm5hbF9naWZfYnlfaWQmY3Q9Zw/IwTWTsUzmIicM/giphy.webp)


# Whats a Webserver ?

*   A web server is a program that delivers web content over HTTP by handling requests from clients. It works by listening on a specific port for incoming connections. When a client, like a web browser, wants to access a website, it first sends a request to a DNS server to find the website’s address. This request then travels through the internet, passing through routers and reaching the web server.
![](https://media4.giphy.com/media/v1.Y2lkPTc5MGI3NjExa20zcGE2Z3Q1cDBkN2FkbjM0dnYzdG84aHFldnV0b2k3Y3Z3NGgzeiZlcD12MV9pbnRlcm5hbF9naWZfYnlfaWQmY3Q9Zw/d2BCGOX8usZK0zdxw5/giphy.webp)

*   So, how does the server handle this? It uses its operating system and network hardware to manage connections and process requests. Once the server receives the request, it generates a response, which may include the web page or other content. This response is then sent back through the same network route to the client's browser, allowing the user to view the requested content.

![](https://media0.giphy.com/media/v1.Y2lkPTc5MGI3NjExNnJ2ZXhwZmYycDc4dzM4Y21zejNyMm15bjQyZTg5NmptYmRtMWI3bSZlcD12MV9pbnRlcm5hbF9naWZfYnlfaWQmY3Q9Zw/AqV8uSb8ptxyo7Wyog/giphy.webp)