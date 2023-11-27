#!/bin/bash

if [[ $# -lt 1 ]]; then
	echo "Provide file name in the argument"
	exit 1
fi

file=$(basename $1 .hhl)

hhlog -in $file.hhl -out "hhl" -tpl '%my_call %my_sota_ref %d %t %f %m %c' > $file.tsv
