package utils

import (
	"github.com/google/uuid"

	"gerador-id/models" // usa o nome que apareceu no seu go.mod
)

// GerarID cria um novo ID usando UUID e devolve um models.ID
func GeraRID() models.ID {
	novoID := uuid.New().String() // gera um ID único em formato de texto

	return models.ID{
		Codigo: novoID,
	}
}
