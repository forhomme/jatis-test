package config

import (
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

type DB struct {
	DbPostgres *gorm.DB
}

var (
	onceDbPostgres sync.Once
	instanceDB     *DB
)

func GetInstancePostgresDb() *gorm.DB {
	onceDbPostgres.Do(func() {
		postgreInfo := Config.Database
		logs := fmt.Sprintf("[INFO] Connected to POSTGRES TYPE = %s | LogMode = %+v", postgreInfo.Host, postgreInfo.LogMode)

		dbConfig := "host=" + postgreInfo.Host + " port=" + fmt.Sprintf("%d", postgreInfo.Port) + " user=" + postgreInfo.Username + " dbname=" + postgreInfo.Name + " search_path=" + postgreInfo.Schema + " sslmode=" + postgreInfo.SSLMode
		if postgreInfo.Password != "" {
			dbConfig += " password=" + postgreInfo.Password
		}
		db, err := gorm.Open(postgres.Open(dbConfig), &gorm.Config{})
		if err != nil {
			logs = fmt.Sprintf("[ERROR] Failed to connect to POSTGRES with err %s. Config=%s", err.Error(), postgreInfo.Host)
			log.Fatalln(logs)
		}

		sqlDB, err := db.DB()
		if err != nil {
			logs = fmt.Sprintf("[ERROR] Failed to connect to POSTGRES with err %s. Config=%s", err.Error(), postgreInfo.Host)
			log.Fatalln(logs)
		}

		sqlDB.SetMaxIdleConns(10)
		sqlDB.SetMaxOpenConns(20)
		sqlDB.SetConnMaxLifetime(10 * time.Minute)
		dialect := postgres.New(postgres.Config{Conn: sqlDB})
		loggerLevel := logger.Error
		if postgreInfo.LogMode {
			loggerLevel = logger.Info
		}
		dbConnection, err := gorm.Open(dialect, &gorm.Config{
			NamingStrategy: schema.NamingStrategy{
				SingularTable: true,
			},
			Logger: logger.New(
				log.New(os.Stdout, "\r\n", log.LstdFlags),
				logger.Config{
					SlowThreshold: time.Second,
					LogLevel:      loggerLevel,
				},
			),
		})
		if err != nil {
			logs = fmt.Sprintf("[ERROR] Failed to connect to POSTGRES with err %s. Config=%s", err.Error(), postgreInfo.Host)
			log.Fatalln(logs)
		}
		fmt.Println(logs)
		instanceDB = &DB{DbPostgres: dbConnection}
	})
	return instanceDB.DbPostgres
}
