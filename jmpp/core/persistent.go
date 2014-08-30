package core


import (
	"net"
	"jmpp/nfnet"
	"encoding/json"
	"labix.org/v2/mgo/bson"
	"jmpp/tools"
	"time"
)

var (
	ListAA = make(map[string]net.Conn)
	ListAB = make(map[string]net.Conn)
	ListAC = make(map[string]net.Conn)
)

type Person struct {
	Id_       bson.ObjectId `bson:"_id"`
	Code string `bson:"code"`
	Imei string `bson:"imei"`
	Phone string `bson:"phone"`
	CreateTime int64 `bson:"createTime"`
	UpdateTime int64 `bson:"updateTime"`
}
type Message struct {
	Id_ bson.ObjectId `bson:"_id"`
	UserId string `bson:userid`
	Msg string `bson:Msg`
	CreateTime int64 `bson:"createTime"`
	UpdateTime int64 `bson:"updateTime"`
}
func (p *Persistent)Welcome(conn net.Conn,server *Server){
	msgChan := make(chan ChanMs)
	defer func() {
		conn.Close()
	}()
	var (
		nw nfnet.NetWrite
		nr nfnet.NetRead
		uid string
	)
	nw.Conn = conn
	nr.Conn = conn
	//welcome
	var mw01 nfnet.MS01
	mw01.RB = false
	mw01.Msg = "welcome"
	nw.SendToClient("MW01",mw01)
	var tryCount int
	for {
		tryCount++
		if tryCount>2{
			goto DISCONN
		}
		bfl,msg,err :=nr.ReadFromClient()
			if err !=nil{
		return
	}
		switch string(bfl){
			case "MI01":
			var mi01 nfnet.MI01
			err =json.Unmarshal(msg,&mi01)
			if err!=nil{
				continue
			}

			var register tools.MongodbOperate
				register.Operate = "Insert"
				register.Collection = "goServer"
				register.Data = &Person{bson.NewObjectId(),mi01.Id,mi01.Imei,mi01.Phone,time.Now().Unix(),time.Now().Unix()}
				uid = mi01.Id
			server.RegisterChannel<-register
			id := mi01.Id
				//check id is exist
			exisInfo := p.IdList[id]
				var connInfo tools.ConnInfo
				connInfo.Conn = conn
				connInfo.PhoneNum = mi01.Phone
				connInfo.Imei = mi01.Imei
			if exisInfo.Conn == nil{
				defer func() {
					delete(p.IdList,mi01.Id)
				//	delete(ListAA,mi01.Id)
				}()
			//ListAA[id] = conn
				p.IdList[mi01.Id] = connInfo
			}else{
				defer func() {
					delete(p.IdList,mi01.Id)
					//	delete(ListAA,mi01.Id)
				}()
				//check the same device
				if(mi01.Imei!=exisInfo.Imei){
					p.IdList[id].Conn.Close();
					p.IdList[mi01.Id] = connInfo
				}else{
					//
					conn.Close();
				}
			}
				goto CONTN
		default:
			goto DISCONN
		}
	}
CONTN:
	go func() {
		for {
		bfl,msg,err :=nr.ReadFromClient()
		if err !=nil{
			break
		}
		switch string(bfl){
		case "MC01":
			var ms01 nfnet.MS01
			err =json.Unmarshal(msg,&ms01)
			if err==nil{
				var mr01 nfnet.MR01
				mr01.Msg = "R"
				mr01.RL = len(msg)
				mr01.V = ms01.V
				var cm  ChanMs
				cm.bfl = "MR01"
				cm.chM = mr01
				msgChan <-cm
			}
		case "HC01":
			var hc01 nfnet.HC01
			err =json.Unmarshal(msg,&hc01)
			if err==nil{
				var hs01 nfnet.HS01
				hs01.H = "y"
				var cm  ChanMs
				cm.bfl = "HS01"
				cm.chM = hs01
				msgChan <-cm
			}
		case  "MS01":
			var ms01 nfnet.MS01;
			err =json.Unmarshal(msg,&ms01)
			if err==nil{
				var register tools.MongodbOperate
				register.Operate = "Insert"
				register.Collection = "Message"
				register.Data = &Message{bson.NewObjectId(),uid,ms01.Msg,int64(ms01.MST),time.Now().Unix()}
				server.RegisterChannel<-register
			}
		}
		}
		var cm  ChanMs
		cm.bfl = "ME02"
		cm.CMD = "stop"
		msgChan <-cm
	}()

	for {
		mc := <-msgChan
		if mc.bfl=="ME02" && mc.CMD=="stop"{
			goto DISCONN
		}
		err :=nw.SendToClient(mc.bfl,mc.chM)
		if err!=nil {
			goto DISCONN
		}
	}
DISCONN:
}
