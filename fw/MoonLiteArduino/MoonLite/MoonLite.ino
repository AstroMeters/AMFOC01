#include <Adafruit_TinyUSB.h>
#include <Wire.h>
#include <SPI.h>
#include <TMCStepper.h>
#include <Adafruit_GFX.h>
#include <Adafruit_SH110X.h>
#include "RP2040_PWM.h"



#define SCREEN_WIDTH 128
#define SCREEN_HEIGHT 64

#define OLED_RESET 11  // GPIO11 (Reset pin)
#define OLED_DC    10  // GPIO10 (Data/Command pin)
#define OLED_CS     9  // GPIO9 (Chip Select pin)
#define OLED_CLK   6
#define OLED_MOSI   7
Adafruit_SH1106G display = Adafruit_SH1106G(128, 64, &SPI, OLED_DC, OLED_RESET, OLED_CS);



#define MOTOR_CS  5
#define MOTOR_EN  3
TMC5130Stepper driver(MOTOR_CS, 1.0f);


#define BTN_BACK  12
#define BTN_SET   13
#define BTN_L     8
#define BTN_R     15

#define BTN_EXT   26

RP2040_PWM* PWM_Instance;

#define PWR_PIN_USB 21
#define PWR_PIN_EXT 1
bool pwr_state_usb;
bool pwr_state_usb_last;
bool pwr_state_ext;
bool pwr_state_ext_last;

#define BUZZER    14

#define INT_TEMP = 27;


#define LED1_T 20
#define LED1_B 22

Adafruit_USBD_CDC USBSer;


void beep(float freq, uint32_t duration){
  PWM_Instance->setPWM(BUZZER, freq, 50.0f);
  delay(duration);
  PWM_Instance->setPWM(BUZZER, freq, 0.0f);
}


void motor_init(){

  digitalWrite(MOTOR_EN, LOW);

  driver.begin();
  //driver.defaults();
  //driver.push();
  //driver.toff(0);
  //delay(100);


  driver.GCONF(0);
  driver.rms_current(600); // mA
  driver.CHOPCONF(0x000101D5);
  driver.TPOWERDOWN(10);
  driver.PWMCONF(0x000401C8);


  driver.en_pwm_mode(1);
  //driver.TCOOLTHRS(0xFFFFF); // 20bit max
  driver.microsteps(128);
  //driver.microsteps(64);
  //driver.XACTUAL(0);
  //driver.VACTUAL(0);
  //driver.toff()

  driver.RAMPMODE(0);
  driver.v1(100000);
  driver.a1(1000);
  driver.d1(1400);
  driver.AMAX(10000);
  driver.VMAX(10000000);
  driver.VSTOP(10);
  driver.XACTUAL(0);
  driver.XTARGET(0);
  //driver.push();

  //driver.XTARGET(10); // Move 100mm
  

  //driver.en_pwm_mode(true);       // Toggle stealthChop on TMC2130/2160/5130/5160
  //driver.en_spreadCycle(false);   // Toggle spreadCycle on TMC2208/2209/2224
  driver.pwm_autoscale(true);     // Needed for stealthChop
}

void setup() {

  // Vypni motor!
  pinMode(MOTOR_EN, OUTPUT);
  digitalWrite(MOTOR_EN, HIGH);

  pinMode(LED1_T, OUTPUT);
  pinMode(LED1_B, OUTPUT);
  analogWriteResolution(12);
  digitalWrite(LED1_B, HIGH);


  // CHECK powering methods
  pinMode(PWR_PIN_USB, INPUT);
  pinMode(PWR_PIN_EXT, INPUT);
  pwr_state_usb = digitalRead(PWR_PIN_USB);
  pwr_state_usb_last = digitalRead(PWR_PIN_USB);
  pwr_state_ext = digitalRead(PWR_PIN_EXT);
  pwr_state_ext_last = digitalRead(PWR_PIN_EXT);


  pinMode(BUZZER, OUTPUT);
  PWM_Instance = new RP2040_PWM(BUZZER, 400000, 50.0f);
  PWM_Instance->setPWM();

  beep(1000.0f, 500);

  Serial.begin(115200);
  Serial.setTimeout(100); // 100 ms


  
  SPI.setRX(4);
  //SPI.setCS(OLED_CS);
  SPI.setSCK(OLED_CLK);
  SPI.setTX(OLED_MOSI);

  SPI.begin();

  display.begin(0,1);
  display.clearDisplay();
  display.display();
  
  display.setTextSize(1);
  display.setTextColor(SH110X_WHITE);

  display.setCursor(10, 10);
  display.println(F("Hello,"));

  display.setTextColor(SH110X_BLACK, SH110X_WHITE);
  display.println(F("Arduino!"));
  display.display();  // Zobrazí zapsaný text


  attachInterrupt(digitalPinToInterrupt(BTN_SET), btn_set, FALLING);
  attachInterrupt(digitalPinToInterrupt(BTN_BACK), btn_back, FALLING);
  attachInterrupt(digitalPinToInterrupt(BTN_L), btn_l, FALLING);
  attachInterrupt(digitalPinToInterrupt(BTN_R), btn_r, FALLING);

  motor_init();
  driver.XTARGET(100000);

}

void update_led(){
  int duration = 3;
  analogWrite(LED1_B, int(4095*sin(millis()/1000.0*2*3.1415/duration)) );
}

int LEDstate = 0;



#define BUFFER_SIZE 128

char buffer[BUFFER_SIZE];
int bufferIndex = 0;
bool newData = false;

long FuturePos = 0;



void loop() {
  int ch;
  update_led();

  //Serial.println( driver.XACTUAL() );

  if(pwr_state_usb != digitalRead(PWR_PIN_USB)){
    pwr_state_usb = digitalRead(PWR_PIN_USB);

    if(pwr_state_usb){
      motor_init();
      beep(900.0f, 150);
      beep(1100.0f, 100);
    }else{
      beep(1100.0f, 100);
      beep(900.0f, 150);
    }
  }

  if(pwr_state_ext != digitalRead(PWR_PIN_EXT)){
    pwr_state_ext = digitalRead(PWR_PIN_EXT);

    if(pwr_state_ext){
      motor_init();
      beep(800.0f, 150);
      beep(1000.0f, 100);
    }else{
      beep(1000.0f, 100);
      beep(800.0f, 150);
    }

  }


  if (Serial.available() > 0) {
    char incomingChar = Serial.read();
    
    if (incomingChar == '#') {
      buffer[bufferIndex] = '#';
      buffer[bufferIndex+1] = '\0';
      newData = true;
      bufferIndex = 0;
    } else {
      buffer[bufferIndex] = incomingChar;
      bufferIndex++;
      
      if (bufferIndex >= BUFFER_SIZE) {
        bufferIndex = BUFFER_SIZE - 1;
      }
    }
  }

  if (newData) {
    //Serial.println(String(buffer));

    // display.clearDisplay();
    // display.setTextSize(1);
    // display.setTextColor(SH110X_WHITE);
    // display.setCursor(10, 10);
    // display.println(String(buffer));
    // display.display();
    
    newData = false;
    if (buffer[0] == ':') {
      String command = String(buffer).substring(1, 3); // Vybereme třetí znak

      // display.setTextColor(SH110X_BLACK, SH110X_WHITE);
      // display.setCursor(10, 18);
      // display.println(command);
      // display.display();

      if (command.equals("C#")) {
        // Initialize temperature conversion
      } else if (command.equals("FG")) {
        // Go to position
        driver.XTARGET(FuturePos);
      } else if (command.equals("FQ")) {
        // Immediately stop movement
        driver.XTARGET(driver.XACTUAL());
      } else if (command.equals("GC")) {
        // Return temperature coefficient
      } else if (command.equals("GD")) {
        // Return current stepping delay
      } else if (command.equals("GH")) {
        // Return if it is halfstep "FF#", other "00#"
        Serial.println("FF#");
      } else if (command.equals("GI")) {
        // Is motor moving? "00#" - stop, "01#" - in movement
        if (driver.VACTUAL() == 0) {
          Serial.println("00#");
        } else {
          Serial.println("01#");
        }
      } else if (command.equals("GN")) {
        // Return new target position
        Serial.printf("%04x#\n", driver.XTARGET());
      } else if (command.equals("GP")) {
        // Return current position

        Serial.printf("%04x#\n", driver.XACTUAL());
      } else if (command.equals("GT")) {
        // Return current temperature
        Serial.println("0000#"); // Nahraďte skutečnými daty teploty
      } else if (command.equals("GV")) {
        // Return firmware version
        Serial.println("22#");
      } else if (command.equals("SC")) {
        // Set temperature coefficient
        long val = strtol(String(buffer).substring(3, 8).c_str(), NULL, 16);
      } else if (command.equals("SD")) {
        motor_init();
        // Set stepping delay -- tohle je asi rychlost
        long val = strtol(String(buffer).substring(3, 8).c_str(), NULL, 16);
      } else if (command.equals("SN")) {
        // Set new position
        long val = strtol(String(buffer).substring(3, 8).c_str(), NULL, 16);
        FuturePos = val;

        
        display.clearDisplay();
        display.setCursor(10, 10);
        display.println(String(buffer));
        display.println(String(buffer).substring(3, 8).c_str());
        display.println(val);
        display.display();

        //driver.XTARGET(val);
      } else if (command.equals("SP")) {
        // Set current position
        long val = strtol(String(buffer).substring(3, 8).c_str(), NULL, 16);
        driver.XACTUAL(val);
        FuturePos = val;
      } else if (command.equals("+#")) {
        // Enable temperature compensation
      } else if (command.equals("-#")) {
        // Disable temperature compensation
      } else if (command.equals("PO")) {
        // Temperature calibration offset
        // Hex number, halfdegree increment
        // PO02 - 1 degree, POFB - -2.5 degree
        long val = strtol(String(buffer).substring(3, 6).c_str(), NULL, 16);
      } else {
        Serial.println("Unknown command");
      }
            
  }
  }


}




void btn_set(){
  Serial.println("Zmacknuto SET");
}
void btn_back(){
  Serial.println("Zmacknuto BACK");
}
void btn_l(){
  Serial.println("Zmacknuto <");
}
void btn_r(){
  Serial.println("Zmacknuto >");
}
