#!/bin/bash
set -e
PREFIX=$1

NS="$PREFIX"-vpn-ns
VPN=$2
#argument
# {path/to/vpn/conf}

ip netns exec $NS cp $VPN /etc/wireguard/"$PREFIX"-wg.conf
ip netns exec $NS wg-quick up "$PREFIX"-wg