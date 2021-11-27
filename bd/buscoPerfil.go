package bd

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"time"
	"twittor/models"
)

func BuscoPerfil(ID string) (models.Usuario, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()
	bd := MongoCN.Database("twittor")
	col := bd.Collection("usuarios")

	var perfil models.Usuario
	objId, _ := primitive.ObjectIDFromHex(ID)

	condicion := bson.M{
		"_id": objId,
	}

	err := col.FindOne(ctx, condicion).Decode(&perfil)
	perfil.Password = ""
	if err != nil {
		log.Println("Registro no encontrado " + err.Error())
		return perfil, err
	}

	return perfil, nil
}