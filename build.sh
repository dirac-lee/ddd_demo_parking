#!/bin/bash

RUN_NAME="ddd.demo.parking"

mkdir -p output/bin output/conf
cp script/bootstrap.sh output 2>/dev/null
chmod +x output/bootstrap.sh
cp script/bootstrap.sh output/bootstrap_staging.sh
chmod +x output/bootstrap_staging.sh
find conf/ -type f ! -name "*_local.*" | xargs -I{} cp {} output/conf/

wire ./...
go build -o output/bin/${RUN_NAME}
