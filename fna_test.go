package gomathx

import "testing"

func TestAdd(t *testing.T) {
	f := New(1, 1, 4)
	f.Add(2, 3, 5)
	t.Log(f)
	t.Skipped()

}

func TestGCD(t *testing.T) {
	gcd := GCD(5, 40)
	t.Log(gcd)
}
