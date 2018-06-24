package recursive_staircase_problem

import (
	"fmt"
	"strings"
)

type StepSequence []int

func (seq *StepSequence) String() string {
	builder := strings.Builder{}

	builder.WriteString("[")
	for _, step := range *seq {
		builder.WriteString(fmt.Sprintf("%v ", step))
	}
	builder.WriteString("]")

	return builder.String()
}

func (seq *StepSequence) Copy() StepSequence {
	copy := make(StepSequence, len(*seq))
	for i := 0; i < len(*seq); i++ {
		copy[i] = (*seq)[i]
	}

	return copy
}

func NumWaysWithStepSequences(N int, X []int, stepSequence StepSequence) (numberOfPathes int, stepSequences []StepSequence) {
	if N == 0 {
		return 1, []StepSequence{stepSequence}
	}

	var stepSequencesToAdd []StepSequence
	var numberOfPathesFromHere int
	numberOfPathes = 0

	for _, x := range X {
		if x <= N {
			numberOfPathesFromHere, stepSequencesToAdd = NumWaysWithStepSequences(N - x, X, append(stepSequence.Copy(), x))
			numberOfPathes += numberOfPathesFromHere
			for _, stepSequenceToAdd := range stepSequencesToAdd {
				stepSequences = append(stepSequences, stepSequenceToAdd)
			}
		}
	}

	return
}