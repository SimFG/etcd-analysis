package cmd

import (
	"fmt"
	"github.com/SimFG/etcd-analysis/core"
	"github.com/spf13/cobra"
	"os"
)

var (
	rootCmd = &cobra.Command{
		Use:   "etcdctl+",
		Short: "short tips for the command",
		Long: `long long
long tips for the command
	`,
	}
)

func Start() {
	if err := rootCmd.Execute(); err != nil {
		if rootCmd.SilenceErrors {
			fmt.Fprintln(os.Stderr, err)
		}
		os.Exit(-1)
	}
}

func init() {
	cobra.EnablePrefixMatching = true

	rootCmd.PersistentFlags().StringSliceVar(&core.C.Endpoints, "Endpoints", []string{"127.0.0.1:2379"}, "etcd connect Endpoints")

	rootCmd.AddCommand(NewDistributeCmd())
	rootCmd.AddCommand(NewLookCmd())
}