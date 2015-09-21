#!/bin/bash

enab(){
	sudo update-rc.d isc-dhcp-server defaults
	sudo update-rc.d hostapd defaults
}

disable(){
	sudo update-rc.d -f isc-dhcp-server remove
	sudo update-rc.d -f hostapd remove
}

case $1 in
	"enable")
		enab
	;;
	
	"disable")
		disable
	;;

	*)
		echo "usage: enable or disable"
	;;
esac

exit 0
