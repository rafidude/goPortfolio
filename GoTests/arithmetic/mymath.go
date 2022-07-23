package arithmetic

type Val interface {
	int | float64
}

func Add[T Val](x T, y T) T {
	return x + y
}

func Sub[T Val](x T, y T) T {
	return x - y
}

func Div[T Val](x T, y T) T {
	return x / y
}

func Mul[T Val](x T, y T) T {
	return x * y
}
