package main

import (
	"context"
	"github.com/VlaDkoR-280/VPN-Checker/internal/adapters/vpn"
	"github.com/VlaDkoR-280/VPN-Checker/internal/bot/telegram"
	"github.com/VlaDkoR-280/VPN-Checker/internal/panels"
	"log"
	"os"
	"time"
)

var (
	baseBotName string
	botToken    string
	nsName      string
)

func main() {
	vpnAdapter := vpn.InitAdapter(nsName)

	bot := telegram.InitBot(baseBotName, botToken)

	panelInfo := panels.InitPanel(bot)

	vpnStatus := false
	tCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err := vpnAdapter.CheckStatus(tCtx)
	if err != nil {
		log.Printf("Error checking VPN status: %+v", err)
	} else {
		vpnStatus = true
	}

	panelCtx, panelCancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer panelCancel()
	if errSetVpnStatus := panelInfo.SetVpnStatus(panelCtx, vpnStatus); errSetVpnStatus != nil {
		log.Printf("Error setting VPN status: %+v", errSetVpnStatus)
		os.Exit(1)
	}
	log.Printf("VPN status set to %t", vpnStatus)
}
