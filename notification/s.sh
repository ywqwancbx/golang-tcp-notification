#!/bin/sh
killall Notification

rm -f Notification

go build Notification.go

sh startNoti.sh
