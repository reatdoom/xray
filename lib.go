package xray

import (
    "fmt"
)

func AddInbound(port int) {
    var (
        xrayCtl *XrayController
        cfg     = &BaseConfig{
            APIAddress: "127.0.0.1",
            APIPort:    10085,
        }
    )

    xrayCtl = new(XrayController)
    err := xrayCtl.Init(cfg)
    defer xrayCtl.CmdConn.Close()
    if err != nil {
        fmt.Println("Failed %s", err)
    }

    fmt.Println("添加Inbound, 端口：", port)

    err = addSSInbound(xrayCtl.HsClient, port)
    if err != nil {
        fmt.Println("Failed ", err)
    }
}

func AddUser(userid string, encryptType string) {
    var (
        xrayCtl *XrayController
        cfg     = &BaseConfig{
            APIAddress: "127.0.0.1",
            APIPort:    10085,
        }
        user = UserInfo{
            Uuid:       userid,
            AlertId:    0,
            Level:      0,
            InTag:      "ss",
            Email:      userid,
            CipherType: encryptType,
            Password:   userid,
        }
    )

    xrayCtl = new(XrayController)
    err := xrayCtl.Init(cfg)
    defer xrayCtl.CmdConn.Close()
    if err != nil {
        fmt.Println("Failed %s", err)
    }

    fmt.Println("添加用户: ", userid)

    err = addSSUser(xrayCtl.HsClient, &user)
    if err != nil {
        fmt.Println("Failed ", err)
    }
}

func DelUser(userid string) {
    var (
        xrayCtl *XrayController
        cfg     = &BaseConfig{
            APIAddress: "127.0.0.1",
            APIPort:    10085,
        }
        user = UserInfo{
            InTag:      "ss",
            Email:      userid,
        }
    )

    xrayCtl = new(XrayController)
    err := xrayCtl.Init(cfg)
    defer xrayCtl.CmdConn.Close()
    if err != nil {
        fmt.Println("Failed %s", err)
    }
    err = removeUser(xrayCtl.HsClient, &user)
    if err != nil {
        fmt.Println("Failed %s", err)
    }
}
