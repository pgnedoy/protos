#!/bin/bash


while getopts s:v: flag
do
    case "${flag}" in
        s) service=${OPTARG};;
        v) api_version=${OPTARG};;
    esac
done

path=$service/$api_version
files=""

for file in $(find ./$path -name "*.proto"  -type f);
do
  files="$files "protos/$path/$(echo $file | rev | cut -d'/' -f-1 | rev)""
done

echo $files


