// handlers.go
package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"microservice/services"
	"net/http"
)

func GetTopTrackInfo(w http.ResponseWriter, r *http.Request) {
	region := r.URL.Query().Get("region")

	fmt.Println("region: ", region)
	w.Header().Set("Content-Type", "application/json")
	// Call services to get information
	topTrackInfo, err := services.GetTopTrack(region)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		io.WriteString(w, fmt.Sprintf("{\"message\":\"%v\"}", err.Error()))
		return
	}

	// Return the response
	w.WriteHeader(http.StatusOK)
	resp, err := json.Marshal(&topTrackInfo)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		io.WriteString(w, fmt.Sprintf("{\"message\":\"%v\"}", err.Error()))
		return
	}
	io.WriteString(w, string(resp))
}
