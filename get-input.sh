#!/bin/bash

usage() {
	echo "$0 <year> <day>"
}

if [ "$#" -ne 2 ]; then
	usage
	exit 1
fi

year=$1
day=$2
zday=$(printf "%02d" ${day})

dir="input/${year}/${zday}"

mkdir -p "${dir}"

if [ -z "${SESSION_COOKIE}" ]; then
	echo "SESSION_COOKIE not defined"
	exit 2
fi

curl -b "session=$SESSION_COOKIE" "https://adventofcode.com/${year}/day/${day}/input" -o "${dir}/input.txt"
