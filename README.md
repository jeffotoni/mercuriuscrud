# mercuriuscrud

Um Crud Usando MongoDb e Mercurius como Framework ...

# Install Mongo

Install [https://docs.mongodb.com/v3.4/tutorial/install-mongodb-on-ubuntu/]

go get -v gopkg.in/mgo.v2/bson

go get -v github.com/satori/go.uuid

go get -v github.com/jeffotoni/mercuriuscrud

# Start app
go run main.go 

# Compila app
go build main.go

# Start app compilado
./main


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

// testando o server
http://localhost:8080/v1/public/ping

// inserindo perguntas na base
http://localhost:8080/v1/pergunta/insert

// deletando a pergunta com seu id
http://localhost:8080/v1/pergunta/delete/id

// Atualizando a pergunta com seu id
http://localhost:8080/v1/pergunta/update/id

// Buscando registro especifico
http://localhost:8080/v1/pergunta/find/id

// Buscando todos registros
http://localhost:8080/v1/pergunta/findall

```

# Example Curl Ping

```

curl -X POST localhost:8080/v1/public/ping
```

# return

```
pong
```

# Example Curl Width Json :: pergunta/insert 

```

curl -X POST localhost:8080/v1/pergunta/insert \
-H "Content-Type: application/json" \
-d @perguntas.json
```
# OR

```

curl -X POST localhost:8080/v1/pergunta/insert \
-H "Content-Type: application/json" \
-d '{"ppr_cod":100,"ppr_ppq_cod":6,"ppr_per_cod":5,"ppr_ordem":3,"ppr_dtcadastro":"10/07/2017","ppr_dtaltera":"12/08/2017"}'
```

# return

```

'{"status'":"error|ok","msg":"aqui estara a mensagem de error ou sucesso", "id":"aqui ira conter o id"}' 
```

# Example Curl Width Json :: pergunta/delete/id

```

curl -X DELETE localhost:8080/v1/pergunta/delete/id
```

# return

```

'{"status":"ok","msg":"removido com sucesso!"}'
```

# Example Curl Width Json :: pergunta/update/id

```

curl -X DELETE localhost:8080/v1/pergunta/update/id
-H "Content-Type: application/json" \
-d @perguntas-update.json
```

# return

```

'{"status":"ok","msg":"Atualizado com sucesso!"}'
```

# Example Curl :: pergunta/find/id

```

curl -X GET localhost:8080/v1/pergunta/find/id
```

# return

```

'{
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
}'
```

# Example Curl :: pergunta/findall

```

curl -X GET localhost:8080/v1/pergunta/findall
```

# return

```

'{
	"status":"ok",
	"msg":"Encontrou o id na base de dados!", 
	"data":"
	[{
		"ppr_uid":"9a3e781b-fc03-46d2-96a0-d9c1456a61e6",
		"ppr_cod":15,
		"ppr_ppq_cod":5,
		"ppr_per_cod":2,
		"ppr_ordem":19,
		"ppr_dtcadastro":"10/01/2017",
		"ppr_dtaltera":"12/02/2017",
		"ppr_datetime":"2017-12-14 16:15:55"
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
}'