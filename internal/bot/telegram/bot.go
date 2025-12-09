package telegram

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"io"
	"log"
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
	bot.baseURL = fmt.Sprintf("https://api.telegram.org/bot%s", token)
	return bot
}

func (b *Bot) SetStatusVPN(ctx context.Context, status bool) error {
	statusTxt := "❌"
	if status {
		statusTxt = "✅"
	}

	timeStr := time.Now().Format("02-01 15:04")

	name := fmt.Sprintf("%s %s %s", b.BotBaseName, statusTxt, timeStr)
	urlReq := fmt.Sprintf("%s/setMyName?name=%s", b.baseURL, url.QueryEscape(name))
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, urlReq, nil)
	if err != nil {
		return errors.WithStack(err)
	}

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return errors.WithStack(err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	log.Println(string(body))
	if err != nil {
		return errors.Wrap(err, "failed to read response")
	}

	answer := Answer{}
	if errUnmarshal := json.Unmarshal(body, &answer); errUnmarshal != nil {
		return errors.WithStack(errUnmarshal)
	}

	if answer.Status {
		return nil
	}

	log.Printf("error {%d}: %s", answer.ErrorCode, answer.Description)
	return errors.WithStack(errors.New("answer.Status"))

}

func (b *Bot) SetInformationVPN(ctx context.Context, info string) error {
	return nil
}
