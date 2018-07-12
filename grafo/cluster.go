package grafo

import (
	"context"
	"sort"

	"github.com/grafodb/grafodb/internal/rpc"
)

type ClusterAPI interface {
	MembersList() ([]Member, error)
}

type clusterAPI struct {
	client *Client
}

func (api clusterAPI) MembersList() ([]Member, error) {
	client := rpc.NewGrafoRPCClient(api.client.conn)

	msg := rpc.NewRequestMessage(rpc.CmdClusterMemberList, ApiV1)
	resp, err := client.SendMessage(context.Background(), msg)
	if err != nil {
		return nil, err
	}

	memberList := rpc.GetRootAsMemberList(resp.Payload, 0)

	count := memberList.MembersLength()
	members := make([]Member, count)
	for i := 0; i < count; i++ {
		var m rpc.Member
		memberList.Members(&m, i)
		members[i] = Member{
			ID:     string(m.Id()),
			Addr:   string(m.Addr()),
			Status: string(m.Status()),
		}
	}
	sort.Sort(memberSortByID(members))

	return members, nil
}
