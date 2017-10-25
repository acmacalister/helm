package helm

import (
	"encoding/json"
	"encoding/xml"
	"net/http"
)

func RespondWithJSON(w http.ResponseWriter, v interface{}, status int) (int, error) {
	json, err := json.Marshal(v)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(json)

	return http.StatusOK, nil
}

func RespondWithXML(w http.ResponseWriter, v interface{}, status int) (int, error) {
	xml, err := xml.Marshal(v)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	w.Header().Set("Content-Type", "application/xml")
	w.WriteHeader(status)
	w.Write(xml)

	return http.StatusOK, nil
}
