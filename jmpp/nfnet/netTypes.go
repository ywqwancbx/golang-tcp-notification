package nfnet

import "net"

type NetWrite struct {
	Conn net.Conn
}

type NetRead struct {
	Conn net.Conn
}

//hearbeat collection
type HeartBeat struct {
	Hs01 HS01
}
//message collection
type Msg struct {
	Mw01 MW01 //welcome msg
	Ms01 MS01 //server msg 01
}
//data collection
type Data struct {

}
//collection
type Befehl struct {
	H HeartBeat
	M Msg
	D Data
}
