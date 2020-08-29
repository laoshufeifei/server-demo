#!/bin/bash
# https://github.com/eranyanay/1m-go-websockets/blob/master/client.go

docker rm -vf $(docker ps -q --filter label=1m-go-websockets)
# docker container prune -f