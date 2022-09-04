#!/bin/sh

# By default run all tests with -verbose mode on and clean all cache with count = 1 
if [ $# -eq 0 ]
then
  go test -v -count=1 ./...
  exit 0
fi


while getopts t:p:h: flag 
do 
  case "${flag}" in 
    t) test_name="${OPTARG}";;
    p) package_path="${OPTARG}";;
    h) help=true;;
    *) help=true;;
  esac
done

if [ $help ]
then
  printf "%s accepts follwoing args: 
    -t name of the test 
    -p path to the package
    
    If you want to run all tests invoke without any args.

    \n" "${0}"
  exit 0
fi

if [ "$test_name" ] && [ "$package_path" ]
then
  go test -v -count=1 -run  "$test_name" "$package_path"
fi

