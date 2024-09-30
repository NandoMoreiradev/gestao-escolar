package servicos

import (
	"fmt"

	"gestaoEscolar/entidades"
)

type DisciplinaServico interface {
	Cadastrar(nome string, professorID int) (*entidades.Disciplina, error)
	MatricularAluno(disciplinaID, alunoID int) error
	Listar() []entidades.Disciplina
}

type disciplinaServico struct {
	disciplinas []entidades.Disciplina
	alunos      AlunoServico
	professores ProfessorServico
}

func NovaDisciplinaServico(as AlunoServico, ps ProfessorServico) DisciplinaServico {
	return &disciplinaServico{
		disciplinas: make([]entidades.Disciplina, 0),
		alunos:      as,
		professores: ps,
	}
}

func (ds *disciplinaServico) Cadastrar(nome string, professorID int) (*entidades.Disciplina, error) {
	professor, err := ds.professores.BuscarPorID(professorID)
	if err != nil {
		return nil, err
	}

	disciplina := entidades.Disciplina{
		ID:        len(ds.disciplinas) + 1,
		Nome:      nome,
		Professor: *professor,
		Alunos:    make([]entidades.Aluno, 0),
	}
	ds.disciplinas = append(ds.disciplinas, disciplina)
	return &disciplina, nil
}

func (ds *disciplinaServico) MatricularAluno(disciplinaID, alunoID int) error {
	aluno, err := ds.alunos.BuscarPorID(alunoID)
	if err != nil {
		return err
	}
	for i, disc := range ds.disciplinas {
		if disc.ID == disciplinaID {
			ds.disciplinas[i].Alunos = append(ds.disciplinas[i].Alunos, *aluno)
			return nil
		}
	}
	return fmt.Errorf("disciplina com ID %d não encontrada", disciplinaID)
}

func (ds *disciplinaServico) Listar() []entidades.Disciplina {
	return ds.disciplinas
}
