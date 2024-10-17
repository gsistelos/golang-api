#!/usr/bin/env bash 

./wait-for-it.sh $MYSQL_ADDR -s -t 0 -- ./bin/api
