#!/usr/bin/env bash

curl -s -X POST -H "Content-Type: application/json" -H "Accept: application/json" "http://localhost:9006/api/$1" -d "$2"
