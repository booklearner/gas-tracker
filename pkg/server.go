package tracker

import (
	"fmt"
	"net/http"

	"github.com/spf13/viper"
)

func RunServer() {
	bind := viper.GetString("bind")

	http.HandleFunc("/gas", func(w http.ResponseWriter, r *http.Request) {
		gas := GetGas()
		fmt.Fprintf(w, gas)
	})

	http.ListenAndServe(bind, nil)
}
