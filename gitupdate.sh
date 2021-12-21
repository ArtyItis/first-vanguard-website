#!/bin/bash

#Alle 15 Minuten per Cronjob
#*/15 * * * * cd masterserver/vanguard-homepage/ && timeout 300 ./gitupdate.sh >/dev/null 2>&1

git pull
