#!/bin/bash

install_utils(){
	sudo apt-get -y install git mecurial vim build-essential
}

install_go(){
	# from source
	cd
	git clone https://github.com/golang/go
	cd go
	git checkout release-branch.go1.4
	cd src
	./all.bash
	cd
	mkdir -p ~/code/go
	echo "export PATH=$PATH:$HOME/go/bin" >> .bashrc
	echo "export GOPATH=$HOME/code/go" >> .bashrc
	echo
	echo "Your code will now sit in ~/code/go"
	source ~/.bashrc
	echo "Go 1.4 should work now..."	
}

install_dhcp_server(){
	echo "Installing the dhcp server..."
	apt-get install isc-dhcp-server
	echo "About to edit /etc/dhcpd.conf - this may or may not be the way you want to"
	echo "configure it."
	cd /etc/dhcp
	sudo cp dhcpd.conf dhcpd.conf.bak
	echo "subnet 192.168.42.0 netmask 255.255.255.0 {
range 192.168.42.10 192.168.42.50;
option broadcast-address 192.168.42.255;
option routers 192.168.42.1;
default-lease-time 600;
max-lease-time 7200;
option domain-name "local";
option domain-name-servers 8.8.8.8, 8.8.4.4;
}" | sudo tee --append /etc/dhcp/dhcpd.conf

}

install_hostapd(){
	echo "Installing hostapd..."
	apt-get -y install hostapd
	cd /etc/hostapd
	sudo mv hostapd.conf hostapd.conf.bak
	echo "About to make a conf file for hostapd. You will need to change this depending"
	echo "on which driver you use. rtl871xdrv also requires a pre-compiled binary which you"
	echo "can get from https://learn.adafruit.com/setting-up-a-raspberry-pi-as-a-wifi-access-point/install-software. This is"
	echo "specifically for the Edimax dongle that I bought of off amazon."
	echo "interface=wlan0
#driver=nl80211
driver=rtl871xdrv
ssid=AdFree
hw_mode=g
channel=6
macaddr_acl=0
auth_algs=1
ignore_broadcast_ssid=0
wpa=2
wpa_passphrase=g3t 0ff my l4wn
wpa_key_mgmt=WPA-PSK
wpa_pairwise=TKIP
rsn_pairwise=CCMP
" | sudo tee /etc/hostapd/hostapd.conf
	echo "Edit /etc/default/hostapd and replace #DAEMON_CONF with the right path (and uncomment it)."
	echo "That right path is /etc/hostapd/hostapd.conf"
}

install_adproxy(){
	echo "About to install the adproxy code..."
	cd
	git clone <repo_name_here>
	cd
}

usage(){
	echo "This will install the toolchain necessary to run adproxy on the raspberry pi."
	echo "It assumes you have an up to date installation of Raspian OS"
	echo "Some things will be installed from source so get yourself a diet coke and relax."
	echo "It will ask for sudo when necessary. DO NOT RUN AS ROOT."
	echo
	echo "You will also need an ethernet connection and a Wifi dongle to access it."
	echo "This script has NOT been tested. Use at your own risk."
	echo "See https://learn.adafruit.com/setting-up-a-raspberry-pi-as-a-wifi-access-point/install-software for info on setting up the RPi as an AP." 
}

setup_iptables(){
	#This needs to be made into an init script.
	sudo sh -c "echo 1 > /proc/sys/net/ipv4/ip_forward"   #use sed to change the /etc/sysctl.conf to make this permanent
	sudo iptables -t nat -A POSTROUTING -o eth0 -j MASQUERADE
    sudo iptables -A FORWARD -i eth0 -o wlan0 -m state --state RELATED,ESTABLISHED -j ACCEPT
    sudo iptables -A FORWARD -i wlan0 -o eth0 -j ACCEPT	
}

finally(){
	echo "Finally you need to start all services at startup. There are some problems with hostapd."
	sudo update-rc.d hostapd enable
	sudo update-rc.h isc-dhcp-server enable
	setup_iptables
}

usage

if [ "$EUID" -eq 0 ]
	then echo -e "\nI said DO NOT RUN AS ROOT, bro..."
	exit 1
fi

install_utils
install_go
install_dhcp_server
install_hostapd
finally
