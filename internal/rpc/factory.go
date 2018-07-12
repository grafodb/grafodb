package rpc

import (
	"github.com/google/flatbuffers/go"
)

func NewRequestMessage(cmd uint16, version byte) *Message {
	builder := flatbuffers.NewBuilder(0)

	RequestStart(builder)
	RequestAddCommand(builder, cmd)
	RequestAddVersion(builder, version)
	request := RequestEnd(builder)

	builder.Finish(request)

	return &Message{Payload: builder.FinishedBytes()}
}

func NewRequestMessageWithPayload(cmd uint16, version byte, payload []byte) *Message {
	builder := flatbuffers.NewBuilder(0)

	RequestStartPayloadVector(builder, len(payload))
	for i := len(payload) - 1; i >= 0; i-- {
		builder.PrependByte(payload[i])
	}
	payloadOffset := builder.EndVector(len(payload))

	RequestStart(builder)
	RequestAddCommand(builder, cmd)
	RequestAddVersion(builder, version)
	RequestAddPayload(builder, payloadOffset)

	request := RequestEnd(builder)

	builder.Finish(request)

	return &Message{Payload: builder.FinishedBytes()}
}

