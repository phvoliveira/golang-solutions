package main

import "fmt"

func main() {
    output := [][]int{}
    numbers := []int{2, 7, 6, 5, 12, 4}
    value := 24
    auxMap := map[int]bool{
        int(numbers[0]): true,
    }

    for idx := 1; idx < len(numbers); idx++ {
        loopValue := int(numbers[idx])

        if value%loopValue == 0 {
            num := value / loopValue
            if auxMap[num] {
                output = append(output, []int{num, loopValue})
            }
        }

        auxMap[loopValue] = true
    }

    fmt.Println(output)
}