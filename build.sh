#!/usr/bin/env bash

BIN_NAME="esm"

mkdir -p bin/
rm -rf bin/*
go build -o bin/${BIN_NAME}