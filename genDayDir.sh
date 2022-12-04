#!/bin/bash

# Generates a directory for a new AoC day

dayName=$1

: ${1?"Usage: $0 Name of the day"}

echo "Generating $dayName dir"

mkdir "$dayName"

touch "$dayName"/input.txt
touch "$dayName"/example.txt
touch "$dayName"/README.md

echo -e "package main\nimport (\n\t\"fmt\"\n)\n\nfunc main() {\n\tfmt.Println(\"$dayName\")\n}" >> $dayName"/"$dayName".go"
echo -e "package main\nimport (\n\t\"testing\"\n)\n\nfunc BenchmarkP1(b *testing.B){\n\tpart1()\n}\n\nfunc BenchmarkP2(b *testing.B){\n\tmain_old()\n}"

