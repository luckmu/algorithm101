package graph

var _ = canFinish

// No.207
func canFinish(numCourses int, prerequisites [][]int) bool {
	indeg := make([]int, numCourses)       // A -> B
	outdeg := make([][]int, 0, numCourses) // A -> [B,C,D]
	queue := make([]int, 0)
	learned := 0

	for i := 0; i < len(prerequisites); i++ {
		course, precourse := prerequisites[i][0], prerequisites[i][1]
		indeg[course]++
		outdeg[precourse] = append(outdeg[precourse], course)
	}

	for i := 0; i < len(indeg); i++ {
		if indeg[i] == 0 {
			queue = append(queue, indeg[i])
		}
	}

	for len(queue) > 0 {
		course := queue[0]
		queue = queue[1:]
		learned++
		for _, postcourse := range outdeg[course] {
			indeg[postcourse]--
			if indeg[postcourse] == 0 {
				queue = append(queue, postcourse)
			}
		}
	}
	return learned == numCourses
}
