package main

type Set map[[2]int]bool

func cleanRoom(robot *Robot) {
	directions := [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
	visited := make(Set)
	backTrack(0, 0, 0, visited, directions, robot)
}

func goBack(r *Robot) {
	r.TurnRight()
	r.TurnRight()
	r.Move()
	r.TurnRight()
	r.TurnRight()
}

func backTrack(row int, col int, d int, visited Set, directions [][]int, robot *Robot) {
	point := [2]int{row, col}
	visited[point] = true
	robot.Clean()
	for i := 0; i < 4; i++ {
		newDirection := (d + i) % 4
		r := row + directions[newDirection][0]
		c := col + directions[newDirection][1]
		if visited[[2]int{r, c}] == false && robot.Move() {
			backTrack(r, c, newDirection, visited, directions, robot)
			goBack(robot)
		}
		robot.TurnRight()
	}
}
