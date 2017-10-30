#!/usr/bin/env sh

go build -o build/apiserver ./server && docker build -t apiserver .