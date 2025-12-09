prefix ?= vpn-test
base-bot-name ?= "VPN Status"

chmod:
	chmod +x ./scripts/base-setup.sh
	chmod +x ./scripts/vpn-setup.sh

setup: chmod
	./scripts/base-setup.sh $(prefix)
	./scripts/vpn-setup.sh $(prefix) $(vpn-conf)

	go mod tidy
	go build -ldflags="-X 'main.botToken=$(bot-token)' -X 'main.nsName=$(prefix)-vpn-ns -X 'main.baseBotName=$(base-bot-name)'" -o /tmp/vpn-checker-program ./cmd

	ip netns exec $(ns-name) cp /tmp/vpn-checker-program /etc/







