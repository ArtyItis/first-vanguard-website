#!/bin/bash

#Alle 15 Minuten per Cronjob
#*/15 * * * * cd masterserver/golangforgottennw/ && timeout 300 ./gitupdate.sh >/dev/null 2>&1

git pull