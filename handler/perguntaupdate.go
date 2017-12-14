/*
* Go Library (C) 2017 Inc.
*
* @project     Project Mercurius test
* @author      @jeffotoni
* @size        14/12/2017
 */

package handler

import (
	"encoding/json"
	// "fmt"
	// "github.com/jeffotoni/mercuriuscrud/conf"
	"github.com/jeffotoni/mercuriuscrud/lib/context"
	"github.com/jeffotoni/mercuriuscrud/model"
	"github.com/jeffotoni/mercuriuscrud/repo"
	bson "gopkg.in/mgo.v2/bson"
	// "io/ioutil"
	"net/http"
	// "strings"
)

// Update na base de dados
func PerguntaUpdate(ctx *context.Context) {

	// collection
	collection := "tpesqperguntas"

	// mensagem json
	var msgJson string

	// struct model data
	var Tp model.TPesqPerguntas

	// chave para atualizacao
	Uuid := ctx.Params(":id")

	if Uuid != "" {

		// capturando json findo da requisicao
		// estamos pegando em bytes para ser
		// usado no Unmarshal que recebe bytes
		byteJson, err := ctx.Req.Body().Bytes()

		// fechando Req.Body
		defer ctx.Req.Body().ReadCloser()

		err = json.Unmarshal(byteJson, &Tp)

		if err != nil {

			msgJson := `{"status":"error","msg":"O json passado esta com error:` + err.Error() + `"}`

			//user is the struct you want to return
			ctx.JSON(http.StatusUnauthorized, msgJson)

		} else {

			// chave para fazer update
			KeyUp := bson.M{"ppr_uuid": Uuid}

			// campos a serem realizado o update
			SetFields := bson.M{"$set": bson.M{"ppr_cod": Tp.Ppr_cod,
				"ppr_ppq_cod": Tp.Ppr_ppq_cod, "ppr_per_cod": Tp.Ppr_per_cod,
				"ppr_ordem": Tp.Ppr_ordem, "ppr_dtcadastro": Tp.Ppr_dtcadastro,
				"ppr_dtaltera": Tp.Ppr_dtaltera,
			}}

			// para atualizacao temos o nome do collection a chave para efetuar o update e
			// os campose que sera feita o set update
			status, err := repo.Update(collection, KeyUp, SetFields)

			// testando se tudo
			// correu bem
			if err == nil && status {

				// Uuid
				msgJson = `{"status":"ok","msg":"Atualizado com sucesso!"}`

				// send write to client
				ctx.JSON(http.StatusOK, msgJson)

			} else {

				msgJson = `{"status":"error","msg":"Algo estranho ocorreu na sua atualizacao, foi usado o id: ` + Uuid + `"}`

				ctx.JSON(http.StatusUnauthorized, msgJson)
			}
		}

	} else {

		msgJson = `{"status":"error","msg":"id é obrigatório!"}`

		ctx.JSON(http.StatusUnauthorized, msgJson)
	}
}
