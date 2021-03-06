#!/bin/sh

sudo iptables -t nat -A POSTROUTING -o eth0 -j MASQUERADE
sudo iptables -A FORWARD -i eth0 -o wlan0 -m state --state RELATED,ESTABLISHED -j ACCEPT
sudo iptables -A FORWARD -i wlan0 -o eth0 -j ACCEPT

#redirect http to adproxy
iptables -t nat -A PREROUTING -p tcp --dport 80 -j REDIRECT --to 9000
