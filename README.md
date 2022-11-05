# go-charmagic

> Expect heavy refactoring of the top-level API.

Some encodings are still missing and/or not supported.

Tests are missing, encodings not supported and the top-level API not stable. This project is far from done and still a WIP.

## CLI

The project provides a basic CLI, which you can use to guess a file's encoding or transform a file with arbitrary encoding to utf-8.

### Usage

#### guess
Displays a ranking of all supported encodings and a confidence value.

`charmagic guess <flags>` 

##### Flags
- `-i <file>` or `--input <file>` source file

#### transform
Takes a input file and transforms it into utf-8.

`charmagic transform <flags>` 

##### Flags
- `-i <file>` or `--input <file>` source file
- `-o <file>` or `--output <file>` destination file 
- `-e <name>` or `--encoding <name>` encoding to use
  - if omitted: tries to guess encoding
