package core

import (
	"jmpp/tools"
)

//server
type Server  struct{
	CurPath string
	NotificationServer  tools.NotificationServerS
	RPCServer tools.RPCServerS
	Cluser tools.CluserS
	Persistenter Persistent
	MongodbConf tools.Mongodb
	RegisterChannel chan tools.MongodbOperate
	MongodbOtherChannel chan tools.MongodbOperate
}

//rpc call type
type SendMsg struct {
	Persistenter *Persistent
}

type Msg struct {
	Id      string
	MsgType int
	Msg     string
}


//Persistent struct
type Persistent struct {
	IdList map[string]tools.ConnInfo
}



type  ChanMs struct {
	bfl string
	chM interface {}
	CMD string //command
}

