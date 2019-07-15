package usuarios

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type StUsersNovoIn struct {
	ID    primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Nome  *string            `bson:"nome" json:"nome"`
	Email *string            `bson:"email" json:"email"`
	Senha *string            `bson:"senha" json:"senha"`
}

func (s *UsuarioST) NovoUnico(u StUsersNovoIn) (string, error) {
	return "Usuario cadastrado com sucesso.", nil
}

func (s *UsuarioST) InserirVarios(u []UsuarioDadosST) (string, error) {
	return "Usuario cadastrado com sucesso.", nil
}
