package core

import (
	"errors"
	"time"
	"jmpp/nfnet"
)

func (sm *SendMsg) Send(msg *Msg, reply *bool) error {
	ConnInfo := sm.Persistenter.IdList[msg.Id]
	if ConnInfo.Conn==nil{
		*reply = false
		return errors.New("Id Not Found")
	}
	var ms01 nfnet.MS01
	ms01.Msg = msg.Msg
	ms01.MST = time.Now().UTC().Nanosecond()/1000000
	var nw nfnet.NetWrite
	nw.Conn = ConnInfo.Conn
	nw.SendToClient("MS01",ms01)
	*reply = true
	return nil
}

func(sm *SendMsg) GetAllClient(msg string,reply *[]string)error{
	var temp = make([]string,0)
	for k,_ :=range sm.Persistenter.IdList {
		temp = append(temp,k)
	}
	*reply = temp
	return nil
}

func (sm *SendMsg)GetOnlineTotal(msg string,relay *int)error{
	*relay = len(sm.Persistenter.IdList)
	return nil
}
