#!/bin/bash
reg=docker.navicore.cn
tag=1.0.0

home=`dirname $(realpath $0)`
cd $home
cd ..

git pull -p
go build -o build/server-demo main.go

cd $home
if [[ -d _build ]];then
	rm -rf _build
fi
mkdir -p _build/server-demo

chmod +x server-demo
cp server-demo _build/server-demo
cp config.json  _build/server-demo
cp start.sh _build/

for f in `find _build * -type f`;do
	file $f | grep executable > /dev/null
	if [[ $? != 0 ]];then
		continue
	fi
	chmod +x $f
done

docker build --pull -t $reg/test/server-demo:$tag .

sleep 1
docker push $reg/test/server-demo:$tag
