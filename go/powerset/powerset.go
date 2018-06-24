package powerset

import (
	"strings"
	"fmt"
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