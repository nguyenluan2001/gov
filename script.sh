#! /bin/bash

command=$1
gov_path="./.gov"


if [[ ! -e $gov_path ]];then
    mkdir $gov_path
fi

function run(){
    docker rm gov-container
    docker run \
    -w /app \
    --mount type=bind,source=./main.go,target=/app/main.go \
    --mount type=bind,source=./go.mod,target=/app/go.mod \
    --mount type=bind,source=./go.sum,target=/app/go.sum \
    --mount type=bind,source=./utils/,target=/app/utils \
    --mount type=bind,source=./controller/,target=/app/controller \
    --mount type=bind,source=./${gov_path}/,target=/app/.gov \
    -it --name gov-container \
    gov-img
}

function build(){
    docker build -t gov-img .
}

case $command in
    run) run ;;
    build) build ;;
    *) echo "No command found.";;
esac