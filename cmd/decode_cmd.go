package cmd

import (
	"encoding/base64"
	"fmt"

	"github.com/spf13/cobra"
)

var encodedValue string

func NewDecodeCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "decode",
		Short: "decode the base64-encoded value",
		Run:   decodeFunc,
	}

	cmd.Flags().StringVar(&encodedValue, "value", "", "the base64-encoded value")
	return cmd
}

func decodeFunc(cmd *cobra.Command, args []string) {
	v, err := base64.StdEncoding.DecodeString(encodedValue)
	if err != nil {
		fmt.Println("decode err:", err)
	} else {
		fmt.Println("decode value:\n", string(v))
	}
}
