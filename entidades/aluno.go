package entidades

import "time"

type Aluno struct {
	ID        int
	Nome      string
	DataNasc  time.Time
	Matricula string
}
