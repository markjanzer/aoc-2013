package lib

/*
Search algorithm that tries to find the shortest path

initialState is a list of states to start from

done is a function that returns true if the state is the goal

next is a function that returns the next states from a given state. It returns a
map of states and the cost of getting to that state

estimate is a function that returns the estimated cost of getting to the goal from a given state
*/
func AStar[S comparable](
	initialState []S,
	done func(state S) bool,
	next func(state S) map[S]int,
	estimate func(state S) int,
) int {

	upNext := NewHeap(func(a, b statePriority[S]) bool {
		return a.priority < b.priority
	})
	costs := make(map[S]int)

	for _, init := range initialState {
		upNext.Insert(statePriority[S]{init, 0})
		costs[init] = 0
	}

	for upNext.Size() > 0 {
		nextNode := upNext.Pop()
		state := nextNode.state

		if done(state) {
			return costs[state]
		}

		for next, cost := range next(state) {
			newCost := costs[state] + cost

			if oldCost, ok := costs[next]; !ok || newCost < oldCost {
				costs[next] = newCost
				priority := newCost + estimate(next)
				upNext.Insert(statePriority[S]{next, priority})
			}
		}
	}

	panic("No result found")
}

type statePriority[S comparable] struct {
	state    S
	priority int
}
