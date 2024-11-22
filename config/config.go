package config

type MongoConfig struct {
	MongoDBURL            string `env:"MONGODB_URL" envDefault:"mongodb://localhost:27017"`
	MongoDBName           string `env:"MONGODB_DB_NAME" envDefault:"mongo_db"`
	SiteManagerCollection string `env:"MONGODB_SITE_MANAGER_COLLECTION" envDefault:"site_manager_details"`
	WorkerCollection      string `env:"MONGODB_WORKER_COLLECTION" envDefault:"workers"`
}

type PostgresConfig struct {
	PostgresURL string `env:"POSTGRES_URL" envDefault:"localhost"`
	DBName      string `env:"POSTGRES_DB_NAME" envDefault:"postgres_db"`
	User        string `env:"POSTGRES_USER" envDefault:"postgres_user"`
	Password    string `env:"POSTGRES_PASSWORD" envDefault:"postgres_password"`
}
