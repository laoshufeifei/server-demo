#!/bin/bash
# git pull -p

bin=server-demo
go build -o $bin main.go

ps aux | grep -v grep | grep -E -i $bin | awk '{print $2}' | xargs kill -9 > /dev/null 2>&1

if [[ -f access.log ]]; then
	rm access.log
fi

if [[ -f nohup.out ]]; then
	rm nohup.out
fi

# ulimit -n 1000000
nohup ./$bin &

sleep 1
echo tail -f access.log nohup.out
tail -f access.log nohup.out
