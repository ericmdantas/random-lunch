package placecontroller

import (
	"encoding/json"
	"github.com/ericmdantas/random-lunch/server/api/places/dao"
	place "github.com/ericmdantas/random-lunch/server/api/places/model"
	"github.com/julienschmidt/httprouter"
	"io/ioutil"
	"net/http"
)

func GetAll(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var pls []place.Place

	pls, err := placedao.All()

	w.Header().Set("Content-Type", "application/json")

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	plsm, err := json.Marshal(pls)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(plsm)
}

func New(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var pl place.Place

	info, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()

	w.Header().Set("Content-Type", "application/json")

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	err = json.Unmarshal(info, &pl)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Error trying to parse place info."))
		return
	}

	err = placedao.CreateNew(&pl)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	plm, err := json.Marshal(pl)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(plm)
}

func Update(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var pl place.Place

	w.Header().Set("Content-Type", "application/json")

	info, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	err = json.Unmarshal(info, &pl)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	err = placedao.Update(&pl)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	plm, err := json.Marshal(pl)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(plm)
}
