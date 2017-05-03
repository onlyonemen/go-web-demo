#! /bin/bash
echo "start to rebuild main.go\n"
cd ../
go build
echo "build done\n"
cd docker
rm -rdf demo
mkdir demo
cp -rf ../demo ./demo/
cp -rf ../public ./demo/
cp -rf ../view ./demo/
cp -rf ../files ./demo/


DOCKER_NAME="demo:latest"
PORT="8088:8088"
SHARE_FOLDER="/home/local/Work/go/src/demo/files:/app/demo/files"
docker build -t "$DOCKER_NAME" .

echo "start to run countaner\n"
docker run -p "$PORT" -v "$SHARE_FOLDER" "$DOCKER_NAME"

rm -rdf demo