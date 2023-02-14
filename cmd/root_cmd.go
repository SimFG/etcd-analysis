package cmd

import (
	"fmt"
	"github.com/SimFG/etcd-analysis/core"
	"github.com/spf13/cobra"
	cobracompletefig "github.com/withfig/autocomplete-tools/integrations/cobra"
	"os"
)

var (
	rootCmd = &cobra.Command{
		Use:   "etcdctl+",
		Short: "etcd data analysis tool",
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
	rootCmd.PersistentFlags().StringVar(&core.C.TLS.CertFile, "cert", "", "identify secure client using this TLS certificate file")
	rootCmd.PersistentFlags().StringVar(&core.C.TLS.KeyFile, "key", "", "identify secure client using this TLS key file")
	rootCmd.PersistentFlags().StringVar(&core.C.TLS.TrustedCAFile, "cacert", "", "verify certificates of TLS-enabled secure servers using this CA bundle")

	rootCmd.AddCommand(NewDistributeCmd())
	rootCmd.AddCommand(NewLookCmd())
	rootCmd.AddCommand(NewLeaderCmd())
	rootCmd.AddCommand(NewClearCmd())
	rootCmd.AddCommand(NewFindCmd())
	rootCmd.AddCommand(NewDecodeCmd())
	rootCmd.AddCommand(cobracompletefig.CreateCompletionSpecCommand())
}
