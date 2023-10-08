package main

import (
	"flag"
	"os"
)

var flagRunAddr string
var flagLogLevel string
var flagDatabaseURI string

func parseFlags() {
	flag.StringVar(&flagRunAddr, "a", ":8080", "address and port to run server")
	flag.StringVar(&flagLogLevel, "l", "info", "log level")
	flag.StringVar(&flagDatabaseURI, "d", "", "database URI")
	flag.Parse()

	// для случаев, когда в переменной окружения RUN_ADDR присутствует непустое значение,
	// переопределим адрес запуска сервера,
	// даже если он был передан через аргумент командной строки
	if envRunAddr := os.Getenv("RUN_ADDR"); envRunAddr != "" {
		flagRunAddr = envRunAddr
	}
	if envLogLevel := os.Getenv("LOG_LEVEL"); envLogLevel != "" {
		flagLogLevel = envLogLevel
	}
	if envDatabaseURI := os.Getenv("DATABASE_URI"); envDatabaseURI != "" {
		flagDatabaseURI = envDatabaseURI
	}
}
