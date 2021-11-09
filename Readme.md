# convertHIBP

convertHIBP is a golang tool to convert a hash ordered txt file from [HIBP](https://haveibeenpwned.com/Passwords) to a binary file.
This will cut the file size in half and enable other tools to do a proper binary search over the file.

## Installation

```
git clone https://github.com/fblz/convertHIBP.git
cd convertHIBP
go build
#  <or>
go install
```

## Usage

```
convertHIBP
  -InputFile string
        Specify the ascii hibp file (default "none")
  -OutputFile string
        Specify the name for the binary target file (default "none")
```
