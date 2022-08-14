package main

import (
	"database/sql"
	"flag"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"satotory.dyplom.go/pkg/models/mysql"
)

type neuteredFileSystem struct {
	fs http.FileSystem
}

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
	lessons  *mysql.LessonModel
}

func main() {
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Llongfile)

	addr := flag.String("addr", ":4000", "web addr HTTP")
	dsn := flag.String("dsn", "web:123@/dyplom_rework?parseTime=true", "Название MySQL источника данных")
	flag.Parse()

	db, err := openDB(*dsn)
	if err != nil {
		errorLog.Fatal(err)
	}
	defer db.Close()

	infoLog.Printf("Server start on %s", *addr)

	app := &application{
		errorLog: errorLog,
		infoLog:  infoLog,
		lessons:  &mysql.LessonModel{DB: db},
	}

	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  app.routes(),
	}

	err = srv.ListenAndServe()
	errorLog.Fatal(err)
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}

func (nfs neuteredFileSystem) Open(path string) (http.File, error) {
	f, err := nfs.fs.Open(path)
	if err != nil {
		return nil, err
	}

	s, err := f.Stat()
	if s.IsDir() {
		return nil, err
	}

	return f, nil
}
