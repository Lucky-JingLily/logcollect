package tailf

import (
	"github.com/hpcloud/tail"
	"sync"
)

const (
	StatusNormal = 1
	StatusDelete = 2
)

type CollectConf struct {
	LogPath string `json:"log_path"`
	Topic string `json:"topic"`
}

type TailObj struct {
	tail *tail.Tail
	conf CollectConf
	status int
	exitChan chan int
}

type TextMsg struct {
	Msg string
	Topic string
}

type TailObjMgr struct {
	tailObj []*TailObj
	msgChan chan *TextMsg
	lock sync.Mutex
}