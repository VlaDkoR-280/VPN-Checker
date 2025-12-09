#!/bin/bash

PREFIX=$1

NS="$PREFIX"-vpn-ns
BR="$PREFIX"-br-vpn

VETH1="$PREFIX"-veth1
VETH2="$PREFIX"-veth2
VETH3="$PREFIX"-veth3
VETH4="$PREFIX"-veth4

iptables -t nat -D POSTROUTING -o eth0 -j MASQUERADE 2>/dev/null || true
iptables -D FORWARD -i $BR -o eth0 -j ACCEPT 2>/dev/null || true
iptables -D FORWARD -i eth0 -o $BR -j ACCEPT 2>/dev/null || true

ip link del VETH1
ip link del VETH3

ip link del $BR
ip netns del $NS
