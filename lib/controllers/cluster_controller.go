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

func (cluster ClusterController) Index(c web.C, w http.ResponseWriter, r *http.Request) {
	cluster.API(w, Cluster)
}
