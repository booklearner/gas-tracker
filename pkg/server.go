package tracker

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/spf13/viper"
)

type Response struct {
	Gas  int    `json:"gas"`
	Node string `json:"node"`
}

func RunServer() {
	bind := viper.GetString("bind")

	http.HandleFunc("/gas", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		gas := Response{Gas: GetGas(), Node: viper.GetString("node")}
		resp, err := json.Marshal(gas)
		if err != nil {
			w.WriteHeader(500)
		}

		fmt.Fprintf(w, string(resp))
	})

	http.ListenAndServe(bind, nil)
}
