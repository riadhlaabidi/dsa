#!/bin/bash

usage() {
    echo "Usage: `basename $0` day year"
    echo -e "\tday: between 1..25 (if omitted year should also be omitted)"
    echo -e "\tyear: between 2015..$current_year (if only day is specified, year is the current year)"
    echo
    echo -e "\tNo arguments: setup the current day if it's already December this year"
    exit 1
}

current_day=$(date +%-d)
current_month=$(date +%-m)
current_year=$(date +%Y)

day=${1:-$current_day}
year=${2:-$current_year}

if [ -z "$1" ]; then
    if [[ $current_month -ne 12 || $current_day -gt 25 ]]; then
        echo "It's not December yet, You should provide a day number to setup"    
        usage 
    fi
fi

if [[ "$day" -lt 1 || "$day" -gt 25 ]]; then
    echo "Invalid day number"    
    usage 
fi

if [[ $year -lt 2015 || $year -gt $current_year ]]; then
    echo "Invalid year"    
    usage
fi

# create year folder if it doesn't exists
day_path="$year/day$day"
mkdir -p $day_path

src_file_path="$day_path/day$day.go"
test_file_path="$day_path/day$day.test"
input_file_path="$day_path/day$day.input"

if [ -f "$src_file_path" ]; then
    echo "Source code file \"$src_file_path\" exists, Skipping..."
else
    touch "$day_path/day$day.go"
fi

if [ -f "$test_file_path" ]; then
    echo "Test input file \"$test_file_path\" exists, Skipping..."
else
    touch "$day_path/day$day.test"
fi

if [ -f "$input_file_path" ]; then
    echo "Input file \"$input_file_path\" exists, Skipping..."
else
    touch "$day_path/day$day.input"
    input_url="https://adventofcode.com/$year/day/$day/input"
    if [ ! -f ".env" ]; then
        echo ".env file not found!"
        exit 1 
    fi
    COOKIE_SESSION=$(cat .env)
    if [ -z "$COOKIE_SESSION" ]; then
        echo "No session cookie found in .env"
        exit 1
    fi
    curl "$input_url" -H "Cookie: session=$COOKIE_SESSION" > $input_file_path
fi

exit 0
