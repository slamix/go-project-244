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
[![asciicast](https://asciinema.org/a/u9fB0s48oqfEoD4k.svg)](https://asciinema.org/a/u9fB0s48oqfEoD4k)
### Data type - YML, format - stylish 
[![asciinema](https://asciinema.org/a/hPBmKWvdkJIxK7eX.svg)](https://asciinema.org/a/hPBmKWvdkJIxK7eX)
## Examples with nested files
### Data type - JSON, format - stylish
[![asciinema](https://asciinema.org/a/Ya9nLF1WXEb45RVd.svg)](https://asciinema.org/a/Ya9nLF1WXEb45RVd)
### Data type - JSON, format - plain
[![asciinema](https://asciinema.org/a/H5ETPGuXenvqljj3.svg)](https://asciinema.org/a/H5ETPGuXenvqljj3)
### Data type - JSON, format - json
[![asciinema](https://asciinema.org/a/XW9KmjhBHYCZhugG.svg)](https://asciinema.org/a/XW9KmjhBHYCZhugG)
### Data type - YML, format - stylish
[![asciinema](https://asciinema.org/a/QpM2D0JXgcDwQJHM.svg)](https://asciinema.org/a/QpM2D0JXgcDwQJHM)

