
$fn=60;

lift_diameter = 76.5;
height = 22;

around_thickness = 10;

tube_offset = [0, -lift_diameter/2+2, 0];

axis_shifh = 31;

difference(){
    union(){
        translate(tube_offset) cylinder(d=lift_diameter+2*around_thickness, h=height-0.1, $fn=240);
        translate(tube_offset) translate([-65/2-around_thickness, 0, 0]) cube([65+2*around_thickness, lift_diameter/2+1.6+around_thickness, height], $fn=240);
        
        
        translate([30, -6.4, 0]+[0, 0, 0]) cube([120-30, 20, height-2]);
        
    }
    translate(tube_offset) cylinder(d=lift_diameter, h=height);
    translate(tube_offset) translate([-65/2, 0, 7]) cube([65, lift_diameter/2+1.6, height]);
    
    translate([-100-30, 0.5, 0]) cube([100, 3, height]);
    translate([-2.5-30, -26.4, 0]) cube([3, 30, height]);
    
    
    
    for(z=[height/4, height/4*3]) translate([-38, 20, z]) rotate([90, 0, 0]){
            cylinder(d=3.6, h=30, $fn=30);
            cylinder(d=7, h=9, $fn=30);
            
            
            translate([0, 0, 9+12]) hull(){
                    cylinder(d=6.5, h=3, $fn=6);
                    translate([-20, 0, 0]) cylinder(d=6.5, h=3, $fn=6);
                }
          
          }
}


// motor flange



    hull(){
        translate([115, -6.4, 0]) cube([5, 20, 10]);
        translate([115, -42/2, 40-42/2]) cube([5, 42, 1]);
    }
difference(){
    
    translate([115, -42/2, 40-42/2])
    difference(){
        cube([5, 42, 42]);
        translate([0, 42/2, 42/2]) rotate([0, 90, 0]){
            cylinder(d=22.5, h=10);
            for(x=[-0.5, 0.5], y=[0.5, -0.5]) translate([x*31, y*31, -0.1]) cylinder(d=3.3, h=10, $fn=60);
        }
    }


}