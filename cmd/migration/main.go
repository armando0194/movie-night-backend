package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/armando0194/movie-night-backend/pkg/utl/model"
	"github.com/armando0194/movie-night-backend/pkg/utl/secure"
	"github.com/go-pg/pg/orm"

	"github.com/go-pg/pg"
)

func main() {
	dbInsert := `INSERT INTO public.roles VALUES (100, 100, 'SUPER_ADMIN');
	INSERT INTO public.roles VALUES (110, 110, 'ADMIN');
	INSERT INTO public.roles VALUES (200, 200, 'USER');`
	var psn = "postgres://armando:dev@localhost:5432/movie"
	queries := strings.Split(dbInsert, ";")

	u, err := pg.ParseURL(psn)
	checkErr(err)
	db := pg.Connect(u)
	_, err = db.Exec("SELECT 1")
	checkErr(err)
	createSchema(db, &model.Role{}, &model.User{}, &model.Ratings{}, &model.Movie{}, &model.MovieNight{})

	for _, v := range queries[0 : len(queries)-1] {
		_, err := db.Exec(v)
		checkErr(err)
	}

	sec := secure.New(1, nil)

	userInsert := `INSERT INTO public.users (id, created_at, updated_at, first_name, last_name, username, password, email, active, role_id) VALUES (1, now(),now(),'Admin', 'Admin', 'admin', '%s', 'johndoe@mail.com', true, 100);`
	_, err = db.Exec(fmt.Sprintf(userInsert, sec.Hash("admin")))
	checkErr(err)
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func createSchema(db *pg.DB, models ...interface{}) {
	for _, model := range models {
		checkErr(db.CreateTable(model, &orm.CreateTableOptions{
			FKConstraints: true,
		}))
	}
}
