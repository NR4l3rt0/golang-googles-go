// Package dog helps synchronizing the life between dogs and humans
package dog

// Dog identifies a dog by name and age
type Dog struct {
	Name string
	Age  int
}

// Years turns human years into dog ones
func Years(hy int) int {
	return hy * 7
}
