#include <SPI.h>
#include <Arduino.h>

void setup()
{
    //assigns pin 25 (built in LED), with frequency of 20 KHz and a duty cycle of 0%
    //PWM_Instance = new RP2040_PWM(25, 2, 50);

    pinMode(10, OUTPUT); // set the SS pin as an output
    SPI.beginTransaction(SPISettings(100000, MSBFIRST, SPI_MODE3)) ;
    //SPI.setClockDivider(SPI);
    //SPI.setDataMode(SPI_MODE3);
    //SPI.setRX(16);
    //SPI.setCS(17);
    //SPI.setSCK(18);
    //SPI.setTX(19);
    SPI.begin(true);         // initialize the SPI library
    Serial.begin(9600);
}

int a;
int b;
int c;
int d;
int e;

void loop()
{

  a = SPI.transfer(0x21);
  b = SPI.transfer(0x00);
  c = SPI.transfer(0x00);
  d = SPI.transfer(0x00);
  e = SPI.transfer(0x00);

  Serial.println(a);
  Serial.println(b);
  Serial.println(c);
  Serial.println(d);
  Serial.println(e);

  Serial.println("-----------"); // read the value from analog pin A0 and send it to serial
  delay(100); 
  



}
