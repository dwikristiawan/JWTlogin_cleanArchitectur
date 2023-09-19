package config

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Root struct {
	App      App
	Postgres Postgres
}

func Load(filenames ...string) Root {
	// we do not care if there is no .env file.
	_ = godotenv.Overload(filenames...)

	r := Root{
		App:      App{},
		Postgres: Postgres{},
	}

	mustLoad("APP", &r.App)
	mustLoad("POSTGRES", &r.Postgres)
	return r
}

// mustLoad require env vars to satisfy spec interface rules.
func mustLoad(prefix string, spec interface{}) {
	err := envconfig.Process(prefix, spec)
	if err != nil {
		panic(err)
	}
}

// mayLoad assume env vars can to satisfy spec interface rules.
func mayLoad(prefix string, spec interface{}) {
	_ = envconfig.Process(prefix, spec)
}
