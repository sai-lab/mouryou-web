package controllers

import (
	"net/http"
	"strconv"

	"github.com/sai-lab/mouryou/lib/models"
	"github.com/zenazn/goji/web"
)

type VirtualMachinesController struct {
	controller
}

func (machines VirtualMachinesController) Index(c web.C, w http.ResponseWriter, r *http.Request) {
	hypervisor := HypervisorsController{}.get(c.URLParams["hid"], w)
	if hypervisor == nil {
		return
	}

	machines.API(w, hypervisor.VirtualMachines)
}

func (machines VirtualMachinesController) Show(c web.C, w http.ResponseWriter, r *http.Request) {
	hypervisor := HypervisorsController{}.get(c.URLParams["hid"], w)
	if hypervisor == nil {
		return
	}

	machine := machines.get(hypervisor, c.URLParams["vid"], w)
	if machine == nil {
		return
	}

	machines.API(w, machine)
}

func (machines VirtualMachinesController) get(hypervisor *models.HypervisorStruct, vid string, w http.ResponseWriter) *models.VirtualMachineStruct {
	id, err := strconv.Atoi(vid)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return nil
	} else if id < 0 || id >= len(hypervisor.VirtualMachines) {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return nil
	}

	return &hypervisor.VirtualMachines[id]
}
