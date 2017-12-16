/*
* Go Library (C) 2017 Inc.
*
* @project     Project Mercurius test
* @author      @jeffotoni
* @size        14/12/2017
 */

package handler

import (
	"github.com/jeffotoni/mercuriuscrud/lib/context"
	"github.com/jeffotoni/mercuriuscrud/repo"
	"log"
	"net/http"
	"strings"
)

var msgerror string

// inserindo perguntas na base de dados
func QuestionsCreate(ctx *context.Context) {

	// bytes json body
	var byteJson []byte

	// define
	// var error
	var err error

	// mensagem json
	var msgJson, Uuid string

	// tipo aceito de content-type
	cTypeAceito := "application/json"

	// capturando content-type
	cType := ctx.Req.Header.Get("Content-Type")

	// validando content-type
	if strings.ToLower(strings.TrimSpace(cType)) == cTypeAceito {

		// capturando json findo da requisicao
		// estamos pegando em bytes para ser
		// usado no Unmarshal que recebe bytes
		byteJson, err = ctx.Req.Body().Bytes()

		// fechando Req.Body
		defer ctx.Req.Body().ReadCloser()

		// caso ocorra
		// erro ao ler
		// envia mensagem
		// de error
		if err != nil {
			msgerror = "[QuestionsCreate] Erro ao capturar Json: " + err.Error()
			log.Println(msgerror)
			msgJson = `{"status":"error","msg":"` + msgerror + `}`
			ctx.JSON(http.StatusUnauthorized, msgJson)
			return
		} else {

			// se nao tiver nem um valor
			// no json, emite um erro
			// para o usuario
			if string(byteJson) != "" {

				// Criar registro no banco de dados
				Uuid, err = repo.AddQuestion(byteJson)

				// tratando o erro
				if err != nil {
					log.Println(err.Error())
					msgJson = `{"status":"error","msg":"` + err.Error() + `"}`
					ctx.JSON(http.StatusUnauthorized, msgJson)
					return
				} else {
					// sucesso
					msgJson = `{"status":"ok","msg":"seus dados foram inseridos com sucesso!", "uuid":"` + Uuid + `"}`
					ctx.JSON(http.StatusOK, msgJson)
				}
			} else {
				log.Println("[QuestionsCreate] Erro em sua string json nao pode ser vazia!")
				msgJson = `{"status":"error","msg":"Erro em sua string json"}`
				ctx.JSON(http.StatusUnauthorized, msgJson)
				return
			}
		} // fim else
	} else {
		log.Println("[QuestionsCreate] Erro Content-Type: aceitamos somente " + cTypeAceito)
		msgJson = `{"status":"error","msg":"error no Content-Type: ` + cType + `, aceitamos somente [Content-Type: ` + cTypeAceito + `]"}`
		ctx.JSON(http.StatusUnauthorized, msgJson)
		return
	}
}

// removendo registro da base de dados
func QuestionsDelete(ctx *context.Context) {

	// mensagem json
	var msgJson string

	// Chave unica
	Uuid := ctx.Params(":id")

	// testando
	// o Uuid
	if Uuid != "" {

		// del question
		err := repo.DelQuestion(Uuid)

		// if tudo correr bem
		// registro foi removido
		// com sucesso
		if err == nil {

			// Uuid
			msgJson = `{"status":"ok","msg":"removido com sucesso seu Uuid: ` + Uuid + `!"}`
			ctx.JSON(http.StatusOK, msgJson)

		} else {
			msgerror = `[QuestionsDelete] Algo estranho ocorreu em sua remocao Uuid: ` + Uuid
			log.Println(msgerror)
			msgJson = `{"status":"error","msg":"` + msgerror + `]"}`
			ctx.JSON(http.StatusUnauthorized, msgJson)
			return
		}

	} else {
		msgerror = "[QuestionsDelete] Uuid é obrigatório!"
		log.Println(msgerror)
		msgJson = `{"status":"error","msg":"` + msgerror + `]"}`
		ctx.JSON(http.StatusUnauthorized, msgJson)
		return
	}
}

// Update na base de dados
func QuestionsUpdate(ctx *context.Context) {

	// define err
	var err error

	// mensagem json
	var msgJson string

	// byjson
	var byteJson []byte

	// Chave unica
	Uuid := ctx.Params(":id")

	// testando
	// o Uuid
	if Uuid != "" {

		// capturando json findo da requisicao
		// estamos pegando em bytes para ser
		// usado no Unmarshal que recebe bytes
		byteJson, err = ctx.Req.Body().Bytes()

		if err != nil {
			msgJson = `{"status":"error","msg":"` + err.Error() + `"}`
			ctx.JSON(http.StatusUnauthorized, msgJson)
			return
		}

		// fechando Req.Body
		defer ctx.Req.Body().ReadCloser()

		// del question
		err := repo.UpQuestion(Uuid, byteJson)

		// if tudo correr bem
		// registro foi atualizado
		// com sucesso
		if err == nil {
			// Uuid
			msgJson = `{"status":"ok","msg":"atualizado com sucesso seu Uuid: ` + Uuid + `!"}`
			ctx.JSON(http.StatusOK, msgJson)

		} else {

			log.Println(err.Error())
			msgJson = `{"status":"error","msg":"` + err.Error() + `]"}`
			ctx.JSON(http.StatusUnauthorized, msgJson)
			return
		}

	} else {
		msgerror = "[QuestionsUpdate] Uuid é obrigatório!"
		log.Println(msgerror)
		msgJson = `{"status":"error","msg":"` + msgerror + `]"}`
		ctx.JSON(http.StatusUnauthorized, msgJson)
		return
	}
}

// Busca Questions especifico na base de dados
func QuestionsFind(ctx *context.Context) {

	// mensagem json
	var msgJson string

	// chave para atualizacao
	Uuid := ctx.Params(":id")

	if Uuid != "" {

		// para atualizacao temos o nome do collection a chave para efetuar o update e
		// os campose que sera feita o set update
		strJson, err := repo.GetQuestion(Uuid)

		// testando se tudo
		// correu bem
		if err == nil {

			// Uuid
			msgJson = `{"status":"ok","msg":"Encontrou o id na base de dados!", "data":"` + strJson + `"}`
			// send write to client
			ctx.JSON(http.StatusOK, msgJson)

		} else {
			msgerror = "[QuestionsFind] " + err.Error()
			log.Println(msgerror)
			msgJson = `{"status":"error","msg":"` + msgerror + `]"}`
			ctx.JSON(http.StatusUnauthorized, msgJson)
			return
		}

	} else {

		msgerror = "[QuestionsFind] Uuid é obrigatorio!"
		log.Println(msgerror)
		msgJson = `{"status":"error","msg":"` + msgerror + `]"}`
		ctx.JSON(http.StatusUnauthorized, msgJson)
		return
	}
}

// Busca Questions especifico na base de dados
func QuestionsFindAll(ctx *context.Context) {

	// mensagem json
	var msgJson string

	// para atualizacao temos o nome do collection a chave para efetuar o update e
	// os campose que sera feita o set update
	strJson, err := repo.GetAllQuestion()

	// testando se tudo
	// correu bem
	if err == nil {

		// Uuid
		msgJson = `{"status":"ok","msg":"encontrado com sucesso", "data":"` + strJson + `"}`
		// send write to client
		ctx.JSON(http.StatusOK, msgJson)

	} else {
		msgerror = "[QuestionsFind] " + err.Error()
		log.Println(msgerror)
		msgJson = `{"status":"error","msg":"` + msgerror + `]"}`
		ctx.JSON(http.StatusUnauthorized, msgJson)
		return
	}
}
