// Code generated by "stringer -type=Status -output=stringer.go"; DO NOT EDIT

package vagrantutil

import "fmt"

const _Status_name = "UnknownNotCreatedRunningSavedPowerOffAbortedPreparing"

var _Status_index = [...]uint8{0, 7, 17, 24, 29, 37, 44, 53}

func (i Status) String() string {
	if i < 0 || i >= Status(len(_Status_index)-1) {
		return fmt.Sprintf("Status(%d)", i)
	}
	return _Status_name[_Status_index[i]:_Status_index[i+1]]
}
