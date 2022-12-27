package graph

var _ = findOrder

// 210
func findOrder(numCourses int, prerequisites [][]int) []int {
	in := make([]int, numCourses)
	out := make([][]int, numCourses)

	for i := 0; i < len(prerequisites); i++ {
		course, precourse := prerequisites[i][0], prerequisites[i][1]
		// precourse -> course
		in[course]++
		out[precourse] = append(out[precourse], course)
	}

	queue := make([]int, 0)

	for course := 0; course < len(in); course++ {
		if in[course] == 0 {
			queue = append(queue, course)
		}
	}

	ans := make([]int, 0)

	for len(queue) > 0 {
		course := queue[0]
		queue = queue[1:]
		ans = append(ans, course)
		for _, postcourse := range out[course] {
			in[postcourse]--
			if in[postcourse] == 0 {
				queue = append(queue, postcourse)
			}
		}
	}

	if len(ans) == numCourses {
		return ans
	}
	return []int{}
}
