package main

import (
	"infra-base-go/internal/config"
	"infra-base-go/internal/database"
	"log"
	"os"
	"strconv"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

// Essa função vai ser reponsavél por criar e gerenciar as migrações via comandos cli
func main() {

	// carregando conf de ambiente
	config.LoadConfig()
	dbConfig := &config.GetConfig().DB

	// criando instancia de conexao com banco de dados(gorm)
	db, err := database.New(dbConfig)
	if err != nil {
		log.Fatal(err)
	}

	sqlInstance, err := db.DB.DB()
	if err != nil {
		log.Fatal(err)
	}

	// criando instancia de migracao postgres
	instance, err := postgres.WithInstance(sqlInstance, &postgres.Config{})
	if err != nil {
		log.Fatal(err)
	}

	// instanciando gerenciador de migracao(pasta de destino, driverName, instancia)
	m, err := migrate.NewWithDatabaseInstance("file://cmd/migrate/migrations", "postgres", instance)
	if err != nil {
		log.Fatal(err)
	}

	command := os.Args[1]

	switch command {
	case "up":
		if err := m.Up(); err != nil && err != migrate.ErrNoChange {
			log.Fatal(err)
		}
		log.Println("Migration up executed")
	case "down":
		if err := m.Down(); err != nil && err != migrate.ErrNoChange {
			log.Fatal(err)
		}
		log.Println("Migration down executed")
	case "force":
		if len(os.Args) < 3 {
			log.Fatal("Informe a versão a forçar (ex: force 20251013185523)")
		}
		mVersion, _ := strconv.Atoi(os.Args[2])
		if err := m.Force(mVersion); err != nil {
			log.Fatal(err)
		}
	default:
		log.Fatalf("Comando desconhecido: %s", command)
	}
}
