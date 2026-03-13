package httperror

import (
	"encoding/json"
	"net/http"
)

// WriteJSONError is a helper function to simplify sending back the data to the client
// with an optional httpcode. Defaults to Bad request (400)
func (h *HttpError) WriteJSONError(w http.ResponseWriter, httpCode ...int) {
	w.Header().Set("Content-Type", "application/json")
	responseHttpCode := http.StatusBadRequest
	if len(httpCode) > 0 && httpCode[0] > 0 {
		responseHttpCode = httpCode[0]
	}
	w.WriteHeader(responseHttpCode)

	json.NewEncoder(w).Encode(h.PublicError)
}
