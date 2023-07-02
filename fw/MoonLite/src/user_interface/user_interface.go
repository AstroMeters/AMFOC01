package user_interface

import (
	//"errors"
	"strconv"
	//"strings"
	//"machine"
	//"time"
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
	Cursor    int
	MaxCursor int
	Name      string
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

func (t *Tree) Insert(parent *Node, data *DataObject) {
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

	data.Cursor = 1
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
	default:
		// Specifická akce pro hodnotu kurzoru
		if t.CurrentNode != nil {
			t.CurrentNode.Data.Action(t.Cursor - 2)
		}
	}
}

func (t *Tree) SetDefaultPosition(node *Node) {
	t.CurrentNode = node
	t.Cursor = 1
}

func (t *Tree) findNodeByCursor(nodes []*Node, cursor int) *Node {
	for _, node := range nodes {
		if node.Data.Cursor == cursor {
			return node
		}
		if foundNode := t.findNodeByCursor(node.Children, cursor); foundNode != nil {
			return foundNode
		}
	}
	return nil
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

type Device struct {
	display sh1106.Device
	state   int
	tree    Tree

	temperature_in  float64
	temperature_out float64
	motor_position  int
	ext_pwr         bool
	//last_update time.time
}

func New(display sh1106.Device) *Device {
	d := &Device{
		display: display,
		state:   Init,
	}

	root_1 := &DataObject{
		Cursor:    1,
		MaxCursor: 2,
		Name:      "Init",
		Action:    nil,
	}

	d.tree.Insert(nil, root_1)
	d.tree.Cursor = 1

	main_1 := &DataObject{
		Cursor:    1,
		MaxCursor: 2,
		Name:      "Home page",
		Action: func(cursor int) {
		},
		Draw: func() {
			d.DrawHomeScreen()
		},
	}

	main_2 := &DataObject{
		Cursor:    1,
		MaxCursor: 3,
		Name:      "Ocular page",
		Action:    nil,
	}

	main_3 := &DataObject{
		Cursor:    1,
		MaxCursor: 3,
		Name:      "Nastavení page",
		Action:    nil,
	}

	d.tree.Insert(d.tree.Roots[0], main_1)
	d.tree.Insert(d.tree.Roots[0], main_2)
	d.tree.Insert(d.tree.Roots[0], main_3)

	d.tree.SetDefaultPosition(d.tree.Roots[0].Children[0])

	// }
	// d.tree = append(d.tree,

	return d
}

func (d *Device) SetData(temp_in float64, temp_out float64, pos int, ext_pwr bool) {
	d.temperature_in = temp_in
	d.temperature_out = temp_out
	d.motor_position = pos
	d.ext_pwr = ext_pwr
}

func (d *Device) SetScreen(page int) {
	d.state = page
	d.Action()
}

func (d *Device) Btn_set() {
	// println(d.tree.GetNodeInfo())
	d.tree.ProcessCursor()
	// println(d.tree.GetNodeInfo())

	if d.tree.Cursor < 2 {
		d.DrawMenuInfo()
	} else {
		d.tree.CurrentNode.Data.Draw()
	}

}

func (d *Device) Btn_up() {
	// println(d.tree.GetNodeInfo())
	d.tree.IncreaseCursor()
	// //d.tree.ProcessCursor()
	// println(d.tree.GetNodeInfo())

	if d.tree.Cursor < 2 {
		d.DrawMenuInfo()
	}
}

func (d *Device) Btn_down() {
	// println(d.tree.GetNodeInfo())
	d.tree.DecreaseCursor()
	// //d.tree.ProcessCursor()
	// println(d.tree.GetNodeInfo())

	if d.tree.Cursor < 2 {
		d.DrawMenuInfo()
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
	tinyfont.WriteLine(&d.display, &freesans.Regular9pt7b, 0, 13, "AMFOC01", color.RGBA{100, 100, 100, 100})
	tinyfont.WriteLine(&d.display, &freesans.Regular9pt7b, 0, 30, "Temp: "+strconv.Itoa(int(d.temperature_in))+" C", color.RGBA{255, 255, 255, 255})
	tinyfont.WriteLine(&d.display, &freesans.Regular9pt7b, 0, 47, "pos: "+strconv.Itoa(d.motor_position)+"", color.RGBA{255, 255, 255, 255})
	tinyfont.WriteLine(&d.display, &freesans.Regular9pt7b, 0, 63, "12V: "+BoolToString(d.ext_pwr)+"", color.RGBA{255, 255, 255, 255})

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
