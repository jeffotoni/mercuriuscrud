/*
* Go Library (C) 2017 Inc.
*
* @project     Project Mercurius test
* @author      @jeffotoni
* @size        14/12/2017
 */

package repo

import (
	"os"
	"strings"
)

// criando path absoluto apartir
// de onde encontra-se a executacao
// do script
func GetWdLocal(level int) string {

	pathNow, _ := os.Getwd()

	pathV := strings.Split(pathNow, "/")

	pathNew := pathV[:len(pathV)-level]

	pathNewOrg := strings.Join(pathNew, "/")

	return pathNewOrg
}

// checando se existe o path
func GetExistFile(pathNewOrg string) string {

	_, err := os.Stat(pathNewOrg)

	if err == nil {

		return pathNewOrg

	} else if os.IsNotExist(err) {

		return ""

	} else {

		return ""
	}
}

// criando o path absoluto
func TmplSqlGwd(NameTmpl string, level int) string {

	pathNewOrg := GetWdLocal(level)

	pathNewOrg = pathNewOrg + "/sqltabletmpl/" + NameTmpl

	return GetExistFile(pathNewOrg)
}
