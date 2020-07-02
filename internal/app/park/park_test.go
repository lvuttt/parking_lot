package park

import (
	"fmt"
	"testing"
)

func TestNewPark(t *testing.T) {
	park := New(2)
	if park.Slots[0].Number != 1 {
		t.Errorf("number of slot1 should be 1 but have %v\n", park.Slots[0].Number)
	}
	if park.Slots[1].Number != 2 {
		t.Errorf("number of slot2 should be 2 but have %v\n", park.Slots[1].Number)
	}
	if park.Stdout[len(park.Stdout)-1] != "Created a parking lot with 2 slots\n" {
		t.Errorf("message should be Created a parking lot with 2 slots\n but have %s\n", park.Stdout[len(park.Stdout)-1])
	}
}

func TestParkIsAvailable(t *testing.T) {
	found := false
	park := New(2)
	park.Park("red", "abc-123")
	park.Park("yellow", "abc-100")
	for index := 0; index < len(park.Slots); index++ {
		slot := park.Slots[index]
		if slot.Color == "red" && slot.Empty == false && slot.RegistryNum == "abc-123" {
			found = true
			break
		}
	}
	if !found {
		t.Error("park should be available but it is unavailable")
	}
	if park.Stdout[len(park.Stdout)-1] != "Allocated slot number: 2\n" {
		t.Errorf("message should be Allocated slot number: 2\n but have %s\n", park.Stdout[len(park.Stdout)-1])
	}
}

func TestParkIsUnAvailable(t *testing.T) {
	park := New(2)
	park.Park("red", "abc-123")
	park.Park("yellow", "abc-100")
	park.Park("black", "abc-101")
	for index := 0; index < len(park.Slots); index++ {
		slot := park.Slots[index]
		if slot.Color == "black" && slot.Empty == false && slot.RegistryNum == "abc-101" {
			t.Error("park should be unavailable but it is available")
		}
	}
	if park.Stdout[len(park.Stdout)-1] != "Sorry, parking lot is full\n" {
		t.Errorf("message should be Sorry, parking lot is full\n but have %s\n", park.Stdout[len(park.Stdout)-1])
	}
}

func TestLeave(t *testing.T) {
	park := New(2)
	park.Park("red", "abc-123")
	park.Park("yellow", "abc-100")
	park.Leave(1)
	if !park.Slots[0].Empty {
		t.Error("slot 1 should be empty but it is not empty")
	}
	if park.Stdout[len(park.Stdout)-1] != "Slot number 1 is free\n" {
		t.Errorf("message should be Slot number 1 is free\n but have %s\n", park.Stdout[len(park.Stdout)-1])
	}
}

func TestStatus(t *testing.T) {
	park := New(2)
	park.Park("red", "abc-123")
	park.Status()
	expected := fmt.Sprintf("Slot No.    Registration No    Colour\n" + "1           abc-123      red\n")
	if park.Stdout[len(park.Stdout)-1] != expected {
		t.Errorf("message should be %s but have %s\n", expected, park.Stdout[len(park.Stdout)-1])
	}
}

func TestGetAllSlotNumberByColor(t *testing.T) {
	park := New(2)
	park.Park("red", "abc-123")
	park.Park("yellow", "abc-100")
	park.GetAllSlotNumberByColor("red")
	if park.Stdout[len(park.Stdout)-1] != "1\n" {
		t.Errorf("message should be 1\n but have %s\n", park.Stdout[len(park.Stdout)-1])
	}
	park.GetAllSlotNumberByColor("blue")
	if park.Stdout[len(park.Stdout)-1] != "Not found\n" {
		t.Errorf("message should be Not found\n but have %s\n", park.Stdout[len(park.Stdout)-1])
	}
}

func TestGetAllRegistryNumberByColor(t *testing.T) {
	park := New(2)
	park.Park("red", "abc-123")
	park.Park("yellow", "abc-100")
	park.GetAllRegistryNumberByColor("red")
	if park.Stdout[len(park.Stdout)-1] != "abc-123\n" {
		t.Errorf("message should be abc-123\n but have %s\n", park.Stdout[len(park.Stdout)-1])
	}
	park.GetAllRegistryNumberByColor("blue")
	if park.Stdout[len(park.Stdout)-1] != "Not found\n" {
		t.Errorf("message should be Not found\n but have %s\n", park.Stdout[len(park.Stdout)-1])
	}
}

func TestGetSlotNumberByRegistryNumber(t *testing.T) {
	park := New(2)
	park.Park("red", "abc-123")
	park.Park("yellow", "abc-100")
	park.GetSlotNumberByRegistryNumber("abc-123")
	if park.Stdout[len(park.Stdout)-1] != "1\n" {
		t.Errorf("message should be 1\n but have %s\n", park.Stdout[len(park.Stdout)-1])
	}
	park.GetAllRegistryNumberByColor("abc-234")
	if park.Stdout[len(park.Stdout)-1] != "Not found\n" {
		t.Errorf("message should be Not found\n but have %s\n", park.Stdout[len(park.Stdout)-1])
	}
}