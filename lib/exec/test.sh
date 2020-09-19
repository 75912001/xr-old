#!/bin/bash
p1=$1

echo ${p1}

for((i=0;i<10;i++))
do
  sleep 1
  echo $(date +"%Y-%m-%d %H:%M:%S")
  cdd
  #exit 1
done