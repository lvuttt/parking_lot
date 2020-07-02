package translate

import (
	"testing"
)

func TestOperateCreate(t *testing.T) {
	var functionName string
	cmd := []string{}
	trans := &Translation{}

	cmd = []string{"create_parking_lot", "2"}
	functionName, _ = trans.Operate(cmd)
	if functionName != "park.New" {
		t.Errorf("functionName should be park.New but got: %s\n", functionName)
	}

	cmd = []string{"park", "red", "abc-123"}
	functionName, _ = trans.Operate(cmd)
	if functionName != "tran.Park.Park" {
		t.Errorf("functionName should be tran.Park.Park but got: %s\n", functionName)
	}

	cmd = []string{"leave", "1"}
	functionName, _ = trans.Operate(cmd)
	if functionName != "tran.Park.Leave" {
		t.Errorf("functionName should be tran.Park.Leave but got: %s\n", functionName)
	}

	cmd = []string{"status"}
	functionName, _ = trans.Operate(cmd)
	if functionName != "tran.Park.Status" {
		t.Errorf("functionName should be tran.Park.Status but got: %s\n", functionName)
	}

	cmd = []string{"registration_numbers_for_cars_with_colour", "red"}
	functionName, _ = trans.Operate(cmd)
	if functionName != "tran.Park.GetAllRegistryNumberByColor" {
		t.Errorf("functionName should be tran.Park.GetAllRegistryNumberByColor but got: %s\n", functionName)
	}

	cmd = []string{"slot_numbers_for_cars_with_colour", "red"}
	functionName, _ = trans.Operate(cmd)
	if functionName != "tran.Park.GetAllSlotNumberByColor" {
		t.Errorf("functionName should be tran.Park.GetAllSlotNumberByColor but got: %s\n", functionName)
	}

	cmd = []string{"slot_number_for_registration_number", "1"}
	functionName, _ = trans.Operate(cmd)
	if functionName != "tran.Park.GetSlotNumberByRegistryNumber" {
		t.Errorf("functionName should be tran.Park.GetSlotNumberByRegistryNumber but got: %s\n", functionName)
	}
}