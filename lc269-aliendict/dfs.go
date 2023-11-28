package lc269aliendict

import "fmt"

/* https://www.youtube.com/watch?v=ZU7fiX0WCCY
Input: words = ["wrt","wrf","er","ett","rftt"]
Output: "wertf"

t -> f

w -> e

r -> t

e -> r

w - e - r - t - f
*/

func AlienOrder(words []string) string {

	child := make(map[string]map[string]bool)
	pare := make(map[string]map[string]bool)

	for _, word := range words {
		for _, char := range word {
			child[string(char)] = make(map[string]bool)
			pare[string(char)] = make(map[string]bool)
		}
	}

	for i := 0; i < len(words)-1; i++ {
		w1 := words[i]
		w2 := words[i+1]

		if len(w1) > len(w2) && w1[:len(w2)] == w2 {
			return ""
		}

		// fmt.Println(w1)
		// fmt.Println(w2)

		for j := 0; j < len(w1); j++ {
			if w1[j] != w2[j] {
				// child[[]string(w1)[j]] = append(child[[]string(w1)[j]], []string(w2)[j])
				// pare[[]string(w2)[j]] = append(pare[[]string(w2)[j]], []string(w1)[j])

				// fmt.Println(string(w1[j]))
				// fmt.Println(string(w2[j]))

				child[string(w1[j])][string(w2[j])] = true
				pare[string(w2[j])][string(w1[j])] = true
				break
			}
		}
	}

	// for c, cs := range pare {
	// 	fmt.Printf("parent = %v, chilren = ", c)
	// 	for _, p := range cs {
	// 		fmt.Printf(" %v,", string(p))
	// 	}
	// 	fmt.Println()
	// }

	// fmt.Println("----------")

	// for c, cs := range child {
	// 	fmt.Printf("child = %v, parents = ", c)
	// 	for _, p := range cs {
	// 		fmt.Printf(" %v,", string(p))
	// 	}
	// 	fmt.Println()
	// }

	res := []string{}
	queue := []string{}

	no_par_node := []string{}
	for node, cs := range pare {
		// fmt.Println(string(node))
		if len(cs) == 0 {
			no_par_node = append(no_par_node, node)
		}
	}

	fmt.Println(no_par_node)

	for _, node := range no_par_node {
		delete(pare, node)
		queue = append(queue, node)
	}

	for len(queue) != 0 {
		node := queue[0]
		queue = queue[1:]
		res = append(res, node)
		for ch := range child[node] {
			delete(pare[ch], node)
			if len(pare[ch]) == 0 {
				queue = append(queue, ch)
				delete(pare, ch)
			}
		}
	}

	if len(pare) != 0 {
		return ""
	}
	str := ""
	for _, s := range res {
		str += s
	}
	return str
}
