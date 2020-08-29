#!/bin/bash
## This script helps regenerating multiple client instances in different network namespaces using Docker
## This helps to overcome the ephemeral source port limitation
## Usage: ./connect <connections> <number of clients> <test url>
## Number of clients helps to speed up connections establishment at large scale, in order to make the demo faster
## https://github.com/eranyanay/1m-go-websockets/blob/master/client.go

REPLICAS=10
if [[ ! -z $1 ]];then
	REPLICAS=$1
fi

CONNECTIONS=20000
URL="ws://172.17.0.1:9090/echo"

echo REPLICAS is $REPLICAS
echo CONNECTIONS is $CONNECTIONS
echo URL is $URL

go build --tags "static netgo" -o client client.go
for i in `seq 1 $REPLICAS`; do
    docker run -l 1m-go-websockets -v $(pwd)/client:/client -d ubuntu:16.04 /client -conn=${CONNECTIONS} -url=${URL}
done
