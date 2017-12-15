/*
* Go Library (C) 2017 Inc.
*
* @project     Project Mercurius test
* @author      @jeffotoni
* @size        14/12/2017
 */

package repo

import (
	"bytes"
	"html/template"
	"log"
)

type TemplateSql struct{}

// retornando o sql na pasta template do sql
func SqlTmplString(nameTemplate string) (str string, err error) {

	// struct do template
	var tmpSql TemplateSql

	// buffer do template
	var tpl bytes.Buffer

	// criando referencia e abrindo o file no diretorio correspondente onde encontra-se
	// o template
	t := template.Must(template.New(nameTemplate).ParseFiles(TmplSqlGwd(nameTemplate, 0)))

	if err = t.Execute(&tpl, tmpSql); err != nil {

		log.Println("[SqlTmplString] Erro ao criar instancia para tempplate: " + err.Error())
		return
	}

	str = tpl.String()
	return
}
