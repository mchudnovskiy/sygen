package sender

import (
	"fmt"
	"strings"
)

// Sender is a default interface for sending messager over deffirent transports
type Sender interface {
	Send(payload string, headers map[string]string) error
}

// NewSender is a factory method creates a concret sender by endpoint parsing
func NewSender(endpoint string) (Sender, error) {
	protocol := strings.Split(endpoint, "||")[0]
	switch strings.ToLower(protocol) {
	case "http":
		s := &httpSender{
			e: strings.Split(endpoint, "||")[1],
		}
		return s, nil
	case "mq":
		s := &mqSender{
			e: strings.Split(endpoint, "||")[1],
		}
		return s, nil
	}

	return nil, nil
}

type httpSender struct {
	e string
}

//Send method sends a message via http connection
func (hs *httpSender) Send(payload string, headers map[string]string) error {
	fmt.Printf("HTTP Sender is sending payload: %s at endpoint %s \n\n", payload, hs.e)
	return nil
}

type mqSender struct {
	e string
}

//Send method sends a message via mq connection
func (ms *mqSender) Send(payload string, headers map[string]string) error {
	fmt.Printf("MQ Sender is sending payload: %s at endpoint %s \n\n", payload, ms.e)
	return nil
}
