package v1

import (
	"fmt"
)

func ValidateProductID(productID int32) error {
	// Check if the product ID is within a valid range (you can customize this)
	if productID <= 0 {
		return fmt.Errorf("product ID must be a positive integer")
	}

	// Optionally, you can add additional checks (e.g., ensuring the product ID is within a certain range)
	// if productID > 10000 {
	//     return fmt.Errorf("product ID is too large")
	// }

	return nil
}
