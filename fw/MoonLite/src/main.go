package main

import (
	//"tinygo.org/x/drivers"

	"fmt"
	"machine"
	"time"

	//"tmc5130"
	//"tmc5130"
	//"image/color"

	"tinygo.org/x/drivers/tmc5130"

	//"tinygo.org/x/drivers/st7789"
	"math/big"
	//"strconv"
	"strings"

	//"tinygo.org/x/drivers/ds18b20"
	"tinygo.org/x/drivers/sh1106"

	//
	// "tinygo.org/x/tinyfont/freeserif"
	// "tinygo.org/x/tinyfont/gophers"

	// "tinygo.org/x/tinyfont"
	// "tinygo.org/x/tinyfont/freesans"
	// "tinygo.org/x/tinyfont/freemono"
	// "tinygo.org/x/tinyfont/freesans"
	// "tinygo.org/x/tinyfont/freeserif"
	// "tinygo.org/x/tinyfont/gophers"

	"github.com/roman-dvorak/user_interface"
)

const (
	TEMP_INTERNAL = machine.ADC1
	TEMP_EXTERNAL = machine.ADC2
	BTN_EXT       = machine.GPIO26
	BTN_EXT_ADC   = machine.ADC0

	OLED_CS  = machine.GPIO9
	OLED_DC  = machine.GPIO10
	OLED_RST = machine.GPIO11
	//IPS_BL = machine.GPIO8

	TMC_CS = machine.GPIO5

	SPI_CLK  = machine.GPIO6
	SPI_MOSI = machine.GPIO7
	SPI_MISO = machine.GPIO4
)

type focuser_status struct {
	new_pos          int
	actual_pos       int
	temp_coef        int
	temp_coef_offset int
	steeper_speed    uint8
	step_mode        uint8
	scale            int
}

func int_external_button(p machine.Pin) {
	println("Bylo stisknuto tlacitko ... ", p)
	analog_value := machine.ADC{BTN_EXT_ADC}.Get()

	if analog_value > (2 ^ 16/10*9) {

	} else if analog_value > (2 ^ 16/10*8) {

	} else if analog_value > (2 ^ 16/10*7) {

	} else if analog_value > (2 ^ 16/10*6) {

	} else if analog_value > (2 ^ 16/10*5) {

	}

}

func init_motor(motor *tmc5130.Device, display sh1106.Device) {

	// display.ClearBuffer()
	// tinyfont.WriteLine(&display, &freesans.Regular9pt7b, 0, 13, "AMFOC01", color.RGBA{100, 100, 100, 100})
	// tinyfont.WriteLine(&display, &freesans.Regular9pt7b, 0, 30, "PWR UP", color.RGBA{255, 255, 255, 255})
	// tinyfont.WriteLine(&display, &freesans.Regular9pt7b, 0, 47, "Initializing motor", color.RGBA{255, 255, 255, 255})
	// display.Display()

	motor.SetRegister(tmc5130.GCONF|tmc5130.WRITE, 0x00000000)    //GCONF
	motor.SetRegister(tmc5130.CHOPCONF|tmc5130.WRITE, 0x000101D5) //CHOPCONF: TOFF=5, HSTRT=5, HEND=3, TBL=2, CHM=0 (spreadcycle)
	//motor.SetRegister(0x90,0x00070603); //IHOLD_IRUN: IHOLD=3, IRUN=10 (max.current), IHOLDDELAY=6
	//motor.SetRegister(0x90,(2 &0b11111)<<0|(2 &0b11111)<<8|(1&0b1111)<<16);
	motor.SetCurrent(0, 50, 0)
	motor.SetRegister(tmc5130.TPOWERDOWN|tmc5130.WRITE, 10)
	motor.SetRegister(tmc5130.PWM_CONF|tmc5130.WRITE, 0x00000000)
	//PWM_CONF: autoscale=1, 2/1024 Fclk, Switch amp limit=200, grad=1
	motor.SetRegister(tmc5130.PWM_CONF|tmc5130.WRITE, 0x000401C8)

	motor.SetRegister(tmc5130.A1|tmc5130.WRITE, 1000)
	motor.SetRegister(tmc5130.V1|tmc5130.WRITE, 100000)
	motor.SetRegister(tmc5130.AMAX|tmc5130.WRITE, 5000)
	motor.SetRegister(tmc5130.VMAX|tmc5130.WRITE, 100000)
	motor.SetRegister(tmc5130.D1|tmc5130.WRITE, 1400)
	motor.SetRegister(tmc5130.VSTOP|tmc5130.WRITE, 10)

	motor.SetRegister(tmc5130.RAMPMODE|tmc5130.WRITE, 0)

	motor.SetRegister(tmc5130.XACTUAL|tmc5130.WRITE, 0)
	motor.SetRegister(tmc5130.XTARGET|tmc5130.WRITE, 0)
	//	motor.SetRegister(tmc5130.XTARGET|tmc5130.WRITE,10000);
}

func main() {

	machine.SPI0.Configure(machine.SPIConfig{
		SCK:  SPI_CLK,
		SDO:  SPI_MOSI,
		SDI:  SPI_MISO,
		Mode: 3})
	machine.SPI0.SetBaudRate(115200 * 32)
	//machine.SPI0.SetBaudRate(115200*64)

	display := sh1106.NewSPI(machine.SPI0, OLED_DC, OLED_RST, OLED_CS)

	display.Configure(sh1106.Config{
		Width:  128,
		Height: 64,
	})

	OLED_CS.Configure(machine.PinConfig{Mode: machine.PinOutput})

	//drawDisp(&display)

	ui := user_interface.New(display)
	ui.Configure()
	//ui.DrawSplashScreen()

	time.Sleep(2000 * time.Millisecond)

	ser := machine.Serial

	var FocuserStatus focuser_status
	FocuserStatus.scale = 1

	machine.InitADC()
	//thermistor := machine.ADC{machine.THERMISTOR}

	ext_pwr_det := machine.EXT_PWR_DET
	ext_pwr_det.Configure(machine.PinConfig{Mode: machine.PinInput})

	pin_btn_a := machine.BTN_A
	pin_btn_a.Configure(machine.PinConfig{Mode: machine.PinInput})
	pin_btn_a.SetInterrupt(machine.PinRising, func(p machine.Pin) {
		ui.Btn_down()
	})
	pin_btn_b := machine.BTN_B
	pin_btn_b.Configure(machine.PinConfig{Mode: machine.PinInput})
	pin_btn_b.SetInterrupt(machine.PinRising, func(p machine.Pin) {
		ui.Btn_up()
	})
	pin_btn_c := machine.BTN_C
	pin_btn_c.Configure(machine.PinConfig{Mode: machine.PinInput})
	pin_btn_c.SetInterrupt(machine.PinRising, func(p machine.Pin) {
		ui.Btn_set()
	})

	BTN_EXT.Configure(machine.PinConfig{Mode: machine.PinOutput})
	BTN_EXT.SetInterrupt(machine.PinRising, int_external_button)

	pwm3 := machine.PWM3
	pwm2 := machine.PWM2

	mot_en := machine.GPIO3
	mot_en.Configure(machine.PinConfig{Mode: machine.PinOutput})
	mot_en.Low()

	motor := tmc5130.New(machine.SPI0, machine.GPIO5)
	motor.Configure()
	init_motor(motor, display)

	pwm3.Configure(machine.PWMConfig{Period: 10})
	pwm2.Configure(machine.PWMConfig{Period: 9})

	ch3, _ := pwm3.Channel(22)
	pwm3.Set(ch3, pwm3.Top()/2)
	ch2, _ := pwm2.Channel(20)
	pwm2.Set(ch2, pwm2.Top()/3)

	var input_string string

	var temp int

	var last_ext_pwr bool
	last_ext_pwr = false

	ui.PrepareScreen()
	time.Sleep(time.Second * 3)

	// ow := onewire.New(machine.ADC1)
	// romIDs, err := ow.Search(onewire.SEARCH_ROM_COMMAND)
	// if err != nil {
	// 	println(err)
	// }
	// ds := ds18b20.New(ow)

	// for {
	// 	time.Sleep(3 * time.Second)

	// 	println()
	// 	println("Device:", machine.Device)

	// 	println()
	// 	println("Request Temperature.")
	// 	for _, romid := range romIDs {
	// 		println("Sensor RomID: ", hex.EncodeToString(romid))
	// 		ds.RequestTemperature(romid)
	// 	}

	// 	// wait 750ms or more for DS18B20 convert T
	// 	time.Sleep(1 * time.Second)

	// 	println()
	// 	println("Read Temperature")
	// 	for _, romid := range romIDs {
	// 		raw, err := ds.ReadTemperatureRaw(romid)
	// 		if err != nil {
	// 			println(err)
	// 		}
	// 		println()
	// 		println("Sensor RomID: ", hex.EncodeToString(romid))
	// 		println("Temperature Raw value: ", hex.EncodeToString(raw))

	// 		t, err := ds.ReadTemperature(romid)
	// 		if err != nil {
	// 			println(err)
	// 		}
	// 		println("Temperature in celsius milli degrees (°C/1000): ", t)
	// 	}
	// }

	ui.SetScreen(user_interface.Home)

	for {

		if last_ext_pwr != ext_pwr_det.Get() && ext_pwr_det.Get() {
			init_motor(motor, display)
			last_ext_pwr = ext_pwr_det.Get()
		}

		FocuserStatus.actual_pos = int(motor.GetXACTUAL().XACTUAL / int32(FocuserStatus.scale))
		//println("POS>", FocuserStatus.actual_pos)

		// display.ClearBuffer()
		// tinyfont.WriteLine(&display, &freesans.Regular9pt7b, 0, 13, "AstroMeters", color.RGBA{100, 100, 100, 100})
		// tinyfont.WriteLine(&display, &freesans.Regular9pt7b, 0, 30, "Temp: "+strconv.Itoa(temp)+" C", color.RGBA{255, 255, 255, 255})
		// tinyfont.WriteLine(&display, &freesans.Regular9pt7b, 0, 47, "pos: "+strconv.Itoa(FocuserStatus.actual_pos/1000)+"", color.RGBA{255, 255, 255, 255})
		// display.Display()
		ui.SetData(float64(temp), -100, FocuserStatus.actual_pos, ext_pwr_det.Get())
		//ui.Action()

		if ser.Buffered() > 0 {
			for ser.Buffered() > 0 {
				s, _ := ser.ReadByte()
				input_string += string(s)
			}
			for (strings.Index(input_string, ":") + strings.Index(input_string, "#")) > 0 {
				cmd := input_string[strings.Index(input_string, ":")+1 : strings.Index(input_string, "#")]
				input_string = input_string[strings.Index(input_string, "#")+1:]
				if len(cmd) == 1 {
					cmd += " "
				}
				switch cmd[0:2] {
				case "C ":
					// Initiate temperature conversion
					// ds.Reset()
					// ds.SkipRom()
					// ds.ConvertT()

				case "FG":
					//Go to positition set by SNYYYY
					motor.SetXTARGET(FocuserStatus.new_pos * FocuserStatus.scale)

				case "FQ":
					motor.SetXTARGET(int(motor.GetXACTUAL().XACTUAL))
					//STOP ALL MOTION

				case "GC":
					// return temperature coeficient
					fmt.Printf("%02x#\n\r", FocuserStatus.temp_coef)

				case "GD":
					// return current stepping delay
					fmt.Printf("%02x#\n\r", FocuserStatus.steeper_speed)

				case "GH":
					// return half or full step mode
					fmt.Printf("FF#\n\r") // halfstep
					// fmt.Printf("00#\n\r") // fullstep

				case "GI":
					// return if motor is in movement
					println(motor.GetVACTUAL().VACTUAL)
					if motor.GetVACTUAL().VACTUAL > 1 || motor.GetVACTUAL().VACTUAL < -1 {
						fmt.Printf("01#\n\r") // moving
					} else {
						fmt.Printf("00#\n\r") // stale
					}

				case "GN":
					// return new target position
					fmt.Printf("%04x#\n\r", FocuserStatus.new_pos)

				case "GP":
					// return current position
					fmt.Printf("%04x#\n\r", FocuserStatus.actual_pos)

				case "GT":
					// return current temperature position
					//temp := int((input0.Get()-500)/100)
					var sign int
					if temp < 0 {
						temp = -temp
						sign = 0b1 << 15
					}
					fmt.Printf("%04x#\n\r", int(temp)|sign)

				case "GV":
					// return firmware version
					fmt.Printf("10#\n\r")

				case "SC":
					// Set temperature coeficient
					tmp, _ := new(big.Int).SetString(cmd[2:], 16)
					num := tmp.Int64()
					println("KOEF", num)

				case "SD":
					// Set speed
					tmp, _ := new(big.Int).SetString(cmd[2:], 16)
					num := int(tmp.Int64())
					FocuserStatus.steeper_speed = uint8(num)
					motor.SetRegister(tmc5130.VMAX|tmc5130.WRITE, (25-num)*10000)

				case "SF":
					// Set fullstep mode
					//FocuserStatus.scale = 100000
					//println("FS")

				case "SH":
					// Set halfstep mode
					//FocuserStatus.scale = 100000
					//println("HS")

				case "SN":
					// Set position
					tmp, _ := new(big.Int).SetString(cmd[2:], 16)
					num := int(tmp.Int64())
					//println("SetPos", num)
					FocuserStatus.new_pos = num

				case "SP":
					// Set current position
					tmp, _ := new(big.Int).SetString(cmd[2:], 16)
					num := int(tmp.Int64())
					motor.SetXACTUAL(num * FocuserStatus.scale)
					motor.SetXTARGET(num * FocuserStatus.scale)
					//println("SetPos", num)

				case "+ ":
					// Enable temperature coeif
					//println("Activate temperature coeficient")

				case "- ":
					// Disable temperature coeif
					//println("Disable temperature coeficient")

				case "PO":
					// Temperature calibration offset
					tmp, _ := new(big.Int).SetString(cmd[2:], 16)
					num := int(tmp.Int64())
					FocuserStatus.temp_coef_offset = num
					//println("TempOffset", num)

				default:
					println("Unknown command...", input_string)

				}
			}
		}
		//println("TEXT>", "  ADC: ", input0.Get())
		//println(pin_btn_a.Get(), pin_btn_b.Get(), pin_btn_c.Get())

		time.Sleep(time.Millisecond * 10)
	}
}
