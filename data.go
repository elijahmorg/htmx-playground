package main

import (
	"fmt"
	"strconv"
)

var (
	devices []Device

)

type Device struct {
	ID                 string
	Hostname           string
	IPAddress          string
	State              string
	UserName           string
	Notes              string
	DeviceStateOptions []DeviceState
}

type DeviceState struct {
	Value string
	State string
}

func init() {
	devices = []Device{
		{
			ID:        "1",
			Hostname:  "lab-label1",
			IPAddress: "10.80.80.10",
			State:     "available",
			UserName:  "",
			Notes:     "Custom config for hardware bug",
		},
		{
			ID:        "2",
			Hostname:  "lab-label2",
			IPAddress: "10.80.80.11",
			State:     "in use",
			UserName:  "elimorga",
			Notes:     "",
		},
		{
			ID:        "3",
			Hostname:  "lab-label3",
			IPAddress: "10.80.80.12",
			State:     "offline",
			UserName:  "jenkins",
			Notes:     "under maintenance",
		},
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
	result.DeviceStateOptions = SetDeviceStateOptions(result.State)
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

func SetDeviceStateOptions(state string) []DeviceState {
	// <!-- <option value="In Use" selected="selected">In Use</option> -->
	// <!-- <option value="Under Maintenance">Under Maintenance</option> -->
	// <!-- <option value="Available">Available</option> -->
	// <!-- <option value="Needs Attention/broken">Needs Attention/broken</option> -->
	// <!-- <option value="Offline">Offline</option> -->
	s := []DeviceState{}
	for _, value := range []string{"In Use", "Under Maintenance", "Available", "Needs Attention/broken", "Offline"} {
		fmt.Printf("Value: %s, State: %s\n", value, state)
		s = append(s, DeviceState{Value: value, State: state})
	}
	fmt.Printf("DeviceStateOptons %v\n", s)
	return s
}
