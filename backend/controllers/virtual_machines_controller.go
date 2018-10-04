package controllers

import (
	"net/http"

	"github.com/sai-lab/mouryou/lib/models"
	"github.com/zenazn/goji/web"
)

type VirtualMachinesController struct {
	controller
}

func (machines VirtualMachinesController) IndexAPI(c web.C, w http.ResponseWriter, r *http.Request) {
	vendor := VendorsController{}.get(c.URLParams["vid"], w)
	if vendor == nil {
		return
	}

	machines.JSON(w, vendor.VirtualMachines)
}

func (machines VirtualMachinesController) ShowAPI(c web.C, w http.ResponseWriter, r *http.Request) {
	vendor := VendorsController{}.get(c.URLParams["vid"], w)
	if vendor == nil {
		return
	}

	machine := machines.get(vendor, c.URLParams["vmid"], w)
	if machine == nil {
		return
	}

	machines.JSON(w, machine)
}

func (machines VirtualMachinesController) get(vendor *models.VendorStruct, vmid string, w http.ResponseWriter) *models.VirtualMachine {
	machine, ok := vendor.VirtualMachines[vmid]
	if !ok {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return nil
	}

	return &machine
}
