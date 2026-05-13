package main

import "fmt"

func main() {
	numbers := [3]int{10, 20, 30}
	fmt.Println("array:", numbers)
	fmt.Println("first number:", numbers[0])

	names := []string{"Alice", "Bob", "Carol"}
	fmt.Println("orginal slice:", names)
	names = append(names, "Dave")

	fmt.Println("slice:", names)
	fmt.Println("slice length:", len(names))

	scores := map[string]int{
		"Alice": 90,
		"Bob":   75,
		"Chanhyuk": 100, //추가됨
	}

	aliceScore, ok := scores["Alice"]
	if ok {
		fmt.Println("Alice score:", aliceScore)
	}
	
	for name, score := range scores {
		fmt.Printf("%s: %d점\n", name, score)
	}
	//평균 구하기
	totalScore := 0
	for _, score := range scores {
		totalScore += score
	}
	fmt.Printf("전체 평균: %d점\n", totalScore / len(scores))
}
