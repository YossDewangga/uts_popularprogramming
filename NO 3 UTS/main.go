package main

import (
	"biodata"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func VeiwDataDiri(w http.ResponseWriter, r *http.Request) {
	nama := r.FormValue("Nama")
	NIM := 0
	Alamat := ""

	if nama == "gusti" {
		nama = "Gusti"
		NIM = 1234
		Alamat = "Malang"
	} else if nama == "pangestu" {
		nama = "Pangestu"
		NIM = 321
		Alamat = "Kabupaten Malang"
	}
	datadiri := biodata.ShowDataDiri(nama, NIM, Alamat)

	str, err := json.Marshal(datadiri)

	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		fmt.Print(err)
	} else {
		io.WriteString(w, string(str))
	}
}

func main() {
	var mux = http.NewServeMux()

	mux.HandleFunc("/nama", VeiwDataDiri)

	http.ListenAndServe(":8080", mux)
}
