package services

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"

	"github.com/glaubergoncalves/analise/api/models"
)

const (
	API_URL   = "http://api.ipstack.com"
	API_TOKEN = "1ba3f62891a5cd6e23c948130cbb4b90"
)

func origemDaCompraMesmoDoCadastro(venda models.Venda) (bool, error) {

	const PAIS = "country_name"
	dispositivo := venda.Dispositivo
	paisCadastro := venda.Cliente.EnderecoPrincipal.Pais

	resp, err := http.Get(API_URL + "/" + dispositivo.IP + "?access_key=" + API_TOKEN + "&format=1&language=pt-br")
	if err != nil {
		return true, nil
	}
	defer resp.Body.Close()

	var result map[string]string
	json.NewDecoder(resp.Body).Decode(&result)

	if strings.ToUpper(paisCadastro) != strings.ToUpper(result[PAIS]) {
		return false, errors.New("País de origem do IP é '" + result[PAIS] + "'e a do país cadastrado é '" + paisCadastro)
	}
	return true, nil
}
