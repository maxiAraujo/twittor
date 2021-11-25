package bd

import (
    "context"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
    "log"
)

/*MongoCN es el objeto de conexion de la bd*/
var MongoCN=conectarBD()
var clientOptions = options.Client().ApplyURI("mongodb+srv://maxi:<password>@cluster0.isqpa.mongodb.net/twittor?retryWrites=true&w=majority")

/*conectarBD es una funcion que permite conectar la bd*/
func conectarBD() *mongo.Client {
    client, err := mongo.Connect(context.TODO(), clientOptions)

    if err != nil{
        log.Fatal(err.Error())
        return client
    }
    err = client.Ping(context.TODO(), nil)
    if err != nil {
        log.Fatal(err.Error())
        return client
    }

    log.Println("Conexion Exitosa con la BD")
    return client
}

/*ChequeoConnection es una funcion que devuelve un entero*/
func ChequeoConnection() int {
    err := MongoCN.Ping(context.TODO(), nil)
    if err != nil {
        log.Fatal(err.Error())
        return 0
    }
    return 1
}