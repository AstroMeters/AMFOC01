
$fn=120;
tube_d = 91;
ring_thickness = 20;

ribbon_th = 15.5;
ribbon_depth = 9;
ribbon_base = 30;

ribbon_height = 40;


module ring_a(){
    difference(){
        union(){
            scale([1, 1.05, 1]) cylinder(d=tube_d+20, h=ring_thickness);
            translate([0, -ribbon_base/2, 0]) cube([tube_d/2+10, ribbon_base, ribbon_height]);
            translate([0, -ribbon_th/2, 0]) cube([tube_d/2+10+ribbon_depth, ribbon_th, ribbon_height]);
        }
        cylinder(d=tube_d, h=ring_thickness+ribbon_height);
        
        
        for(x=[1, -1], y=[5, 15]) translate([0, x*(tube_d/2+5), y]) rotate([0, 90, 0]){
                cylinder(d = 3.3, h=100, $fn=40);
                translate([0, 0, 10]) rotate(30) cylinder(d = 6.5, h=100, $fn=6);
            }
        
        translate([1-tube_d*2, -tube_d, 0]) cube([tube_d*2, tube_d*2, ring_thickness]);
        
        translate([0, 0, 15]) rotate([0, 90, 0]){
            cylinder(d=6.5, h = 100);
            cylinder(d=11, h = tube_d/2+8);
        }

    }
}



module ring_b() rotate([0, 0, 180]) {
    difference(){
        union(){
            scale([1, 1.05, 1]) cylinder(d=tube_d+20, h=ring_thickness);
            //translate([0, -ribbon_base/2, 0]) cube([tube_d/2+10, ribbon_base, ribbon_height]);
            //translate([0, -ribbon_th/2, 0]) cube([tube_d/2+10+ribbon_depth, ribbon_th, ribbon_height]);
        }
        cylinder(d=tube_d, h=ring_thickness+ribbon_height);
        
        
        for(x=[1, -1], y=[5, 15]) translate([0, x*(tube_d/2+5), y]) rotate([0, 90, 0]){
                cylinder(d = 3.3, h=100, $fn=40);
                translate([0, 0, 10]) rotate(30) cylinder(d = 6.5, h=100, $fn=6);
            }
        
        translate([1-tube_d*2, -tube_d, 0]) cube([tube_d*2, tube_d*2, ring_thickness]);
        

    }
}


ring_a();
ring_b();