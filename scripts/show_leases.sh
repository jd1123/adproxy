#!/bin/bash

grep "^lease" /var/lib/dhcp/dhcpd.leases | sort | uniq | wc -l
