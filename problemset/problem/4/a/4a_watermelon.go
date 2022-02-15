/* Problemset:
A. Watermelon
time limit per test1 second
memory limit per test64 megabytes
inputstandard input
outputstandard output
One hot summer day Pete and his friend Billy decided to buy a watermelon. They chose the biggest and the ripest one, in their opinion. After that the watermelon was weighed, and the scales showed w kilos. They rushed home, dying of thirst, and decided to divide the berry, however they faced a hard problem.

Pete and Billy are great fans of even numbers, that's why they want to divide the watermelon in such a way that each of the two parts weighs even number of kilos, at the same time it is not obligatory that the parts are equal. The boys are extremely tired and want to start their meal as soon as possible, that's why you should help them and find out, if they can divide the watermelon in the way they want. For sure, each of them should get a part of positive weight.

Input
The first (and the only) input line contains integer number w (1 ≤ w ≤ 100) — the weight of the watermelon bought by the boys.

Output
Print YES, if the boys can divide the watermelon into two parts, each of them weighing even number of kilos; and NO in the opposite case.

Examples
inputCopy
8
outputCopy
YES
Note
For example, the boys can divide the watermelon into two parts of 2 and 6 kilos respectively (another variant — two parts of 4 and 4 kilos).

---

What we learned:
CodeForces takes input with fmt.Scanln and not os.Args. 1 is an odd number of slices. Each submission shows the full results of tests, including input, in the submission number link.
*/

package main

import "fmt"

func input_calc(n int) string {
	var answer string
	fmt.Scanln(&n)
	if n%2 != 0 {
		answer = "NO"
	} else if n <= 2 {
		answer = "NO"
	} else {
		answer = "YES"
	}

	return answer
}

func main() {
	fmt.Println(input_calc(0))
}
