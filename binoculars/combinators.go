package binoculars

func identity() func(Any) Any {
	return func(x Any) Any {
		return x
	}
}

func constant(x Any) func() Any {
	return func() Any {
		return x
	}
}

func compose(f func(x Any) Any) func(func(Any) Any) func(Any) Any {
	return func(g func(Any) Any) func(Any) Any {
		return func(a Any) Any {
			return f(g(a))
		}
	}
}
