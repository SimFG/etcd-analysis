package cmd

import (
	"errors"
	"fmt"

	"github.com/SimFG/etcd-analysis/core"
	"github.com/golang/protobuf/proto"
	"github.com/jhump/protoreflect/desc"
	"github.com/jhump/protoreflect/desc/protoparse"
	"github.com/jhump/protoreflect/dynamic"
	"github.com/spf13/cobra"
)

var (
	unmarshallKey            string
	unmarshalImportPaths     []string
	unmarshalFileNames       []string
	unmarshalFullMessageName string
)

func NewUnmarshalCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "unmarshal",
		Short: "unmarshal the etcd value",
		Run:   unmarshalFunc,
	}

	cmd.Flags().StringVar(&unmarshallKey, "key", "", "the proto.marshal value of the full key")
	cmd.Flags().StringArrayVar(&unmarshalImportPaths, "import-path", []string{}, "the proto.marshal value of the full key")
	cmd.Flags().StringArrayVar(&unmarshalFileNames, "proto", []string{}, "the proto.marshal value of the full key")
	cmd.Flags().StringVar(&unmarshalFullMessageName, "full-message-name", "", "the proto.marshal value of the full key")

	return cmd
}

func unmarshalFunc(cmd *cobra.Command, args []string) {
	client := core.InitClient()
	resp, err := core.EtcdGet(client, unmarshallKey)
	if err != nil {
		core.Exit(err)
	}
	if len(resp.Kvs) == 0 {
		core.Exit(errors.New("not found the key"))
	}
	bytesValue := resp.Kvs[0].Value
	decodeBytes(bytesValue)
	return
}

func decodeBytes(bytesValue []byte) {
	fileNames, err := protoparse.ResolveFilenames(unmarshalImportPaths, unmarshalFileNames...)
	if err != nil {
		fmt.Println("counld not resolve file names")
		core.Exit(err)
	}
	p := protoparse.Parser{
		ImportPaths:           unmarshalImportPaths,
		InferImportPaths:      len(unmarshalImportPaths) == 0,
		IncludeSourceCodeInfo: true,
	}

	fds, err := p.ParseFiles(fileNames...)
	if err != nil {
		fmt.Println("could not parse given files")
		core.Exit(err)
	}
	var messageType *desc.MessageDescriptor
	for _, protoDesc := range fds {
		messageType = protoDesc.FindMessage(unmarshalFullMessageName)
		if messageType != nil {
			break
		}
	}
	if messageType == nil {
		fmt.Println("could not find message type, please check the message name is full, name: " + unmarshalFullMessageName)
		return
	}
	message := dynamic.NewMessage(messageType)
	err = proto.Unmarshal(bytesValue, message)
	if err != nil {
		fmt.Println("could not unmarshal bytes")
		core.Exit(err)
	}
	fmt.Println()
	fields := message.GetKnownFields()
	for _, field := range fields {
		fieldValue := message.GetField(field)
		fmt.Printf("%s: %v\n", field.GetName(), fieldValue)
	}
}
