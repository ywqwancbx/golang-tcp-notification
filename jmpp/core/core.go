package core

import (
	"fmt"
	"os"
	"net"
	"jmpp/tools"
	"net/rpc"
	"net/http"
	"labix.org/v2/mgo"
)

func (server *Server)Init(confParam tools.ConfParam){
	fmt.Println("Server init")
	server.NotificationServer = confParam.NotificationServer
	server.RPCServer = confParam.RPCServer
	server.Cluser = confParam.Cluser
	server.MongodbConf = confParam.Mongodb
}

func (server *Server)Start(){
	tcpAddr, err := net.ResolveTCPAddr("tcp",server.NotificationServer.Port)
	if err!=nil{
		fmt.Println("TCPAddr Resolve Fail...",err)
		os.Exit(10004)
	}
	listener, err := net.ListenTCP("tcp", tcpAddr)
	if err!=nil{
		fmt.Println("TCP Listen Fail...",err)
		os.Exit(10004)
	}
	server.Persistenter.IdList =  make(map[string]tools.ConnInfo)
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		go server.Persistenter.Welcome(conn,server)
	}
}

func (server *Server)RPCStart(){
	sdm := new(SendMsg)
	sdm.Persistenter = &server.Persistenter
	rpc.Register(sdm)
	rpc.HandleHTTP()
	l, e := net.Listen("tcp",server.RPCServer.Port)
	if e != nil {
		fmt.Println("RCP Server start Fail...",e)
	}
	go http.Serve(l, nil)
}

func (server *Server)MongoDBInit() *mgo.Collection{
	session, err := mgo.Dial(server.MongodbConf.Conn) // mongodb运行的server名字，可以多个server
	if err != nil {
		fmt.Println(err)
	}
	defer session.Close()
	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)
	db := session.DB(server.MongodbConf.Db)
	for {
		var t = <-server.RegisterChannel
		c := db.C(t.Collection)
		if(t.Operate=="Insert"){
		err = c.Insert(t.Data)
		if err != nil {
			fmt.Println(err)
		}
		}
	}
}
