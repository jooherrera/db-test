package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

const (
    host     = "postgres"
    port     = 5432
    user     = "postgres"
    password = ""
    dbname   = "postgres"
)

func main() {
    for {
        fmt.Println("Menu:")
        fmt.Println("1. Crear base de datos")
		fmt.Println("2. Crear tablas")
        fmt.Println("3. Salir")
        fmt.Print("Seleccione una opción: ")

        var choice int
        fmt.Scanln(&choice)

        switch choice {
        case 1:
            fmt.Println("Creando base de datos...")
            createDatabase()
        case 2:
            fmt.Println("Creando tablas...")
		case 3:
            fmt.Println("Saliendo..")
            return
        default:
            fmt.Println("Opción no válida")
        }
    }
}

func createDatabase() {
    psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
        host, port, user, password, dbname)

    db, err := sql.Open("postgres", psqlInfo)

    if err != nil {
        log.Fatalf("Error al conectar con la base de datos: %v", err)
    }

    defer db.Close()

    sqlFile, err := os.ReadFile("../scripts/crear_db.sql")
    if err != nil {
        log.Fatalf("Error al leer el archivo SQL: %v", err)
    }

    _, err = db.Exec(string(sqlFile))
    if err != nil {
        log.Fatalf("Error al ejecutar el archivo SQL: %v", err)
    }

    fmt.Println("Base de datos creada exitosamente")
}
