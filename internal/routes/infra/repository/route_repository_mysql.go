package repository

import (
	"database/sql"

	"github.com/gustavobcx/fullcycle-imersao14/desafio-go/internal/routes/entity"
)

type RouteRepositoryMysql struct {
	db *sql.DB
}

func NewRouteRepositoryMysql(db *sql.DB) *RouteRepositoryMysql {
	return &RouteRepositoryMysql{
		db: db,
	}
}

func (r *RouteRepositoryMysql) Create(route *entity.Route) (int64, error) {
	sql := "INSERT INTO routes (name, source, destination) VALUES (?, ?, ?)"
	result, err := r.db.Exec(sql, route.Name, route.Source, route.Destination)

	if err != nil {
		return 0, err
	}

	lid, err := result.LastInsertId()

	if err != nil {
		return 0, err
	}

	return lid, nil
}

func (r *RouteRepositoryMysql) Find() (*[]entity.Route, error) {
	rows, err := r.db.Query("SELECT id, name, source, destination FROM routes")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var routes []entity.Route

	for rows.Next() {
		var route entity.Route
		if err := rows.Scan(&route.ID, &route.Name, &route.Source, &route.Destination); err != nil {
			return &routes, err
		}
		routes = append(routes, route)
	}

	if err = rows.Err(); err != nil {
		return &routes, err
	}

	return &routes, nil
}
