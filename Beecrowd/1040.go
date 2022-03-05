package main

import "fmt"

const alunoAprovado = "aprovado."
const alunoReprovado = "reprovado."
const alunoEmExame = "em exame."

func checkStudentResult(average float64) (resultado string) {
	if average >= 7.0 {
		resultado = alunoAprovado
	} else if average < 5.0 {
		resultado = alunoReprovado
	} else {
		resultado = alunoEmExame
	}
	return
}

func checkStudentFinalResult(average float64) (resultado string) {
	if average >= 5.0 {
		resultado = alunoAprovado
	} else {
		resultado = alunoReprovado
	}
	return
}

func main() {
	var N1, N2, N3, N4, exame float64
	_, _ = fmt.Scanf("%f %f %f %f", &N1, &N2, &N3, &N4)
	average := (N1*2 + N2*3 + N3*4 + N4*1) / 10.0

	if checkStudentResult(average) == alunoEmExame {
		_, _ = fmt.Scanf("%f", &exame)
		fmt.Printf("Media: %.1f\n", average)
		fmt.Printf("Aluno %s\n", checkStudentResult(average))
		fmt.Printf("Nota do exame: %.1f\n", exame)
		mediaFinal := (average + exame) / 2.0
		fmt.Printf("Aluno %s\n", checkStudentFinalResult(mediaFinal))
		fmt.Printf("Media final: %.1f\n", mediaFinal)
	} else {
		fmt.Printf("Media: %.1f\n", average)
		fmt.Printf("Aluno %s\n", checkStudentFinalResult(average))
	}
}
