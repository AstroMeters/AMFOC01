#pragma once
#ifndef DATASTRUCT_H_
#define DATASTRUCT_H_

#include <Adafruit_SH110X.h>


struct Data{
   bool motor_movement;
   int motor_position;
   bool pwr_ext;
   bool pwr_usb;
   float temp;
   String text;
};

struct ToolSet{
  Adafruit_SH1106G* display;
  void* buzzer;
  Data* status;
};


#endif