package cmd

import (
	"errors"
	"fmt"

	"github.com/SimFG/etcd-analysis/core"
	"github.com/spf13/cobra"
)

var (
	renameSourceKey = ""
	renameTargetKey = ""
	renameIsBak     = false
)

func NewRenameCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "rename",
		Short: "rename the key from all etcd data",
		Run:   renameFunc,
	}

	cmd.Flags().StringVar(&renameSourceKey, "source-key", "", "The key of rename data")
	cmd.Flags().StringVar(&renameTargetKey, "target-key", "", "The target key of rename data")
	cmd.Flags().BoolVar(&renameIsBak, "bak", true, "whether or not back the origin key-value")
	return cmd
}

func renameFunc(cmd *cobra.Command, args []string) {
	if renameSourceKey == "" {
		core.Exit(errors.New("should set the source-key param"))
	}
	if renameTargetKey == "" {
		core.Exit(errors.New("should set the target-key param"))
	}
	client := core.InitClient()
	resp, err := core.EtcdGet(client, renameSourceKey)
	if err != nil {
		core.Exit(err)
	}
	if len(resp.Kvs) == 0 {
		core.Exit(errors.New("not found the key"))
	}
	if renameIsBak {
		err = core.EtcdPut(client, "etcd-bak/"+renameSourceKey, string(resp.Kvs[0].Value))
		if err != nil {
			core.Exit(err)
		}
	}
	err = core.EtcdPut(client, renameTargetKey, string(resp.Kvs[0].Value))
	if err != nil {
		core.Exit(err)
	}
	err = core.EtcdDelete(client, renameSourceKey)
	if err != nil {
		core.Exit(err)
	}
	fmt.Println("success rename the key, the bak data:" + "etcd-bak/" + renameSourceKey)
}
