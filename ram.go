package main

type Ram2KBSlot struct {
	b [2000]byte
}

type Ram2MBSlot struct {
	b [1999]Ram2KBSlot
	d uint8 // 0x0 => used //\\ 0x1 => free
}

type Ram2GBSlot struct {
	b [1000]Ram2MBSlot
}

type Slot2KB_Pointer *Ram2KBSlot
type Slot2MB_Pointer *Ram2MBSlot
type Slot2GB_Pointer *Ram2GBSlot

var memp = []Slot2GB_Pointer{}

var test2KB = make([]byte, 2000)

const gb = 2

func InitMemp() {
	n := gb / 2
	for x := 0; x < n; x++ {
		gbSlot := Ram2GBSlot{}
		memp = append(memp, &gbSlot)
	}
}
