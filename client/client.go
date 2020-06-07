package client

import (
	"time"

	"github.com/yzhaoyu/yRPC/codec"
	"github.com/yzhaoyu/yRPC/protocol"
	"github.com/yzhaoyu/yRPC/transport"
)

type Option struct {
	ProtocolType  protocol.ProtocolType
	SerializeType codec.SerializeType
	CompressType  protocol.CompressType
	TransportType transport.TransportType

	RequestTimeout time.Duration
}

var DefaultOption = Option{
	ProtocolType:  protocol.Default,
	SerializeType: codec.MessagePack,
	CompressType:  protocol.CompressTypeNone,
	TransportType: transport.TCPTransport,
}
