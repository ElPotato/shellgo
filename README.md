# shellgo
Parse shellcode from `go tool objdump -S <file>` output.

## How to?

* `make && make install`
* `go tool compile -S -N file.go`
* `go tool objdump -S file.o | shellgo`