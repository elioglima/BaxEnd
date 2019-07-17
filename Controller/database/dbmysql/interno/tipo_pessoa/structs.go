package tipo_pessoa

import (
	"errors"
	"strings"
)

type TipoPessoaDadosST struct {
	Id        int64
	Descricao string
}

type TipoPessoaST struct {
	RecordCount int
	Rows        []TipoPessoaDadosST
	Field       *TipoPessoaDadosST
	Fields      *[]TipoPessoaDadosST
}

func New() TipoPessoaST {
	s := TipoPessoaST{}
	s.Ini()
	return s
}

func (s *TipoPessoaST) Ini() {
	if s.RecordCount == 0 {
		s.AddRows(0, "Pessoa Física")
		s.AddRows(1, "Pessoa Jurídica")
	}
}

func (s *TipoPessoaST) PesquisaID(id int64) error {
	s.Clear()
	for _, tipo := range s.Rows {
		if tipo.Id == id {
			s.Add(tipo)
			return nil
		}
	}
	return errors.New("Registro não localizado")
}

func (s *TipoPessoaST) PesquisaDesc(descricao string) error {
	s.Clear()
	for _, tipo := range s.Rows {
		if strings.Contains(tipo.Descricao, descricao) {
			s.Add(tipo)
		}
	}
	if s.RecordCount >= 0 {
		return nil
	}
	return errors.New("Registro não localizado")
}

func (s *TipoPessoaST) AddRows(id int64, descricao string) {
	temp := TipoPessoaDadosST{}
	temp.Id = id
	temp.Descricao = descricao
	s.Rows = append(s.Rows, temp)
}

func (s *TipoPessoaST) Clear() {
	s.Field = &TipoPessoaDadosST{}
	s.Fields = &[]TipoPessoaDadosST{}
	s.RecordCount = 0
}

func (s *TipoPessoaST) Add(field TipoPessoaDadosST) {
	s.Field = &field
	*s.Fields = append(*s.Fields, field)
	s.RecordCount = len(*s.Fields)
}
