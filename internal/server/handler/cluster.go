package handler

import (
	"github.com/google/flatbuffers/go"
	"github.com/grafodb/grafodb/internal/rpc"
)

func (h Handler) clusterMemberList(version byte, payload []byte) (*rpc.Message, error) {
	builder := flatbuffers.NewBuilder(0)

	// serialize members
	members := h.cluster.Members()
	membersCount := len(members)
	memberOffsets := make([]flatbuffers.UOffsetT, membersCount)

	for i, m := range members {
		id := builder.CreateString(m.ID)
		addr := builder.CreateString(m.Addr)
		status := builder.CreateString(m.Status)
		rpc.MemberStart(builder)
		rpc.MemberAddId(builder, id)
		rpc.MemberAddAddr(builder, addr)
		rpc.MemberAddStatus(builder, status)
		member := rpc.MemberEnd(builder)
		memberOffsets[i] = member
	}

	rpc.MemberListStartMembersVector(builder, membersCount)
	for _, offset := range memberOffsets {
		builder.PrependUOffsetT(offset)
	}
	membersVector := builder.EndVector(membersCount)

	// build response
	rpc.MemberListStart(builder)
	rpc.MemberListAddMembers(builder, membersVector)
	resp := rpc.MemberListEnd(builder)

	builder.Finish(resp)

	return &rpc.Message{Payload: builder.FinishedBytes()}, nil
}

func (h Handler) clusterStatus(version byte, payload []byte) (*rpc.Message, error) {
	return nil, nil
}
