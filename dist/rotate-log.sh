#!/bin/bash
HERE=`cd -P $(dirname $0) && pwd`

DATE=`date +"%Y-%m-%d_%Hh%M"`
cp $HERE/log/log.txt $HERE/log/${DATE}_log.txt 

echo "" > $HERE/log/log.txt
