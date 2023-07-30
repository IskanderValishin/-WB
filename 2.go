package main

import (
	"fmt"
	"os"
	"runtime"
	"sync"
)

func sqrt1(num int) {
	fmt.Fprintf(os.Stdout, "Квадрат %v равно %v \n", num, num*num)
}
func Square() {
	nums := [5]int{2, 4, 6, 8, 10}
	for _, num := range nums {
		go sqrt1(num)
		runtime.Gosched()
		//fmt.Printf("Квадрат %v равно: %v \n", num, num*num)
		//num := num * num
		//io.WriteString(os.Stdout, strconv.Itoa(num))
		//fmt.Fprintf(os.Stdout, "Квадрат %v равно %v \n", num, num*num)
	}

}

func main() {
	var wg sync.WaitGroup // создание очереди ожидания из пакета синк
	//wg.add(5)
	nums := [5]int{2, 4, 6, 8, 10} // Создание массива чисел 2,4,6,8,10
	for _, num := range nums {     // Цикл с использованием range длины масива
		wg.Add(1)          // добавление задачи по количеству итераций в очердь ожидания
		go func(num int) { // горутина для расчета квадратов
			fmt.Fprintf(os.Stdout, "Квадрат %v равно %v \n", num, num*num) //запись значений квадратов в Stdout
			wg.Done()                                                      // оконачание задачи
		}(num)

	}
	wg.Wait() // ожидание когда все задачи выполняться
	fmt.Println("Расчёт закончен.")
}
