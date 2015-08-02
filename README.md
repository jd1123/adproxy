# Ad Proxy in Go for Xfinity

This is a program to block ads for the Xfninity app on the iPad. It uses an HTTP and CONNECT proxy to filter adserver requests.
It also blocks analytics requests from the app. I don't know what these are but they can't be good.

## How did I do it?

I started by intercepting app traffic using wireshark and monitored the traffic. It turns out that most video on demand is streamed via http.
The app hits an adserver to download the ad manifest for a tv show via HTTP, and I filter these requests and return an empty response body.
Right now, it works for Xfinity, but likely can be extended for Hulu as well.

## Can it be broken?

If Xfinity ever decides to implment ads over TLS, it would probably make it much harder to do this, but not impossible. One would need the TLS keys or implement sslstrip over the protocol and it would possibly still work. This would require much more technical knowhow than I have at the moment. Would be a fun project though!

## Shout outs

Thank you to elazarl (https://github.com/elazarl) for the goproxy project, on which this is based (http://github.com/elazarl/goproxy). 

## How do I get it to work?

Right now, you need a linux box with go installed. Build it, then run the server:
```
cd /path/to/project
go build
./adproxy
```
Point your device to your server as an http proxy, and presto, it works.

## Why do this?

A friend asked me this once. I'll give you the same answer I gave him. I don't know!

## TODO

* Make it work for Hulu
* Allow devices without an HTTP proxy setting to use this as well
