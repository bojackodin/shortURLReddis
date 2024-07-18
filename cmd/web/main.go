package main

import (
	"context"
	"flag"
	"log"
	"log/slog"
	"os"
	"os/signal"

	redismodel "github/my-project/URL/shortURLReddis/internal/model"

	"github.com/redis/go-redis/v9"
)

type config struct {
	port int
}

type application struct {
	config   config
	infoLog  *log.Logger
	errorLog *log.Logger
	URLmodel redismodel.URLRepository
}

func main() {
	var cfg config

	flag.IntVar(&cfg.port, "port", 8080, "server port")

	flag.Parse()

	ctx := context.Background()

	db, err := connectReddisDB(ctx)
	if err != nil {
		slog.Error("connect db", "error", err)
		os.Exit(-1)
	}
	defer db.Close()

	slog.Info("redis db connected succesfully")

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	app := &application{
		config:   cfg,
		infoLog:  infoLog,
		errorLog: errorLog,
		URLmodel: redismodel.NewRedis(db),
	}

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	err = app.run(ctx)
	if err != nil {
		log.Fatal(err)
	}
}

func connectReddisDB(ctx context.Context) (*redis.Client, error) {
	// Set client options
	db := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	// Check the connection
	err := db.Ping(ctx).Err()
	if err != nil {
		return nil, err
	}

	return db, nil
}
