package models

type Venda struct {
	Valor           float64     `json:"valor"`
	Cliente         Cliente     `json:"cliente"`
	CartaoCredito   Cartao      `json:"cartao_credito"`
	EnderecoEntrega Endereco    `json:"endereco_entrega"`
	Dispositivo     Dispositivo `json:"dispositivo"`
}
