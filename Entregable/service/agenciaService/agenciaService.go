package agenciaService

import (
	"database/sql"
	"errors"
	"strconv"

	"github.com/Juanscabu/SeminarioGoLang/Entregable/entity"
)

// Servicio de agencia
type ServiceAgencia interface {
	Save(entity.Agencia) (entity.Agencia, error)
	FindByID(int) (entity.Agencia, error)
	FindAll() []entity.Agencia
	Remove(int) error
	Update(entity.Agencia) (entity.Agencia, error)
}

type service struct {
	db *sql.DB
}

// New ...
func New(db *sql.DB) (ServiceAgencia, error) {
	return service{db}, nil
}

func (s service) Save(a entity.Agencia) (entity.Agencia, error) {
	query := "INSERT INTO agencia(name) VALUES (?)"
	prepare, err := s.db.Prepare(query)
	if err != nil {
		return entity.Agencia{}, err
	}

	row, err := prepare.Exec(a.Nombre)

	if err != nil {
		return entity.Agencia{}, err
	}

	a.ID, err = row.LastInsertId()

	return a, err
}

func (s service) FindByID(ID int) (entity.Agencia, error) {
	rows, err := s.db.Query("SELECT * FROM agency WHERE id_agency = ?", ID)
	if err != nil {
		return entity.Agencia{}, err
	}

	rows.Next()

	var agencia entity.Agencia
	var id int64
	var nombre string
	err2 := rows.Scan(&id, &nombre)

	if err2 != nil {
		return entity.Agencia{}, err
	} else {
		agencia = entity.Agencia{id, nombre}
	}

	return agencia, nil
}

func (s service) FindAll() []entity.Agencia {
	rows, err := s.db.Query("SELECT * FROM agencia")
	if err != nil {
		return nil
	} else {
		agencias := []entity.Agencia{}
		for rows.Next() {
			var id int64
			var nombre string
			err2 := rows.Scan(&id, &nombre)

			if err2 != nil {
				return nil
			} else {
				agencia := entity.Agencia{id, nombre}
				agencias = append(agencias, agencia)
			}
		}
		return agencias
	}
}

func (s service) Remove(ID int) error {
	result, err := s.db.Exec("DELETE FROM agencia WHERE id_agencia = ?", ID)
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
		return errors.New("agency with id=" + strconv.Itoa(ID) + " no fue eliminada")
	}
}

func (s service) Update(a entity.Agencia) (entity.Agencia, error) {
	result, err := s.db.Exec("UPDATE agencia SET nombre = ? WHERE id_agencia = ?", a.Nombre, a.ID)
	if err != nil {
		return entity.Agencia{}, err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return entity.Agencia{}, err
	}

	if rows > 0 {
		return a, nil
	} else {
		return entity.Agencia{}, errors.New("No se puede actualizar esa agencia")
	}
}
