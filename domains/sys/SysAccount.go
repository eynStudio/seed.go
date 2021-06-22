package sys

import (
	"time"

	"github.com/eynstudio/seed.go/base/db"
	"github.com/graphql-go/graphql"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type SysAccount struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Username  string             `bson:"username" json:"username,omitempty"`
	Password  string             `bson:"password" json:"password,omitempty"`
	Email     string             `bson:"email" json:"email,omitempty"`
	CreatedAt time.Time          `bson:"createdAt,omitempty" json:"createdAt,omitempty"`
}

var SysAccountType = graphql.NewObject(graphql.ObjectConfig{
	Name: "SysAccount",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: db.ID,
		},
		"username": &graphql.Field{
			Type: graphql.String,
		},
		"password": &graphql.Field{
			Type: graphql.String,
		},
		"email": &graphql.Field{
			Type: graphql.String,
		},
		"createdAt": &graphql.Field{
			Type: graphql.String,
		},
	},
})
