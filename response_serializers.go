package helm

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"
)

func RespondWithJSON(w http.ResponseWriter, v interface{}, status int) {
	json, err := json.Marshal(v)
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to marshal to json: %s", err), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(json)
}

func RespondWithXML(w http.ResponseWriter, v interface{}, status int) {
	xml, err := xml.Marshal(v)
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to marshal to xml: %s", err), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/xml")
	w.WriteHeader(status)
	w.Write(xml)
}
