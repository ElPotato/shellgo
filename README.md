# shellgo
Parse bytecode from `go tool objdump` output.

## How to?

* `make`
* `go tool compile -S -N file.go`
* `go tool objdump -S file.o | shellgo`

### Flags

* `-dump <file_name>` / Result in `file_name.in` and `file_name.out` respectively to `stdin` and `stdout`.
* `-version` / Print out build version
* `-0x` / Change format from packed (02) to formatted (0x00, 0x02, ).
