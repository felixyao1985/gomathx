package gomathx

import "testing"

func TestAdd(t *testing.T) {
	f := New(1, 1, 4)
	f.Add(2, 3, 5)
	t.Log(f)
	t.Skipped()

}

func TestAdd2(t *testing.T) {
	f := New(0, 0, 0)
	f.Add(2, 3, 5)
	t.Log(f)
	t.Skipped()

}

func TestSub(t *testing.T) {
	f := New(4, 1, 4)
	f.Sub(2, 3, 5)
	t.Log(f)
	t.Skipped()
}

func TestSub2(t *testing.T) {
	f := New(0, 1, 4)
	f.Sub(2, 3, 5)
	t.Log(f)
	t.Skipped()
}

func TestGCD(t *testing.T) {
	gcd := GCD(5, 40)
	t.Log(gcd)
}
