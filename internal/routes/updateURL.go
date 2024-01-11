package routes

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type UpdateURLRequest struct {
	URL  string `json:"url"`
	Name string `json:"name"`
}

func (h *Routes) updateURLHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")

	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("URL ID is required"))
		return
	}

	body, err := io.ReadAll(r.Body)
	defer r.Body.Close()

	if err != nil {
		println(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Error reading request body"))
		return
	}
	url, err := h.models.URLs.GetById(id)
	var updatedValues UpdateURLRequest

	err = json.Unmarshal(body, &updatedValues)
	if err != nil {
		println(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Error unmarshalling request body"))
		return
	}

	if updatedValues.URL == "" {
		updatedValues.URL = url.URL
	}
	if updatedValues.Name == "" {
		updatedValues.Name = url.Name
	}

	if err != nil {
		println(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Error retrieving URL"))
		return
	}

	err = h.models.URLs.Update(id, updatedValues.URL, updatedValues.Name)
	if err != nil {
		println(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Error updating URL"))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(url.URL))
}
