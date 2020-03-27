package models

type Endereco struct {
	Logradouro  string `json:"logradouro"`
	Numero      uint32 `json:"numero"`
	Complemento string `json:"complemento"`
	Cidade      string `json:"cidade"`
	Estado      string `json:"cidade"`
	Pais      	string `json:"pais"`
	Cep         uint32 `json:cep`
}
