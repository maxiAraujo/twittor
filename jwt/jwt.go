package jwt

import (
	jwt "github.com/golang-jwt/jwt"
	"time"
	"twittor/models"
)

func GeneroJWT(t models.Usuario) (string, error){

	miClave := []byte("MiClave123")

	var payload = jwt.MapClaims{
		"nombre":           t.Nombre,
		"apellidos":        t.Apellidos,
		"fecha_nacimiento": t.FechaNacimiento,
		"biografia":        t.Biografia,
		"ubicacion":        t.Ubicacion,
		"sitioweb":         t.SitioWeb,
		"_id":              t.ID.Hex(),
		"exp":              time.Now().Add(time.Hour * 24).Unix(),
		"email":            t.Email,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	tokenStr, err := token.SignedString(miClave)
	if err != nil {
		return tokenStr, err
	}

	return tokenStr, nil
}


