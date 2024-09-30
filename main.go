package main

import (
	"fmt"
	"time"

	"gestaoEscolar/servicos"
)

func main() {
	as := servicos.NovoAlunoServico()
	ps := servicos.NovoProfessorServico()
	ds := servicos.NovaDisciplinaServico(as, ps)

	_, err := as.Cadastrar("Alice", time.Date(2003, time.January, 15, 0, 0, 0, 0, time.UTC), "M001")
	if err != nil {
		fmt.Println("Erro ao cadastrar aluno:", err)
		return
	}

	_, err = as.Cadastrar("Bob", time.Date(2002, time.May, 5, 0, 0, 0, 0, time.UTC), "M002")
	if err != nil {
		fmt.Println("Erro ao cadastrar aluno:", err)
		return
	}

	_, err = ps.Cadastrar("Dr. Smith", "Matemática")
	if err != nil {
		fmt.Println("Erro ao cadastrar professor:", err)
		return
	}

	_, err = ds.Cadastrar("Cálculo I", 1)
	if err != nil {
		fmt.Println("Erro ao cadastrar disciplina:", err)
		return
	}

	err = ds.MatricularAluno(1, 1)
	if err != nil {
		fmt.Println("Erro ao matricular aluno:", err)
		return
	}

	fmt.Println("Alunos:", as.Listar())
	fmt.Println("Professores:", ps.Listar())
	fmt.Println("Disciplinas:", ds.Listar())
}
