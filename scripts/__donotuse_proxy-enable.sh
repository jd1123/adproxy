#!/bin/bash

sudo iptables -t nat -A PREROUTING -i wlan0 -s ! 192.168.42.1 -p tcp --dport 80 -j DNAT --to 192.168.42.1:9000
sudo iptables -t nat -A POSTROUTING -o eth0 -s 192.168.42.0/24 -d 192.168.42.1 -j SNAT --to 192.168.42.1
sudo iptables -A FORWARD -s 192.168.42.0/24 -d 192.168.42.1 -i wlan0 -o eth0 -p tcp --dport 9000 -j ACCEPT

#redirect http to adproxy

iptables -t nat -A PREROUTING -p tcp --dport 80 -j REDIRECT --to 9000
