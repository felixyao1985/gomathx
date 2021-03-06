package gomathx

type Fractional struct {
	Number      int64
	Molecule    int64
	Denominator int64
}

func New(n, m, d int64) *Fractional {
	return &Fractional{
		Number:      n,
		Molecule:    m,
		Denominator: d,
	}
}

//同化分子分母
func (s *Fractional) ass(f *Fractional) {
	if s.Denominator == f.Denominator {
		if s.Denominator == 0 {
			s.Denominator = 1
			f.Denominator = 1
			s.Molecule = s.Number
			f.Molecule = f.Number
			s.Number = 0
			f.Number = 0
		}
		return
	}

	if s.Denominator == 0 {
		s.Denominator = f.Denominator
	}

	if f.Denominator == 0 {
		f.Denominator = s.Denominator
	}

	if s.Denominator%f.Denominator == 0 || f.Denominator%s.Denominator == 0 {
		if s.Denominator > f.Denominator {
			d := s.Denominator / f.Denominator
			f.Denominator *= d
			f.Molecule *= d
		} else {
			d := f.Denominator / s.Denominator
			s.Denominator *= d
			s.Molecule *= d
		}
		return
	}
	fd := f.Denominator
	sd := s.Denominator
	s.Molecule *= fd
	s.Denominator *= fd
	f.Molecule *= sd
	f.Denominator *= sd
}

//整理
func (s *Fractional) offset() {
	if s.Molecule == 0 || s.Denominator == 0 {
		s.Molecule = 0
		s.Denominator = 0
		return
	}

	if s.Denominator == 1 {
		if s.Number == 0 {
			s.Number = s.Molecule
		} else {
			s.Number = s.Number + s.Molecule*s.Number/s.Number
		}
		s.Denominator = 0
		s.Molecule = 0

		return
	}

	if s.Denominator < s.Molecule {
		m := s.Molecule % s.Denominator
		s.Number = s.Number + (s.Molecule-m)/s.Denominator
		s.Molecule = m
	}

	if s.Denominator == s.Molecule {
		s.Number++
		s.Molecule = 0
		s.Denominator = 0
		return
	}

	gcd := GCD(s.Molecule, s.Denominator)

	s.Molecule /= gcd
	s.Denominator /= gcd

	return
}

// Addition
func (s *Fractional) Add(n, m, d int64) *Fractional {

	f := &Fractional{
		Number:      n,
		Molecule:    m,
		Denominator: d,
	}
	s.ass(f)

	s.Number += f.Number
	s.Molecule += f.Molecule

	s.offset()
	return s
}

// Subtraction
func (s *Fractional) Sub(n, m, d int64) *Fractional {

	f := &Fractional{
		Number:      n,
		Molecule:    m,
		Denominator: d,
	}
	s.ass(f)

	s.Molecule = s.Molecule + s.Number*s.Denominator
	s.Number = 0
	f.Molecule = f.Molecule + f.Number*f.Denominator
	f.Number = 0

	s.Molecule -= f.Molecule

	s.offset()
	return s
}

// Multiplication
func (s *Fractional) Mul(n, m, d int64) *Fractional {
	f := &Fractional{
		Number:      n,
		Molecule:    m,
		Denominator: d,
	}

	s.Number *= f.Number
	s.Molecule *= f.Molecule
	s.Denominator *= f.Denominator

	s.offset()

	return s
}

// Division
func (s *Fractional) Div(n, m, d int64) *Fractional { // 除法
	f := &Fractional{
		Number:      n,
		Molecule:    m,
		Denominator: d,
	}

	if s.Denominator == 0 {
		s.Denominator = 1
	}
	if f.Denominator == 0 {
		f.Denominator = 1
	}

	s.Molecule = s.Molecule + s.Number*s.Denominator
	s.Number = 0
	f.Molecule = f.Molecule + f.Number*f.Denominator
	f.Number = 0

	s.Mul(f.Number, f.Denominator, f.Molecule)

	return s
}
