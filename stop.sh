#!/bin/bash
netstat -lntp |grep 20157
echo "begin to stop proxy.."
netstat -lntp |grep 20157|awk '{print $7}'|awk -F/ '{print $1}'|xargs -I {} kill -9 {}
echo "proxy is stopped"
