/*
* Go Library (C) 2017 Inc.
*
* @project     Project Mercurius test
* @author      @jeffotoni
* @size        14/12/2017
 */

package handler

import (
	// "github.com/jeffotoni/mercuriuscrud/conf"
	"github.com/jeffotoni/mercuriuscrud/lib/context"
	// "github.com/jeffotoni/mercuriuscrud/model"
	"github.com/jeffotoni/mercuriuscrud/repo"
	// bson "gopkg.in/mgo.v2/bson"
	"net/http"
	// "strings"
)

// Update na base de dados
func PerguntaFind(ctx *context.Context) {

	// collection
	collection := "tpesqperguntas"

	// mensagem json
	var msgJson string

	// chave para atualizacao
	Uuid := ctx.Params(":id")

	if Uuid != "" {

		// para atualizacao temos o nome do collection a chave para efetuar o update e
		// os campose que sera feita o set update
		strJson, err := repo.Find(collection, Uuid)

		// testando se tudo
		// correu bem
		if err == nil && strJson != "" {

			// Uuid
			msgJson = `{"status":"ok","msg":"Encontrou o id na base de dados!", "data":"` + strJson + `"}`

			// send write to client
			ctx.JSON(http.StatusOK, msgJson)

		} else {

			msgJson = `{"status":"error","msg":"Algo estranho ocorreu em sua busca usamos o id: ` + Uuid + `"}`

			ctx.JSON(http.StatusUnauthorized, msgJson)
		}

	} else {

		msgJson = `{"status":"error","msg":"id é obrigatório!"}`

		ctx.JSON(http.StatusUnauthorized, msgJson)
	}
}
