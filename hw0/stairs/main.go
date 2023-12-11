package main

// Василь займається будівництвом драбин з блоків. Драбина містить сходинки, при цьому і-та сходинка має складатися рівно з і блоків.

// Драбина, що складається з 10 блоків та 4 сходинок, має наступний вигляд:

// []
// [] []
// [] [] []
// [] [] [] []
// По заданому числу блоків n визначте максимальну кількість сходинок в драбині, яку можна побудувати з цих блоків.

// Input Format

// Вводиться одне число n

// Constraints

// 1 ≤ n ≤ 2^31 - 1

// Output Format

// Виведіть одне число - кількість сходинок в драбині.

// Малюнок відповідає тестовим прикладам - [] показані блоки, використані при побудові драбини, а X — зайві блоки, яких недостатньо для будівництва чергової сходинки.

// []
// [] []
// X  X
// []
// [] []
// [] [] []
// X  X
// Sample Input 0

// 5
// Sample Output 0

// 2
// Sample Input 1

// 8
// Sample Output 1

// 3
import (
	"fmt"
)

func main() {
	var totalBlocks int
	fmt.Scanln(&totalBlocks)
	numOfStairs := 0
	for usedBlocks := 0; usedBlocks < totalBlocks-numOfStairs; {
		numOfStairs++
		usedBlocks += numOfStairs
	}
	fmt.Println(numOfStairs)
}