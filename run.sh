#!/usr/bin/env bash 

make

./wait-for-it.sh $MYSQL_ADDR -s -t 0 -- ./bin/api
