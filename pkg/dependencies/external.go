package dependencies

import (
	"core-users/pkg/configs"
	"core-users/pkg/repository/external/database"

	"go.mongodb.org/mongo-driver/mongo"
)

type External struct {
	mongo *mongo.Database
}

func InitExternal(config configs.Config) External {
	return External{
		mongo: database.NewDatabaseMongo(config.DataBaseMongo),
	}
}
