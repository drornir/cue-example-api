#!/bin/bash
source ".bingo/variables.env"

set -e

function gen_api() {
	(cd api
		$CUE get go --local
	)
}

function gen_gotags() {
	for f in $(find ./api -name '*go') ; do
		echo generating json tags for $f
		$GOMODIFYTAGS -file $f -all -add-tags json -transform snakecase -w  --quiet
	done
}

function gen_govalidate() {
	echo gen_govalidate: go run ./cmd gen go-validators
	go run ./cmd gen go-validators
}

function build () {
	echo gen_gotags
	gen_gotags
	echo gen_api
	gen_api
	echo gen_govalidate
	gen_govalidate
	echo build done
}



function main() {
    local _commands=("build" "gen_api" "gen_gotags" "gen_govalidate")

    local _cmd=$($GUM choose --header='choose a command to run' ${_commands[*]})

    if [[ " ${_commands[*]} " =~ " ${_cmd} " ]]; then
        echo "executing '$_cmd'"
        $_cmd
        exit 0
    else
    	if test "$_cmd" = 'user aborted'; then
     		exit 1
        fi
        echo "Invalid command: '$_cmd'"
        exit 2
    fi
}

main
