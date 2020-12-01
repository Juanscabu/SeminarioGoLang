package autoService

import (
	"database/sql"
	"errors"
	"strconv"

	"github.com/Juanscabu/SeminarioGoLang/Entregable/entity"
)

// ServiceAuto ...
type ServiceAuto interface {
	Save(entity.Auto) (entity.Auto, error)
	FindByID(int) (entity.Auto, error)
	FindAll() []entity.Auto
	Remove(int) error
	Update(entity.Auto) (entity.Auto, error)
}

type service struct {
	db *sql.DB
	//conf *config.Config
}

// New ...
func New(db *sql.DB) (ServiceAuto, error) {
	return service{db}, nil
}

func (s service) Save(a entity.Auto) (entity.Auto, error) {
	query := "INSERT INTO auto(modelo,marca,patente) VALUES (?,?,?)"
	prepare, err := s.db.Prepare(query)
	if err != nil {
		panic(err.Error())
	}

	row, err := prepare.Exec(a.Modelo, a.Marca, a.Patente)

	if err != nil {
		panic(err.Error())
	}

	a.ID, err = row.LastInsertId()

	return a, err
}

func (s service) FindByID(ID int) (entity.Auto, error) {
	rows, err := s.db.Query("SELECT * FROM auto WHERE id_auto = ?", ID)
	if err != nil {
		return entity.Auto{}, err
	}

	rows.Next()

	var auto entity.Auto
	var id int64
	var modelo string
	var marca string
	var patente string
	err2 := rows.Scan(&id, &modelo, &marca, &patente)

	if err2 != nil {
		return auto, err
	} else {
		auto = entity.Auto{ID: id, Modelo: modelo, Marca: marca, Patente: patente}

	}

	return auto, nil
}

func (s service) FindAll() []entity.Auto {
	rows, err := s.db.Query("SELECT * FROM auto")
	if err != nil {
		return nil
	} else {
		autos := []entity.Auto{}
		for rows.Next() {
			var id int64
			var modelo string
			var marca string
			var patente string
			err2 := rows.Scan(&id, &modelo, &marca, &patente)

			if err2 != nil {
				return nil
			} else {
				auto := entity.Auto{ID: id, Modelo: modelo, Marca: marca, Patente: patente}
				autos = append(autos, auto)
			}
		}
		return autos
	}
}

func (s service) Remove(ID int) error {
	result, err := s.db.Exec("DELETE FROM auto WHERE id_auto = ?", ID)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows > 0 {
		return nil
	} else {
		return errors.New("auto con id=" + strconv.Itoa(ID) + " no fue eliminado")
	}
}

func (s service) Update(a entity.Auto) (entity.Auto, error) {
	_, err := s.db.Exec("UPDATE Auto SET modelo = ?, marca = ?, patente = ? WHERE id_auto = ?", a.Modelo, a.Marca, a.Patente, a.ID)
	return a, err
}
