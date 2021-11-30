package routers

import (
	"encoding/json"
	"net/http"
	"twittor/bd"
	"twittor/models"
)

func ConsultaRelacion(w http.ResponseWriter, r *http.Request){
	ID := r.URL.Query().Get("id")

	var t models.Relacion
	t.UsuarioID = IDUsuario
	t.UsuarioRelacionID = ID

	var resp models.RespuestaConsultaRelacion

	status, err := bd.ConsultoRelacion(t)

	if err != nil || status == false {
		resp.Status=false
	} else {
		resp.Status=true
	}

	w.Header().Set("Content-Type", "applicacion/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}