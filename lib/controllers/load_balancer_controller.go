package controllers

import (
	"net/http"

	"github.com/zenazn/goji/web"
)

type LoadBalancerController struct {
	controller
}

func (balancer LoadBalancerController) IndexAPI(c web.C, w http.ResponseWriter, r *http.Request) {
	balancer.JSON(w, Cluster.LoadBalancer)
}
