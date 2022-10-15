package main

import (
	//"tmc5130"
	"machine"
	"time"
  "fmt"
	//"tmc5130"
	//"tmc5130"
	"tinygo.org/x/drivers/tmc5130"
	"strings"
	"math/big"
)

var period uint64 = 1e9 / 500


type focuser_status struct {
		new_pos int
		temp_coef int
		temp_coef_offset int
		steeper_speed	uint8
		step_mode	uint8
		scale int
}

func main() {
    // This program is specific to the Raspberry Pi Pico.
    //
		time.Sleep(time.Millisecond * 1000)
    ser := machine.Serial

    //fmt.Println("start")
		var FocuserStatus focuser_status
		FocuserStatus.scale = 1000

    machine.InitADC()
    input0 := machine.ADC{machine.ADC0}

    pwm3 := machine.PWM4

		machine.SPI0.Configure(machine.SPIConfig{
				SCK: machine.GPIO6,
				SDO: machine.GPIO7,
				SDI: machine.GPIO4,
				Mode: 3})
	  machine.SPI0.SetBaudRate(115200)
		motor := tmc5130.New(machine.SPI0, machine.GPIO5)
		motor.Configure()

	  pwm3.Configure(machine.PWMConfig{ Period: 1e9/4 })

	  ch3, _:= pwm3.Channel(25)
	  pwm3.Set(ch3, pwm3.Top()/2)

	 time.Sleep(time.Millisecond * 1000)

	 motor.SetRegister(tmc5130.GCONF|tmc5130.WRITE,			0x00000000); //GCONF
	 motor.SetRegister(tmc5130.CHOPCONF|tmc5130.WRITE,	0x000101D5); //CHOPCONF: TOFF=5, HSTRT=5, HEND=3, TBL=2, CHM=0 (spreadcycle)
	 //motor.SetRegister(0x90,0x00070603); //IHOLD_IRUN: IHOLD=3, IRUN=10 (max.current), IHOLDDELAY=6
	 //motor.SetRegister(0x90,(2 &0b11111)<<0|(2 &0b11111)<<8|(1&0b1111)<<16);
	 motor.SetCurrent(1,10, 0)
	 motor.SetRegister(tmc5130.TPOWERDOWN|tmc5130.WRITE,10)
	 motor.SetRegister(tmc5130.PWM_CONF|tmc5130.WRITE,	0x00000000)
	 //PWM_CONF: autoscale=1, 2/1024 Fclk, Switch amp limit=200, grad=1
	 motor.SetRegister(tmc5130.PWM_CONF|tmc5130.WRITE, 	0x000401C8)

	 motor.SetRegister(tmc5130.A1|tmc5130.WRITE,				1000)
	 motor.SetRegister(tmc5130.V1|tmc5130.WRITE,				100000)
	 motor.SetRegister(tmc5130.AMAX|tmc5130.WRITE,			5000)
	 motor.SetRegister(tmc5130.VMAX|tmc5130.WRITE,			100000)
	 motor.SetRegister(tmc5130.D1|tmc5130.WRITE,				1400)
	 motor.SetRegister(tmc5130.VSTOP|tmc5130.WRITE, 		10)

	 motor.SetRegister(tmc5130.RAMPMODE|tmc5130.WRITE,	0)

	 motor.SetRegister(tmc5130.XACTUAL|tmc5130.WRITE,		0)
	 motor.SetRegister(tmc5130.XTARGET|tmc5130.WRITE,		0)
 	//	motor.SetRegister(tmc5130.XTARGET|tmc5130.WRITE,10000);

	//trg := 0
	var input_string string
	//var cmd [][]byte

	for {
		//println("")
		//println("CMD", input_string)

		if(ser.Buffered()>0){
				for(ser.Buffered()>0){
					s, _ := ser.ReadByte()
					input_string += string(s)
				}
				for ( (strings.Index(input_string, ":") + strings.Index(input_string, "#"))>0 ) {
					cmd := input_string[strings.Index(input_string, ":")+1:strings.Index(input_string, "#")]
					input_string = input_string[strings.Index(input_string, "#")+1:]
					//println("CMD:", cmd)
					if(len(cmd) == 1){cmd+=" "}
					switch cmd[0:2] {
						case "C ":
							// Initiate temperature conversion

						case "FG":
							//Go to positition set by SNYYYY
							motor.SetXTARGET(FocuserStatus.new_pos*FocuserStatus.scale)

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
							if(motor.GetVACTUAL().VACTUAL > 1 || motor.GetVACTUAL().VACTUAL < -1){
								fmt.Printf("01#\n\r") // moving
							} else {
								fmt.Printf("00#\n\r") // stale
							}

						case "GN":
							// return new target position
							fmt.Printf("%04x#\n\r", FocuserStatus.new_pos)

						case "GP":
							// return current position
							fmt.Printf("%04x#\n\r", int(motor.GetXACTUAL().XACTUAL/int32(FocuserStatus.scale)))

						case "GT":
							// return current temperature position
							temp := int((input0.Get()-500)/100)
							var sign int
							if(temp < 0){
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
	 						motor.SetRegister(tmc5130.VMAX|tmc5130.WRITE,	(25-num)*10000)

						case "SF":
							// Set fullstep mode
							FocuserStatus.scale = 1000
							//println("FS")

						case "SH":
							// Set halfstep mode
							FocuserStatus.scale = 500
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
							motor.SetXACTUAL(num*FocuserStatus.scale)
							motor.SetXTARGET(num*FocuserStatus.scale)
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
							println("NEZNAMY PRIKAZ...")

					}
				}
		}
		//println(input_string)

		// switch os := runtime.GOOS; os {
		// 	case "darwin":
		// 		fmt.Println("OS X.")
		// 	case "linux":
		// 		fmt.Println("Linux.")
		// 	default:
		// 		// freebsd, openbsd,
		// 		// plan9, windows...
		// 		fmt.Printf("%s.\n", os)
		// 	}
								//motor.GetRegister(0x01)
//								xa := motor.GetXACTUAL()
								//println(xa.String())
								//println(xa.XACTUAL)

								//println("rychlost")
								//motor.GetVACTUAL()
								//println(xb.String())
								//println(xb.VACTUAL)

								// if (motor.GetVACTUAL().VACTUAL == 0){
								// 	time.Sleep(time.Millisecond*500)
								// 	trg += 0xfff;
								// 	motor.SetXTARGET(int(trg))
								// 	//motor.SetRegister(0xAD,trg); //XACTUAL=0
								// }

								//a = motor.GetRegister(0x04)
								//println(a)

								//mstat := motor.InputStatus()
								//println(mstat.String())

      //println("TEXT>", "  ADC: ", input0.Get())

			time.Sleep(time.Millisecond * 10)
		//}
	}
}
