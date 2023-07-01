module astrometers.eu/amfoc01

go 1.18

replace github.com/roman-dvorak/tinygo-drivers => ./tinygo.org/x/drivers

require (
	tinygo.org/x/drivers v0.25.0
	tinygo.org/x/tinyfont v0.3.0
)

require github.com/google/shlex v0.0.0-20191202100458-e7afc7fbc510 // indirect

replace tinygo.org/x/drivers/tmc5130 => ./tinygo.org/x/drivers/tmc5130

replace tinygo.org/x/drivers => ./tinygo.org/x/drivers
