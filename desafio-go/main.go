package main

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	_ "github.com/go-sql-driver/mysql"

	"github.com/gustavobcx/fullcycle-imersao14/desafio-go/internal/routes/entity"
	"github.com/gustavobcx/fullcycle-imersao14/desafio-go/internal/routes/infra/repository"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/api/routes", getRoutes)
	r.Post("/api/routes", postRoutes)

	http.ListenAndServe(":8080", r)
}

func getDb() *sql.DB {
	db, err := sql.Open("mysql", "root:root@tcp(database:3306)/routes?parseTime=true")
	if err != nil {
		panic(err)
	}

	return db
}

func getRoutes(w http.ResponseWriter, r *http.Request) {
	db := getDb()

	repository := repository.NewRouteRepositoryMysql(db)

	routes, _ := repository.Find()

	defer db.Close()

	response, _ := json.Marshal(routes)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func postRoutes(w http.ResponseWriter, r *http.Request) {
	var newRoute entity.Route

	json.NewDecoder(r.Body).Decode(&newRoute)

	db := getDb()

	repository := repository.NewRouteRepositoryMysql(db)

	lid, err := repository.Create(&newRoute)

	defer db.Close()

	if err != nil {
		panic(err)
	}

	newRoute.ID = lid

	response, _ := json.Marshal(newRoute)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(response)
}
