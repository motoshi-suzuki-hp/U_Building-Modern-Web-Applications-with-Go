package config

import "html/template"

// AppConfig hodls the application config
type AppConfig struct {
	UseCache bool
	TemplateCache map[string]*template.Template
}