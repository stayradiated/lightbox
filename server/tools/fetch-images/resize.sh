#!/bin/sh

mkdir -p $2

for file in $1/*.jpg; do
  outfile=$2/`basename $file`
  echo convert "'$file'" -resize $3 -unsharp 2x0.5+0.7+0 -quality 98 "'$outfile'"
done | gm batch -
