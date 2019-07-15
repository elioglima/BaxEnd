package usuarios

type StUsersAlteraUnicoIn struct {
	ID    string  `bson:"id" json:"id"`
	Nome  *string `bson:"nome" json:"nome"`
	Email *string `bson:"email" json:"email"`
}

func (s *UsuarioST) AlteraUnico(u StUsersAlteraUnicoIn) (string, error) {

	smsg := "Usuario cadastrado com sucesso."
	return smsg, nil
}
