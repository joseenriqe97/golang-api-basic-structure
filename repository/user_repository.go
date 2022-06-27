package repository

import (
	"fmt"
	"time"

	"github.com/joseenriqe97/golang-init/config"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type userRepository struct {
	db *mongo.Database
}
type UserRepository interface {
	GetUser(idUser string) (string, error)
}

//BORRAR LUEGO
type ColectionMongo struct {
	ID        primitive.ObjectID `bson:"_id"`
	Email     string             `bson:"email"`
	Pass      string             `bson:"pass"`
	NameStore string             `bson:"nameStore"`
	CreatedAt time.Time          `bson:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at"`
	V         int64              `bson:"__v"`
	Image     string             `bson:"image"`
}

func NewUserRepository() *userRepository {
	return &userRepository{
		db: config.MongoDB,
	}
}

func (repository *userRepository) GetUser(idUser string) (string, error) {

	// config.GetCollection(repository.db, "use")
	// test := repository.db.FindOne(context.TODO(), bson.D{primitive.E{Key: "email", Value: "admin@gmail.com"}}).Decode(&user)
	// fmt.Print(test)
	// repository.db.Collection()
	// colle := config.GetCollection(config.MongoDb, "test")
	fmt.Println("colle")
	return "ss", nil
}
