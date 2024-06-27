package main

import (
	"context"
	"log"
	"log/slog"
	"net/http"
	"os"

	redismodel "github/my-project/URL/shortURLReddis/internal/model"

	"github.com/redis/go-redis/v9"
)

// var ctx = context.Background()

type application struct {
	infoLog  *log.Logger
	errorLog *log.Logger
	URLmodel redismodel.URLRepository
	ctx      context.Context
}

func main() {
	ctx := context.Background()
	db, err := connectReddisDB(ctx)
	if err != nil {
		slog.Error("connect db", "error", err)
		os.Exit(-1)
	}
	defer db.Close()
	slog.Info("redis db connected succesfully")

	addr := ":8080"

	infoLogs := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLogs := log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	app := &application{
		infoLog:  infoLogs,
		errorLog: errorLogs,
		URLmodel: redismodel.NewRedis(db),
		ctx:      ctx,
	}
	srv := &http.Server{
		Addr:    addr,
		Handler: app.routes(),
	}

	infoLogs.Printf("Starting server on %s", addr)

	err = srv.ListenAndServe()
	log.Fatal(err)
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
