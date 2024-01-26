#!/bin/bash
while true
do
    cat ../cmd/reports/output.json | jq -C .
    sleep 1
    clear
done
