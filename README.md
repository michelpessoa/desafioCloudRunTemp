0# Go Expert - Desfio Clean Architecture
# cloudRunDesafioGoExpert

Go | REST | Wire | Docker  

# Sistema em GO de temperatura por CEP

 O sistema tem como entrada o CEP de uma cidade (somente números), consulta a cidade, e após saber a cidade consulta a temperatura, retornando a mesma em Celsius, Fahrenheit, e Kelvin.


### Variável obrigatória de preenchimento

O sistema consome a API ViaCEP e WeatherAPI, esta última necessita de uma API necessita de uma chave WEATHER_API_KEY no arquivo .env, que deve ser preenchida para o funcionamento. 
É possível obte-lá atráves do cadastro em (https://www.weatherapi.com)

### Para executar o projeto localmente siga os seguintes passos:


1. `git clone https://https://github.com/michelpessoa/desafioCloudRun`
2. `go mod tidy` para instalar todas as dependências
3. `make run` execute o comando para iniciar os serviços 


### Utilize algum client http para realizar as seguintes chamadas abaixo

#### 200

curl -X GET http://localhost:8080?cep=74740250

#### 404

curl -X GET http://localhost:8080?cep=74000000

#### 422

curl -X GET http://localhost:8080?cep=74000

### Cloud run

A apliação está executando em no GCP pelo seguinte endereço
`https://desafio-cloud-run-e4qa7deoaq-uc.a.run.app`

#### 200

curl -X GET https://desafio-cloud-run-e4qa7deoaq-uc.a.run.app?cep=74740250

#### 404

curl -X GET https://desafio-cloud-run-e4qa7deoaq-uc.a.run.app?cep=74000000

#### 422

curl -X GET https://desafio-cloud-run-e4qa7deoaq-uc.a.run.app?cep=74000

#### Responses

- In case of success:

  - HTTP Code: 200
  - Response Body: `{ "temp_C": 28.5, "temp_F": 28.5, "temp_K": 28.5 }`

- In case of failure, if the ZIP code is not valid (with correct format):

  - HTTP Code: 422
  - Message: `invalid zipcode`

- In case of failure, if the ZIP code is not found:
  - HTTP Code: 404
  - Message: `zipcode can not found`
