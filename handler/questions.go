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

// inserindo perguntas na base de dados
func QuestionsCreate(ctx *context.Context) {

	// bytes json body
	var byteJson []byte

	// define
	// var error
	var err error

	// mensagem json
	var msgJson string

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
			log.Println("[QuestionsCreate] Erro ao capturar Json: " + err.Error())
			msgJson := `{"status":"error","msg":"Nao foi possivel ler seu json error: " ` + err.Error() + `}`
			ctx.JSON(http.StatusUnauthorized, msgJson)
			return
		} else {

			// se nao tiver nem um valor
			// no json, emite um erro
			// para o usuario
			if string(byteJson) != "" {

				// Criar registro no banco de dados
				Uuid, err := repo.AddQuestion(byteJson)

				// tratando o erro
				if err != nil {
					log.Println("[QuestionsCreate] Erro ao tentar inserir registro: " + err.Error())
					msgJson := `{"status":"error","msg":"Erro ao tentar inserir registro"}`
					// send write to client
					ctx.JSON(http.StatusUnauthorized, msgJson)
					return
				} else {
					// sucesso
					msgJson = `{"status":"ok","msg":"seus dados foram inseridos com sucesso!", "Uuid":"` + Uuid + `"}`
					// send write to client
					ctx.JSON(http.StatusOK, msgJson)
				}
			} else {
				log.Println("[QuestionsCreate] Erro em sua string json nao pode ser vazia!")
				msgJson := `{"status":"error","msg":"Erro em sua string json"}`
				// send write to client
				ctx.JSON(http.StatusUnauthorized, msgJson)
				return
			}
		} // fim else
	} else {
		log.Println("[QuestionsCreate] Erro Content-Type: aceitamos somente " + cTypeAceito)
		msgJson := `{"status":"error","msg":"error no Content-Type: ` + cType + `, aceitamos somente [Content-Type: ` + cTypeAceito + `]"}`
		// send write to client
		ctx.JSON(http.StatusUnauthorized, msgJson)
		return
	}
}
