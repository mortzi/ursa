package data

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	urlCollectionName = "url"
	defaultTimeout    = 10 * time.Second
)

//UrsaRepo used to query data
type UrsaRepo struct {
	Db *mongo.Database
}

//AddURL adds a url to database
func (repo *UrsaRepo) AddURL(ctx context.Context, url *UrsaURL) (<-chan string, <-chan error) {
	resultChan := make(chan string)
	errChan := make(chan error, 1)

	go func() {
		defer close(resultChan)
		defer close(errChan)

		db := repo.Db
		urlBson := &bson.D{
			{Key: "url", Value: url.URL},
			{Key: "category", Value: url.Category},
			{Key: "tag", Value: url.Tag},
		}

		res, err := db.Collection(urlCollectionName).InsertOne(ctx, urlBson)
		if err != nil {
			errChan <- err
			return
		}

		select {
		case resultChan <- res.InsertedID.(primitive.ObjectID).Hex():
		case <-ctx.Done():
			return
		}
	}()

	return resultChan, errChan
}

//DeleteURL removes a url
func (repo *UrsaRepo) DeleteURL(ctx context.Context, ids ...primitive.ObjectID) (<-chan int64, <-chan error) {
	resultChan := make(chan int64)
	errChan := make(chan error, 1)

	go func() {
		defer close(resultChan)
		defer close(errChan)

		db := repo.Db
		toRemoveBson := bson.D{}

		for _, id := range ids {
			toRemoveBson = append(toRemoveBson, bson.E{Key: "_id", Value: id})
		}

		res, err := db.Collection(urlCollectionName).DeleteMany(ctx, toRemoveBson)
		if err != nil {
			errChan <- err
			return
		}

		select {
		case resultChan <- res.DeletedCount:
		case <-ctx.Done():
			return
		}
	}()

	return resultChan, errChan
}

//DeleteURLByURL deletes a url
func (repo *UrsaRepo) DeleteURLByURL(ctx context.Context, urls ...string) (<-chan int64, <-chan error) {
	resultChan := make(chan int64)
	errChan := make(chan error, 1)

	go func() {
		defer close(resultChan)
		defer close(errChan)

		db := repo.Db
		toRemoveBson := bson.D{}

		for _, url := range urls {
			toRemoveBson = append(toRemoveBson, bson.E{Key: "url", Value: url})
		}

		res, err := db.Collection(urlCollectionName).DeleteMany(ctx, toRemoveBson)
		if err != nil {
			errChan <- err
			return
		}

		select {
		case resultChan <- res.DeletedCount:
		case <-ctx.Done():
			return
		}
	}()

	return resultChan, errChan
}

//GetURLByID get a url by it's id
func (repo *UrsaRepo) GetURLByID(ctx context.Context, ids ...primitive.ObjectID) (<-chan *UrsaURL, <-chan error) {
	resultChan := make(chan *UrsaURL)
	errChan := make(chan error, 1)

	go func() {
		defer close(resultChan)
		defer close(errChan)

		db := repo.Db
		toFindBson := bson.D{}

		for _, id := range ids {
			toFindBson = append(toFindBson, bson.E{Key: "_id", Value: id})
		}

		cur, err := db.Collection(urlCollectionName).Find(ctx, toFindBson)
		if err != nil {
			errChan <- err
			return
		}
		if err := cur.Err(); err != nil {
			errChan <- err
			return
		}
		defer cur.Close(ctx)

		for cur.Next(ctx) {
			elem := &bson.D{}

			if err := cur.Decode(elem); err != nil {
				errChan <- err
				return
			}

			m := elem.Map()

			url := createUrsaurlFromBsonMap(m)

			select {
			case resultChan <- url:
			case <-ctx.Done(): //operation was canceled
				return
			}
		}
	}()

	return resultChan, errChan
}

//GetAllURLs gets all the urls
func (repo *UrsaRepo) GetAllURLs(ctx context.Context) (<-chan *UrsaURL, <-chan error) {
	return repo.GetURLByID(ctx)
}

func createUrsaurlFromBsonMap(m bson.M) *UrsaURL {
	return &UrsaURL{
		ID:       m["_id"].(primitive.ObjectID),
		URL:      m["url"].(string),
		Category: m["category"].(string),
		Tag:      m["tag"].(string),
	}
}
