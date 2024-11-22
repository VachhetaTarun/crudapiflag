package main

import (
	"context"
	"log"
	"net/http"

	"crudecho/config"
	model "crudecho/models"
	"crudecho/routes"
	service "crudecho/services"

	"github.com/caarlos0/env/v6"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	pgDB    *gorm.DB
	mongoDB *mongo.Database
)

var dsn = "host=localhost user=postgres password=root dbname=manager port=5434 sslmode=disable TimeZone=Asia/Kolkata"

func migrate() {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	// Auto migrate the Worker and SiteManagerDetails models
	db.AutoMigrate(&model.Worker{}, &model.SiteManagerDetails{})

	service.SetPostgresDB(db) // Set the postgres DB in the service package
}

func init() {
	var err error

	// Load MongoConfig and PostgresConfig from environment variables
	mongoCfg := &config.MongoConfig{}
	if err := env.Parse(mongoCfg); err != nil {
		log.Fatalf("Failed to parse MongoConfig: %v", err)
	}
	pgCfg := &config.PostgresConfig{}
	if err := env.Parse(pgCfg); err != nil {
		log.Fatalf("Failed to parse PostgresConfig: %v", err)
	}

	// Initialize PostgreSQL connection using GORM
	pgDB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to PostgreSQL: %v", err)
	} else {
		log.Println("Connected to PostgreSQL using GORM successfully.")
		service.SetPostgresDB(pgDB) // Pass the GORM instance to the service layer
	}

	// Initialize MongoDB connection using configurations
	clientOptions := options.Client().ApplyURI(mongoCfg.MongoDBURL)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	} else {
		log.Println("Connected to MongoDB successfully.")
		service.SetMongoDB(client.Database(mongoCfg.MongoDBName))
	}
}

func main() {
	migrate()
	router := routes.RegisterRoutes()
	log.Fatal(http.ListenAndServe(":8082", router))
}
