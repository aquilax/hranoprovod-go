# Hranoprovod [![Build Status](https://travis-ci.org/aquilax/hranoprovod-go.svg)](https://travis-ci.org/aquilax/hranoprovod-go)

**Note: This is the legacy verion of the tool and as such, must be treated as deprecated. The new version can be found [here](https://github.com/aquilax/hranoprovod-cli)**

Command-line life tracker.

## Description

Hranoprovod is command line tracking tool. It supports recipies, which makes it 
perfect for tracking calories and other nutionin data.

## Installation

First make sure you have go (golang) installed.

    http://golang.org/

Download the source code.
  
    git clone git://github.com/aquilax/hranoprovod-go.git
    cd hranoprovod-go
	go build

## Requirements

Hranoprovod uses two files.

* database file (default: food.yaml) contains the recipies for the tracked items:

The file format is:

    element_name/measure_name[:]
      ingredient1[:] quantity
      ingredient2[:] quantity

[:] are optional (I use them for yaml highlight support)

Example:

    pie/apple/100g:
      calories: 265
      fat: 13
      carb: 37
      protein: 2
    pie/apple/slice:
      pie/apple/100g: 1.55

* log file (default: log.yaml) contains the daily log:

The file format is:

    date[:]
      element_name1/measure_name[:] quantity
      element_name2/measure_name[:] quantity

[:] are optional (I use them for yaml highlight support)

**Note**: The date format is (YYYY/MM/DD)

Example:

    2011/08/08:
      walking/slow/km: 5
      egg/boiled/pcs: 3
      brad/slice: 1

## Command-line options

### Files

  -d="food.yaml": Specifies the database file name

  -f="log.yaml": Specifies log file name

### Filtering

  -b="": Beginning of date interval (YYYY/MM/DD)

  -e="": Ending of date interval (YYYY/MM/DD)

### Output

  -single="": Show only single element
  
  -food="": Shows single food

  -total=true: Shows totals for each day

  -unresolved=false: Shows unresolved elements

  -csv=false: Shows data in CSV format (works only with --single)

### Misc

  -help=false: Shows this message

  -version=false: Shows version

# Changelog
0.1.2 (2013-07-11) Added comments support in data files;
To addd a comment start the line with # (pound sign) 

0.1.1 - First stable version
