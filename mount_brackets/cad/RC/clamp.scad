
$fn=60;

diameter = 82.3;
clamp_thickness = 10;
clamp_height = 17;


module bracket(){

    translate([diameter/2+clamp_thickness, 0, 0]) difference(){
        union(){
            cylinder(d = diameter+clamp_thickness*2, h = clamp_height);
            translate([-diameter/2-clamp_thickness,-diameter/2-clamp_thickness,0]) cube([diameter/2+clamp_thickness, diameter/2+clamp_thickness, clamp_height]);
        }
        
        translate([0, 0, -0.1]) cylinder(d = diameter, h = clamp_height+0.2);

        
        translate([0, 0, 3]) for(r = [0, 180, 180-60, 180+60]) rotate([0, 0, r]) rotate([0, -90, 0]) hull(){
            cylinder(d=4, h = 100);
            translate([-5, 0, 0]) cylinder(d=4, h = 100);
            }
        
        translate([-diameter/2, -diameter/2-clamp_thickness, -0.1]) cube([2, diameter/2+clamp_thickness, clamp_height+0.2]);
        
        for(z = [20/4-1,-20/4+1])
            translate([-diameter/2, -diameter/2-clamp_thickness+15, z+clamp_height/2]){
                rotate([0, 90, 0]) cylinder(d = 3.5, h = 30, center=true, $fn=30);
                translate([-10, 0, 0]) rotate([0, 90, 0]) cylinder(d = 7, h = 4, $fn=30);
                
                translate([5, -5.6/2, -10]) cube([3, 5.6, 20]);
            
            }
                
        
        
    }
    
    translate([0, -110, 0]) cube([clamp_thickness, 105-40, clamp_height]);
    translate([0, -110, clamp_height/2]) scale([1, 1, 1.2]) rotate([0, 45, 0]) cube([10, 105-47, 10]);
    
    
    
    // motormount
    translate([0, -110, 0]){
        
        cube([42, 6, 20]);
        translate([0, 0, 17]) difference(){
            cube([42, 6, 42]);
            translate([42/2, 0, 42/2]) rotate([-90, 0, 0]) cylinder(d = 9.8, h= 10);
            translate([42/2, 0, 42/2]) rotate([-90, 0, 0]) cylinder(d = 22.5, h= 2);
            for(x=[-31/2, 31/2], y=[-31/2, 31/2])translate([x, 0, y])
                translate([42/2, 0, 42/2]) rotate([-90, 0, 0]) cylinder(d = 3.3, h= 10);
        }
        
    }
}

bracket();