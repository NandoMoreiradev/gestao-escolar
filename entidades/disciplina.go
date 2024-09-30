package entidades

type Disciplina struct {
	ID        int
	Nome      string
	Professor Professor
	Alunos    []Aluno
}
