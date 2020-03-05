package api

import (
	"crypto/sha1"

	"github.com/armando0194/movie-night-backend/pkg/api/movie"

	"github.com/armando0194/movie-night-backend/pkg/api/auth"
	al "github.com/armando0194/movie-night-backend/pkg/api/auth/logging"
	at "github.com/armando0194/movie-night-backend/pkg/api/auth/transport"
	ml "github.com/armando0194/movie-night-backend/pkg/api/movie/logging"
	mt "github.com/armando0194/movie-night-backend/pkg/api/movie/transport"
	"github.com/armando0194/movie-night-backend/pkg/api/password"
	pl "github.com/armando0194/movie-night-backend/pkg/api/password/logging"
	pt "github.com/armando0194/movie-night-backend/pkg/api/password/transport"
	"github.com/armando0194/movie-night-backend/pkg/api/user"
	ul "github.com/armando0194/movie-night-backend/pkg/api/user/logging"
	ut "github.com/armando0194/movie-night-backend/pkg/api/user/transport"
	"github.com/armando0194/movie-night-backend/pkg/utl/config"
	jwt "github.com/armando0194/movie-night-backend/pkg/utl/middleware/jwt"
	"github.com/armando0194/movie-night-backend/pkg/utl/postgres"
	"github.com/armando0194/movie-night-backend/pkg/utl/rbac"
	"github.com/armando0194/movie-night-backend/pkg/utl/secure"
	"github.com/armando0194/movie-night-backend/pkg/utl/server"
	"github.com/armando0194/movie-night-backend/pkg/utl/zap"
)

func Serve(cfg *config.Configuration) error {
	db, err := postgres.New(cfg.DB.PSN, cfg.DB.Timeout, cfg.DB.LogQueries)
	if err != nil {
		return err
	}

	sec := secure.New(cfg.App.MinPasswordStr, sha1.New())
	rbac := rbac.New()
	jwt := jwt.New(cfg.JWT.Secret, cfg.JWT.SigningAlgorithm, cfg.JWT.Duration)
	log := zap.New()

	e := server.New()
	e.Static("/swaggerui", cfg.App.SwaggerUIPath)

	at.NewHTTP(al.New(auth.Initialize(db, jwt, sec, rbac), log), e, jwt.MWFunc())

	v1 := e.Group("/v1")
	v1.Use(jwt.MWFunc())

	ut.NewHTTP(ul.New(user.Initialize(db, rbac, sec), log), v1)
	pt.NewHTTP(pl.New(password.Initialize(db, rbac, sec), log), v1)
	mt.NewHTTP(ml.New(movie.Initialize(db), log), v1)

	server.Start(e, &server.Config{
		Port:                cfg.Server.Port,
		ReadTimeoutSeconds:  cfg.Server.ReadTimeout,
		WriteTimeoutSeconds: cfg.Server.WriteTimeout,
		Debug:               cfg.Server.Debug,
	})

	return nil
}

func checkErr(err error) {
	if err != nil {
		panic(err.Error())
	}
}
