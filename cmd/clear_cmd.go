package cmd

import (
	"fmt"

	"github.com/SimFG/etcd-analysis/core"
	"github.com/spf13/cobra"
	clientv3 "go.etcd.io/etcd/client/v3"
)

func NewClearCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "clear",
		Short: "Clear all etcd data",
		Run:   clearFunc,
	}

	return cmd
}

func clearFunc(cmd *cobra.Command, args []string) {
	var (
		client  = core.InitClient()
		enter   string
		getResp *clientv3.GetResponse
		err     error
	)
	getResp, err = core.EtcdGet(client, core.EmptyChar(), clientv3.WithFromKey(), clientv3.WithCountOnly(), clientv3.WithSerializable())
	if err != nil {
		core.Exit(err)
	}
	fmt.Println("Current Data Count:", getResp.Count)
	fmt.Print("Clear All Data, (Y/n):")
	fmt.Scan(&enter)
	if enter != "Y" {
		return
	}
	err = core.EtcdDelete(client, core.EmptyChar(), clientv3.WithFromKey())
	if err != nil {
		core.Exit(err)
	}
}
