package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	target := 10
	fmt.Println(twoSum(nums, target))
}

func twoSum(nums []int, target int) []int {
	hm := make(map[int]int)
	for i, num := range nums {
		_, ok := hm[num]
		if ok {
			return []int{i, hm[num]}
		}
		hm[target-num] = i
	}
	return nil
}

//Preciso saber a soma de dois números que vão retornar o target
// Vou criar um hashmap de inteiro com inteiro
//Vou criar um for para iterar os números
//A cada iteração vou checar se já existe uma chave com o número anterior
//Se já existir vou retornar a chave e o conteúdo
//Se não existir, vou adicionar o valor atual como valor e chave vai ser o target - valor atual
