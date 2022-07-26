package cmd

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"github.com/SimFG/etcd-analysis/core"
	"github.com/spf13/cobra"
	"go.etcd.io/etcd/api/v3/mvccpb"
	clientv3 "go.etcd.io/etcd/client/v3"
	"io"
	"os"
	"time"
)

var (
	showValue    bool
	writeOut     string
	hang         bool
	hangInterval int64

	filterAttribute string
	filterMax       int
	filterMin       int
)

func NewLookCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "look",
		Short: "Look all etcd data",
		Long: `
Look all etcd data.

Considering that the value is generally encrypted and difficult to read and is relatively long, the value is not displayed by default.
If you want to display, you can set <show-value> to true.

By default, the command will output all the results to the console, which may be more practical in combination with some text viewing tools, such as 'more' or 'vim', like:
$ look | more
$ look | vim
Of course, you can also save the output as a file by setting <write-out> to 'file', and the generated file is named 'analysis.txt'.

If you want to continuously observe all keys, you can set <hang> to true, and you can use <hang-interval> to set the update interval. 
This function only works when <write-out> is set to 'file'.

Sometimes, you may only want to observe data of a certain range size, you can set <filter> related parameters.
<filter> can set how to calculate the data size, including 'none', 'key', 'value' and 'kv';
<filter-max> and <filter-min> are used to specify the maximum and minimum values of the data, respectively.
`,
		Run: lookFunc,
	}

	cmd.Flags().BoolVar(&showValue, "show-value", false, "Get keys with matching prefix")
	cmd.Flags().StringVar(&writeOut, "write-out", "stdout", "The looking type")
	cmd.Flags().BoolVar(&hang, "hang", false, "Get updates periodically, only '--write-out=file' takes effect")
	cmd.Flags().Int64Var(&hangInterval, "hang-interval", 2, "Update interval, and the unit is 's'")
	cmd.Flags().StringVar(&filterAttribute, "filter", "none", "The filter attribute")
	cmd.Flags().IntVar(&filterMax, "filter-max", -1, "The filter max value")
	cmd.Flags().IntVar(&filterMin, "filter-min", -1, "The filter min value")

	cmd.RegisterFlagCompletionFunc("write-out", func(_ *cobra.Command, _ []string, _ string) ([]string, cobra.ShellCompDirective) {
		return []string{"stdout", "file"}, cobra.ShellCompDirectiveDefault
	})
	cmd.RegisterFlagCompletionFunc("filter", func(_ *cobra.Command, _ []string, _ string) ([]string, cobra.ShellCompDirective) {
		return []string{"none", "key", "value", "kv"}, cobra.ShellCompDirectiveDefault
	})
	return cmd
}

func lookFunc(cmd *cobra.Command, args []string) {
	core.InitClient()
	resp := core.GetAllData()

	var writer io.Writer
	switch writeOut {
	case "file":
		f := GetFileWriter()
		defer f.Close()
		writer = f
	case "stdout":
		fallthrough
	default:
		writer = os.Stdout
	}
	appendBuffer(resp, writer)
	if hang && writeOut == "file" {
		ct := time.Tick(time.Second * time.Duration(hangInterval))
		i := 0
		for {
			select {
			case <-ct:
				resp = core.GetAllData()
				appendBuffer(resp, writer)
				fmt.Println(i, "flush...")
				i++
			}
		}
	}
}

func appendBuffer(resp *clientv3.GetResponse, writer io.Writer) {
	if f, ok := writer.(*os.File); ok {
		f.Truncate(0)
		f.Seek(0, 0)
	}
	var buffer bytes.Buffer
	buffer.WriteString("\nCurrent Stage\n")
	buffer.WriteString(fmt.Sprintf("  %s", resp.Header.String()))
	buffer.WriteString("\nKv List\n")
	buffer.WriteString("| Key | Value | CreateRevision | ModRevision | Version | Lease |\n")
	if len(resp.Kvs) > 0 {
		for _, kv := range resp.Kvs {
			if filter(kv) {
				continue
			}
			v := "-"
			if showValue {
				v = string(kv.Value)
			}
			buffer.WriteString(fmt.Sprintf("| %s | %s | %d | %d | %d | %d |\n",
				string(kv.Key), v,
				kv.CreateRevision, kv.ModRevision, kv.Version,
				kv.Lease))
		}
	} else {
		buffer.WriteString("Empty Data\n")
	}

	buffer.WriteTo(writer)
}

func filter(kv *mvccpb.KeyValue) bool {
	size := -1
	switch filterAttribute {
	case "key":
		size = binary.Size(kv.Key)
	case "value":
		size = binary.Size(kv.Value)
	case "kv":
		size = binary.Size(kv.Key) + binary.Size(kv.Value)
	default:
		size = -1
	}
	if size < 0 || (filterMin < 0 && filterMax < 0) || (size >= filterMin && size <= filterMax) {
		return false
	}
	return true
}

func GetFileWriter() *os.File {
	f, err := os.OpenFile("analysis.txt", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		core.Exit(err)
	}
	return f
}
