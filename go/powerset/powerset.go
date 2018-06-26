package powerset

import (
	"strings"
	"fmt"
	"math"
)

type Set []interface{}

func SetOf(items []interface{}) Set {
	contained := map[interface{}]bool{}
	set := Set{}

	for _, item := range items {
		if(!contained[item]) {
			set = append(set, item)
			contained[item] = true
		}
	}

	return set
}

func (set *Set) Copy() Set {
	copy := make(Set, len(*set))
	for i := 0; i < len(*set); i++ {
		copy[i] = (*set)[i]
	}

	return copy
}

func (set *Set) String() string {
	builder := strings.Builder{}
	builder.WriteString("{")
	for _, item := range *set {
		builder.WriteString(fmt.Sprintf("%v ", item))
	}
	builder.WriteString("}")

	return builder.String()
}

func PowerSet(set Set) Set {
	return powerSetHelper(set, Set{})
}

func powerSetHelper(set Set, setComponent Set) Set {
	if len(set) == 0 {
		return setComponent
	}

	var resultingSet Set
	setsWithItem := powerSetHelper(set[1:], append(setComponent.Copy(), set[0]))
	setsWithoutItem := powerSetHelper(set[1:], setComponent)

	if len(set) > 1 {
		resultingSet = Set{}
		for _, item := range setsWithItem {
			resultingSet = append(resultingSet, item.(Set))
		}

		for _, item := range setsWithoutItem {
			resultingSet = append(resultingSet, item.(Set))
		}
	} else {
		resultingSet = Set{setsWithItem, setsWithoutItem}
	}

	return resultingSet
}

func PowerSetIterative(set Set) Set {
	cardinalityOfResultingSet := int(math.Pow(2, float64(len(set))))
	resultingSet := make(Set, cardinalityOfResultingSet)

	for i := 0; i < cardinalityOfResultingSet; i++ {
		resultingSet[i] = Set{}
	}

	if cardinalityOfResultingSet > 1 {
		divisions := 2
		elementsPerDivision := cardinalityOfResultingSet / 2
		for _, element := range set {
			index := 0
			bound := elementsPerDivision
			for division := 0; division < divisions; division += 2 {
				for ; index < bound; index++ {
					resultingSet[index] = append(resultingSet[index].(Set), element)
				}

				bound += 2 * elementsPerDivision
				index += elementsPerDivision
			}

			divisions = divisions * 2
			elementsPerDivision = elementsPerDivision / 2
		}
	}

	return resultingSet
}
