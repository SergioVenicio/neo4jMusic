package repositories

import (
	"context"

	"github.com/SergioVenicio/neo4jMusic/entities"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

type AlbumRepository struct {
	Driver neo4j.DriverWithContext
}

func (r *AlbumRepository) FindAll(ctx context.Context) ([]*entities.Album, error) {
	session := r.Driver.NewSession(
		ctx,
		neo4j.SessionConfig{AccessMode: neo4j.AccessModeRead},
	)
	defer session.Close(ctx)

	albuns, err := session.ExecuteRead(ctx, func(tx neo4j.ManagedTransaction) (interface{}, error) {
		var albuns []*entities.Album
		result, err := tx.Run(
			ctx,
			`MATCH(a:Album)-[r:HAS_MUSIC]->(m)
			WITH a.name AS album, {order: r.order, name: m.name} AS music
			ORDER BY r.order
			WITH album, COLLECT(music) AS musics
			RETURN album, musics`,
			nil,
		)
		if err != nil {
			return nil, err
		}

		records, _ := result.Collect(ctx)
		for _, r := range records {
			var musics []*entities.Music
			for _, m := range r.Values[1].([]interface{}) {
				musicValue := m.(map[string]any)
				musics = append(musics, entities.NewMusic(
					musicValue["name"].(string),
					musicValue["order"].(float64),
				))
			}
			m := &entities.Album{
				Name:   r.Values[0].(string),
				Musics: musics,
			}
			albuns = append(albuns, m)
		}

		return albuns, nil
	})

	if err != nil {
		return nil, err
	}

	return albuns.([]*entities.Album), nil
}

func (r *AlbumRepository) FindByName(ctx context.Context, name string) ([]*entities.Album, error) {
	session := r.Driver.NewSession(
		ctx,
		neo4j.SessionConfig{AccessMode: neo4j.AccessModeRead},
	)
	defer session.Close(ctx)

	albuns, err := session.ExecuteRead(ctx, func(tx neo4j.ManagedTransaction) (interface{}, error) {
		var albuns []*entities.Album
		result, err := tx.Run(
			ctx,
			`MATCH(a:Album{name: $name})-[r:HAS_MUSIC]->(m)
			WITH a.name AS album, {order: r.order, name: m.name} AS music
			ORDER BY r.order
			WITH album, COLLECT(music) AS musics
			RETURN album, musics`,
			map[string]interface{}{"name": name},
		)
		if err != nil {
			return nil, err
		}

		records, _ := result.Collect(ctx)
		for _, r := range records {
			album := &entities.Album{
				Name: r.Values[0].(string),
			}
			var musics []*entities.Music
			for _, m := range r.Values[1].([]interface{}) {
				musicValue := m.(map[string]any)
				musics = append(musics, entities.NewMusic(
					musicValue["name"].(string),
					musicValue["order"].(float64),
				))
			}
			album.Musics = musics
			albuns = append(albuns, album)
		}

		return albuns, nil
	})

	if err != nil {
		return nil, err
	}

	return albuns.([]*entities.Album), nil
}

func NewAlbumRepository(driver neo4j.DriverWithContext) *AlbumRepository {
	return &AlbumRepository{
		Driver: driver,
	}
}
