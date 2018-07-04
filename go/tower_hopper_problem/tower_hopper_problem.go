package tower_hopper_problem

import "math"

type TowerHeight int
type Towers []int
type HopPosition int
type HopSequence []HopPosition

func IsHoppable(towers Towers) (isPossible bool, hopsTaken HopSequence)  {
	numberOfTowers := len(towers)
	doneWithHopping := false
	currentTower := 0
	hopsTaken = HopSequence{0}

	for !doneWithHopping {
		maxHopsPossible := towers[currentTower]
		currentBestNextTower := currentTower
		currentWidestHopDistance := maxHopsPossible
		for hopLength := 1; hopLength <= maxHopsPossible; hopLength++ {
			if nextTower := currentTower + hopLength; nextTower < numberOfTowers {
				hopsFromThere := towers[nextTower]
				hopSum := hopLength + hopsFromThere
				if hopsFromThere > 0 &&  hopSum > currentWidestHopDistance {
						currentWidestHopDistance = hopSum
						currentBestNextTower = nextTower
				}
			} else {
				isPossible = true
				doneWithHopping = true
				currentBestNextTower = math.MinInt64
				break
			}
		}

		if currentBestNextTower == currentTower {
			isPossible = false
			break
		}

		hopsTaken = append(hopsTaken, HopPosition(currentBestNextTower))

		currentTower = currentBestNextTower
	}

	return
}
