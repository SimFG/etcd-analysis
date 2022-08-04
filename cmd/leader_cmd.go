package cmd

import (
	"context"
	"fmt"
	"github.com/SimFG/etcd-analysis/core"
	"github.com/spf13/cobra"
	clientv3 "go.etcd.io/etcd/client/v3"
)

func NewLeaderCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "leader",
		Short: "Get the leader node info",
		Run:   leaderFunc,
	}
	return cmd
}

func leaderFunc(cmd *cobra.Command, args []string) {
	var (
		c    *clientv3.Client
		err  error
		resp *clientv3.MemberListResponse
	)

	c = core.InitClient()

	resp, err = c.MemberList(context.TODO())
	if err != nil {
		core.Exit(err)
	}

	leaderId := uint64(0)
	for _, ep := range c.Endpoints() {
		if sresp, serr := c.Status(context.TODO(), ep); serr == nil {
			leaderId = sresp.Leader
			break
		}
	}

	for _, m := range resp.Members {
		if m.ID == leaderId {
			fmt.Println("Name:", m.Name)
			fmt.Println("ClientUrls:", m.ClientURLs)
			return
		}
	}
	fmt.Println("error")
}
