package greedy

func canCompleteCircuit(gas []int, cost []int) int {
	curOil, sumOil, start := 0, 0, 0
	for i := 0; i < len(cost); i++ {
		curOil += gas[i] - cost[i]
		sumOil += gas[i] - cost[i]
		if curOil < 0 {
			start = i + 1
			curOil = 0
		}
	}
	if sumOil < 0 {
		return -1
	}
	return start
}
