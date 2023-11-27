package main

import (
	"os"
	"tg_bot/internal/logger"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		logger.Error("Error loading .env file")
	}

	g := gen.NewGenerator(gen.Config{
		OutPath:       "./internal/database/dal",
		Mode:          gen.WithDefaultQuery | gen.WithQueryInterface | gen.WithoutContext,
		FieldNullable: true,
	})

	gormdb, _ := gorm.Open(mysql.Open(os.Getenv("DB_DSN")))

	g.UseDB(gormdb)

	g.ApplyBasic(g.GenerateAllTable()...)

	g.Execute()
}
