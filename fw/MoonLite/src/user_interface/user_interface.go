package user_interface

import (
	//"errors"
	"strconv"
	//"strings"
	//"machine"
	//"time"
	//"reflect"
	"fmt"
	"image/color"
	"tinygo.org/x/drivers/sh1106"
	"tinygo.org/x/tinyfont"
	"tinygo.org/x/tinyfont/freesans"
)

////////
///////
///////

type DataObject struct {
	MaxCursor int
	Name      string
	Type      int
	Action    func(cursor int)
	Draw      func()
	// Další atributy datové struktury
}

type Node struct {
	Data     *DataObject
	Parent   *Node
	Children []*Node
	Previous *Node
	Next     *Node
}

type Tree struct {
	Roots       []*Node
	CurrentNode *Node
	Cursor      int
}

var Cursor int

func (t *Tree) Insert(parent *Node, data *DataObject) *Node {
	newNode := &Node{Data: data, Parent: parent}

	if parent == nil {
		t.Roots = append(t.Roots, newNode)
	} else {
		parent.Children = append(parent.Children, newNode)

		// Nastavení odkazu na předchozí prvek
		if len(parent.Children) > 1 {
			newNode.Previous = parent.Children[len(parent.Children)-2]
			parent.Children[len(parent.Children)-2].Next = newNode
		}
	}

	// Nastavení odkazu Next pro poslední prvek ve stejné úrovni
	if parent != nil && newNode.Next == nil {
		newNode.Next = parent.Children[0]
		parent.Children[0].Previous = newNode
	}

	return newNode
}

func (t *Tree) IncreaseCursor() {
	if t.Cursor < t.CurrentNode.Data.MaxCursor {
		t.Cursor++
	} else {
		t.Cursor = 0
	}
}

func (t *Tree) DecreaseCursor() {
	if t.Cursor > 0 {
		t.Cursor--
	} else {
		t.Cursor = t.CurrentNode.Data.MaxCursor
	}
}

func (t *Tree) ProcessCursor() {
	if t.CurrentNode == nil {
		return
	}

	switch t.Cursor {
	case 0:
		// Vrátit se do nadřazené úrovně
		if t.CurrentNode.Parent != nil {
			t.CurrentNode = t.CurrentNode.Parent
			t.Cursor = 1
		}
	case 1:
		// Další prvek v řadě
		if t.CurrentNode.Next != nil {
			t.CurrentNode = t.CurrentNode.Next
			t.Cursor = 1
		}
	case 2:
		// Vyber pod prvek
		if t.CurrentNode.Children != nil {
			t.CurrentNode = t.CurrentNode.Children[0]
			t.Cursor = 1
		}
	default:
		// Specifická akce pro hodnotu kurzoru

		// if t.CurrentNode.Data.Draw != nil {
		// 	t.CurrentNode.Data.Draw()
		// }
	}
}

func (t *Tree) SetDefaultPosition(node *Node) {
	t.CurrentNode = node
	t.Cursor = 1
}

func (t *Tree) GetNodeInfo() string {
	if t.CurrentNode == nil {
		return "Aktuální pozice: Není nastavena"
	}

	info := fmt.Sprintf("Kurzor: %d, Objekt: %s", t.Cursor, t.CurrentNode.Data.Name)
	return info
}

////////
//
/////////
//
///////////

func BoolToInt(b bool) int {
	if b {
		return 1
	}
	return 0
}

func BoolToString(b bool) string {
	if b {
		return "1"
	}
	return "0"
}

const (
	Init int = iota
	Home
	Auto
	Ocular
	Manual
	SetupTemReg
	SetupRange
)

const (
	NODE_TREE   = iota // Rozcestnik
	NODE_ACTION        // Akce vykonavana strankou
)

type Device struct {
	display sh1106.Device
	state   int
	tree    Tree

	currentPage int

	temperature_in  float64
	temperature_out float64
	motor_position  int
	ext_pwr         bool
	usb_pwr         bool
	//last_update time.time
}

func New(display sh1106.Device) *Device {
	d := &Device{
		display:        display,
		state:          Init,
		currentPage:    Home,
		motor_position: 0,
	}

	root_1 := &DataObject{
		MaxCursor: 1,
		Name:      "Init",
		Action:    nil,
		Type:      NODE_TREE,
	}

	d.tree.Insert(nil, root_1)
	d.tree.Cursor = 1

	main_1 := &DataObject{
		MaxCursor: 3,
		Name:      "Home page",
		Type:      NODE_TREE,
	}

	main_2 := &DataObject{
		MaxCursor: 3,
		Name:      "Manual ctrl",
		Type:      NODE_TREE,
	}

	main_3 := &DataObject{
		MaxCursor: 3,
		Name:      "Ocular page",
		Action:    nil,
		Type:      NODE_TREE,
	}

	main_4 := &DataObject{
		MaxCursor: 3,
		Name:      "Nastavení page",
		Action:    nil,
		Type:      NODE_TREE,
	}

	node_home := d.tree.Insert(d.tree.Roots[0], main_1)
	d.tree.Insert(d.tree.Roots[0], main_2)
	d.tree.Insert(d.tree.Roots[0], main_3)
	d.tree.Insert(d.tree.Roots[0], main_4)

	home_view := &DataObject{
		MaxCursor: 3,
		Name:      "Home view",
		Draw:      d.DrawHomeScreen,
		Type:      NODE_ACTION,
	}

	home_status := &DataObject{
		MaxCursor: 3,
		Name:      "Motor status",
		Draw:      d.DrawHomeScreen,
		Type:      NODE_ACTION,
	}

	d.tree.Insert(node_home, home_view)
	d.tree.Insert(node_home, home_status)

	ocular_list := &DataObject{
		MaxCursor: 3,
		Name:      "Ocular list",
		Type:      NODE_TREE,
		//Draw:      d.DrawHomeScreen,
	}

	ocular_select := &DataObject{
		MaxCursor: 3,
		Name:      "Select Ocular",
		Type:      NODE_TREE,
		Draw:      d.PrepareScreen,
	}

	d.tree.Insert(d.tree.Roots[0].Children[2], ocular_list)
	d.tree.Insert(d.tree.Roots[0].Children[2], ocular_select)

	d.tree.SetDefaultPosition(d.tree.Roots[0].Children[0])

	// }
	// d.tree = append(d.tree,

	return d
}

func (d *Device) SetData(temp_in float64, temp_out float64, pos int, ext_pwr bool, usb_pwr bool) {
	d.temperature_in = temp_in
	d.temperature_out = temp_out
	d.motor_position = pos
	d.ext_pwr = ext_pwr
	d.usb_pwr = usb_pwr
}

func (d *Device) GetDevice() *Device {
	return d
}

func (d *Device) SetScreen(page int) {
	d.state = page
	d.Action()
}

func (d *Device) Btn_set() {
	//println("Nastavit")
	//println("SET")

	if len(d.tree.CurrentNode.Children) > 0 {
		d.tree.CurrentNode = d.tree.CurrentNode.Children[0]
	} else {
		//println("Není následovník")
	}
	//d.DrawMenuInfo()
	d.Process()

}

func (d *Device) Btn_back() {
	//println("BACK")
	if d.tree.CurrentNode.Parent != nil {
		d.tree.CurrentNode = d.tree.CurrentNode.Parent
	}

	d.Process()
}

func (d *Device) Btn_up() {

	d.tree.CurrentNode = d.tree.CurrentNode.Next
	//d.DrawMenuInfo()
	d.Process()

}

func (d *Device) Btn_down() {

	d.tree.CurrentNode = d.tree.CurrentNode.Previous
	//d.DrawMenuInfo()
	d.Process()

}

func (d *Device) Process() {
	if d.tree.CurrentNode.Data.Draw == nil {
		//println("Drawer není")
		d.DrawMenuInfo()
	} else {
		//println("DRAW>")
		d.tree.CurrentNode.Data.Draw()
		//d.DrawMenuInfo()
	}
}

func (d *Device) Configure() {
	d.Clear()
	d.Display()
	d.Action()
}

func (d *Device) Display() {
	d.display.Display()
}

func (d *Device) Clear() {
	d.display.ClearBuffer()
}

func (d *Device) DrawMenuInfo() {
	d.Clear()
	tinyfont.WriteLine(&d.display, &freesans.Regular9pt7b, 0, 13, strconv.Itoa(d.tree.Cursor), color.RGBA{100, 100, 100, 100})
	tinyfont.WriteLine(&d.display, &freesans.Regular9pt7b, 0, 43, d.tree.CurrentNode.Data.Name, color.RGBA{100, 100, 100, 100})
	//tinyfont.WriteLine(&d.display, &freesans.Regular9pt7b, 0, 50, strconv.Itoa(len(d.tree.CurrentNode.Children)), color.RGBA{100, 100, 100, 100})
	d.Display()
}

func (d *Device) DrawSplashScreen() {
	d.Clear()
	tinyfont.WriteLine(&d.display, &freesans.Regular9pt7b, 0, 13, "AMFOC01", color.RGBA{100, 100, 100, 100})
	tinyfont.WriteLine(&d.display, &freesans.Regular9pt7b, 0, 30, "astrometers.eu", color.RGBA{100, 100, 100, 100})
	//tinyfont.WriteLine(&d.display, &tinyfont.Org01, 0, 56, "FW: 00000", color.RGBA{100, 100, 100, 100})
	tinyfont.WriteLine(&d.display, &tinyfont.Org01, 0, 63, "HW: 00000, FW: 00000", color.RGBA{100, 100, 100, 100})

	d.Display()
}

func (d *Device) DrawHomeScreen() {
	d.Clear()
	dev := d.GetDevice()
	//println(dev.motor_position)
	tinyfont.WriteLine(&d.display, &freesans.Regular9pt7b, 0, 13, "AMFOC01", color.RGBA{100, 100, 100, 100})
	//tinyfont.WriteLine(&d.display, &freesans.Regular9pt7b, 0, 30, "Temp: "+strconv.Itoa(int(dev.temperature_in))+" C", color.RGBA{255, 255, 255, 255})
	//str_motor := reflect.ValueOf(dev.motor_position).String()
	//println(str_motor)
	//tinyfont.WriteLine(&d.display, &freesans.Regular9pt7b, 0, 47, "str_motor", color.RGBA{255, 255, 255, 255})
	tinyfont.WriteLine(&d.display, &freesans.Regular9pt7b, 0, 63, BoolToString(dev.ext_pwr), color.RGBA{255, 255, 255, 255})
	tinyfont.WriteLine(&d.display, &freesans.Regular9pt7b, 0, 47, BoolToString(dev.ext_pwr), color.RGBA{255, 255, 255, 255})

	//println(&d.motor_position)

	d.Display()
}

func (d *Device) PrepareScreen() {
	d.Clear()
	tinyfont.WriteLine(&d.display, &tinyfont.Org01, 0, 7, "1/5, USB, PWR, usb, pwr", color.RGBA{100, 100, 100, 255})

	d.Display()
}

func (d *Device) Action() {
	switch d.state {
	case Init:
		d.DrawSplashScreen()
	case Home:
		d.DrawHomeScreen()
	}
}
