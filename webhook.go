package df2discord

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"net/http"
	"net/url"
)

func NewWebhook(url string) Webhook {
	return Webhook{url: url}
}

type Webhook struct {
	logger logrus.FieldLogger
	url    string
}

func (w Webhook) WithLogger(logger logrus.FieldLogger) Webhook {
	w.logger = logger
	return w
}

func (w Webhook) postMessage(m string) error {
	data := url.Values{
		"content": {m},
	}
	resp, err := http.PostForm(w.url, data)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return nil
}

// Message sends a message to the channel where the webhook is set up.
func (w Webhook) Message(a ...interface{}) {
	go func() {
		err := w.postMessage(filterColor(fmt.Sprint(a...)))
		if err != nil {
			w.logger.Errorln(err)
		}
	}()
}
