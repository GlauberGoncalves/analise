package servicestest

import (
	"fmt"
	"testing"

	"github.com/glaubergoncalves/analise/api/models"
	"github.com/glaubergoncalves/analise/api/services"
)

func TestValidaCartaoCredito(t *testing.T) {

	cartao := models.Cartao{
		NomeImpresso:    "glauber goncalves",
		Numero:          "4012001037141112",
		CodigoSeguranca: "123",
		Bandeira:        "MasterCard",
		MesExpiracao:    "12",
		AnoExpiracao:    "2023",
	}

	validacao, err := services.ValidaCartaoCredito(cartao)

	if validacao != true {
		fmt.Println(err.Error())
		t.Errorf("falha na validação do cartão de credito")
	}
}
