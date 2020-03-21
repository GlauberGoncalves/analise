package models

type Cliente struct {
	ID                uint32   `json:"id"`
	Nome              string   `json:"nome"`
	Email             string   `json:"email"`
	TipoDocumento     string   `json:"tipo_documento"`
	NumeroDocumento   string   `json:"numero_documento"`
	Celular           string   `json:"celular"`
	EnderecoPrincipal Endereco `json:"endereco"`
}
