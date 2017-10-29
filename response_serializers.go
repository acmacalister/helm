package helm

import (
	"encoding/json"
	"encoding/xml"
	"net/http"
)

func RespondWithJSON(w http.ResponseWriter, v interface{}, status int) {
	json, _ := json.Marshal(v)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(json)
}

func RespondWithXML(w http.ResponseWriter, v interface{}, status int) {
	xml, _ := xml.Marshal(v)

	w.Header().Set("Content-Type", "application/xml")
	w.WriteHeader(status)
	w.Write(xml)
}
