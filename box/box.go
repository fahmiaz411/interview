package box

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	colorRed = "\033[31m"
	colorYellow = "\033[33m"
	colorCyan = "\033[36m"
	colorWhite = "\033[37m"
)

type Object struct {
	Box []string
	Index int
}

func NewObject(box []string) *Object {

	obj := &Object{}

	obj.Box = box

	rangeLower := 0
    rangeUpper := 4

	rand.Seed(time.Now().Unix())
    obj.Index = rangeLower + rand.Intn(rangeUpper-rangeLower+1)

	fmt.Printf("Object first in Box : %s\n", colorRed + obj.Box[obj.Index] + colorWhite)

	return obj
}

func (obj *Object) moveObject() int {
	fmt.Println(colorYellow + "Move Object" + colorWhite)

	var newObjIndex int
	rangeLower := 1
	rangeUpper := 2

	rand.Seed(time.Now().UnixNano())
	switch rangeLower + rand.Intn(rangeUpper-rangeLower+1) {
	case 1:
		newObjIndex = -1
	case 2:
		newObjIndex = 1
	}

	obj.Index += newObjIndex

	if obj.Index < 0 {
		obj.Index = 0
	} else if obj.Index > 4 {
		obj.Index = 4
	}

	fmt.Printf("Object in Box : %s\n", colorCyan + obj.Box[obj.Index] + colorWhite)

	return obj.Index
} 

func (obj *Object) FindObject() string {
	index := -1

    for i := 0; i < 7; i++ {
		index = obj.moveObject()
	}

	if index == -1 {
		return "Not found"
	}
	return obj.Box[index]
}

func Boxes() {
	box := []string{
		"A",
		"B",
		"C",
		"D",
		"E",
	}

	object := NewObject(box)

	fmt.Printf("Object Found in Box : %s\n", colorRed + object.FindObject() + colorWhite)
}