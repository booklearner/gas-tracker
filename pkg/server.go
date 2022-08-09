package tracker

import (
	"encoding/json"
	"net/http"

	"github.com/spf13/viper"
)

type Response struct {
	Gas  int    `json:"gas"`
	Node string `json:"node"`
}

func gasHandlerFunc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	resp := Response{
		Gas:  GetGas(),
		Node: viper.GetString("node"),
	}
	json, err := json.Marshal(resp)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(json)
	return
}

func RunServer() {
	http.HandleFunc("/gas", http.HandlerFunc(gasHandlerFunc))
	bind := viper.GetString("bind")
	http.ListenAndServe(bind, nil)
}
