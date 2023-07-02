module astrometers.eu/amfoc01

go 1.18

replace github.com/roman-dvorak/tinygo-drivers => ./tinygo.org/x/drivers

replace github.com/roman-dvorak/user_interface => ./user_interface

require (
	github.com/roman-dvorak/user_interface v0.0.0-00010101000000-000000000000
	tinygo.org/x/drivers v0.25.0
	tinygo.org/x/tinyfont v0.4.0

)

replace tinygo.org/x/drivers/tmc5130 => ./tinygo.org/x/drivers/tmc5130

replace tinygo.org/x/drivers => ./tinygo.org/x/drivers
