# Ad Proxy

This is the code and setup scripts to turn your Raspberry pi into a hardware adblocker. Connect to the AP and it blocks all ads according to the filtes you set

## How did I do it?

I started by intercepting app traffic using wireshark and monitored the traffic. It turns out that most video on demand is streamed via http.
The app hits an adserver to download the ad manifest for a tv show via HTTP, and I filter these requests and return an empty response body.
Right now, it works for Xfinity, but likely can be extended for Hulu as well.

## Hardware
To get this to work you need:
* 1 Raspberry Pi
* 1 Wifi Dongle (this assumes the Edimax dongle from Amazon)
* 1 Ethernet Cable

## Can it be broken?

Yes, TLS will break this. Need to work on this.

## Shout outs

Thank you to elazarl (https://github.com/elazarl) for the goproxy project, on which this is based (http://github.com/elazarl/goproxy). 

## How do I get it to work?

It's not done. Run the setup script it will turn your pi into a shitty router. It needs to be looked at because it will only work for certain wifi dongles. You also need to follow additional instructions about getting hostapd to work for the RASOMETHINGOROTHER chipset.

## Why do this?

A friend asked me this once. I'll give you the same answer I gave him. I don't know!

## TODO

* Make it work for Hulu
* Allow devices without an HTTP proxy setting to use this as well
