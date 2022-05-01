#!/bin/sh
set -e
for lispfile in tests/*.zy
do
    cmd/zygo/zygo -demo -exitonfail "${lispfile}" || (echo "${lispfile} failed" && exit 1)
    echo "${lispfile} passed"        
done
echo 
echo "good: all tests/ scripts passed."
