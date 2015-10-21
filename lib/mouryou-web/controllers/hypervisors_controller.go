package controllers

import (
	"net/http"
	"strconv"

	"github.com/sai-lab/mouryou/lib/models"
	"github.com/zenazn/goji/web"
)

type HypervisorsController struct {
	controller
	VirtualMachines *VirtualMachinesController
}

func (hypervisors HypervisorsController) Index(c web.C, w http.ResponseWriter, r *http.Request) {
	hypervisors.API(w, Cluster.Hypervisors)
}

func (hypervisors HypervisorsController) Show(c web.C, w http.ResponseWriter, r *http.Request) {
	hypervisor := hypervisors.get(c.URLParams["hid"], w)
	if hypervisor == nil {
		return
	}

	hypervisors.API(w, hypervisor)
}

func (hypervisors HypervisorsController) get(hid string, w http.ResponseWriter) *models.HypervisorStruct {
	id, err := strconv.Atoi(hid)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return nil
	} else if id < 0 || id >= len(Cluster.Hypervisors) {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return nil
	}

	return &Cluster.Hypervisors[id]
}
