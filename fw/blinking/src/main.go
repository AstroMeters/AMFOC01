package main

import (
	//"tmc5130"


	"machine"
	"time"
  "fmt"
	//"tmc5130"
	//"tmc5130"
	"tinygo.org/x/drivers/tmc5130"

)

var period uint64 = 1e9 / 500

func main() {
    // This program is specific to the Raspberry Pi Pico.
    //
		time.Sleep(time.Millisecond * 1000)
    //ser := machine.Serial

    text2 := ""
    fmt.Scanln(text2)
    fmt.Println("start")


    machine.InitADC()
    input0 := machine.ADC{machine.ADC0}


    //pwm := machine.PWM5 // Pin 25 (LED on pico) corresponds to PWM4
    //pwm2 := machine.PWM6

    pwm3 := machine.PWM4


		machine.SPI0.Configure(machine.SPIConfig{
			  //Frequency: 10000, // fungovalo /
			  Frequency: 8000,
		    //Frequency: 1000,
				//SCK: machine.GPIO18,
				//SDO: machine.GPIO19,
				//SDI: machine.GPIO16,
				SCK: machine.GPIO6,
				SDO: machine.GPIO7,
				SDI: machine.GPIO8,
				Mode: 3})
	  //machine.SPI0.SetBaudRate(48000)
		spi_cs := machine.GPIO9
		motor := tmc5130.New(machine.SPI0, spi_cs)
		motor.Configure()

	// Configure the PWM with the given period.
        //pwm.Configure(machine.PWMConfig{ Period: period })
        //pwm2.Configure(machine.PWMConfig{ Period: period+1000 })
        pwm3.Configure(machine.PWMConfig{ Period: 1e9/4 })

				//ch, _ := pwm.Channel(10)
        //ch2, _:= pwm2.Channel(12)
        ch3, _:= pwm3.Channel(25)

        //pwm.Set(ch, pwm.Top()/2)
        //pwm2.Set(ch2, pwm.Top()/2)
        pwm3.Set(ch3, pwm3.Top()/2)
	//if err != nil {
	//	println(err.Error())
	//	return
	//}
  a := []byte{1, 1, 1, 1, 1}
								//a := motor.GetXACT(0x01)
								//println(a[0], a[1], a[2], a[3])
  motor.SetRegister(0xEC, []byte{0x00, 0x01, 0x00, 0xc3})
  motor.SetRegister(0x90, []byte{0x00, 0x06, 0x1f, 0x0a})
  motor.SetRegister(0x91, []byte{0x00, 0x00, 0x00, 0x0a})
  motor.SetRegister(0x80, []byte{0x00, 0x00, 0x00, 0x04})
  motor.SetRegister(0x93, []byte{0x00, 0x00, 0x01, 0xf4})
  motor.SetRegister(0xf0, []byte{0x00, 0x04, 0x01, 0xc8})
	for {
		println("")
								// a = motor.GetXACT(0x01)
								// time.Sleep(time.Millisecond * 5)
								// println(",", a[0], a[1], a[2], a[3], a[4])

								a = motor.GetXACT(0x21)
								println(":", a[0], a[1], a[2], a[3], a[4])

//		if ser.Buffered() > 0 {
//			data, _ := ser.ReadB
//    fmt.Scanf("%s", &text2)
    println("TEXT>", text2, "  ADC: ", input0.Get())

		//for i := 1; i < 255; i++ {
            // This performs a stylish fade-out blink
		//	pwm.Set(ch, pwm.Top()/uint32(i))
			time.Sleep(time.Millisecond * 200)
		//}
	}
}
