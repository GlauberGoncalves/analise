package models

type Cartao struct {
	NomeImpresso    string `json:"numero_impresso"`
	Numero          string `json:"numero"`
	Bandeira        string `json:"bandeira"`
	CodigoSeguranca string `json:"codigo_seguranca"`
	MesExpiracao    string `json:"mes_expiracao"`
	AnoExpiracao    string `json:"ano_expiracao"`
	Titular         string `json:"titular"`
}
