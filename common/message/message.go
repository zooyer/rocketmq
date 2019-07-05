package message

import "fmt"

type Message struct {
	topic         string
	flag          int
	properties    map[string]string
	body          []byte
	transactionId string
}

func New() *Message {
	return &Message{
		properties: make(map[string]string),
	}
}

func WithBody(topic string, body []byte) *Message {
	return WithAll(topic, "", "", 0, body, true)
}

func WithTags(topic, tags string, body []byte) *Message {
	return WithAll(topic, tags, "", 0, body, true)
}

func WithKeys(topic, tags, keys string, body []byte) *Message {
	return WithAll(topic, tags, keys, 0, body, true)
}

func WithAll(topic, tags, keys string, flag int, body []byte, waitStoreMsgOK bool) *Message {
	var msg = New()
	msg.topic = topic
	msg.flag = flag
	msg.body = body

	if tags != "" && len(tags) > 0 {
		msg.SetTags(tags)
	}
	if keys != "" && len(keys) > 0 {
		msg.SetKeys(keys)
	}
	msg.SetWaitStoreMsgOK(waitStoreMsgOK)

	return msg
}

func (m *Message) SetTags(tags string) {
	m.properties[PropertyTags] = tags
}

func (m *Message) SetKeys(keys string) {
	m.properties[PropertyKeys] = keys
}

func (m *Message) SetWaitStoreMsgOK(flag bool) {
	m.properties[PropertyWaitStoreMsgOk] = fmt.Sprintf("%v", flag)
}
