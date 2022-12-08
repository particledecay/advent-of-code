#!/bin/bash

# create the structure
while read -r line
do
  if [[ "$line" == "$"* ]]
  then
    if [[ "$line" == "$ cd"* ]]
    then
      dirname="${line/$ cd /}"
      [ -d "$dirname" ] || mkdir "$dirname"
      cd "$dirname"
    else
      command="${line/$ /}"
      eval "$command"
    fi
  elif [[ "$line" != "dir "* ]]
  then
    filesize="$( echo "$line" | cut -d' ' -f1 )"
    filename="$( echo "$line" | cut -d' ' -f2 )"
    touch "${filesize}-${filename}"
  fi
done < "$1"
