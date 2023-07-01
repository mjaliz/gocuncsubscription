package main

import (
	"database/sql"
	"github.com/alexedwards/scs/v2"
	"github.com/mjaliz/gocuncsubcription/data"
	"log"
	"sync"
)

type Config struct {
	Session  *scs.SessionManager
	DB       *sql.DB
	InfoLog  *log.Logger
	ErrorLog *log.Logger
	Wait     *sync.WaitGroup
	Models   data.Models
	Mailer   Mail
}
