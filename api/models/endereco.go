package models

type Endereco struct {
	Logradouro  string `json:"logradouro"`
	Numero      uint32 `json:"numero"`
	Complemento string `json:"complemento"`
	Cidade      string `json:"cidade"`
	estado      string `json:"cidade"`
	cep         string `json:cep`
}
