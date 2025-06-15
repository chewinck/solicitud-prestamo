#!/bin/bash

container='solicitud-prestamo-container'

if [ "$(docker ps -aq -f name=$container -f status=running)" ]; then
    docker stop $container
    docker rm $container
fi

