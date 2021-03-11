#!/bin/bash

PWD=`pwd`
DIR=`dirname $PWD`
API=$DIR/"api"

if [ ! -d "$API" ]; then 
    echo "you need run this file in scripts dir..."
    exit 1
fi 

FILENAMES=`find $API -type f -name "*.proto" -print`

LocalBuild() {
    if [ ! -d "$API" ]; then 
       mkdir "$API"
    fi 

    if [ ! -f "$1" ]; then
        echo "$1 is not here..."
        exit 1
    fi

	protoc $1 -I $DIR/third_party -I $API \
		--go_out=plugins=grpc:$DIR \
		--grpc-gateway_out=logtostderr=true:$DIR \
		--swagger_out=logtostderr=true:$DIR \
		--validate_out=lang=go:$DIR
}

for file in $FILENAMES; do 
    LocalBuild $file
done
