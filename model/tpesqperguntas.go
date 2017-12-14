/*
* Go Library (C) 2017 Inc.
*
* @project     Project Mercurius test
* @author      @jeffotoni
* @size        14/12/2017
 */

package model

// struct TPesqPerguntas
type TPesqPerguntas struct {
	Ppr_uuid       string `json:"ppr_uid"`
	Ppr_cod        int    `json:"ppr_cod"`
	Ppr_ppq_cod    int    `json:"ppr_ppq_cod"`
	Ppr_per_cod    int    `json:"ppr_per_cod"`
	Ppr_ordem      int    `json:"ppr_ordem"`
	Ppr_dtcadastro string `json:"ppr_dtcadastro"`
	Ppr_dtaltera   string `json:"ppr_dtaltera"`
	Ppr_datetime   string `json:"ppr_datetime"`
}
