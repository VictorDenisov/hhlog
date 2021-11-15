#!/bin/bash

# This script converts an hhl file to a file suitable for uploading to lotw.

if [[ $# -lt 1 ]]; then
	echo "Provide file name in the argument"
	exit 1
fi

file=$(basename $1 .hhl)

hhlog -in $file.hhl -out "adi" -tpl "%c %b %m %d %t %f" > $file.adi
