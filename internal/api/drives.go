package api

import (
	"encoding/json"
	"net/http"
)

type Drive struct {
	Name string `json:"name"`
	Size string `json:"size"`
}

var drives = []Drive{
	{Name: "sda", Size: "500GB"},
	{Name: "sdb", Size: "1TB"},
}

func GetDrives(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(drives)
}

func AddDrive(w http.ResponseWriter, r *http.Request) {
	var d Drive
	if err := json.NewDecoder(r.Body).Decode(&d); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	drives = append(drives, d)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(d)
}
