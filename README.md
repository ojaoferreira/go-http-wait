# HTTP WAIT

## A simple Go web server
So for this example I created a simple web server which has 2 different endpoints.

Namely:
- /timeout/second/{second}
- /health

## Building the container

The Dockerfile is included in the repo. Building it, is simple. Simply run:
`docker build . -t http-wait`

To run the container, simply run:
`docker run -p 8080:8080 http-wait`

Or just use the container which is on my dockerhub:
`docker run -P -p 8080:8080 johnjohnofficial/http-wait`