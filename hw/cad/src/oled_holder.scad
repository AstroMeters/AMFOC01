// LCD spacer


oled_x = 35.1;
oled_y = 23.2;
oled_z = 1.5;
oled_cable_width = 14;

wall = 0.4;
space = 3; // Mezera pod displayem


difference(){
translate([-oled_x/2-wall, -oled_y/2-wall, 0]) cube([oled_x+2*wall, oled_y+2*wall, space+oled_z]);

translate([-oled_x/2, -oled_y/2, space]) cube([oled_x, oled_y, oled_z+0.1]);

translate([-oled_cable_width/2, -oled_y/2-10, space-space/2]) cube([oled_cable_width, oled_y-10, 10]);

}