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
	FindAllByAgencia(int) []entity.Auto
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
	query := "INSERT INTO auto(modelo,marca,patente,id_agencia) VALUES (?,?,?,?)"
	prepare, err := s.db.Prepare(query)
	if err != nil {
		panic(err.Error())
	}

	row, err := prepare.Exec(a.Modelo, a.Marca, a.Patente, a.IdAgencia)

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
	var idAgencia int64
	err2 := rows.Scan(&id, &modelo, &marca, &patente, &idAgencia)

	if err2 != nil {
		return auto, err
	} else {
		auto = entity.Auto{ID: id, Modelo: modelo, Marca: marca, Patente: patente, IdAgencia: idAgencia}

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
			var idAgencia int64
			err2 := rows.Scan(&id, &modelo, &marca, &patente, &idAgencia)

			if err2 != nil {
				return nil
			} else {
				auto := entity.Auto{ID: id, Modelo: modelo, Marca: marca, Patente: patente, IdAgencia: idAgencia}
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
	_, err := s.db.Exec("UPDATE Auto SET modelo = ?, marca = ?, patente = ?, id_agencia = ? WHERE id_auto = ?", a.Modelo, a.Marca, a.Patente, a.IdAgencia, a.ID)
	return a, err
}

func (s service) FindAllByAgencia(idAgencia int) []entity.Auto {
	rows, err := s.db.Query("SELECT auto.* FROM auto JOIN agencia ON(auto.id_agencia = agencia.id_agencia) WHERE agencia.id_agencia = ?", idAgencia)
	if err != nil {
		return nil
	} else {
		autos := []entity.Auto{}
		for rows.Next() {
			var id int64
			var modelo string
			var marca string
			var patente string
			var idAgencia int64
			err2 := rows.Scan(&id, &modelo, &marca, &patente, &idAgencia)

			if err2 != nil {
				panic(err2.Error())
			} else {
				auto := entity.Auto{ID: id, Modelo: modelo, Marca: marca, Patente: patente, IdAgencia: idAgencia}
				autos = append(autos, auto)
			}
		}
		return autos
	}
}
