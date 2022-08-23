package tracker

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/spf13/viper"
)

type Response struct {
	Node string `json:"node"`
	Gas  Gas    `json:"gas"`
}

func gasHandlerFunc(c *ethclient.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		g, err := GetGas(c)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		resp := Response{
			Gas:  g,
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
}

func RunServer() {
	c, err := InitClient()
	if err != nil {
		log.Fatal(err)
	}
	http.HandleFunc("/gas", gasHandlerFunc(c))
	bind := viper.GetString("bind")
	http.ListenAndServe(bind, nil)
}
