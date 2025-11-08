package envfetcher

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Data struct {
	port int
	host string
}

func ReadEnv() (values Data) {
	godotenv.Load()
	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		panic(values)
	}
	host := os.Getenv("HOST")
	values.port = port
	values.host = host
	return
}
