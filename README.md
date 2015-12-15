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

## How do I get it to work?

1. Set up your Pi with Raspbian. Update it (apt-get update, apt-get upgrade)
2. Clone this repo on the pi
3. Run setup\_env.sh in the scripts folder (this will also clone this repo into the correct directory)
4. Enjoy.

## Why do this?

A friend asked me this once. I'll give you the same answer I gave him. I don't know!

## Notes

This is a very buggy project. It sucessfully blocks ads for what my wife watches on a regular basis. If you like it, please feel free to hack on it and send a pull request.
