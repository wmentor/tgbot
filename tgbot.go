package tgbot

import (
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/wmentor/log"
)

const (
	OptDisablePreview int = iota + 1
	OptHTML
	OptSilently
)

type Bot struct {
	token   string
	timeout time.Duration
}

func New(token string) *Bot {

	return &Bot{
		token: token,
	}

}

func (b *Bot) Send(cid int64, msg string, opt ...int) {

	opts := make(map[int]bool)

	for _, v := range opt {
		opts[v] = true
	}

	vals := url.Values{}

	vals.Add("chat_id", fmt.Sprintf("-100%d", cid))
	vals.Add("text", msg)

	if opts[OptDisablePreview] {
		vals.Add("disable_web_page_preview", "1")
	}

	if opts[OptHTML] {
		vals.Add("parse_mode", "HTML")
	}

	if opts[OptSilently] {
		vals.Add("disable_notification", "1")
	}

	client := http.Client{
		Timeout: b.timeout,
	}

	uri := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", b.token)

	if _, err := client.PostForm(uri, vals); err != nil {
		log.Error(err.Error())
	}
}
