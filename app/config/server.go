package config

import "github.com/joho/godotenv"

type Server struct {
	Address string `default:":8080"`
}

func LoadForServer(filenames ...string) Server {
	// we do not care if there is no .env file.
	_ = godotenv.Overload(filenames...)

	r := Server{}

	mustLoad("SERVER", &r)

	return r
}
