$fn=80;

dh = 27;
osad = 5.25;

difference(){
union(){
cylinder(d = dh+3, h = 10);
translate([0, 0, 10]) cylinder(d1 = dh+3, d2 = dh, h=3);
translate([0, 0, 10+3]) cylinder(d1 = 9+3, d2 = 9, h=3);
cylinder(d = 9, h = 30);
}

    cylinder(d=dh, h = 6);
    translate([0, 0, 6.2]) difference(){
        cylinder(d = osad,  h = 50, $fn=60);
        translate([-osad/2, osad/2-0.5, 0]) cube([5, 5, 50]);
    }
}