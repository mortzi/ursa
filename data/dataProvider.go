package data

import (
	"encoding/json"
	"net/url"
	"os"

	"github.com/google/uuid"
)

//Category something
type Category string

//UrsaURL something
type UrsaURL struct {
	ID       uuid.UUID
	URL      url.URL
	Name     string
	Category Category
}

func (ursaURL *UrsaURL) toJSON() string {
	enc := json.NewEncoder(os.Stdout)
	enc.Encode(&ursaURL)

	return ""
}

//Provider something
type Provider interface {
	Add(url *UrsaURL) error
	AddCategory(cat *Category) error
	Remove(url *UrsaURL) error
	RemoveCategory(cat *Category) error
	GetAll() ([]*UrsaURL, error)
	FindID(id *uuid.UUID) (*UrsaURL, error)
	FindURL(url *url.URL) (*UrsaURL, error)
	FindName(name string) (*UrsaURL, error)
	FindCategory(cat *Category) ([]*UrsaURL, error)
	FindCloseTo(a ...string) ([]*UrsaURL, error)
	FindCloseToURL(a ...string) ([]*UrsaURL, error)
	FindCloseToName(a ...string) ([]*UrsaURL, error)
	FindCloseToCategory(a ...string) ([]*UrsaURL, error)
}
