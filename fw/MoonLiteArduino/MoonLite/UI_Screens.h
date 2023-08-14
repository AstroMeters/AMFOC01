// #pragma once

// #include <Adafruit_SH110X.h>
// #include "UI_Tree.h"



// void draw_top_row(Adafruit_SH1106G *display, String content){
//   display->fillRoundRect(5, 0, 118, 11, 3, SH110X_WHITE);
//   display->setTextColor(SH110X_BLACK, SH110X_WHITE);
//   display->setTextSize(1);
//   display->setCursor(9, 2);
//   display->print( content );
//   display->setTextColor(SH110X_WHITE);

// }

// void draw_default(ToolSet* toolset){

//   toolset->display->clearDisplay();
  
//   draw_top_row(toolset->display, toolset->current_node->name);

//   toolset->display->setCursor(1, 15);
//   toolset->display->printf("<    > ");
//   // display->setCursor(65, 15);
//   // display->printf("12V: %s", status->pwr_ext?"Conn":"N/C");
//   // display->setCursor(1, 27);
//   // display->printf("Mot: %s %d", status->motor_movement ? "In move" : "Steady ", status->motor_position);
//   // display->setCursor(1, 39);
//   // display->printf("Tep: %d C", status->temp);

//   toolset->display->display();  // Zobrazí zapsaný text

// }



// void draw_init_page(Adafruit_SH1106G *display){
//   display-pclearDisplay();
  
//   display->setTextSize(2);
//   display->fillRoundRect(19, 26, 92, 23, 6, SH110X_WHITE);

//   display->setTextColor(SH110X_BLACK, SH110X_WHITE);
//   display->setCursor(25, 31);
//   display->println(F("AMFOC01"));

//   display->setTextColor(SH110X_WHITE);
//   display->setTextSize(1);
//   display->setCursor(23, 53);
//   display->println(F("AstroMeters.eu"));

//   draw_top_row(display, "Initialization ...");

//   display->display();  // Zobrazí zapsaný text
// }



// void draw_home_page(ToolSet* toolset){
//   display->clearDisplay();
  
//   draw_top_row(display, "AMFOC01 - home");

//   display->setCursor(1, 15);
//   display->printf("USB: %s", status->pwr_usb?"Conn":"N/C");
//   display->setCursor(65, 15);
//   display->printf("12V: %s", status->pwr_ext?"Conn":"N/C");
//   display->setCursor(1, 27);
//   display->printf("Mot: %s %d", status->motor_movement ? "In move" : "Steady ", status->motor_position);
//   display->setCursor(1, 39);
//   display->printf("Tep: %d C", status->temp);

//   display->display();  // Zobrazí zapsaný text
// }