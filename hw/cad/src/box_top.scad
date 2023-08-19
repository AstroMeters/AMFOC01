// Top box

$fn = 60;
pcb_x = 71;
pcb_y = 61;

wall_th = 2;

rantl_th = 0.8;

oled_x = 32;
oled_y = 17;
oled_y_dis = -2.5;   // displacement
oled_x_dis = -0;
oled_out_x = 36.2;
oled_out_y = 24.3;

//OLED na PCB (stred oled vs stred PCB) 
oled_pcb_x = 71/2 - (17.4+oled_out_x/2);
oled_pcb_y = 71/2 - (20.7+oled_out_y/2);
oled_pcb_z = 4.4; // Vyska oled nad PCB

pcb_max_z = oled_pcb_z;

oled_border_th = 0.8;

button_pcb_z = 2.45; // vyska tlacitka nad PCB
button_pattern = 8;

LED_POS = [[50-0.25, pcb_y-3.5, 0], [56-0.25, pcb_y-3.5, 0]];

module box(){

difference(){
    translate([-wall_th, -wall_th, -wall_th-oled_pcb_z]) cube([pcb_x+2*wall_th, pcb_y+2*wall_th, wall_th+oled_pcb_z+1]);
    
    // ODebrani hlavniho objemu 
    translate([0, 0, -oled_pcb_z]) cube([pcb_x, pcb_y, 20]);
    
    
    
    translate([pcb_x/2 + oled_pcb_x, pcb_y/2+oled_pcb_y, -oled_pcb_z-0.9]) hull() {
        // cube([oled_out_x, oled_out_y, 10], center=true);
        translate([oled_x_dis, oled_y_dis, 0]) cube([oled_x, oled_y, 2], center=true);
        translate([oled_x_dis, oled_y_dis, -3]) cube([oled_x+4, oled_y+4, 1], center=true);

        //translate([0, 0, -3]) cube([oled_out_x+4, oled_out_y+4, 1], center=true);
    
    }
    
    
   translate([pcb_x-13, pcb_y-16.2-8, -oled_pcb_z-wall_th]) for(x=[0, 1], y=[0, 1]) translate([x*button_pattern, y*button_pattern, 0]){
        cylinder(d = 3, h = 10);
        translate([0, 0, -0.1]) cylinder(d1=3.6, d2=3, h=0.6, $fn=60);
        }
        
        
        
   
// Pruzor pro LEDky   
   for(p = LED_POS) translate([0, 0, -oled_pcb_z-wall_th]) translate(p){
        cylinder(d1 = 3, d2=2.5, h = 10, center=true, $fn=60);
        }
        
        
        
     // Srazeni hran
     for(y=[-wall_th, pcb_y+wall_th]) translate([0, y, -pcb_max_z-wall_th]) rotate([45, 0, 0]) translate([-wall_th, -1, -1]) cube([pcb_x+2*wall_th, 2, 2]);
    
     for(x=[-wall_th, pcb_x+wall_th]) translate([x, 0, -pcb_max_z-wall_th]) rotate([0, 45, 0]) translate([-1, -wall_th, -1]) cube([2,pcb_y+2*wall_th, 2]);
     
     for(x = [1, 0], y=[1, 0]) translate([(pcb_x+wall_th*2)*x-wall_th, (pcb_y+2*wall_th)*y-wall_th, 0]) rotate([0, 0, 45]) cube([2, 2, 50], center=true);
    
}

for(x=[-6/2, 6/2], y=[2.5,-2.5]) translate([pcb_x/2+10.16*x, pcb_y/2+10.16*y, -oled_pcb_z-1]) difference(){
    union(){
        cylinder(d = 8, h = 1+oled_pcb_z);
        translate([0, 0, 1+oled_pcb_z]) cylinder(d = 3, h = 0.6);
    }
    cylinder(d = 2.3, h = 10);
    
    }
    
for(pos=[[10, pcb_y/2, 0], [pcb_x-5, pcb_y/2-5, 0], [pcb_x/2, 5, 0]]) translate(pos+[0, 0, -oled_pcb_z]) difference(){
    
    cylinder(d = 4, h = oled_pcb_z);
    
    }
    
 difference(){
 
 
   translate([pcb_x-13, pcb_y-16.2-8, -oled_pcb_z]) for(x=[0, 1], y=[0, 1]) translate([x*button_pattern, y*button_pattern, 0]){
        cylinder(d = 5, h = button_pcb_z-1 );
        }
        
        
   translate([pcb_x-13, pcb_y-16.2-8, -oled_pcb_z-wall_th]) for(x=[0, 1], y=[0, 1]) translate([x*button_pattern, y*button_pattern, 0]){
        cylinder(d = 3, h = 10, center=true);
        }
        
 }
 
 
 
 
    
    translate([pcb_x/2 + oled_pcb_x, pcb_y/2+oled_pcb_y, -oled_pcb_z-0.9]) difference() {
        translate([0, 0, oled_pcb_z/2]) cube([oled_out_x+4*oled_border_th, oled_out_y+4*oled_border_th, oled_pcb_z-1], center=true);
        
        translate([0, 0, 5]) cube([oled_out_x, oled_out_y, 10], center=true);
        translate([0, 0, 2]) hull(){
            translate([0, 0, 5]) cube([oled_out_x, oled_out_y, 10], center=true);
            translate([0, 0, 5+oled_pcb_z-2]) cube([oled_out_x+3*oled_border_th, oled_out_y+3*oled_border_th, 10], center=true);
        }
    
    }
    
 
 
 
 for(p=LED_POS) translate(p+[0, 0,  -oled_pcb_z]) difference(){
 
  
        cylinder(d = 5, h = oled_pcb_z);
        cylinder(d = 3, h = oled_pcb_z);
  
}
}

box();