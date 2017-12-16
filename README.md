# mercuriuscrud

An example of crud using MongoDb and Posgresql and as Framework the Mercurius. The project is an example of how to develop a CRUD using the Mercurius framework. A REST API was built to be consumed.

# Install Mongo and Postgresql

```
Install [https://docs.mongodb.com/v3.4/tutorial/install-mongodb-on-ubuntu/]

Install [https://www.postgresql.org/download/linux/ubuntu/]
```

# Create Database Postgresql

```
createdb mercuriuscrud -U postgres -O postgres -E UTF-8 -T template0
```

# Install Dependencies

```
go get -v gopkg.in/mgo.v2/bson

go get -v github.com/satori/go.uuid

go get -v github.com/jeffotoni/mercuriuscrud
```

# Start App with Run or Compile

```
go run main.go 

go build main.go

./main
```

# Structure of a Project

```
/conf 
Application configuration including environment-specific configs

/conf/app
Middlewares and routes configuration

/handler
HTTP handlers

/locale
Language specific content bundles

/lib
Common libraries to be used across your app

/model
Models

/public
Web resources that are publicly available

/public/templates
Jade templates

/repository
Database comunication following repository pattern

main.go
Application entry
```


# Routes 

```

# test the server

- POST /v1/public/ping


# Data Base using Mongo

- POST /v1/questions 		- Cria uma nova questao

- GET /v1/questions 		- Recupera a lista de questoes

- GET /v1/questions/12 		- Recupera uma questao específica

- PUT /v1/questions/23 		- Atualiza a questao #23

- DELETE /v1/questions/33 	- Deleta a questao #33


# Data Base using Postgresql

- POST /v1/answers 		- Cria uma nova answers

- GET /v1/answers 		- Recupera a lista de answers

- GET /v1/answers/12 		- Recupera uma answers específica

- PUT /v1/answers/23 		- Atualiza a answers #23

- DELETE /v1/answers/33 	- Deleta a answers #33

- POST /v1/answers/tables 	- Cria as tabelas dinamicamente


```

# Example Curl - POST /v1/public/ping

```

curl -X POST localhost:8080/v1/public/ping

```

# Response

A successful request returns pong

```
pong

```

# Example Curl - POST /v1/questions

```

curl -v -X POST localhost:8080/v1/questions \
-H "Content-Type: application/json" \
-d @questionss.json

```
# OR

```

curl -X POST localhost:8080/v1/questions \
-H "Content-Type: application/json" \
-d '{"ppr_cod":100,"ppr_ppq_cod":6,"ppr_per_cod":5,"ppr_ordem":3,"ppr_dtcadastro":"10/07/2017","ppr_dtaltera":"12/08/2017"}'

```

# Sample Response

A successful request returns the HTTP 200 OK status code and a JSON response body

```json

{
	"status'":"ok",
	"msg":"seus dados foram inseridos com sucesso!", 
	"uuid":"238d377b-5db7-45be-ba71-858bed7b4cfc"
}

```

# Example Curl - DELETE /v1/questions/1234

```

curl -X DELETE localhost:8080/v1/questions/1234
```

# Sample Response

```json

{"status":"ok","msg":"removido com sucesso!"}

```

# Example Curl - PUT /v1/questions/1234

```

curl -X PUT localhost:8080/v1/questions/1234
-H "Content-Type: application/json" \
-d @questionss-update.json

```

# Sample Response

```json

{"status":"ok","msg":"Atualizado com sucesso!"}

```

# Example Curl - GET /v1/questions/1235

```

curl -X GET localhost:8080/v1/questions/1235

```

# Sample Response

```json

{
	"status":"ok",
	"msg":"Encontrou o id na base de dados!", 
	"data":
		"{
			"ppr_cod":"234",
			"ppr_dtaltera":"12/03/2017",
			"ppr_dtcadastro":"10/02/2017",
			"ppr_ordem":"120",
			"ppr_per_cod":"1001",
			"ppr_ppq_cod":"1500"
		}"
}

```

# Example Curl - GET /v1/questions

```

curl -X GET localhost:8080/v1/questions

```

# Sample Response

```json

{
	"status":"ok",
	"msg":"Encontrou o id na base de dados!", 
	"data":"
	[{
		"ppr_uid": "9a3e781b-fc03-46d2-96a0-d9c1456a61e6",
		"ppr_cod": 15,
		"ppr_ppq_cod": 5,
		"ppr_per_cod":2 ,
		"ppr_ordem": 19,
		"ppr_dtcadastro": "10/01/2017",
		"ppr_dtaltera": "12/02/2017",
		"ppr_datetime": "2017-12-14 16:15:55"
	}
	,
	
	{
		"ppr_uid":"3a0f8f03-d5c0-4ace-9c0a-91f7f810fdf1",
		"ppr_cod":30,
		"ppr_ppq_cod":32,
		"ppr_per_cod":20,
		"ppr_ordem":24,
		"ppr_dtcadastro":"10/01/2017",
		"ppr_dtaltera":"12/02/2017",
		"ppr_datetime":"2017-12-14 16:37:59"
	}
	,

	{
		"ppr_uid":"6d9090a3-6b14-498f-8d83-0ae6b2401ab1",
		"ppr_cod":123,
		"ppr_ppq_cod":132,
		"ppr_per_cod":3050,
		"ppr_ordem":24,
		"ppr_dtcadastro":"10/01/2017",
		"ppr_dtaltera":"12/02/2017",
		"ppr_datetime":"2017-12-14 16:38:09"
	}
	,

	{
		"ppr_uid":"723ad212-f309-42e3-bcc0-dc3720827e0e",
		"ppr_cod":234,
		"ppr_ppq_cod":1500,
		"ppr_per_cod":1001,
		"ppr_ordem":120,
		"ppr_dtcadastro":"10/02/2017",
		"ppr_dtaltera":"12/03/2017",
		"ppr_datetime":"2017-12-14 16:38:18"
	}]"
}

```

# Example Curl  - POST /v1/answers/tables

```

curl -X POST localhost:8080/v1/answers/tables

```

# Example Curl - POST /v1/answers

```

curl -X POST localhost:8080/v1/answers \
-H 'Content-Type: application/json'
-d @answers.json

```
