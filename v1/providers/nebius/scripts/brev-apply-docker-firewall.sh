#!/bin/sh

iptables -N DOCKER-USER 2>/dev/null || true
iptables -F DOCKER-USER || true
iptables -A DOCKER-USER -m conntrack --ctstate ESTABLISHED,RELATED -j ACCEPT
iptables -A DOCKER-USER -i docker0 ! -o docker0 -j ACCEPT
iptables -A DOCKER-USER -i br+     ! -o br+     -j ACCEPT
iptables -A DOCKER-USER -i cni+    ! -o cni+    -j ACCEPT
iptables -A DOCKER-USER -i cali+   ! -o cali+   -j ACCEPT
iptables -A DOCKER-USER -i docker0 -o docker0 -j ACCEPT
iptables -A DOCKER-USER -i br+     -o br+     -j ACCEPT
iptables -A DOCKER-USER -i cni+    -o cni+    -j ACCEPT
iptables -A DOCKER-USER -i cali+   -o cali+   -j ACCEPT
iptables -A DOCKER-USER -i lo -j ACCEPT
iptables -A DOCKER-USER -j DROP
iptables -A DOCKER-USER -j RETURN

exit 0
