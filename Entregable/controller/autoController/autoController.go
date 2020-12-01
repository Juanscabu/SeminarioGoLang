package autocontroller

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Juanscabu/SeminarioGoLang/Entregable/entity"
	"github.com/Juanscabu/SeminarioGoLang/Entregable/service/autoService"
)

var serviceAuto autoService.ServiceAuto

func Start(db *sql.DB) {
	serviceAuto, _ = autoService.New(db)
}

func SaveAutoHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var a entity.Auto
	_ = json.NewDecoder(r.Body).Decode(&a)
	response, _ := serviceAuto.Save(a)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}

func FindByIdAutoHandler(w http.ResponseWriter, r *http.Request, params map[string]string) {
	w.Header().Set("Content-Type", "application/json")
	param := params["id"]
	id, err := strconv.Atoi(param)
	itsID := err == nil

	if !itsID {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	a, _ := serviceAuto.FindByID(id)
	if a.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(a)
}

func FindAllAutoHandler(w http.ResponseWriter, r *http.Request, params map[string]string) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(serviceAuto.FindAll())
}

func UpdateAutoHandler(w http.ResponseWriter, r *http.Request, params map[string]string) {
	w.Header().Set("Content-Type", "application/json")
	param := params["id"]
	id, err := strconv.Atoi(param)
	itsID := err == nil
	if !itsID {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var a entity.Auto
	_ = json.NewDecoder(r.Body).Decode(&a)

	auto, _ := serviceAuto.FindByID(id)
	if auto.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	auto.ID = int64(id)

	response, _ := serviceAuto.Update(a)
	if response.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
	} else {
		json.NewEncoder(w).Encode(response)
	}

}

func RemoveAutotHandler(w http.ResponseWriter, r *http.Request, params map[string]string) {
	param := params["id"]
	id, err := strconv.Atoi(param)
	itsID := err == nil

	if !itsID {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = serviceAuto.Remove(id)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
	} else {
		w.WriteHeader(http.StatusNoContent)
	}
}
