# Example of Server and Proxy in Go (Gerpo)

This is an example of a [websocket](/server) with its own [websocket proxy](/server-proxy). In general terms, the interaction between the two is demonstrated. It should be noted that the websocket is fully functional without the proxy and that the [gorilla websocket](https://github.com/gorilla/websocket) and [gin](https://github.com/gin-gonic/gin) dependencies are used.

## Special Thanks
- On the [lwebapp page](https://lwebapp.com/en/post/go-websocket-simple-server) you can see how to create a websocket in a simple way.
- The [websocketproxy](https://github.com/koding/websocketproxy) repository contains a library for a websocket proxy agnostic to the framework you want to use.
- I find the following [page](https://dev.to/justlorain/how-to-apply-reverse-proxy-over-websocket-27ml) to be a valuable resource if you are trying to create a reverse proxy.