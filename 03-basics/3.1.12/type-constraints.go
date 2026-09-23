package main

type additive interface {
	int | ~uint64
}

func add[T additive](a, b T) T {
	return a + b
}

func add2[T int | ~uint64](a, b T) T {
	return a + b
}

type iAmUnit64Alias uint64

func typeConstraints() {
	_ = add(1, 2)
	_ = add2(3, 4)

	_ = add(1, iAmUnit64Alias(2))
	_ = add2(3, iAmUnit64Alias(4))
}
