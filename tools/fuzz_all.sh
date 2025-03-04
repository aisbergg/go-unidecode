#!/bin/bash
set -euo pipefail

fuzz_time=${1:-1m}

files=$(grep -r --include='**_test.go' --files-with-matches 'func Fuzz' .)
for file in ${files}; do
    funcs=$(grep '^func Fuzz' "$file" | sed s/func\ // | sed 's/(.*$//')
    for func in ${funcs}; do
        echo "Fuzzing $func in $file for $fuzz_time"
        parent_dir=$(dirname "$file")
        go test "$parent_dir" -run="$func" -fuzz="$func" -fuzztime="${fuzz_time}"
    done
done
