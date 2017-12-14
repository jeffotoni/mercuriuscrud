# mercuriuscrud


go get -v gopkg.in/mgo.v2/bson

go get -v https://github.com/satori/go.uuid

go get -v github.com/jeffotoni/mercuriuscrud


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