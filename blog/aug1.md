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

*   1. it serves something that a 
