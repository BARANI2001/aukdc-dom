package main

import (
	"crypto/tls"
	"database/sql"
	"flag"
	"html/template"
	"log"
	"net/http"
	"os"
	"time"
	
	"aukdc.dom.com/internal/models"

	"github.com/alexedwards/scs/postgresstore"
	"github.com/alexedwards/scs/v2"
	"github.com/go-playground/form/v4"
	_ "github.com/lib/pq"
)

type application struct{
	errorLog *log.Logger
	infoLog *log.Logger
	honoraria *models.HonorariumModel
	faculty *models.FacultyModel
	templateCache map[string]*template.Template
	formDecoder *form.Decoder
	sessionManager *scs.SessionManager
}
func main(){
	addr := flag.String("addr", ":4000", "HTTP network address")
	dsn :=flag.String("dsn", "postgres://webaukdcdom:neodom@localhost/aukdcdom?sslmode=disable","PostgreSQL Data Source Name")


	flag.Parse()

	
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	db, err := openDB(*dsn)
	if err != nil {
		errorLog.Fatal(err)
	}
	defer db.Close()

	templateCache, err:=newTemplateCache()
	if err != nil {
		errorLog.Fatal(err)
	}

	formDecoder:= form.NewDecoder()

	sessionManager:=scs.New()
	sessionManager.Store=postgresstore.New(db)
	sessionManager.Lifetime=12*time.Hour

	sessionManager.Cookie.Secure = true

	app:=&application{
		errorLog:errorLog,
		infoLog:infoLog,
		honoraria:&models.HonorariumModel{DB: db},
		faculty:&models.FacultyModel{DB: db},
		templateCache: templateCache,
		formDecoder:formDecoder,
		sessionManager:sessionManager,
	}
	
	tlsConfig:=&tls.Config{
		CurvePreferences: []tls.CurveID{tls.X25519, tls.CurveP256},
	}

	
	srv := &http.Server{
		Addr:*addr,
		ErrorLog: errorLog,
		Handler: app.routes(),
		TLSConfig: tlsConfig,
		IdleTimeout:time.Minute,
		ReadTimeout:5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}


	infoLog.Printf("Starting server on %s",*addr)
	err= srv.ListenAndServeTLS("./tls/cert.pem", "./tls/key.pem")
	errorLog.Fatal(err)
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
