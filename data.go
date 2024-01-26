package main

import (
	"fmt"
	"strconv"

	"github.com/jritsema/go-htmx-starter/lib"
)

var (
	devices []Device
	AddDevicePageInfo AddDevicePage
	PossibleDeviceStates []DeviceState
)

type Device struct {
	ID                 string
	Hostname           string
	IPAddress          string
	State              lib.StateEnum
	UserName           string
	Notes              string
	DeviceStateOptions []DeviceState
}

type AddDevicePage struct {
	Devices []Device
	PossibleDeviceStates []DeviceState
}


//  DeviceState is for tracking a device's current state
// Value is the string value of State
// DisplayValue is the string that should be displayed to the end user
type DeviceState struct {
	DisplayValue string
	State lib.StateEnum
}



func init() {
	fmt.Println("Init Data")
	devices = []Device{
		{
			ID:        "1",
			Hostname:  "lab-label1",
			IPAddress: "10.80.80.10",
			State:     lib.StateEnumAvailable,
			UserName:  "",
			Notes:     "Custom config for hardware bug",
		},
		{
			ID:        "2",
			Hostname:  "lab-label2",
			IPAddress: "10.80.80.11",
			State:     lib.StateEnumInUse,
			UserName:  "elimorga",
			Notes:     "",
		},
		{
			ID:        "3",
			Hostname:  "lab-label3",
			IPAddress: "10.80.80.12",
			State:     lib.StateEnumOffline,
			UserName:  "jenkins",
			Notes:     "under maintenance",
		},
	}
	MakePossibleDeviceStates()
	AddDevicePageInfo.Devices = devices
	AddDevicePageInfo.PossibleDeviceStates = PossibleDeviceStates
}

func MakePossibleDeviceStates() {
	// I need a range of values with state key, and display value
	for _, state := range lib.StateEnumValues() {
		displayValue := lib.StateDisplay[state]
		PossibleDeviceStates = append(PossibleDeviceStates, DeviceState{DisplayValue: displayValue, State: state})
	}
}

func getDeviceByID(id string) Device {
	var result Device
	for _, i := range devices {
		if i.ID == id {
			result = i
			break
		}
	}
	result.DeviceStateOptions = GetDeviceStateOptions()
	return result
}

func updateDevice(device Device) {
	result := []Device{}
	for _, i := range devices {
		if i.ID == device.ID {
			i.Hostname = device.Hostname
			i.IPAddress = device.IPAddress
			i.UserName = device.UserName
			i.State = device.State
			i.Notes = device.Notes
		}
		result = append(result, i)
	}
	devices = result
}

func addDevice(device Device) {
	max := 0
	for _, i := range devices {
		n, _ := strconv.Atoi(i.ID)
		if n > max {
			max = n
		}
	}
	max++
	id := strconv.Itoa(max)

	devices = append(devices, Device{
		ID:        id,
		Hostname:  device.Hostname,
		UserName:  device.UserName,
		IPAddress: device.IPAddress,
		Notes:     device.Notes,
		State:     device.State,
	})
}

func deleteDevice(id string) {
	result := []Device{}
	for _, i := range devices {
		if i.ID != id {
			result = append(result, i)
		}
	}
	devices = result
}

func GetDeviceStateOptions() []DeviceState {
	s := make([]DeviceState, len(PossibleDeviceStates))
	copy(s, PossibleDeviceStates)
	return s
}
