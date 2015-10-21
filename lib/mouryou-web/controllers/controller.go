package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/sai-lab/mouryou/lib/models"
	"github.com/yosssi/ace"
)

var (
	Asset   func(string) ([]byte, error)
	Cluster *models.ClusterStruct
)

type controller struct{}

func (c controller) Render(w http.ResponseWriter, view string) {
	options := &ace.Options{Asset: Asset}

	tpl, err := ace.Load("lib/mouryou-web/views/layout", view, options)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	err = tpl.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (c controller) API(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	encoder.Encode(data)
}
