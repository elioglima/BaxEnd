package views

import (
	"html/template"
)

type stEstrutura struct {
	Data        string
	Versao      string
	MenuLateral template.HTML
	Rodape      template.HTML
	CodeStatus  int
	Mensagem    template.HTML
	Hash        string
}
