package main

import (
	"database/sql"
	"flag"
	"log"
	"study-event-api/config"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

func main() {
	// コマンドライン引数を定義
	migrateUp := flag.Bool("up", false, "Apply all up migrations")
	migrateDown := flag.Bool("down", false, "Apply all down migrations")
	forceVersion := flag.Int("force", -1, "Force set version")
	flag.Parse()

	cfg, err := config.Load()
	if err != nil {
		log.Fatal("Failed to load config:", err)
	}

	db, err := sql.Open("postgres", cfg.GetDSN())
	if err != nil {
		log.Fatal(err)
	}

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		log.Fatal(err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://ddl/migrations",
		"postgres", driver)
	if err != nil {
		log.Fatal(err)
	}

	// 強制的にバージョンを設定
	if *forceVersion >= 0 {
		if err := m.Force(*forceVersion); err != nil {
			log.Fatal(err)
		}
		log.Printf("Database version forced to %d\n", *forceVersion)
		return
	}

	// up または down のオプションに基づいてマイグレーションを実行
	if *migrateUp {
		if err := m.Up(); err != nil && err != migrate.ErrNoChange {
			log.Fatal(err)
		}
		log.Println("Migrations applied successfully!")
	} else if *migrateDown {
		if err := m.Down(); err != nil && err != migrate.ErrNoChange {
			log.Fatal(err)
		}
		log.Println("Migrations rolled back successfully!")
	} else {
		log.Println("Please specify either -up or -down to run migrations.")
	}
}
