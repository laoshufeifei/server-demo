#!/bin/bash

cd /etc/ncserver/server-demo
export GIN_MODE=release

nohup ./server-demo &

sleep 2
tail -f access.log
