package park

import (
	"fmt"
	"strconv"
	"strings"
)

type (
	Park struct {
		Slots []*Slot
		Stdout []string
	}
	Slot struct {
		Empty bool
		Number int
		Color string
		RegistryNum string
	}
)

func New(n int) *Park {
	park := Park{Slots: make([]*Slot, n)}
	for index:=0; index<n; index++ {
		park.Slots[index] = &Slot{
			Number: index + 1,
			Empty: true,
		}
	}
	park.Stdout = append(park.Stdout, fmt.Sprintf("Created a parking lot with %v slots\n", n))
	return &park
}

func (p *Park) Park(color, registryNum string) {
	for index:=0; index < len(p.Slots); index++ {
		if p.Slots[index].Empty {
			p.Slots[index].Empty = false
			p.Slots[index].Color = color
			p.Slots[index].RegistryNum = registryNum
			p.Stdout = append(p.Stdout, fmt.Sprintf("Allocated slot number: %v\n", p.Slots[index].Number))
			return
		}
	}
	p.Stdout = append(p.Stdout, fmt.Sprintf("Sorry, parking lot is full\n"))
}

func (p *Park) Leave(n int) {
	slot := p.Slots[n-1]
	if !slot.Empty {
		slot.Empty = true
		slot.Color = ""
		slot.RegistryNum = ""
		p.Stdout = append(p.Stdout, fmt.Sprintf("Slot number %v is free\n", n))
		return
	}
	// TODO: check is empty
	p.Stdout = append(p.Stdout, fmt.Sprintf("Slot number %v is free\n", n))
}

func (p *Park) Status() {	
	stdout := []string{}
	stdout = append(stdout, fmt.Sprintf("Slot No.    Registration No    Colour\n"))
	for _, slot := range p.Slots {
		if !slot.Empty {
			stdout = append(stdout, fmt.Sprintf("%v           %s      %s\n", slot.Number, slot.RegistryNum, slot.Color))
		}
	}
	concatStdout := strings.Join(stdout, "")
	p.Stdout = append(p.Stdout, concatStdout)
}

func (p *Park) GetAllSlotNumberByColor(color string) {
	slotNum := []string{}
	for _, slot := range p.Slots {
		if slot.Color == color {
			slotNum = append(slotNum, strconv.Itoa(slot.Number))
		}
	}
	message := strings.Join(slotNum, ", ")
	if message == "" {
		p.Stdout = append(p.Stdout, "Not found\n")
		return
	}
	p.Stdout = append(p.Stdout, message + "\n")
}

func (p *Park) GetAllRegistryNumberByColor(color string) {
	registryNumbers := []string{}
	for _, slot := range p.Slots {
		if slot.Color == color {
			registryNumbers = append(registryNumbers, slot.RegistryNum)
		}
	}
	message := strings.Join(registryNumbers, ", ")
	if message == "" {
		p.Stdout = append(p.Stdout, "Not found\n")
		return
	}
	p.Stdout = append(p.Stdout, message + "\n")
}

func (p *Park) GetSlotNumberByRegistryNumber(registryNumber string) {
	for _, slot := range p.Slots {
		if slot.RegistryNum == registryNumber {
			p.Stdout = append(p.Stdout, fmt.Sprintf("%v\n", slot.Number))
			return
		}
	}
	p.Stdout = append(p.Stdout, fmt.Sprintf("Not found\n"))
	return
}


