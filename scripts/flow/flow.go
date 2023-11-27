package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)


type Step struct {
	Name          string
	IsRepeatable  bool
	NextSteps     map[string]*Step
}

stateMap := make([string]*Step) 

func (s *Step) Create(name string, isRepeatable bool) *Step {
	return &Step{
		Name:          name,
		IsRepeatable:  isRepeatable,
		NextSteps:     make(map[string]*Step),
	}
}

// AddNextSteps добавляет ссылки на допустимые следующие шаги
func (s *Step) AddNextSteps(steps ...*Step) {
	for _, step := range steps {
		s.NextSteps[step.Name] = step
	}
}


func main() {
	// Пример использования структуры Step

	// Создаем шаги
	step1 := new(Step).Create("Step1", false)
	step2 := new(Step).Create("Step2", true)
	step3 := new(Step).Create("Step3", false)

	// Добавляем допустимые следующие шаги
	step1.AddNextSteps(step2, step3)
	step2.AddNextSteps(step3)

	// Получаем шаг по имени
	targetStep := step1.GetStep("Step2")

	if targetStep != nil {
		fmt.Printf("Найден шаг: %s\n", targetStep.Name)
	} else {
		fmt.Pri


	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Введите число (или 'exit' для выхода): ")
		scanner.Scan()
		input := scanner.Text()

		if input == "exit" {
			fmt.Println("Программа завершена.")
			break
		}

		num, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Ошибка при чтении числа. Попробуйте снова.")
			continue
		}

		square := num * num
		fmt.Printf("Квадрат введенного числа: %d\n", square)
	}
}
