# Ad Proxy

This is the code and setup scripts to turn your Raspberry pi into a hardware adblocker. 

This package turns your Raspberry Pi into a wireless access point that has an HTTP filter that blocks ads for various services. You can write a "module" to block ads for a particular service, or something more general like import a list of blocked domains or URLs. It is trying to be as flexible as possible to suit your needs.

Currently, there are two modules: Hulu and Xfinity. It blocks ads for both of these services, so it makes a nice router if you use either of theseand hate ads. It has been tested on the iPad clients for both, so YMMV.

This was created because my wife watches a lot of television on her iPad, and hates ads. She is somewhat pleased with it, so I a happy husband :)

## How did I do it?

I started by intercepting app traffic using wireshark and monitored the traffic. It turns out that most video on demand is streamed via http.
The app hits an adserver to download the ad manifest for a tv show via HTTP, and I filter these requests and return an empty response body. This can be extended to return any type of response you like.

The basis of this software is goproxy (http://github.com/elazarl/goproxy). 

## Things you need to get started:
To get this to work you need:
* 1 Raspberry Pi
* 1 Wifi Dongle (this assumes the Edimax dongle from Amazon)
* 1 Ethernet Cable

## How do I get it to work?

1. Hook up your Pi to your home router with an ethernet cord. Power it up. 
2. Set up your Pi with Raspbian. Update it (apt-get update, apt-get upgrade)
3. Clone this repo on the pi
4. Run setup\_env.sh in the scripts folder (this will also clone this repo into the correct directory)
5. Enjoy.

## Notes

This is a very buggy project. It sucessfully blocks ads for what my wife watches on a regular basis. If you like it, please feel free to contribute to it and send a pull request.

I have some decent documentation that needs to be moved to github. I will do this if anyone cares.
