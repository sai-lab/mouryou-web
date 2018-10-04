package controllers

import (
	"net/http"
	"strconv"

	"github.com/sai-lab/mouryou/lib/models"
	"github.com/zenazn/goji/web"
)

type VendorsController struct {
	controller
	VirtualMachines *VirtualMachinesController
}

func (vendors VendorsController) IndexAPI(c web.C, w http.ResponseWriter, r *http.Request) {
	vendors.JSON(w, Cluster.Vendors)
}

func (vendors VendorsController) ShowAPI(c web.C, w http.ResponseWriter, r *http.Request) {
	vendor := vendors.get(c.URLParams["vid"], w)
	if vendor == nil {
		return
	}

	vendors.JSON(w, vendor)
}

func (vendors VendorsController) get(vid string, w http.ResponseWriter) *models.VendorStruct {
	id, err := strconv.Atoi(vid)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return nil
	} else if id < 0 || id >= len(Cluster.Vendors) {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return nil
	}

	return &Cluster.Vendors[id]
}
