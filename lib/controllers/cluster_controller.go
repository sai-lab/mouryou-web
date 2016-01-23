package controllers

import (
	"net/http"

	"github.com/zenazn/goji/web"
)

type ClusterController struct {
	controller
	LoadBalancer *LoadBalancerController
	Hypervisors  *HypervisorsController
}

func (cluster ClusterController) IndexAPI(c web.C, w http.ResponseWriter, r *http.Request) {
	cluster.JSON(w, Cluster)
}
