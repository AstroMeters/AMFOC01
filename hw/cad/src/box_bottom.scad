// Top bottom

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

pcb_max_z = 13.5;

oled_border_th = 0.8;

button_pcb_z = 2.45; // vyska tlacitka nad PCB
button_pattern = 8;

LED_POS = [[50, pcb_y-3.5, 0], [56, pcb_y-3.5, 0]];

//btn_boot_pos = [pcb_x/2+10.16*-3+9.82, pcb_y/2+10.16*-2.5+27.751, -pcb_max_z-wall_th];
btn_boot_pos = [pcb_x/2+10.16*3-9.82, pcb_y/2+10.16*-2.5+27.751-1.5, -pcb_max_z-wall_th];


fastscrew_head_d = 6;


difference(){
    translate([-wall_th, -wall_th, -wall_th-pcb_max_z]) cube([pcb_x+2*wall_th, pcb_y+2*wall_th, wall_th+pcb_max_z+0.6]);
    
    // ODebrani hlavniho objemu 
    difference() { union() {
            translate([0, 0, 0]) cube([pcb_x, pcb_y, pcb_max_z]);
            translate([+1, +1, -pcb_max_z]) cube([pcb_x-2, pcb_y-2, 20]);
            }
     }
     
     // Srazeni hran
     for(y=[-wall_th, pcb_y+wall_th]) translate([0, y, -pcb_max_z-wall_th]) rotate([45, 0, 0]) translate([-wall_th, -1, -1]) cube([pcb_x+2*wall_th, 2, 2]);
     for(x=[-wall_th, pcb_x+wall_th]) translate([x, 0, -pcb_max_z-wall_th]) rotate([0, 45, 0]) translate([-1, -wall_th, -1]) cube([2,pcb_y+2*wall_th, 2]);
     for(x = [1, 0], y=[1, 0]) translate([(pcb_x+wall_th*2)*x-wall_th, (pcb_y+2*wall_th)*y-wall_th, 0]) rotate([0, 0, 45]) cube([2, 2, 50], center=true);
    
    
    for(x=[-6/2, 6/2], y=[2.5,-2.5]) translate([pcb_x/2+10.16*x, pcb_y/2+10.16*y, -pcb_max_z-wall_th-0.1]){
            cylinder(d = 2.3, h = 20);
            cylinder(d=fastscrew_head_d, h = wall_th+pcb_max_z-3);
            translate([0, 0, 0]) cylinder(d1=fastscrew_head_d+1, d2=fastscrew_head_d, h = 0.5);
        }
        
        
        
        // Napajeci konektor 
        translate([-2.5, 15.5, -11]) cube([5, 9, 11]);
        // MiniDIN
        translate([11.2-0.2, -2.5, -13.5+0.5]) cube([14+.4, 5, 13.5+1]);
        // RJ12
        translate([27.6, -2.5, -12]) cube([13.5+0.4, 5, 12+1]);
        // USB
        //#translate([51.5, -2.5, -4]) cube([10, 5, 4]);
        translate([51.5+1+0.1, -2.5, -4+1+0.5]) minkowski() {
            cube([10-2, 5, 4-2+2]);
            rotate([90, 0, 0]) cylinder(d=2, h=1);
            }
        
        
        // Thermistor
        translate([pcb_x-2.5, 14.4-3.5, -8]) cube([5, 6, 5]);
        translate([pcb_x-5, 14.4-3.5, -10]) cube([5, 6, 10]);
        
        // Okenko pro boot tlacitko
        translate(btn_boot_pos) cylinder(d=2.3, h=20);
        translate(btn_boot_pos) cylinder(d1=2.8, d2=2.3, h=0.5);
    
}
    
for(x=[-6/2, 6/2], y=[2.5,-2.5]) translate([pcb_x/2+10.16*x, pcb_y/2+10.16*y, -pcb_max_z-1]) difference(){
    union(){
        cylinder(d = 8, h = 1+pcb_max_z);
        translate([0.5*x, 0.5*y, 0]) translate([-4, -4, 0]) cube([8, 8, 1+pcb_max_z-1]);
        //translate([0, 0, 1+pcb_max_z]) cylinder(d = 2.9, h = 0.6);
    }
    union(){
    translate([0, 0, wall_th+pcb_max_z-3.5]) cylinder(d = 2.7, h = 20);
    //cylinder(d=fastscrew_head_d, h = wall_th+pcb_max_z-3);
    
translate([0, 0, wall_th+pcb_max_z-3.5]) difference(){
    translate([0, 0, -wall_th-pcb_max_z+3]) cylinder(d=fastscrew_head_d, h = wall_th+pcb_max_z-3);
    union() for(r=[[0, 0], [0.2, -60], [0.4, 60]]) rotate([0, 0, r[1]]) {
            translate([2.7/2, -5, -r[0]] ) cube([10, 10, 0.61]);
            translate([-10-2.7/2, -5, -r[0]] ) cube([10, 10, 0.61]);
        }
}
    }
    
    }
    
    
    translate(btn_boot_pos) difference() {
        translate([0, 0, 1]) cylinder(d=2.3+1.8, h=1+pcb_max_z-2.8-1);
        cylinder(d=2.3, h=20);
    }