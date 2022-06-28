/*
Паттерн стратегия определяет семейство алгоритмов, инкапсулирует каждый из них
и делает их взаимозаменяемым. Стратегия позволяет изменять алгоритмы независимо от клиентов,
которые ими пользуются.
Применение:
- имеется много родственных классов, отличающихся только поведением
- нужно иметь несколько разных вариантов алгоритма
- в алгоритме содержатся данные, о которых клиент не должен «знать»
- в классе определено много поведений, что представлено разветвленными условными операторами.
  В этом случае проще перенести код из ветвей в отдельные классы стратегий
Плюсы:
 - изолирует код и данные алгоритмов от остальных классов
 - уход от наследования к делегированию
 - реализует принцип открытости/закрытости
 Минусы:
 -  усложняет программу за счёт дополнительных классов
 -  клиент должен знать, в чём состоит разница между стратегиями, чтобы выбрать подходящую.
*/

package main

import "fmt"

type strategySort interface {
	sort([]int)
}

type bubbleSort struct {
}

func (bubbleSort) sort(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] > arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
}

type insertionSort struct {
}

func (insertionSort) sort(arr []int) {
	for i := 0; i < len(arr); i++ {
		for j := i; j >= 1 && arr[j] < arr[j-1]; j-- {
			arr[j], arr[j-1] = arr[j-1], arr[j]
		}
	}
}

type sortContext struct {
	sortMethod strategySort
}

func (s *sortContext) sort(arr []int) {
	s.sortMethod.sort(arr)
}

func main() {
	task := &sortContext{sortMethod: bubbleSort{}}
	arr := []int{5, 4, 3, 1, 6, 9, 12}
	task.sort(arr)
	fmt.Println(arr)

	arr = []int{5, 4, 3, 1, 6, 9, 12}
	task.sortMethod = insertionSort{}
	task.sort(arr)
	fmt.Println(arr)
}
