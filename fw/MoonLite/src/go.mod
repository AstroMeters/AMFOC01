module astrometers.eu/amfoc01

go 1.18

replace github.com/roman-dvorak/tinygo-drivers => ./tinygo.org/x/drivers

replace github.com/roman-dvorak/user_interface => ./user_interface

require (
	github.com/roman-dvorak/user_interface v0.0.0-00010101000000-000000000000
	tinygo.org/x/drivers v0.25.0

)

require (
	github.com/13rac1/fastmath v1.0.0 // indirect
	github.com/aykevl/go-wasm v0.0.2-0.20220616010729-4a0a888aebdc // indirect
	github.com/blakesmith/ar v0.0.0-20190502131153-809d4375e1fb // indirect
	github.com/creack/goselect v0.1.2 // indirect
	github.com/gofrs/flock v0.8.1 // indirect
	github.com/google/shlex v0.0.0-20191202100458-e7afc7fbc510 // indirect
	github.com/inhies/go-bytesize v0.0.0-20220417184213-4913239db9cf // indirect
	github.com/marcinbor85/gohex v0.0.0-20210308104911-55fb1c624d84 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.19 // indirect
	github.com/mattn/go-tty v0.0.5 // indirect
	github.com/sigurn/crc16 v0.0.0-20211026045750-20ab5afb07e3 // indirect
	github.com/tinygo-org/tinygo v0.28.1 // indirect
	go.bug.st/serial v1.5.0 // indirect
	golang.org/x/sys v0.10.0 // indirect
	golang.org/x/tools v0.11.0 // indirect
	tinygo.org/x/go-llvm v0.0.0-20230505123812-8e7ec80422a4 // indirect
	tinygo.org/x/tinyfont v0.4.0 // indirect
)

replace tinygo.org/x/drivers/tmc5130 => ./tinygo.org/x/drivers/tmc5130

replace tinygo.org/x/drivers => ./tinygo.org/x/drivers
