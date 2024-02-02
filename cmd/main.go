package main

import (
	srv "DATest"
	"DATest/pkg/handler"
	"DATest/pkg/repository"
	"DATest/pkg/service"
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	//Инициализируем логгер
	logrus.SetFormatter(new(logrus.JSONFormatter))
	logrus.SetLevel(logrus.DebugLevel)

	//Инициализируем конфиг данных через енв переменные
	/*
		if err := initConfig(); err != nil {
			logrus.Fatalf("error initializing configs: %s", err.Error())
		}
	*/
	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("unable to open get env variables (%s)", err.Error())
	}
	//Инициализируем БД
	db, err := repository.NewPostgresDB(repository.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Username: os.Getenv("DB_USERNAME"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_DBNAME"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
		/*
			Host:     viper.GetString("db.host"),
			Port:     viper.GetString("db.port"),
			Username: viper.GetString("db.username"),
			Password: viper.GetString("db.password"),
			DBName:   viper.GetString("db.dbname"),
			SSLMode:  viper.GetString("db.sslmode"),
		*/
	})
	if err != nil {
		logrus.Fatalf("Failed to initialize db (%s)", err.Error())
	}
	logrus.Debug("Successful on DB connection")

	//Инициализируем структуру зависимостей
	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	//Запуск сервера
	srv := new(srv.Server)
	go func() {

		if err := srv.Run("8080", handlers.InitRoutes()); err != nil {
			log.Fatalf("error occured while running http server: %s", err.Error())
		}
	}()
	logrus.Print("Server Started")

	//Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit
	if err := srv.Shutdown(context.Background()); err != nil {
		logrus.Errorf("error occured on server shutting (%s)", err.Error())
	}
	if err := db.Close(); err != nil {
		logrus.Errorf("Error occured on db while closing (%s)", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
