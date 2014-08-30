package tools
import (
	"net"
	"fmt"
	"strings"
	"strconv"
	"encoding/json"
)
//fetch local ip
func GetLocalIP() string{
	conn, err := net.Dial("udp", "google.com:80")
	if err != nil {
		fmt.Println("GetLocalIp Fail.."+err.Error())
		return "Unknow"
	}
	defer conn.Close()
	return strings.Split(conn.LocalAddr().String(), ":")[0]
}
//fetch ip type
func GetIPType() string{
	ip :=GetLocalIP()
	if(len(ip)>4){
		return "v4"
	}else{
		return "v6"
	}
}
//fetch TCP connection mode
func GetTCPConnType() string{
	tcptype :=GetIPType()
	if tcptype=="v4"{
		return "tcp"
	}
		return "tcp6"
}

//Converted into a header length
func LengthToSix(l int)string{

	if(l>=10 && l<100){
		return "0000"+strconv.Itoa(l)
	}
	if(l>=100 && l<1000){
		return "000"+strconv.Itoa(l)
	}

	if(l<0){
		return  "000000"
	}
	if(l>=0 && l<10){
		return "00000"+strconv.Itoa(l)
	}

	if(l>=1000 && l<10000){
		return "00"+strconv.Itoa(l)
	}
	if(l>=10000 && l<100000){
		return "0"+strconv.Itoa(l)
	}
	if(l>=100000 && l<1000000){
		return strconv.Itoa(l)
	}
	return  "000000"
}
func DecodeConf(conf []byte,confParam *ConfParam)error{
	return json.Unmarshal(conf,confParam)

}

