package telegram

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"io"
	"net/http"
	"net/url"
	"time"
)

type Bot struct {
	Token       string
	baseURL     string
	BotBaseName string
}

func InitBot(botBaseName, token string) *Bot {
	bot := &Bot{Token: token, BotBaseName: botBaseName}
	bot.baseURL = fmt.Sprintf("https://api.telegram.org/bot%s/", token)
	return bot
}

func (b *Bot) SetStatusVPN(ctx context.Context, status bool) error {
	statusTxt := ""
	if status {
		statusTxt = "✅"
	} else {
		statusTxt = "❌"
	}

	timeStr := time.Now().Format("02-01 15:04")

	name := fmt.Sprintf("%s %s %s", b.BotBaseName, statusTxt, timeStr)
	urlReq := fmt.Sprintf("%s/setMyName?name=%s", b.baseURL, url.QueryEscape(name))
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, urlReq, nil)
	if err != nil {
		return errors.WithStack(err)
	}
	defer req.Body.Close()

	body, _ := io.ReadAll(req.Body)
	answer := Answer{}
	if errUnmarshal := json.Unmarshal(body, &answer); errUnmarshal != nil {
		return errors.WithStack(errUnmarshal)
	}

	if answer.Status == "ok" && answer.Result == "true" {
		return nil
	}
	return errors.WithStack(errors.New(answer.Result))

}

func (b *Bot) SetInformationVPN(ctx context.Context, info string) error {
	return nil
}
