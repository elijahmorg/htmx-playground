package main

import (
	"fmt"
	"net/http"

	"github.com/jritsema/go-htmx-starter/lib"
	"github.com/jritsema/gotoolbox/web"
)

// Delete -> DELETE /company/{id} -> delete, companys.html

// Edit   -> GET /company/edit/{id} -> row-edit.html
// Save   ->   PUT /company/{id} -> update, row.html
// Cancel ->	 GET /company/{id} -> nothing, row.html

// Add    -> GET /company/add/ -> companys-add.html (target body with row-add.html and row.html)
// Save   ->   POST /company -> add, companys.html (target body without row-add.html)
// Cancel ->	 GET /company -> nothing, companys.html

func index(r *http.Request) *web.Response {
	return web.HTML(http.StatusOK, html, "index.html", devices, nil)
}

// GET /company/add
func companyAdd(r *http.Request) *web.Response {
	return web.HTML(http.StatusOK, html, "company-add.html", AddDevicePageInfo, nil)
}

// /GET company/edit/{id}
func companyEdit(r *http.Request) *web.Response {
	id, _ := web.PathLast(r)
	row := getDeviceByID(id)
	fmt.Printf("Device Row state: %v\n", row.State)
	fmt.Printf("Device opts: %v\n", row.DeviceStateOptions)
	return web.HTML(http.StatusOK, html, "row-edit.html", row, nil)
}

// GET /company
// GET /company/{id}
// DELETE /company/{id}
// PUT /company/{id}
// POST /company
func companies(r *http.Request) *web.Response {
	id, segments := web.PathLast(r)
	switch r.Method {

	case http.MethodDelete:
		deleteDevice(id)
		return web.HTML(http.StatusOK, html, "companies.html", devices, nil)

	//cancel +      var cert LCSCert
	case http.MethodGet:
		if segments > 1 {
			//cancel edit
			row := getDeviceByID(id)
			return web.HTML(http.StatusOK, html, "row.html", row, nil)
		} else {
			//cancel add
			return web.HTML(http.StatusOK, html, "companies.html", devices, nil)
		}

	//save edit
	case http.MethodPut:
		row := getDeviceByID(id)
		r.ParseForm()
		row.Hostname = r.Form.Get("hostname")
		row.IPAddress = r.Form.Get("ipaddress")
		row.UserName = r.Form.Get("username")
		row.State, _ = lib.ParseStateEnum(r.Form.Get("state"))
		row.Notes = r.Form.Get("notes")
		updateDevice(row)
		return web.HTML(http.StatusOK, html, "row.html", row, nil)

	//save add
	case http.MethodPost:
		row := Device{}
		r.ParseForm()
		row.Hostname = r.Form.Get("hostname")
		row.IPAddress = r.Form.Get("ipaddress")
		row.UserName = r.Form.Get("username")
		row.State, _ = lib.ParseStateEnum(r.Form.Get("state"))
		row.Notes = r.Form.Get("notes")
		addDevice(row)
		return web.HTML(http.StatusOK, html, "companies.html", devices, nil)
	}

	return web.Empty(http.StatusNotImplemented)
}
