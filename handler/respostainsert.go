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
	//"strconv"
	// "fmt"
	"github.com/jeffotoni/mercuriuscrud/lib/context"
	"github.com/jeffotoni/mercuriuscrud/model"
	"github.com/jeffotoni/mercuriuscrud/repo"
	"net/http"
	"strings"
)

// inserindo perguntas na base de dados
func RepostaInsert(ctx *context.Context) {

	// collection
	table := "trespostas"

	// mensagem json
	var msgJson string

	cTypeAceito := "application/json"

	cType := ctx.Req.Header.Get("Content-Type")

	if strings.ToLower(strings.TrimSpace(cType)) == cTypeAceito {

		// capturando json findo da requisicao
		// estamos pegando em bytes para ser
		// usado no Unmarshal que recebe bytes
		byteJson, err := ctx.Req.Body().Bytes()

		// fechando Req.Body
		defer ctx.Req.Body().ReadCloser()

		// caso ocorra
		// erro ao ler
		// envia mensagem
		// de error
		if err != nil {

			msgJson := `{"status":"error","msg":"Nao foi possivel ler seu json error: " ` + err.Error() + `}`

			ctx.JSON(http.StatusUnauthorized, msgJson)

		} else {

			// se nao tiver nem um valor
			// no json, emite um erro
			// para o usuario
			if string(byteJson) != "" {

				// struct model data
				var Tr model.TRespostas

				// convert json to struct
				err := json.Unmarshal(byteJson, &Tr)

				// caso tenha
				// algum erro
				// enviar para
				// o cliente
				if err != nil {

					msgJson := `{"status":"error","msg":"O json passado esta com error:` + err.Error() + `"}`
					//user is the struct you want to return
					ctx.JSON(http.StatusUnauthorized, msgJson)

				} else {

					// campo obrigatorio string
					if Tr.Rsp_dtcadastro == "" {

						msgJson = `{"status":"error","msg":"A coluna Rsp_dtcadastro obrigatoria!"}`

					} else {

						// validar se existe o dado no banco
						// if exist o dado na base de dados
						exist := repo.GetOne(table, "rsp_cod", Tr.Rsp_cod)

						// caso
						// exista error
						// enviar para o cliente
						//if exist == 0  {

						//msgJson = `{"status":"error","msg":"GetOne Error: ` + err.Error() + `"}`

						//} else {

						// se existe
						// nao faca
						// o insert
						if exist > 0 {

							msgJson = `{"status":"error",msg":"estes dados ja existe na base de dados!"}`

						} else {

							// faca o insert no banco de dados
							// retorne o uuid do salvar
							// caso tenha erro informe ao cliente
							ID, err := repo.InsertResposta(Tr)

							if err != nil {

								msgJson = `{"status":"error","msg":"ocorreu algum erro ao salvar seus dados error: ` + err.Error() + `"}`

							} else {

								msgJson = `{"status":"ok","msg":"seus dados foram inseridos com sucesso!", "id":"` + ID + `"}`
							}
						}
						//}
					}

					// send write to client
					ctx.JSON(http.StatusOK, msgJson)
				}

			} else {

				msgJson := `{"status":"error","msg":"O json passado esta com error"}`

				// send write to client
				ctx.JSON(http.StatusUnauthorized, msgJson)
			}
		}

	} else {

		msgJson := `{"status":"error","msg":"error no Content-Type: ` + cType + `, aceitamos somente [Content-Type: ` + cTypeAceito + `]"}`

		// send write to client
		ctx.JSON(http.StatusUnauthorized, msgJson)
	}
}
