prefix ?= vpn-test
base-bot-name ?= VPN Status
GO ?= /usr/local/go/bin/go
build-path ?= ./build

chmod:
	chmod +x ./scripts/base-setup.sh
	chmod +x ./scripts/vpn-setup.sh
	chmod +x ./scripts/clear.sh

clear:
	./scripts/clear.sh $(prefix)

setup: chmod clear
	./scripts/base-setup.sh $(prefix)
	./scripts/vpn-setup.sh $(prefix) $(vpn-conf)

	$(GO) mod tidy
	$(GO) build -ldflags="-X 'main.botToken=$(bot-token)' -X 'main.nsName=$(prefix)-vpn-ns' -X 'main.baseBotName=$(base-bot-name)'" -o $(build-path) ./cmd


