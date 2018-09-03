//Package arrayutil is just to perform arithmatics operation
package arrayutil

//Sum returns the sum of whole array.
func Sum(a []int) int {
	sum := 0
	for _, e := range a {
		sum += e
	}

	return sum
}

//Mul returns the multiplication of whole array.
func Mul(a []int) int {
	mul := 1
	for _, e := range a {
		mul *= e
	}

	return mul
}
