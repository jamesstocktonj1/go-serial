# Go Serial

This is a lightweight command line based serial console. It is based around Go purely because I want to get to know the language better.


## Usage
```
Usage of Go Serial:
  -baud int
        select Baud Rate (default 9600)
  -d    select Default Serial Port
  -data int
        select number of data bits (default 8)
  -l    list available ports
  -parity string
        select parity type [None, Odd Even, Mark, Space] (default "None")
  -stop int
        select number of stop bits (default 1)
  -v    verbose output e.g. connect/disconnect
```

## Installation
To create an executable run the following command: 
```
go build 
```

To install go-serial to the GOPATH run:
```
go install
```