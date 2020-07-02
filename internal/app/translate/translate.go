package translate

import (
	"errors"
	"fmt"
	"log"
	"parking_lot/internal/app/park"
	"strconv"
)

type (
	Translation struct {
		Park 	*park.Park
		Commands [][]string
	}
)

func (tran *Translation) Run() {
	for _, cmd := range tran.Commands {
		_, err := tran.Operate(cmd)
		if err != nil {
			log.Fatalf("%s\n", err)
		}
	}
}

func (tran *Translation) Operate(cmd []string) (string, error) {
	if cmd[0] == "create_parking_lot" {
		num, err := strconv.Atoi(cmd[1])
		if err != nil {
			log.Fatalf("cannot convert string to int: %s\n", err)
		}
		tran.Park = park.New(num)
		return "park.New", nil
	}
	if cmd[0] == "park" {
		tran.Park.Park(cmd[2], cmd[1])
		return "tran.Park.Park", nil
	}
	if cmd[0] == "leave" {
		num, err := strconv.Atoi(cmd[1])
		if err != nil {
			log.Fatalf("cannot convert string to int: %s\n", err)
		}
		tran.Park.Leave(num)
		return "tran.Park.Leave", nil
	}
	if cmd[0] == "status" {
		tran.Park.Status()
		return "tran.Park.Status", nil
	}
	if cmd[0] == "registration_numbers_for_cars_with_colour" {
		tran.Park.GetAllRegistryNumberByColor(cmd[1])
		return "tran.Park.GetAllRegistryNumberByColor", nil
	}
	if cmd[0] == "slot_numbers_for_cars_with_colour" {
		tran.Park.GetAllSlotNumberByColor(cmd[1])
		return "tran.Park.GetAllSlotNumberByColor", nil
	}
	if cmd[0] == "slot_number_for_registration_number" {
		tran.Park.GetSlotNumberByRegistryNumber(cmd[1])
		return "tran.Park.GetSlotNumberByRegistryNumber", nil
	}
	return "", errors.New(fmt.Sprintf("%s: command not found", cmd[0]))
}