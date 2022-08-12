package cmd

import (
	"encoding/binary"
	"fmt"
	"github.com/SimFG/etcd-analysis/core"
	"github.com/spf13/cobra"
	"go.etcd.io/etcd/api/v3/mvccpb"
)

var (
	distributeType string
	bucketCount    int
)

func NewDistributeCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "distribute",
		Short: "Show the data distribution of etcd",
		Long: `
Show the data distribution of etcd.

According to setting <type>, this command will show the data distribution by the size of the <type>.
The 'kv' means the 'key' and 'value'.

According to setting <bucket>, this command will show the different size histogram.
Each size interval is '(maxSize - minSize) / bucket'.

According to the output below, it means:
when the data size is '0.0 B', the count of this kind of data is 12.
when the data size is greater than '0.0B' and less than or equal to '573.0 B', the count is 275.
'573.0 B' < size <= '1.1 KB', count 80.

Example:
$ distribute --type=value --bucket=8
Summary:
  Count:        399.
  Total:        267.9 KB.
  Smallest:     0.0 Byte.
  Largest:      4.5 KB.
  Average:      687.0 Byte.

Size histogram:
  0.0 B [12]    |∎
  573.0 B [275] |∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎
  1.1 KB [80]   |∎∎∎∎∎∎∎∎∎∎∎
  1.7 KB [0]    |
  2.2 KB [0]    |
  2.8 KB [0]    |
  3.4 KB [0]    |
  3.9 KB [0]    |
  4.5 KB [32]   |∎∎∎∎

Size distribution:
  10% in 3.0 Byte.
  25% in 4.0 Byte.
  50% in 424.0 Byte.
  75% in 854.0 Byte.
  90% in 1.0 KB.
  95% in 4.5 KB.
  99% in 4.5 KB.
`,
		Run: distributeFunc,
	}

	cmd.Flags().StringVar(&distributeType, "type", "key", "Distribution basis; key, value or kv")
	cmd.Flags().IntVar(&bucketCount, "bucket", 5, "Bucket Count")

	cmd.RegisterFlagCompletionFunc("type", func(_ *cobra.Command, _ []string, _ string) ([]string, cobra.ShellCompDirective) {
		return []string{"key", "value", "kv"}, cobra.ShellCompDirectiveDefault
	})
	return cmd
}

func distributeFunc(cmd *cobra.Command, args []string) {
	core.InitClient()
	_, datac := core.GetAllData()

	sizeOf := func(kv *mvccpb.KeyValue) int {
		switch distributeType {
		case "value":
			return binary.Size(kv.Value)
		case "kv":
			return binary.Size(kv.Key) + binary.Size(kv.Value)
		case "key":
			fallthrough
		default:
			return binary.Size(kv.Key)
		}
	}
	r := core.NewReport(bucketCount, sizeOf)
	c1 := r.Results()
	go func() {
		defer close(c1)
		for data := range datac {
			c1 <- data
		}
	}()
	fmt.Println(<-r.Run())
}
