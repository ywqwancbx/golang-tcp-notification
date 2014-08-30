package nfnet

import (
	"fmt"
	"strconv"
	"time"
)

//bfl Befehl instruction，message
func (r *NetRead)ReadFromClient()([]byte,[]byte,error){
	var bfl []byte
	var msg []byte
	headLen := make([]byte, 6)
	r.Conn.SetReadDeadline(time.Now().Add(6 * time.Minute))
	head_len, err := r.Conn.Read(headLen)
	if err != nil {
		fmt.Println("close client:",err)
		return bfl,msg,err
	}
	contentLengString :=ReadMsg(head_len,headLen)
	bflLen := make([]byte, 4)
	bfl_len,err := r.Conn.Read(bflLen)
	if err != nil {
		fmt.Println(err)
		return bfl,msg,err
	}
	bfl =ReadMsg(bfl_len,bflLen)
	if err != nil {
		fmt.Println(err)
		return bfl,msg,err
	}
	contentLeng,err :=strconv.Atoi(string(contentLengString))
	if err != nil {
		fmt.Println(err)
		return bfl,msg,err
	}
	contentLengByte := make([]byte, contentLeng)
	content_lengByte,err := r.Conn.Read(contentLengByte)
	msg = ReadMsg(content_lengByte,contentLengByte)
	return bfl,msg,err
}

func ReadMsg(length int, msg []byte) []byte {
	re:= []byte("")
	if length > 0 {
		re := msg[0:length]
		return re
	}
	return re
}
