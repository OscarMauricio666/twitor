package bd

import(
	"context"
	"log"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoC = ConectarBD()
var clientOptions = options.Client().ApplyURI("mongodb://root:example@localhost:27017/")

func ConectarBD() *mongo.Client{
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	log.Printfln("Conexion exitosa a la base de datos")
	return client
}
func ChequeoConnection() int {
	err = MongoC.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
		return 0
	}
	return 1

}