# Project "Difference Generator"
### Hexlet tests and linter status:
[![Actions Status](https://github.com/slamix/go-project-244/actions/workflows/hexlet-check.yml/badge.svg)](https://github.com/slamix/go-project-244/actions)

### Description
Difference generator determines the difference between two data structures and shows the result in the selected format.
### Features
* Support for JSON and YAML configuration file formats 
* You can select output format: stylish(default), plain or JSON
* Recursive comparison of flat and nested structures

### Usage
```
Usage: gendiff [options] <filepath1> <filepath2>

Compares two configuration files and shows a difference.

Options:
  -V, --version        output the version number
  -f, --format [type]  output format (default: "stylish")
  -h, --help           output usage information
```
## Examples with flat files
### Data type - JSON, format - stylish

### Data type - YML, format - stylish 

## Examples with nested files
### Data type - JSON, format - stylish

### Data type - JSON, format - plain

### Data type - JSON, format - json

### Data type - YML, format - stylish

### Data type - YML, format - plain

### Data type - YML, format - json
