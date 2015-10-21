package main

import (
	"net/http"
	"os"

	assetfs "github.com/elazarl/go-bindata-assetfs"
	"github.com/sai-lab/mouryou-web/lib/mouryou-web/controllers"
	"github.com/sai-lab/mouryou-web/lib/realtime"
	"github.com/sai-lab/mouryou/lib/models"
	"github.com/zenazn/goji"
)

func main() {
	ws := realtime.NewServer("/ws")
	go ws.Listen()

	controllers.Asset = Asset

	config := models.LoadConfig(os.Getenv("HOME") + "/.mouryou.json")
	controllers.Cluster = &config.Cluster

	top := &controllers.TopController{}
	cluster := &controllers.ClusterController{}
	cluster.LoadBalancer = &controllers.LoadBalancerController{}
	cluster.Hypervisors = &controllers.HypervisorsController{}
	cluster.Hypervisors.VirtualMachines = &controllers.VirtualMachinesController{}

	goji.Get("/assets/*", http.FileServer(&assetfs.AssetFS{Asset: Asset, AssetDir: AssetDir, Prefix: "/"}))
	goji.Get("/app/*", http.FileServer(&assetfs.AssetFS{Asset: Asset, AssetDir: AssetDir, Prefix: "/"}))
	goji.Get("/api/cluster", cluster.Index)
	goji.Get("/api/cluster/load_balancer", cluster.LoadBalancer.Index)
	goji.Get("/api/cluster/hypervisors", cluster.Hypervisors.Index)
	goji.Get("/api/cluster/hypervisors/:hid", cluster.Hypervisors.Show)
	goji.Get("/api/cluster/hypervisors/:hid/virtual_machines", cluster.Hypervisors.VirtualMachines.Index)
	goji.Get("/api/cluster/hypervisors/:hid/virtual_machines/:vid", cluster.Hypervisors.VirtualMachines.Show)
	goji.Get("/", top.Index)

	goji.Serve()
}
