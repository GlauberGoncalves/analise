package models

type Cartao struct {
	NomeImpresso    string `json:"numero_impresso"`
	Numero          string `json:"numero"`
	Bandeira        string `json:"bandeira"`
	CodigoSeguranca uint32 `json:"codigo_seguranca"`
	MesExpiracao    uint32 `json:"mes_expiracao"`
	AnoExpiracao    uint32 `json:"ano_expiracao"`
	Titular         string `json:"titular"`
}
