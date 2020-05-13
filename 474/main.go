package main

import "fmt"

func main() {
	strs := []string{"10", "0001", "111001", "1", "0"}
	m := 5
	n := 3
	r := findMaxForm(strs, m, n)
	fmt.Println(r)
}

//m 个 0 和 n 个 1
func findMaxForm(strs []string, m int, n int) int {
	// 计算每个字符串有多少个0和1
	var count [][2]int
	for _, str := range strs {
		m0 := 0
		n1 := 0
		for _, s := range str {
			if s == '0' {
				m0++
			} else {
				n1++
			}
		}
		count = append(count, [2]int{m0, n1})
	}
	cache := make(map[string]int)
	return findMaxFormR(count, cache, m, n, len(strs)-1)
}

func findMaxFormR(count [][2]int, cache map[string]int, m int, n int, i int) int {
	if _, ok := cache[string(m)+string(n)+string(i)]; ok {
		return cache[string(m)+string(n)+string(i)]
	}
	if m < 0 || n < 0 {
		return -1
	}

	if i < 0 {
		return 0
	}
	yes := findMaxFormR(count, cache, m-count[i][0], n-count[i][1], i-1) + 1
	no := findMaxFormR(count, cache, m, n, i-1)
	if yes > no {
		cache[string(m)+string(n)+string(i)] = yes
	} else {
		cache[string(m)+string(n)+string(i)] = no
	}
	return cache[string(m)+string(n)+string(i)]
}
