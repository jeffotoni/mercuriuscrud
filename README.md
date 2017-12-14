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