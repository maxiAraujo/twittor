package routers

import (
	"encoding/json"
	"net/http"
	"twittor/bd"
	"twittor/models"
)

/*Registro es la funcion para crear en la bd el registro de usuario*/
func Registro(w http.ResponseWriter, r *http.Request){

	var t models.Usuario
	err := json.NewDecoder(r.Body).Decode(&t)

	if err != nil {
		http.Error(w, "Error en los datos recibidos "+err.Error(), 400)
		return
	}

	if len(t.Email)==0 {
		http.Error(w, "El email de usuario es requerido", 400)
		return
	}

	if len(t.Password)<6 {
		http.Error(w, "El password debe contener al menos 6 caracteres", 400)
		return
	}

	_,encontrado,_ := bd.ChequeoYaExisteUsuario(t.Email)
	if encontrado == true{
		http.Error(w, "Ya existe un usuario registrado con ese email", 400)
		return
	}

	_, status, _ := bd.InsertoRegistro(t)
	if err != nil {
		http.Error(w, "Ocurrio un error al intentar realizar el registro de usuario " + err.Error(), 400)
		return
	}

	if status == false{
		http.Error(w, "No se ha logrado insentar el registro del usuario", 400)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
