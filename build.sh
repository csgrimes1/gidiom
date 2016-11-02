#!/usr/bin/env bash

set -e
MYDIR=$(cd $(dirname "$0") && pwd)
STRIPPATH="$GOPATH/src/"
LEN=${#STRIPPATH}
MYGOPATH=${MYDIR:LEN}

ARDIRS=( $(ls -1F "$MYDIR" | grep '/') )

go build "$MYGOPATH/cli"
go test "$MYGOPATH/gen"
go install "$MYGOPATH/cli"