package api

import (
	"encoding/json"
	"net/http"

	"github.com/PauWol/ClusterNAS/internal/drives"
)



func GetDrives(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var p drives.Provider
	drives, err := p.GetDrives() // <-- use false here
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(drives)
}

/* func AddDrive(w http.ResponseWriter, r *http.Request) {
	var d util.Drive
	if err := json.NewDecoder(r.Body).Decode(&d); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	drives = append(drives, d)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(d)
}
 */