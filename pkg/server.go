package tracker

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/spf13/viper"
)

type Response struct {
	Node string `json:"node"`
	Gas  Gas    `json:"gas"`
}

func gasHandlerFunc(c *Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		g, err := c.GetGas()
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
	defer c.Close()
	http.HandleFunc("/gas", gasHandlerFunc(c))
	bind := viper.GetString("bind")
	http.ListenAndServe(bind, nil)
}
