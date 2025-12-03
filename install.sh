#!/bin/bash

DAY=$(date +'%-d')
YEAR=$(date +'%Y')
SESSION="${AOC_SESSION}"

usage() {
    echo Usage: $0 [-d day] [-y year] [-s session_cookie]""
    echo "  -d  Day number (default: today's day, ${DAY})"
    echo "  -y  Year number (default: today's year, ${YEAR})"
    echo "  -s  Session cookie (default: value of \$AOC_SESSION)"
}

while getopts "dys:" flag; do
    case "${flag}" in
    d) DAY=$OPTARG ;;
    y) YEAR=$OPTARG ;;
    s) SESSION=$OPTARG ;;
    *)
        usage
        exit 1
        ;;
    esac
done

if [[ -z "${SESSION}" ]]; then
    echo "Error: Session cookie is required."
    usage
    exit 1
fi

PADDED_DAY=$(printf "%02d" "${DAY}")
DIR_PATH="${YEAR}/day${PADDED_DAY}"

mkdir -p "${DIR_PATH}"
echo "Directory: ${DIR_PATH}"

MAIN_FILE="${DIR_PATH}/main.go"
if [[ -f "${MAIN_FILE}" ]]; then
    echo "${MAIN_FILE} already exists. Skipping."
else
    cat <<EOF >"${MAIN_FILE}"
package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	input = strings.TrimSpace(input)
	fmt.Println(input)
}
EOF
    echo "Created file: ${MAIN_FILE}"
fi

INPUT_FILE="${DIR_PATH}/input.txt"
URL="https://adventofcode.com/${YEAR}/day/${DAY}/input"

if [[ -f "${INPUT_FILE}" ]]; then
    echo "${INPUT_FILE} already exists. Skipping download."
else
    curl --cookie "session=${AOC_SESSION}" "${URL}" >"${INPUT_FILE}"
    echo "Downloaded: ${INPUT_FILE}"
fi
