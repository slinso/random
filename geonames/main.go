package main

import (
	"log"
	"strings"

	"github.com/mkrou/geonames"
	"github.com/mkrou/geonames/models"
	"github.com/sanity-io/litter"
)

func main() {
	p := geonames.NewParser()

	// print all cities with a population greater than 5000
	err := p.GetAdminDivisions(func(geoname *models.AdminDivision) error {
		if strings.HasPrefix(geoname.Code, "DE") {
			litter.Dump(geoname)
		}
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}
}
