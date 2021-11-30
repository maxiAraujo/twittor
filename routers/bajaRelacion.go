package routers

import (
	"net/http"
	"twittor/bd"
	"twittor/models"
)

func BajaRelacion(w http.ResponseWriter, r *http.Request)  {
	ID := r.URL.Query().Get("id")
	var t models.Relacion
	t.UsuarioID = IDUsuario
	t.UsuarioRelacionID = ID

	status, err := bd.BorroRelacion(t)
	if err != nil {
		http.Error(w, "Ocurrio un erro al intentar borrar relacion " + err.Error(), http.StatusBadRequest)
		return
	}
	if status == false {
		http.Error(w, "No se ha logrado borrar la relacion "+ err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}