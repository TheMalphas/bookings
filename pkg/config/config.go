package config

import (
	"github.com/alexedwards/scs/v2"
	"html/template"
	"log"
)

// AppConfig holds the application config
type AppConfig struct {
	InfoLog       *log.Logger
	InProduction  bool
	Session       *scs.SessionManager
	TemplateCache map[string]*template.Template
	UseCache      bool
}
