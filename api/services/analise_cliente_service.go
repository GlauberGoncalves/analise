package services

import (
	"errors"
	"github.com/glaubergoncalves/analise/api/models"
	"strings"
)

// AnaliseClienteService : Faz validações relacionadas ao cliente
func AnaliseClienteService(venda models.Venda)(bool, error){

	isPaisValido, err := isEntregaNoMesmoPais(venda)

	if err != nil {
		return false, err
	}
	return isPaisValido, nil
}

func isEntregaNoMesmoPais(venda models.Venda)(bool, error){
	paisCliente, paisEntrega := venda.Cliente.EnderecoPrincipal.Pais, venda.EnderecoEntrega.Pais
	paisCliente, paisEntrega = strings.ToUpper(paisCliente), strings.ToUpper(paisEntrega)
	isMesmoPais := strings.Contains(paisCliente, paisEntrega)

	if isMesmoPais != true {
		return isMesmoPais, errors.New("Pais de origem não é o mesmo de entrega")
	}

	return isMesmoPais, nil
}