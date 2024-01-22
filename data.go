package main

import "strconv"

var (
	data    []Company
	devices []Device
)

type Company struct {
	ID      string
	Company string
	Contact string
	Country string
}

type Device struct {
	ID        string
	Hostname  string
	IPAddress string
	State     string
	UserName  string
	Notes     string
}

func init() {
	data = []Company{
		{
			ID:      "1",
			Company: "Amazon",
			Contact: "Jeff Bezos",
			Country: "United States",
		},
		{
			ID:      "2",
			Company: "Apple",
			Contact: "Tim Cook",
			Country: "United States",
		},
		{
			ID:      "3",
			Company: "Microsoft",
			Contact: "Satya Nadella",
			Country: "United States",
		},
	}
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

func getCompanyByID(id string) Company {
	var result Company
	for _, i := range data {
		if i.ID == id {
			result = i
			break
		}
	}
	return result
}

func updateCompany(company Company) {
	result := []Company{}
	for _, i := range data {
		if i.ID == company.ID {
			i.Company = company.Company
			i.Contact = company.Contact
			i.Country = company.Country
		}
		result = append(result, i)
	}
	data = result
}

func addCompany(company Company) {
	max := 0
	for _, i := range data {
		n, _ := strconv.Atoi(i.ID)
		if n > max {
			max = n
		}
	}
	max++
	id := strconv.Itoa(max)

	data = append(data, Company{
		ID:      id,
		Company: company.Company,
		Contact: company.Contact,
		Country: company.Country,
	})
}

func deleteCompany(id string) {
	result := []Company{}
	for _, i := range data {
		if i.ID != id {
			result = append(result, i)
		}
	}
	data = result
}

func getDeviceByID(id string) Device {
	var result Device
	for _, i := range devices {
		if i.ID == id {
			result = i
			break
		}
	}
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
		Hostname:    device.Hostname,
		UserName:  device.UserName,
		IPAddress: device.IPAddress,
		Notes: device.Notes,
		State: device.State,
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
