#pragma once
#ifndef DATASTRUCT_H_
#define DATASTRUCT_H_

#include <Adafruit_SH110X.h>
#include <vector>


struct Data{
   bool motor_movement;
   int motor_position;
   int motor_manual_movement_speed_step = 5;
   int motor_manual_movement_speed = 5000;
   bool pwr_ext;
   bool pwr_usb;
   float temp;
   String text;
};


struct OcularInfo{
  int offset;
  String Name;

  OcularInfo(int o, const String& n) : offset(o), Name(n) {}
  };

struct OcularsData{
   //OcularInfo *actual_ocular {nullptr};
   //OcularInfo *next_ocular {nullptr};
   int actual_ocular = 0;
   int next_ocular = 0;
   std::vector<OcularInfo*> ocular_list;
};

struct ToolSet{
  Adafruit_SH1106G* display;
  void* buzzer;
  Data* status;
  OcularsData* oculars;
};



void draw_top_row(Adafruit_SH1106G *display, String content){
  display->fillRoundRect(5, 0, 118, 11, 3, SH110X_WHITE);
  display->setTextColor(SH110X_BLACK, SH110X_WHITE);
  display->setTextSize(1);
  display->setCursor(9, 2);
  display->print( content );
  display->setTextColor(SH110X_WHITE);

}



#endif