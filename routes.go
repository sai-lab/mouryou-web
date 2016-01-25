package main

import (
	"net/http"
	"os"

	assetfs "github.com/elazarl/go-bindata-assetfs"
	"github.com/sai-lab/mouryou-web/backend/controllers"
	"github.com/sai-lab/mouryou-web/backend/realtime"
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

	goji.Get("/assets/*", http.FileServer(&assetfs.AssetFS{Asset: Asset, AssetDir: AssetDir, AssetInfo: AssetInfo, Prefix: "/"}))
	goji.Get("/frontend/*", http.FileServer(&assetfs.AssetFS{Asset: Asset, AssetDir: AssetDir, AssetInfo: AssetInfo, Prefix: "/"}))
	goji.Get("/api/cluster", cluster.IndexAPI)
	goji.Get("/api/cluster/load_balancer", cluster.LoadBalancer.IndexAPI)
	goji.Get("/api/cluster/hypervisors", cluster.Hypervisors.IndexAPI)
	goji.Get("/api/cluster/hypervisors/:hid", cluster.Hypervisors.ShowAPI)
	goji.Get("/api/cluster/hypervisors/:hid/virtual_machines", cluster.Hypervisors.VirtualMachines.IndexAPI)
	goji.Get("/api/cluster/hypervisors/:hid/virtual_machines/:vid", cluster.Hypervisors.VirtualMachines.ShowAPI)
	goji.Get("/", top.Index)

	goji.Serve()
}
