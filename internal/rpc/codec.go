package rpc

import (
	"fmt"
)

type MessageCodec struct{}

func (cb *MessageCodec) Marshal(v interface{}) ([]byte, error) {
	p, ok := v.(*Message)
	if !ok {
		return nil, fmt.Errorf("invalid message type: %+v", v)
	}
	return p.Payload, nil
}

func (cb *MessageCodec) Unmarshal(data []byte, v interface{}) error {
	p, ok := v.(*Message)
	if !ok {
		return fmt.Errorf("invalid message type: %+v", v)
	}
	p.Payload = data
	return nil
}

func (cb *MessageCodec) String() string {
	return "rpc.MessageCodec"
}
