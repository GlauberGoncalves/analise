package services

import (
	"github.com/glaubergoncalves/analise/api/models"
)

func AnaliseFraudeService(venda models.Venda) Response {

	response := Response{
		ResultadAnalise: false,
		Motivos:         make([]string, 0),
	}

	// 1 - Analise do cliente
	analiseCliente, errCliente := AnaliseClienteService(venda)
	if errCliente != nil {
		response.Motivos = append(response.Motivos, errCliente.Error())
	}

	// 2 - Analise do dispositivo
	analiseDispositivo, errDispositivo := origemDaCompraMesmoDoCadastro(venda)
	if errDispositivo != nil {
		response.Motivos = append(response.Motivos, errDispositivo.Error())
	}

	// 3 - Analise do cartão de credito
	analiseCartao, errCartao := ValidaCartaoCredito(venda.CartaoCredito)
	if errCartao != nil {
		response.Motivos = append(response.Motivos, errCartao.Error())
	}

	response.ResultadAnalise = analiseDispositivo && analiseCliente && analiseCartao
	return response
}
