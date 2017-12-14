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

http://localhost:8080/v1/public/ping

http://localhost:8080/v1/pergunta/insert

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
-d @pergunta.json
```

# return

```

'{"status'":"error|ok","msg":"aqui estara a mensagem de error ou sucesso", "id":"aqui ira conter o id"}' 
```