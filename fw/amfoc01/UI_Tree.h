#ifndef UI_TREE_H
#define UI_TREE_H

#include <Arduino.h>
#include <Adafruit_SH110X.h>
#include "data_struct.h"
#include <vector>

//extern Adafruit_SH1106G display;



class Node {
public:
  Node* parent;
  std::vector<Node*> child = {};
  Node* left;
  Node* right;
  bool leaf;
  String name;
  void (*drawer)(ToolSet *toolset, Node *currentNode);
  void (*left_action)() = nullptr;
  void (*right_action)() = nullptr;
  void (*left_action_up)() = nullptr;
  void (*right_action_up)() = nullptr;
  void (*set_action)() = nullptr;
  void (*set_action_up)() = nullptr;
  void (*back_action)() = nullptr;

  Node(String n, bool leaf, void (*dr)(ToolSet *toolset, Node *currentNode)) :
          parent(nullptr), left(nullptr), right(nullptr), name(n), leaf(leaf), drawer(dr) {}
  
  Node(String n, bool leaf, void (*dr)(ToolSet *toolset, Node *currentNode), void (*left_action)(), void(*right_action)(), void (*set_action)()) :
          parent(nullptr), left(nullptr), right(nullptr), name(n), leaf(leaf), drawer(dr), left_action(left_action), right_action(right_action), set_action(set_action) {}
  
  Node(String n, bool leaf, void (*dr)(ToolSet *toolset, Node *currentNode), void (*left_action)(), void (*left_action_up)(), void(*right_action)(), void (*right_action_up)(), void (*set_action)(), void (*set_action_up)()) :
          parent(nullptr), left(nullptr), right(nullptr), name(n), leaf(leaf), drawer(dr), left_action(left_action), left_action_up(left_action_up), right_action(right_action), right_action_up(right_action_up), set_action(set_action), set_action_up(set_action_up) {}
  
  Node(String n, bool leaf, void (*dr)(ToolSet *toolset, Node *currentNode), void (*left_action)(), void (*left_action_up)(), void(*right_action)(), void (*right_action_up)(), void (*set_action)(), void (*set_action_up)(), void (*set_action_back)()) :
          parent(nullptr), left(nullptr), right(nullptr), name(n), leaf(leaf), drawer(dr), left_action(left_action), left_action_up(left_action_up), right_action(right_action), right_action_up(right_action_up), set_action(set_action), set_action_up(set_action_up), back_action(set_action_back) {}
  
  Node(String n, bool leaf) :
          parent(nullptr), left(nullptr), right(nullptr), name(n), leaf(leaf), drawer(nullptr) {}
  
  void action() {
    // základní implementace
  }

  void displayContent() {
    // display.clearDisplay();
    // display.setTextSize(1);
    // display.setTextColor(SH110X_WHITE);  // Upřesnil jsem barvu pro SH1106G
    // display.setCursor(0,0);
    // display.println(name);
    // display.display();
  }
};







void draw_default(ToolSet* toolset, Node* current_node){

  toolset->display->clearDisplay();
  
  draw_top_row(toolset->display, current_node->name);

  toolset->display->setCursor(1, 15);
  toolset->display->printf("<    > ");
  // display->setCursor(65, 15);
  // display->printf("12V: %s", status->pwr_ext?"Conn":"N/C");
  // display->setCursor(1, 27);
  // display->printf("Mot: %s %d", status->motor_movement ? "In move" : "Steady ", status->motor_position);
  // display->setCursor(1, 39);
  // display->printf("Tep: %d C", status->temp);

  toolset->display->display();  // Zobrazí zapsaný text

}









class Tree {
public:
  Node *root;
  Node *current;
  Data globalData;
  ToolSet *toolset;
  
  Tree(ToolSet *ts) : root(nullptr), current(nullptr), toolset(ts) {}

  void setRoot(Node *node){
    root = node;
    current = node;
  }

  void setCurrent(Node *node){
    current = node;
  }

  void insertNode(Node *parent, Node *new_node){
    parent->child.push_back(new_node);
    int count = parent->child.size();
    new_node->parent = parent;
    new_node->right = nullptr;

    if(count>1){
      new_node->left = parent->child[count-2];
      new_node->left->right = new_node;
    }
    //if(count>1){
    //  parent->child[count-2]->right = new_node;
    //}

  }

  void navigateLeft() {
    if (current->left_action) {
      current->left_action();
    }
    else if (current->left) {
      Serial.println("NAVIGACE DO STRANY");
      current = current->left;
      current->displayContent();
    }
  }

  void navigateLeftUp() {
    if (current->left_action_up) {
      current->left_action_up();
    }
  }

  void navigateRight() {
    if (current->right_action) {
      current->right_action();
    }
    else if (current->right) {
      Serial.println("NAVIGACE DO STRANY");
      current = current->right;
      current->displayContent();
    }
  }


  void navigateRightUp() {
    if (current->right_action_up) {
      current->right_action_up();
    }
  }

  void navigateBack() {
    if (current->back_action) {
      current->back_action();
    }
    if (current->parent) {
      current = current->parent;
      current->displayContent();
    }
  }
  void navigateBackUp() {
  }

  void navigateSet() {
    if(current->child.size()){
      current = current->child[0];
    }else{
      current->set_action();
    }
  }

  void navigateSetUp() {    
  }

  void Run(){
    if(current->drawer){
      current->drawer(toolset, current);
    } else {
      draw_default(toolset, current);
    }

   // Serial.println(current->parent->child.size());
   //Serial.println(current->name);

  }
};

extern Tree tree;

void handleButtons();

#endif
