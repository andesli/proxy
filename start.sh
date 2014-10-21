#!/bin/bash
num=$(netstat -lntp |grep 20157|wc -l)
if ((num>0))
then
		echo "proxy is running ..."
		exit 2
fi
nohup ./proxy >/dev/null 2>&1 &
sleep 2
netstat -lntp |grep 20157
num=$(netstat -lntp |grep 20157|wc -l)
if ((num>0))
then
		echo "proxy is running"
fi
