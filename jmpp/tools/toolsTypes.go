package tools

import "net"

//Notification server config
type NotificationServerS struct {
	Servername string
	Port string
	Token string
	Welcome string
	Debug string
}
//RPC server config
type RPCServerS struct {
	Port string
	Tonken string
	Connmeth string
}
//Cluser server config
type CluserS struct {

}
type Mongodb struct {
	Conn string
	Username string
	Password string
	KeyFilePath string
	Db string
	Collection string
}
//Json to ConfParam  struct
type ConfParam struct {
	NotificationServer NotificationServerS
	RPCServer RPCServerS
	Cluser CluserS
	Mongodb Mongodb
}
//Mongodb Types
type MongodbOperate struct {
	Operate string //Insert Update Find Delete
    Collection string //集合
	Data interface{} //Operate Data
	ReturnType string //Return ALL ONE Defualt SUCCESS And ONE
}


type ConnInfo struct {
	Conn net.Conn
	Imei string
	PhoneNum string
}
