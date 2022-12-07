package cmd

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"github.com/SimFG/etcd-analysis/core"
	"github.com/spf13/cobra"
	"go.etcd.io/etcd/api/v3/mvccpb"
	clientv3 "go.etcd.io/etcd/client/v3"
	"io"
	"os"
	"strings"
)

var (
	findKey      = ""
	containValue = false
	findLimit    = 10
)

func NewFindCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "find",
		Short: "find the key from all etcd data",
		Run:   findFunc,
	}

	cmd.Flags().StringVar(&findKey, "key", "", "Show the data like the key")
	cmd.Flags().BoolVar(&containValue, "value", false, "Show the value or not")
	cmd.Flags().IntVar(&findLimit, "limit", 10, "The limit of the show keys")
	return cmd
}

func findFunc(cmd *cobra.Command, args []string) {
	core.InitClient()
	resp, datac := core.GetAllData()
	appendBufferForFind(resp, datac, os.Stdout)
}

func appendBufferForFind(resp *clientv3.GetResponse, datac <-chan []*mvccpb.KeyValue, writer io.Writer) {
	var buffer bytes.Buffer
	buffer.WriteString("Kv List\n")
	buffer.WriteString("| Key | Value |\n")

	count := 0
	for data := range datac {
		for _, kv := range data {
			if count >= findLimit {
				buffer.WriteTo(writer)
				return
			}
			key := string(kv.Key)
			if !strings.Contains(key, findKey) {
				continue
			}
			v := ""
			if containValue {
				v = base64.StdEncoding.EncodeToString(kv.Value)
			}
			buffer.WriteString(fmt.Sprintf("| %s | %s |\n",
				key, v))
			count++
		}
	}

	buffer.WriteTo(writer)
}
