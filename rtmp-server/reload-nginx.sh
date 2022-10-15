#!/bin/bash

while true
do
    inotifywait --exclude .swp -e create -e modify -e delete -e move /etc/nginx/include/
    sleep 1
    nginx -t
    if [ $? -eq 0 ]
    then
    echo "Detected Nginx Configuration Change"
    echo "Executing: nginx -s reload"
    nginx -s reload
        if [ $? -ne 0 ]
        then
            rm -rf /etc/nginx/include/*
            rm -rf /etc/nginx/http-include/*
            nginx -s reload
        fi
    else
        rm -rf /etc/nginx/include/*
        rm -rf /etc/nginx/http-include/*
    fi
done