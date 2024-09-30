package servicos

import (
	"fmt"
	"time"

	"gestaoEscolar/entidades"
)

type AlunoServico interface {
	Cadastrar(nome string, dataNasc time.Time, matricula string) (*entidades.Aluno, error)
	BuscarPorID(id int) (*entidades.Aluno, error)
	Listar() []entidades.Aluno
}

type alunoServico struct {
	alunos []entidades.Aluno
}

func NovoAlunoServico() AlunoServico {
	return &alunoServico{
		alunos: make([]entidades.Aluno, 0),
	}
}

func (as *alunoServico) Cadastrar(nome string, dataNasc time.Time, matricula string) (*entidades.Aluno, error) {
	aluno := entidades.Aluno{
		ID:        len(as.alunos) + 1,
		Nome:      nome,
		DataNasc:  dataNasc,
		Matricula: matricula,
	}
	as.alunos = append(as.alunos, aluno)
	return &aluno, nil
}

func (as *alunoServico) BuscarPorID(id int) (*entidades.Aluno, error) {
	for _, a := range as.alunos {
		if a.ID == id {
			return &a, nil
		}
	}
	return nil, fmt.Errorf("aluno com ID %d não encontrado", id)
}

func (as *alunoServico) Listar() []entidades.Aluno {
	return as.alunos
}
