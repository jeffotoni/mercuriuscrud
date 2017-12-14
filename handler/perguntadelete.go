/*
* Go Library (C) 2017 Inc.
*
* @project     Project Mercurius test
* @author      @jeffotoni
* @size        14/12/2017
 */

package handler

import (
	// "encoding/json"
	// "fmt"
	// "github.com/jeffotoni/mercuriuscrud/conf"
	"github.com/jeffotoni/mercuriuscrud/lib/context"
	//"github.com/jeffotoni/mercuriuscrud/model"
	"github.com/jeffotoni/mercuriuscrud/repo"
	bson "gopkg.in/mgo.v2/bson"
	// "io/ioutil"
	"net/http"
	// "strings"
)

// removendo registro da base de dados
func PerguntaDelete(ctx *context.Context) {

	// collection
	collection := "tpesqperguntas"

	// mensagem json
	var msgJson string

	// Chave unica
	Uuid := ctx.Params(":id")

	// testando
	// o Uuid
	if Uuid != "" {

		// chamando o metodo para remover o registro do banco
		status, err := repo.Remove(collection, bson.M{"ppr_uuid": Uuid})

		// if tudo correr bem
		// registro foi removido
		// com sucesso
		if err == nil && status {

			// Uuid
			msgJson = `{"status":"ok","msg":"removido com sucesso!"}`

			// send write to client
			ctx.JSON(http.StatusOK, msgJson)

		} else {

			msgJson = `{"status":"error","msg":"Algo estranho ocorreu sua remocao nao foi realizada no id: ` + Uuid + `"}`

			ctx.JSON(http.StatusUnauthorized, msgJson)
		}

	} else {

		msgJson = `{"status":"error","msg":"id é obrigatório!"}`

		ctx.JSON(http.StatusUnauthorized, msgJson)
	}
}
