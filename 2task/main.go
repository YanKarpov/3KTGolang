// ## Задание 2
// Записать в канал 10 числовых значений в потоке основной горутины. 
// Запустить чтение из канала внутри других 10 горутин. 
// Каждая горутина должна преобразовывать числовое значение в строку и складывать во второй канал.
// Основная горутина должна читать из второго канала строки и выводить их на экран 
// в порядке поступления из канала.
// Строчка с запуском горутины в коде должна быть одна (go ...)
// В результате в консоль должно распечататься 10 чисел в строковом представлении.



package main

import (
	"fmt"
	"strconv"  
	"sync"     
)

func main() {
	numChannel := make(chan int, 10)  
	strChannel := make(chan string, 10) 

	var wg sync.WaitGroup  

	// Заполняем numChannel числами от 1 до 10
	for i := 1; i <= 10; i++ {
		numChannel <- i
	}
	close(numChannel)  // Закрываем канал, так как запись завершена

	// Запускаем 10 горутин для обработки чисел из numChannel
	for i := 0; i < 10; i++ {
		wg.Add(1)  // Увеличиваем счётчик ожидания для каждой горутины
		go func() {
			defer wg.Done()  // Уменьшаем счётчик после завершения горутины
			num := <-numChannel        // Читаем число из канала
			str := strconv.Itoa(num)   // Преобразуем число в строку
			strChannel <- str          
		}()
	}

	wg.Wait()
	close(strChannel)  

	for str := range strChannel {
		fmt.Println(str)  
	}
}
