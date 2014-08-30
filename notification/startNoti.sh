#!/bin/sh
print "This is a Notification start shell"
CURRENT_PATH=$(cd "$(dirname "$0")"; pwd)
cd CURRENT_PATH
./Notification
