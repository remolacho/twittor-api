package postgresSql

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
	"strconv"
	"sync"
	"time"

	_ "github.com/lib/pq"
)

var (
	data *PostgresSql
	once sync.Once
)

type PostgresSql struct {
	DB *gorm.DB
}

func CurrentSession() *PostgresSql {
	once.Do(initDB)
	return data
}

func initDB() {
	db, err := getConnection()

	if err != nil {
		log.Fatal("error connection to DB postgresql " + err.Error())
	}

	log.Println("the connection to postgresql is success " + os.Getenv("DATABASE_NAME_PSQL") + "!!!")

	data = &PostgresSql{DB: db}
}

func getConnection() (*gorm.DB, error) {
	uri := os.Getenv("DATABASE_URI_PSQL")
	db, err := gorm.Open(postgres.Open(uri), &gorm.Config{})

	sqlDB, _ := db.DB()

	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	maxIdl, _ := strconv.Atoi(os.Getenv("MAX_IDLE_CNN"))
	sqlDB.SetMaxIdleConns(maxIdl)

	// SetMaxOpenConns sets the maximum number of open connections to the database.
	maxOpenCnn, _ := strconv.Atoi(os.Getenv("MAX_OPEN_CNN"))
	sqlDB.SetMaxOpenConns(maxOpenCnn)

	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sqlDB.SetConnMaxLifetime(time.Hour)

	return db, err
}
