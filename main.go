package main

import (
	"context"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

type Instrument struct {
	Name string
}

type Musician struct {
	Name       string
	Instrument *Instrument
}

type Band struct {
	Name     string
	Musician *Musician
}

type Album struct {
	Name string
}

type Music struct {
	Name string
}

func main() {
	godotenv.Load()

	dbUri := os.Getenv("NEO4J_HOST")
	user := os.Getenv("NEO4J_USER")
	pwd := os.Getenv("NEO4J_PWD")

	driver, err := neo4j.NewDriverWithContext(dbUri, neo4j.BasicAuth(user, pwd, ""))
	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	defer driver.Close(ctx)

	musicians, err := FindMusicians(driver, ctx)
	if err != nil {
		panic(err)
	}

	for _, m := range musicians {
		fmt.Println(m.Name)
		fmt.Println("\t", m.Instrument.Name)
	}

	fmt.Println("-----------------***-----------------")

	bands, err := FindBands(driver, ctx)
	if err != nil {
		panic(err)
	}

	for _, b := range bands {
		fmt.Println(b.Name)
		fmt.Println("\t", b.Musician.Name)
		fmt.Println("\t\t", b.Musician.Instrument.Name)
	}
}

func FindBands(driver neo4j.DriverWithContext, ctx context.Context) ([]*Band, error) {
	session := driver.NewSession(ctx, neo4j.SessionConfig{AccessMode: neo4j.AccessModeRead})
	defer session.Close(ctx)
	bands, err := session.ExecuteRead(ctx, func(tx neo4j.ManagedTransaction) (interface{}, error) {
		var bands []*Band
		result, err := tx.Run(ctx, "MATCH(b:Band)<-[:Belongs]-(m:Musician)-[:Play]-(i:Instrument) RETURN b.name, m.name, i.name ORDER BY b.name, m.name", nil)
		if err != nil {
			return nil, err
		}

		records, _ := result.Collect(ctx)
		for _, r := range records {
			bands = append(bands, &Band{
				Name: r.Values[0].(string),
				Musician: &Musician{
					Name: r.Values[1].(string),
					Instrument: &Instrument{
						Name: r.Values[2].(string),
					},
				},
			})
		}
		return bands, nil
	})

	if err != nil {
		return nil, err
	}

	return bands.([]*Band), nil
}

func FindMusicians(driver neo4j.DriverWithContext, ctx context.Context) ([]*Musician, error) {
	session := driver.NewSession(ctx, neo4j.SessionConfig{AccessMode: neo4j.AccessModeRead})
	defer session.Close(ctx)
	musicians, err := session.ExecuteRead(ctx, func(tx neo4j.ManagedTransaction) (interface{}, error) {
		var musicians []*Musician
		result, err := tx.Run(ctx, "MATCH(m:Musician)-[:Play]-(i:Instrument) RETURN m.name, i.name", nil)
		if err != nil {
			return nil, err
		}

		records, _ := result.Collect(ctx)
		for _, r := range records {
			m := &Musician{
				Name: r.Values[0].(string),
				Instrument: &Instrument{
					Name: r.Values[1].(string),
				},
			}
			musicians = append(musicians, m)
		}

		return musicians, nil
	})

	return musicians.([]*Musician), err
}
