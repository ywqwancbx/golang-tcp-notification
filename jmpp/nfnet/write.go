package nfnet

import (
	"jmpp/tools"
	"encoding/json"
	"time"
)

//bfl Befehl instruction，message
func (w *NetWrite) SendToClient(bfl string ,msgI interface {})error{
	msgByte ,err := json.Marshal(msgI)
	msg :=string(msgByte)
	length := len(msg)
	wString := tools.LengthToSix(length)+bfl+msg
	w.Conn.SetWriteDeadline(time.Now().Add(1 * time.Minute))
	_,err = w.Conn.Write([]byte(wString));
	if err!=nil{
		return err
	}
	return nil
}
