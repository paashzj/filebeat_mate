#!/bin/bash

mkdir $FILEBEAT_HOME/logs
nohup $FILEBEAT_HOME/filebeat -e -c $FILEBEAT_HOME/filebeat.yml >>$FILEBEAT_HOME/logs/filebeat.stdout.log 2>>$FILEBEAT_HOME/logs/filebeat.stderr.log &

