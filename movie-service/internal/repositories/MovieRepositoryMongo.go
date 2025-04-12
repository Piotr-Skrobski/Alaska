package repositories

import (
	"context"
	"log"

	"github.com/Piotr-Skrobski/Alaska/movie-service/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoMovieRepository struct {
	db *mongo.Collection
}

func NewMovieRepository(db *mongo.Database) *MongoMovieRepository {
	return &MongoMovieRepository{
		db: db.Collection("movies"),
	}
}

func (r *MongoMovieRepository) Create(movie models.Movie) error {
	_, err := r.db.InsertOne(context.Background(), movie)
	if err != nil {
		log.Printf("Error inserting movie: %v", err)
		return err
	}
	return nil
}

func (r *MongoMovieRepository) FindByTitle(title string) (models.Movie, error) {
	var movie models.Movie
	filter := bson.M{"title": title}
	err := r.db.FindOne(context.Background(), filter).Decode(&movie)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return movie, nil
		}
		return movie, err
	}
	return movie, nil
}

func (r *MongoMovieRepository) FindByIMDbID(imdbID string) (models.Movie, error) {
	var movie models.Movie
	filter := bson.M{"imdb_id": imdbID}
	err := r.db.FindOne(context.Background(), filter).Decode(&movie)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return movie, nil
		}
		return movie, err
	}
	return movie, nil
}

func (r *MongoMovieRepository) Update(movie models.Movie) error {
	filter := bson.M{"imdb_id": movie.IMDbID}
	update := bson.M{
		"$set": movie,
	}
	_, err := r.db.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Printf("Error updating movie: %v", err)
		return err
	}
	return nil
}

func (r *MongoMovieRepository) Delete(imdbID string) error {
	filter := bson.M{"imdb_id": imdbID}
	_, err := r.db.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Printf("Error deleting movie: %v", err)
		return err
	}
	return nil
}
