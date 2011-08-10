## Description

Hranoprovod is command line tracking tool. It supports recipies, which makes it 
perfect for tracking calories and other nutionin data.

## Installation

First make sure you have go (golang) installed.

    http://golang.org/

Download the source code.
  
    git clone git://github.com/aquilax/hranoprovod-go.git
    cd hranoprovod-go
    gomake
    sudo gomake install

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

### Misc

  -help=false: Shows this message
  -version=false: Shows version


