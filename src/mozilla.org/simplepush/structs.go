package simplepush

import (
    "code.google.com/p/go.net/websocket"
    "mozilla.org/simplepush/storage"
    "mozilla.org/util"
)

const (
    UNREG = iota
    REGIS
    HELLO
    ACK
    FLUSH
    RETRN
    DIE
)

type PushCommand struct {
    // Use mutable int value
    Command int            //command type (UNREG, REGIS, ACK, etc)
    Arguments interface {} //command arguments
}

type PushWS struct {
    Uaid string             // id
    Socket *websocket.Conn  // Remote connection
    Scmd chan PushCommand   // server command channel
    Ccmd chan PushCommand   // client command channel
    Store *storage.Storage
    Logger *util.HekaLogger
}
