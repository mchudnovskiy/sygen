package sender

import (
	"errors"
	"fmt"
	"github.com/ibm-messaging/mq-golang-jms20/mqjms"
	"net/http"
	"strconv"
	"strings"
	"io"
)

// Sender is a default interface for sending messager over deffirent transports
type Sender interface {
	Send(payload string, headers map[string]string) error
}

// NewSender is a factory method creates a concret sender by endpoint parsing
func NewSender(endpoint string) (Sender, error) {
	protocol := strings.Split(endpoint, "://")[0]
	switch strings.ToLower(protocol) {
	case "http":
		s := &httpSender{
			e: endpoint,
		}
		return s, nil
	case "mq":
		s, err := newMqSender(strings.Split(endpoint, "://")[1])
		return s, err
	}
	return nil, nil
}

type httpSender struct {
	e string
}

//Send method sends a message via http connection
func (hs *httpSender) Send(payload string, headers map[string]string) error {
	fmt.Printf("HTTP Sender is sending payload: %s at endpoint %s \n", payload, hs.e)
	resp, err := http.Get(hs.e)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	
	s, err := copyToString(resp.Body)
	if err!=nil {
		return err
	}
	fmt.Printf("Response: %s \n\n", s)

	return nil
}

func copyToString(r io.Reader) (res string, err error) {
	var sb strings.Builder
	if _, err = io.Copy(&sb, r); err == nil {
		res = sb.String()
	}
	return
}

type mqSender struct {
	connectionParams map[string]string
	cf               *mqjms.ConnectionFactoryImpl
	ctx              *mqjms.ContextImpl
	q                mqjms.QueueImpl
}

func newMqSender(endpoint string) (Sender, error) {
	mqs := &mqSender{}
	err := mqs.initConnection(endpoint)
	return mqs, err
}

func (mqs *mqSender) initConnection(e string) error {
	mqs.connectionParams = make(map[string]string)
	mqs.connectionParams["qm"] = strings.Split(e, "/")[3]
	mqs.connectionParams["channel"] = strings.Split(e, "/")[2]
	mqs.connectionParams["host"] = strings.Split(e, "/")[0]
	mqs.connectionParams["port"] = strings.Split(e, "/")[1]
	mqs.connectionParams["queue"] = strings.Split(e, "/")[4]

	port, _ := strconv.Atoi(mqs.connectionParams["port"])
	mqs.cf = &mqjms.ConnectionFactoryImpl{
		ChannelName: mqs.connectionParams["channel"],
		Hostname:    mqs.connectionParams["host"],
		QMName:      mqs.connectionParams["qm"],
		PortNumber:  port,
	}
	ctx, err := mqs.cf.CreateContext()
	if err == nil {
		mqs.ctx = ctx.(*mqjms.ContextImpl)
	} else {
		return errors.New(err.GetErrorCode())
	}

	mqs.q = (ctx.CreateQueue(mqs.connectionParams["queue"])).(mqjms.QueueImpl)

	return nil
}

//Send method sends a message via mq connection
func (mqs *mqSender) Send(payload string, headers map[string]string) error {

	mqs.ctx.CreateProducer().Send(mqs.q, mqs.ctx.CreateTextMessageWithString(payload))

	return nil
}
