// Utility package for formatting output
package print

import "fmt"

// Format numbers
func Format(num int) string {
	return fmt.Sprintf("The number is %d", num)
}
