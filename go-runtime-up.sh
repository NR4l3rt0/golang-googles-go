#!/bin/bash

workdir=/home/go_user/workspace 

if [ $(docker ps -a --format '{{.Names}}' | grep go | wc -l) -eq 0 ]; then

	docker run -it --name go -w $workdir --mount type=bind,source=$HOME/golang-googles-go,target=$workdir golang

else
	docker start go
	docker attach go
fi

