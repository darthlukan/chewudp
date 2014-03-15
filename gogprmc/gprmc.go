/*
	GoGPRMC

	Description: Contains a single method (for now) that tries to convert
	a GPRMC position float64 to a float64 representation of
	either a latitude or longitude value for use in position
	plotting.
*/
package gogprmc

import (
	"errors"
	"fmt"
)

// ConvertToDegrees takes a float64 value from a GPRMC message
// and converts it to a latitude or longitude float value for use
// in plotting a position.
//
// Example:
//     value = 5755.987654321
//     degrees = 57.93312757201667
//
// Returns an error if the type of the input value is not float64.
func ConvertToDegrees(value float64) (float64, error) {

	if _, ok := value.(float64); ok == false {
		return nil, errors.New("ConvertToDegrees: Expected argument to be type float64, got something else instead.")
	}
	degrees := float64((value / 100) + (value-((value/100)*100))/60)

	return degrees, nil
}
