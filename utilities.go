/*
utilities.go
Description:
*/

package modelchecking

import (
	"fmt"
)

/*
Subset
Description:
	Determines if apSliceA is a subset of apSliceB
*/
func SliceSubset(slice1, slice2 interface{}) (bool, error) {

	switch x := slice1.(type) {
	case []AtomicProposition:

		apSlice1, err1 := ToSliceOfAtomicPropositions(slice1)
		apSlice2, err2 := ToSliceOfAtomicPropositions(slice2)

		if (err1 != nil) || (err2 != nil) {
			return false, fmt.Errorf("Error converting slice1: %v ; Error converting slice2: %v", err1, err2)
		}

		//Iterate through all AtomicPropositions in apSliceA and make sure that they are in B.
		for _, apFrom1 := range apSlice1 {
			if !(apFrom1.In(apSlice2)) {
				return false, nil
			}
		}
		// If all elements of slice1 are in slice2 then return true!
		return true, nil
	case []TransitionSystemState:
		stateSlice1, ok1 := slice1.([]TransitionSystemState)
		stateSlice2, ok2 := slice2.([]TransitionSystemState)

		if (!ok1) || (!ok2) {
			return false, fmt.Errorf("Error converting slice1 (%v) or slice2 (%v).", ok1, ok2)
		}

		//Iterate through all TransitionSystemState in stateSlice1 and make sure that they are in 2.
		for _, stateFrom1 := range stateSlice1 {
			if !(stateFrom1.In(stateSlice2)) {
				return false, nil
			}
		}
		// If all elements of slice1 are in slice2 then return true!
		return true, nil

	default:
		return false, fmt.Errorf("Unexpected type given to SliceSubset(): %v", x)
	}

}

/*
SliceEquals
Description:

*/
func SliceEquals(slice1, slice2 interface{}) (bool, error) {
	//Determine if both slices are of the same type.
	// if slice1.(type) != slice2.(type) {
	// 	fmt.Println("Types of the two slices are different!")
	// 	return false
	// }

	oneSubsetTwo, err := SliceSubset(slice1, slice2)
	if err != nil {
		return false, fmt.Errorf("There was an issue computing SliceSubset(slice1,slice2): %v", err)
	}

	twoSubsetOne, err := SliceSubset(slice2, slice1)
	if err != nil {
		return false, fmt.Errorf("There was an issue computing SliceSubset(slice2,slice1): %v", err)
	}

	return oneSubsetTwo && twoSubsetOne, nil

}

/*
FindInSlice
Description:

*/
func FindInSlice(xIn interface{}, sliceIn interface{}) (int, bool) {

	x := xIn.(string)
	slice := sliceIn.([]string)

	xLocationInSliceIn := -1

	for sliceIndex, sliceValue := range slice {
		if x == sliceValue {
			xLocationInSliceIn = sliceIndex
		}
	}

	return xLocationInSliceIn, xLocationInSliceIn >= 0
}

/*
GetBeverageVendingMachineTS
Description:

*/
func GetBeverageVendingMachineTS() TransitionSystem {

	ts0, err := GetTransitionSystem(
		[]string{"pay", "select", "beer", "soda"}, []string{"", "insert_coin", "get_beer", "get_soda"},
		map[string]map[string][]string{
			"pay": map[string][]string{
				"insert_coin": []string{"select"},
			},
			"select": map[string][]string{
				"": []string{"beer", "soda"},
			},
			"beer": map[string][]string{
				"get_beer": []string{"pay"},
			},
			"soda": map[string][]string{
				"get_soda": []string{"pay"},
			},
		},
		[]string{"pay"},
		[]string{"paid", "drink"},
		map[string][]string{
			"pay":    []string{},
			"soda":   []string{"paid", "drink"},
			"beer":   []string{"paid", "drink"},
			"select": []string{"paid"},
		},
	)

	if err != nil {
		fmt.Println(fmt.Sprintf("There was an issue constructing the beverage vending machine! %v", err.Error()))
	}

	return ts0

}
