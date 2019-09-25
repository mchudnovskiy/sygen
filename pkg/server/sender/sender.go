package sender

import (
	"fmt"
	//	"github.com/ibm-messaging/mq-golang-jms20/mqjms"
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
		s, err := newMqSender(strings.Split(endpoint, "||")[1])
		return s, err
	}

	return nil, nil
}

func newMqSender(endpoint string) (Sender, error) {
	return make(mqSender), nil
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
	connectionParams map[string]string
}

func (mqs *mqSender) initConnection(e string) error {
	mqs.connectionParams = make(map[string]string)
	mqs.connectionParams["qm"] = strings.Split(e, "/")[3]
	mqs.connectionParams["channel"] = strings.Split(e, "/")[2]
	mqs.connectionParams["host"] = strings.Split(e, "/")[0]
	mqs.connectionParams["port"] = strings.Split(e, "/")[1]
	mqs.connectionParams["queue"] = strings.Split(e, "/")[4]
}

//Send method sends a message via mq connection
func (mqs *mqSender) Send(payload string, headers map[string]string) error {

	// cf := &mqjms.ConnectionFactoryImpl {
	// 	ChannelName = mqs.connectionParams["channel"],
	// 	Hostname = mqs.connectionParams["host"],
	// 	QMName = mqs.connectionParams["qm"],
	// 	PortNumber = mqs.connectionParams["port"],
	// }
	// ctx, err := cf.CreateContext()
	// queue = ctx.CreateQueue(mqs.connectionParamsp["queue"])
	// errSend := context.CreateProducer().Send(queue, context.CreateTextMessageWithString(payload))

	return nil
}
