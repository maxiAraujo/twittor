package bd

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
	"twittor/models"
)

func LeoTweets(ID string, pagina int64) ([]*models.DevuelvoTweets, bool) {

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	db := MongoCN.Database("twittor")
	col := db.Collection("tweet")

	var resultados []*models.DevuelvoTweets

	/*es como el where en java ej. WHERE id = :id */
	condicion := bson.M{
		"userid": ID,
	}

	opciones := options.Find()
	/*trae tantos documentos como marca el limite (20)*/
	opciones.SetLimit(20)
	/*ordena los documentos. En este caso por fecha y en orden descendente*/
	opciones.SetSort(bson.D{{Key: "fecha", Value: -1}})
	/*saltea segun pagina -1 * 20*/
	opciones.SetSkip((pagina -1)*20)

	cursor, err := col.Find(ctx, condicion, opciones)
	if err != nil {
		log.Fatal(err.Error())
		return resultados, false
	}

	for cursor.Next(context.TODO()){
		var registro models.DevuelvoTweets
		err := cursor.Decode(&registro)
		if err != nil {
			return resultados, false
		}
		resultados = append(resultados, &registro)
	}

	return resultados, true
}