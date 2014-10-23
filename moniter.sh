#!/bin/bash
num=$(netstat -lntp |grep 20157|wc -l)
if ((num>0))
then
		echo "proxy is running ..."
else
then
		echo "proxy is stopped.."
fi
netstat -lntp |grep 20157
