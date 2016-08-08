#!/bin/sh
if [ -z "$1" ] 
	then
      		echo "type your project's dir name as first argument"
		exit
fi
export GOPATH=$(pwd)/$1
echo "GOPATH configured as " $GOPATH
