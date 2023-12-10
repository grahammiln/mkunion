// Code generated by mkunion. DO NOT EDIT.
package testutils

func MatchAlphabetNumberTupleR0[T0 Alphabet, T1 Number](
	t0 T0,
	t1 T1,
	f0 func(x0 *A1, x1 *N0),
	f1 func(x0 *C3, x1 any),
	f2 func(x0 any, x1 any),
) {
	c0t0, c0t0ok := any(t0).(*A1)
	c0t1, c0t1ok := any(t1).(*N0)
	if c0t0ok && c0t1ok {
		f0(c0t0, c0t1)
		return
	}

	c1t0, c1t0ok := any(t0).(*C3)
	c1t1, c1t1ok := any(t1).(any)
	if c1t0ok && c1t1ok {
		f1(c1t0, c1t1)
		return
	}

	c2t0, c2t0ok := any(t0).(any)
	c2t1, c2t1ok := any(t1).(any)
	if c2t0ok && c2t1ok {
		f2(c2t0, c2t1)
		return
	}

	panic("MatchAlphabetNumberTupleR0 is not exhaustive")
}

func MatchAlphabetNumberTupleR1[T0 Alphabet, T1 Number, TOut1 any](
	t0 T0,
	t1 T1,
	f0 func(x0 *A1, x1 *N0) TOut1,
	f1 func(x0 *C3, x1 any) TOut1,
	f2 func(x0 any, x1 any) TOut1,
) TOut1 {
	c0t0, c0t0ok := any(t0).(*A1)
	c0t1, c0t1ok := any(t1).(*N0)
	if c0t0ok && c0t1ok {
		return f0(c0t0, c0t1)
	}

	c1t0, c1t0ok := any(t0).(*C3)
	c1t1, c1t1ok := any(t1).(any)
	if c1t0ok && c1t1ok {
		return f1(c1t0, c1t1)
	}

	c2t0, c2t0ok := any(t0).(any)
	c2t1, c2t1ok := any(t1).(any)
	if c2t0ok && c2t1ok {
		return f2(c2t0, c2t1)
	}

	panic("MatchAlphabetNumberTupleR0 is not exhaustive")
}

func MatchAlphabetNumberTupleR2[T0 Alphabet, T1 Number, TOut1 any, TOut2 any](
	t0 T0,
	t1 T1,
	f0 func(x0 *A1, x1 *N0) (TOut1, TOut2),
	f1 func(x0 *C3, x1 any) (TOut1, TOut2),
	f2 func(x0 any, x1 any) (TOut1, TOut2),
) (TOut1, TOut2) {
	c0t0, c0t0ok := any(t0).(*A1)
	c0t1, c0t1ok := any(t1).(*N0)
	if c0t0ok && c0t1ok {
		return f0(c0t0, c0t1)
	}

	c1t0, c1t0ok := any(t0).(*C3)
	c1t1, c1t1ok := any(t1).(any)
	if c1t0ok && c1t1ok {
		return f1(c1t0, c1t1)
	}

	c2t0, c2t0ok := any(t0).(any)
	c2t1, c2t1ok := any(t1).(any)
	if c2t0ok && c2t1ok {
		return f2(c2t0, c2t1)
	}

	panic("MatchAlphabetNumberTupleR0 is not exhaustive")
}

func MatchAlphabetNumberTupleR3[T0 Alphabet, T1 Number, TOut1 any, TOut2 any, TOut3 any](
	t0 T0,
	t1 T1,
	f0 func(x0 *A1, x1 *N0) (TOut1, TOut2, TOut3),
	f1 func(x0 *C3, x1 any) (TOut1, TOut2, TOut3),
	f2 func(x0 any, x1 any) (TOut1, TOut2, TOut3),
) (TOut1, TOut2, TOut3) {
	c0t0, c0t0ok := any(t0).(*A1)
	c0t1, c0t1ok := any(t1).(*N0)
	if c0t0ok && c0t1ok {
		return f0(c0t0, c0t1)
	}

	c1t0, c1t0ok := any(t0).(*C3)
	c1t1, c1t1ok := any(t1).(any)
	if c1t0ok && c1t1ok {
		return f1(c1t0, c1t1)
	}

	c2t0, c2t0ok := any(t0).(any)
	c2t1, c2t1ok := any(t1).(any)
	if c2t0ok && c2t1ok {
		return f2(c2t0, c2t1)
	}

	panic("MatchAlphabetNumberTupleR0 is not exhaustive")
}
