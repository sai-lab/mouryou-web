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

	config := new(models.Config)
	config.LoadSetting(os.Getenv("HOME") + "/.mouryou.json")

	controllers.Cluster = &config.Cluster

	top := &controllers.TopController{}
	cluster := &controllers.ClusterController{}
	cluster.LoadBalancer = &controllers.LoadBalancerController{}
	cluster.Vendors = &controllers.VendorsController{}
	cluster.Vendors.VirtualMachines = &controllers.VirtualMachinesController{}

	goji.Get("/assets/*", http.FileServer(&assetfs.AssetFS{Asset: Asset, AssetDir: AssetDir, AssetInfo: AssetInfo, Prefix: "/"}))
	goji.Get("/frontend/*", http.FileServer(&assetfs.AssetFS{Asset: Asset, AssetDir: AssetDir, AssetInfo: AssetInfo, Prefix: "/"}))
	goji.Get("/api/cluster", cluster.IndexAPI)
	goji.Get("/api/cluster/load_balancer", cluster.LoadBalancer.IndexAPI)
	goji.Get("/api/cluster/vendors", cluster.Vendors.IndexAPI)
	goji.Get("/api/cluster/vendors/:vid", cluster.Vendors.ShowAPI)
	goji.Get("/api/cluster/vendors/:vid/virtual_machines", cluster.Vendors.VirtualMachines.IndexAPI)
	goji.Get("/api/cluster/vendors/:vid/virtual_machines/:vmid", cluster.Vendors.VirtualMachines.ShowAPI)
	goji.Put("/api/cluster/vendors/:vid/virtual_machines/:vmid/:operation", cluster.Vendors.VirtualMachines.UpdateAPI)
	goji.Get("/", top.Index)

	goji.Serve()
}
