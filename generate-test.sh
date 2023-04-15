#!/bin/sh
set -e 
mkdir -p inputs
for i in `ls jsonnet/*.jsonnet`; do
    echo $i
    jsonnetfmt --in-place $i
    dir=`dirname $i`
    outputFile=`echo "inputs/${i#*/}" | cut -f 1 -d '.'`
    jsonnet $i --output-file "$outputFile.json"
done
