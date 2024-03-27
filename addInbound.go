package xray

import (
	"context"
	"github.com/xtls/xray-core/app/proxyman"
	"github.com/xtls/xray-core/app/proxyman/command"
	"github.com/xtls/xray-core/common/net"
	"github.com/xtls/xray-core/common/serial"
	"github.com/xtls/xray-core/core"
	"github.com/xtls/xray-core/proxy/shadowsocks"
)

func addSSInbound(client command.HandlerServiceClient, port int) error {
	_, err := client.AddInbound(context.Background(), &command.AddInboundRequest{
		Inbound: &core.InboundHandlerConfig{
			Tag: "ss", // 设置入站代理的标签
			ReceiverSettings: serial.ToTypedMessage(&proxyman.ReceiverConfig{
				//PortRange: net.SinglePortRange(net.Port(port)), // 设置监听的端口
                PortList: &net.PortList{
                    Range: []*net.PortRange{net.SinglePortRange(net.Port(port))},
                },
				Listen:    net.NewIPOrDomain(net.AnyIP),         // 监听所有 IP 地址
			}),
			ProxySettings: serial.ToTypedMessage(&shadowsocks.ServerConfig{
				Network: net.Network_TCP | net.Network_UDP,
			}), // 使用默认配置的 Shadowsocks 服务器
		},
	})
	return err
}

