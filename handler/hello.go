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
	"net/http"
)

func Hello(ctx *context.Context) {

	msgJson := `{"msg":"Hello, Handler Works!"}`
	ctx.JSON(http.StatusOK, msgJson)
}
