package main

import (
	"fmt"
	"os"
	"jmpp/core"
	"io/ioutil"
	"jmpp/tools"
	"path/filepath"
	"os/exec"
	"jmpp/logs"
	"labix.org/v2/mgo/bson"
)
var (
	Confile []byte
	Welcome string
	Log logs.Logs;
)
type Person struct {
	Id_       bson.ObjectId `bson:"_id"`
	Name string `bson:"name"`
	Phone string `bson:"phone"`
}
func main() {
//check is linux
	if(os.PathSeparator==47) {
			if os.Getppid() != 1 {           //
				filePath, _ := filepath.Abs(os.Args[0])  //
				cmd := exec.Command(filePath, os.Args[1:]...)//
				cmd.Stdin = os.Stdin                               //
				cmd.Stdout = os.Stdout
				cmd.Stderr = os.Stderr
				cmd.Start()                                          //
				return
			}
	}
	fmt.Println(`EF Notification Server Start...`)
	var server core.Server //init service
	var persistent core.Persistent //init client persistent
	server.Persistenter = persistent
	path , pathErr := os.Getwd()
	server.CurPath = path
	if pathErr != nil {
		fmt.Println("Get Path Fail...", pathErr)
		os.Exit(10001)
	}
	//read configuration file
	conf , confErr := ioutil.ReadFile(server.CurPath + "/conf/app.conf")
	Confile = conf
	if confErr != nil {
		fmt.Println("Get Path Configure File Fail...", confErr)
		os.Exit(10002)
	}
	var confParam tools.ConfParam
	err := tools.DecodeConf(Confile, &confParam)
		if err != nil {
		fmt.Println("Decode Configure File Fail...", err)
		os.Exit(10003)
	}
	Welcome = confParam.NotificationServer.Welcome
	server.Init(confParam)
	go server.RPCStart()
	server.RegisterChannel = make(chan tools.MongodbOperate)
	go server.MongoDBInit()
	server.Start()
}


