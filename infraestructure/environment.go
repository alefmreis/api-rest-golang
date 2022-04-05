package infraestructure

import "os"

type Env struct {
	APPLICATION_PORT  string
	POSTGRES_DATABASE string
	POSTGRES_USERNAME string
	POSTGRES_PASSWORD string
	POSTGRES_HOST     string
	POSTGRES_PORT     string
}

func NewEnv() Env {
	env := Env{}
	env.loadEnv()
	return env
}

func (env *Env) loadEnv() {
	env.APPLICATION_PORT = os.Getenv("APPLICATION_PORT")
	env.POSTGRES_DATABASE = os.Getenv("POSTGRES_DATABASE")
	env.POSTGRES_USERNAME = os.Getenv("POSTGRES_USERNAME")
	env.POSTGRES_PASSWORD = os.Getenv("POSTGRES_PASSWORD")
	env.POSTGRES_HOST = os.Getenv("POSTGRES_HOST")
	env.POSTGRES_PORT = os.Getenv("POSTGRES_PORT")
}
