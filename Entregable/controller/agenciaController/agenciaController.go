package agenciaController

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Juanscabu/SeminarioGoLang/Entregable/entity"
	"github.com/Juanscabu/SeminarioGoLang/Entregable/service/agenciaService"
	"github.com/dimfeld/httptreemux/v5"
)

var serviceAgencia agenciaService.ServiceAgencia

var find httptreemux.HandlerFunc

func Start(db *sql.DB) {
	serviceAgencia, _ = agenciaService.New(db)
}

func SaveAgenciaHandler(w http.ResponseWriter, r *http.Request, params map[string]string) {
	w.Header().Set("Content-Type", "application/json")
	var a entity.Agencia
	_ = json.NewDecoder(r.Body).Decode(&a)

	response, _ := serviceAgencia.Save(a)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}

func FindByIdAgenciaHandler(w http.ResponseWriter, r *http.Request, params map[string]string) {
	w.Header().Set("Content-Type", "application/json")
	param := params["id"]
	id, err := strconv.Atoi(param)
	itsId := err == nil

	if !itsId {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	a, _ := serviceAgencia.FindByID(id)
	if a.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(a)
}

func FindAllAgenciasHandler(w http.ResponseWriter, r *http.Request, params map[string]string) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(serviceAgencia.FindAll())
}

func UpdateAgenciaHandler(w http.ResponseWriter, r *http.Request, params map[string]string) {
	w.Header().Set("Content-Type", "application/json")
	param := params["id"]
	id, err := strconv.Atoi(param)
	itsId := err == nil

	if !itsId {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var a entity.Agencia
	_ = json.NewDecoder(r.Body).Decode(&a)

	ag, _ := serviceAgencia.FindByID(id)
	if ag.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	a.ID = int64(id)
	serviceAgencia.Update(a)
	json.NewEncoder(w).Encode(a)
}

func RemoveAgenciaHandler(w http.ResponseWriter, r *http.Request, params map[string]string) {
	param := params["id"]
	id, err := strconv.Atoi(param)
	itsId := err == nil

	if !itsId {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = serviceAgencia.Remove(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
	} else {
		w.WriteHeader(http.StatusNoContent)
	}
}
