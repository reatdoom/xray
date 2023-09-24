package xray

import (
	"fmt"
)

func AddUser(userid string) {
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
            CipherType: "aes-256-gcm",
            Password:   userid,
        }
    )

    xrayCtl = new(XrayController)
    err := xrayCtl.Init(cfg)
    defer xrayCtl.CmdConn.Close()
    if err != nil {
        fmt.Println("Failed %s", err)
    }

    fmt.Println("添加用户: %s", userid)

    err = addSSUser(xrayCtl.HsClient, &user)
    if err != nil {
        fmt.Println("Failed %s", err)
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
