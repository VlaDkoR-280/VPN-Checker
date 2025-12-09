ns-name ?= vpn-test
bridge-name ?= br-vpn

chmod:
	chmod +x ./scripts/base-setup.sh
	chmod +x ./scripts/vpn-setup.sh

setup: chmod
	./scripts/base-setup.sh $(ns-name) $(bridge-name)
	./scripts/vpn-setup.sh $(ns-name) $(bridge-name)

	go mod tidy
	go build -o /tmp/vpn-checker-program ./cmd

	ip netns exec $(ns-name) cp /tmp/vpn-checker-program /etc/







