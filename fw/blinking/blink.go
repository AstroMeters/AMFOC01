package main

import (
	//"tmc5130"


	"machine"
	"time"
  "fmt"
	"tmc5130"
	//"tmc5130"

)

var period uint64 = 1e9 / 500

func main() {
    // This program is specific to the Raspberry Pi Pico.
    //
    ser := machine.Serial

    text2 := ""
    fmt.Scanln(text2)
    fmt.Println(text2)


    machine.InitADC()
    input0 := machine.ADC{machine.ADC0}


    pwm := machine.PWM5 // Pin 25 (LED on pico) corresponds to PWM4
    pwm2 := machine.PWM6

	// Configure the PWM with the given period.
        pwm.Configure(machine.PWMConfig{ Period: period })
        pwm2.Configure(machine.PWMConfig{ Period: period+1000 })

	ch, _ := pwm.Channel(10)
        ch2, _:= pwm2.Channel(12)

        pwm.Set(ch, pwm.Top()/2)
        pwm2.Set(ch2, pwm.Top()/2)
	//if err != nil {
	//	println(err.Error())
	//	return
	//}

	type SPIConfig struct {
		Frequency uint32
		SCK       Pin
		SDO       Pin
		SDI       Pin
		Mode      uint8
	}

	machine.SPI0.Configure(machine.SPIConfig{SCK: machine.GPIO6, SDO: machine.GPIO, SDI: machine.GPIO7})
	spi_cs := machine.GPIO5

  //sensor := bmp180.New(machine.I2C0)
	motor := tmc5130.New(machine.SPI0, spi_cs)
  motor.Configure()


	for {
		            motor.GetXACT()
                println(ser.Buffered() )

//		if ser.Buffered() > 0 {
//			data, _ := ser.ReadB
//    fmt.Scanf("%s", &text2)
    fmt.Println("TEXT>", text2, "  ADC: ", input0.Get())

		//for i := 1; i < 255; i++ {
            // This performs a stylish fade-out blink
		//	pwm.Set(ch, pwm.Top()/uint32(i))
			time.Sleep(time.Millisecond * 50)
		//}
	}
}
