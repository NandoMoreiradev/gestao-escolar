package servicos

import (
	"fmt"

	"gestaoEscolar/entidades"
)

type ProfessorServico interface {
	Cadastrar(nome, materia string) (*entidades.Professor, error)
	BuscarPorID(id int) (*entidades.Professor, error)
	Listar() []entidades.Professor
}

type professorServico struct {
	professores []entidades.Professor
}

func NovoProfessorServico() ProfessorServico {
	return &professorServico{
		professores: make([]entidades.Professor, 0),
	}
}

func (ps *professorServico) Cadastrar(nome, materia string) (*entidades.Professor, error) {
	professor := entidades.Professor{
		ID:      len(ps.professores) + 1,
		Nome:    nome,
		Materia: materia,
	}
	ps.professores = append(ps.professores, professor)
	return &professor, nil
}

func (ps *professorServico) BuscarPorID(id int) (*entidades.Professor, error) {
	for _, p := range ps.professores {
		if p.ID == id {
			return &p, nil
		}
	}
	return nil, fmt.Errorf("professor com ID %d não encontrado", id)
}

func (ps *professorServico) Listar() []entidades.Professor {
	return ps.professores
}
