package logs

import (
	"jmpp/tools"
	"fmt"
)
func (l *Logs)Init(ns tools.NotificationServerS){
	l.ns = ns
}
func (l * Logs)Log(msg interface {}){
	if(l.ns.Debug=="true"){
		fmt.Println(msg)
	}
}
