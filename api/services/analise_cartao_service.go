package services

import (
	"errors"

	creditcard "github.com/durango/go-credit-card"
	"github.com/glaubergoncalves/analise/api/models"
)

func ValidaCartaoCredito(cartao models.Cartao) (bool, error) {

	card := creditcard.Card{Number: cartao.Numero, Cvv: cartao.CodigoSeguranca, Month: cartao.MesExpiracao, Year: cartao.AnoExpiracao}
	err := card.Validate()
	if err != nil {
		return false, errors.New("Dados do cartão são invalidos: " + err.Error())
	}

	return true, nil
}
