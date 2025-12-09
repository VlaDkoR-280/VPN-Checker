#!/bin/bash
set -e

# arguments:
# {namespace name} {bridge name}

PREFIX=$1

NS="$PREFIX"-vpn-ns
BR="$PREFIX"-br-vpn

VETH1="$PREFIX"-veth1
VETH2="$PREFIX"-veth2
VETH3="$PREFIX"-veth3
VETH4="$PREFIX"-veth4


ip netns add $NS
echo "netns added"

ip link add $VETH1 type veth peer name $VETH2
ip link add $VETH3 type veth peer name $VETH4
echo "added $VETH1, $VETH2 of peer $VETH2, $VETH4"

ip link set $VETH2 netns $NS
ip link set $VETH4 netns $NS
echo "move $VETH2,4 to ns"

ip link add name $BR type bridge
ip addr add 192.168.89.1/24 dev $BR
ip link set $BR up
echo "up bridge"

ip link set $VETH1 master $BR
ip link set $VETH3 master $BR
ip link set $VETH1 up
ip link set $VETH3 up
echo "up $VETH1, $VETH2"

echo 1 > /proc/sys/net/ipv4/ip_forward
iptables -t nat -A POSTROUTING -o eth0 -j MASQUERADE
iptables -A FORWARD -i %BR -o eth0 -j ACCEPT
iptables -A FORWARD -i eth0 -o $BR -j ACCEPT

ip netns exec $NS ip link set $VETH2 up
ip netns exec $NS ip link set $VETH4 up
ip netns exec $NS ip addr add 192.168.89.2/24 dev $VETH2

ip netns exec $NS ip route add default via 192.168.89.1



