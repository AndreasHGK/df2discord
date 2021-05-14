package df2discord

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"net/http"
	"net/url"
)

func NewWebhook(url string) Webhook {
	w := Webhook{url: url, c: make(chan string)}
	go w.waitForMessage() // start a goroutine to send POST requests
	return w
}

type Webhook struct {
	logger 	logrus.FieldLogger
	url    	string
	c		chan string
}

func (w Webhook) WithLogger(logger logrus.FieldLogger) Webhook {
	w.logger = logger
	return w
}

// Message sends a message to the channel where the webhook is set up.
func (w Webhook) Message(a ...interface{}) {
	w.c <- filterColor(fmt.Sprint(a...))
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

func (w Webhook) waitForMessage() {
	for {
		err := w.postMessage(<-w.c)
		if err != nil {
			w.logger.Errorln(err)
		}
	}
}