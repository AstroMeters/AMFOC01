#ifndef UI_TREE_H
#define UI_TREE_H

#include <Arduino.h>
#include <Adafruit_SH110X.h>
#include "data_struct.h"
#include <vector>

//extern Adafruit_SH1106G display;


void draw_default(ToolSet* toolset){

}


class Node {
public:
  Node* parent;
  std::vector<Node*> child = {};
  Node* left;
  Node* right;
  bool leaf;
  String name;
  void (*drawer)();

  Node(String n, bool leaf, void (*dr)()) : parent(nullptr), left(nullptr), right(nullptr), name(n), leaf(leaf), drawer(dr) {}
  Node(String n, bool leaf) : parent(nullptr), left(nullptr), right(nullptr), name(n), leaf(leaf), drawer(nullptr) {}
  
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
    if (current->left) {
      Serial.println("NAVIGACE DO STRANY");
      current = current->left;
      current->displayContent();
    }
  }

  void navigateRight() {
    if (current->right) {
      Serial.println("NAVIGACE DO STRANY");
      current = current->right;
      current->displayContent();
    }
  }

  void navigateBack() {
    if (current->parent) {
      current = current->parent;
      current->displayContent();
    }
  }

  void navigateSet() {
    if(current->child.size()){
      current = current->child[0];
    }else{
      current->action();
    }
  }

  void Run(){
    if(current->drawer){
      current->drawer();
    }

   // Serial.println(current->parent->child.size());
   //Serial.println(current->name);

  }
};

extern Tree tree;

void handleButtons();

#endif
