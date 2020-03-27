# api analise fraude

api 

Para rodar com docker-compose
```sh
docker-compose up
```

exemplo de post

post localhost:80/analise
```json
{
    "valor": 999.99,
    "cliente": {
        "nome": "Glauber Gonçalves",
        "email": "sistemas.glauber@gmail.com",
        "tipo_documento": "cpf",
        "numero_documento": "11122233338",
        "celular": "5521974067912",
        "endereco":{
            "logradouro": "Rua desenvolvedor G",
            "numero": 49,
            "complemento": "casa2",
            "cidade": "São Gonçalo",
            "estado": "rio de janeiro",
            "pais": "brasil",
            "cep": 24750170
        }
    },
    "cartao_credito": {
        "nome_impresso": "glauber goncalves",
			  "numero": "489283428934898234",
        "codigo_seguranca": "123",
        "bandeira": "visa",
        "mes_expiracao": "12",
        "ano_expiracao": "2023"				
    },

    "endereco_entrega":{
        "logradouro": "Rua Desenvolvedor G",
        "numero": 49,
        "complemento": "casa2",
        "cidade": "São Gonçalo",
        "estado": "rio de janeiro",
        "pais": "Brasil",
        "cep": 24750170
    },

    "dispositivo":{
      "ip": "45.169.172.196",
      "origem": "web"
    }
}
```
Exemplo de saida
```json
{
  "ResultadAnalise": false,
  "Motivos": [
    "Dados do cartão são invalidos: Invalid credit card number"
  ]
}
```
