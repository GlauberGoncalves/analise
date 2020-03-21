package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/glaubergoncalves/analise/api/models"
	"github.com/glaubergoncalves/analise/api/responses"
)

func Analise(w http.ResponseWriter, r *http.Request) {

	venda, _ := extraiDadosDoRequest(w, r)

	responses.JSON(w, http.StatusOK, venda)
}

func extraiDadosDoRequest(w http.ResponseWriter, r *http.Request) (models.Venda, error) {

	venda := models.Venda{}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return venda, err
	}

	err = json.Unmarshal(body, &venda)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return venda, err
	}
	return venda, nil
}
