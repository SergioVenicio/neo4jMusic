package main

import (
	"context"
	"fmt"
	"os"

	"github.com/SergioVenicio/neo4jMusic/entities"
	"github.com/joho/godotenv"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

func mainOld() {
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

func FindBands(driver neo4j.DriverWithContext, ctx context.Context) ([]*entities.Band, error) {
	session := driver.NewSession(ctx, neo4j.SessionConfig{AccessMode: neo4j.AccessModeRead})
	defer session.Close(ctx)
	bands, err := session.ExecuteRead(ctx, func(tx neo4j.ManagedTransaction) (interface{}, error) {
		var bands []*entities.Band
		result, err := tx.Run(ctx, "MATCH(b:Band)<-[:BELONGS]-(m:Musician)-[:PLAYS]-(i:Instrument) RETURN b.name, m.name, i.name ORDER BY b.name, m.name", nil)
		if err != nil {
			return nil, err
		}

		records, _ := result.Collect(ctx)
		for _, r := range records {
			bands = append(bands, &entities.Band{
				Name: r.Values[0].(string),
				Musician: &entities.Musician{
					Name: r.Values[1].(string),
					Instrument: &entities.Instrument{
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

	return bands.([]*entities.Band), nil
}

func FindMusicians(driver neo4j.DriverWithContext, ctx context.Context) ([]*entities.Musician, error) {
	session := driver.NewSession(ctx, neo4j.SessionConfig{AccessMode: neo4j.AccessModeRead})
	defer session.Close(ctx)
	musicians, err := session.ExecuteRead(ctx, func(tx neo4j.ManagedTransaction) (interface{}, error) {
		var musicians []*entities.Musician
		result, err := tx.Run(ctx, "MATCH(m:Musician)-[:PLAYS]-(i:Instrument) RETURN m.name, i.name", nil)
		if err != nil {
			return nil, err
		}

		records, _ := result.Collect(ctx)
		for _, r := range records {
			m := &entities.Musician{
				Name: r.Values[0].(string),
				Instrument: &entities.Instrument{
					Name: r.Values[1].(string),
				},
			}
			musicians = append(musicians, m)
		}

		return musicians, nil
	})

	return musicians.([]*entities.Musician), err
}
