package xray

import (
    "context"
    "github.com/xtls/xray-core/app/proxyman/command"
    "github.com/xtls/xray-core/common/protocol"
    "github.com/xtls/xray-core/common/serial"
    "github.com/xtls/xray-core/proxy/shadowsocks"
)

func addSSUser(client command.HandlerServiceClient, user *UserInfo) error {
    var ssCipherType shadowsocks.CipherType
    switch user.CipherType {
    case "aes-128-gcm":
        ssCipherType = shadowsocks.CipherType_AES_128_GCM
    case "aes-256-gcm":
        ssCipherType = shadowsocks.CipherType_AES_256_GCM
    case "chacha20-ietf-poly1305":
        ssCipherType = shadowsocks.CipherType_CHACHA20_POLY1305
    }

    _, err := client.AlterInbound(context.Background(), &command.AlterInboundRequest{
        Tag: user.InTag,
        Operation: serial.ToTypedMessage(&command.AddUserOperation{
            User: &protocol.User{
                Level: user.Level,
                Email: user.Email,
                Network: 'tcp,udp',
                Account: serial.ToTypedMessage(&shadowsocks.Account{
                    Password:   user.Password,
                    CipherType: ssCipherType,
                }),
            },
        }),
    })
    return err
}

func removeUser(client command.HandlerServiceClient, user *UserInfo) error {
    _, err := client.AlterInbound(context.Background(), &command.AlterInboundRequest{
        Tag: user.InTag,
        Operation: serial.ToTypedMessage(&command.RemoveUserOperation{
            Email: user.Email,
        }),
    })
    return err
}
