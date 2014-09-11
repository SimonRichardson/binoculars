package binoculars

import "reflect"

type Lens struct {
	Run func(Any) Store
}

func NewLens(run func(Any) Store) Lens {
	return Lens{
		Run: run,
	}
}

func (l Lens) Compose(x Lens) Lens {
	return NewLens(func(target Any) Store {
		a := x.Run(target)
		b := l.Run(a.Get())
		return NewStore(
			compose(a.Set)(b.Set),
			b.Get,
		)
	})
}

func (l Lens) AndThen(x Lens) Lens {
	return x.Compose(l)
}

func (l Lens) And(x Lens) Lens {
	return NewLens(func(target Any) Store {
		a := l.Run(target)
		b := x.Run(target)
		return NewStore(
			func(x Any) Any {
				y := x.([]Any)
				return merge(target, a.Set(y[0]), b.Set(y[1]))
			},
			b.Get,
		)
	})
}

func merge(a, b, c Any) Any {
	src := reflect.ValueOf(a)

	x := reflect.ValueOf(b)
	y := reflect.ValueOf(c)

	dst := reflect.New(src.Type()).Elem()
	dst.Set(src)

	for i := 0; i < dst.NumField(); i++ {
		field := dst.Field(i)

		poss_a := x.Field(i).Interface()
		poss_b := y.Field(i).Interface()

		orig := field.Interface()
		if orig != poss_a {
			field.Set(reflect.ValueOf(poss_a))
		} else if orig != poss_b {
			field.Set(reflect.ValueOf(poss_b))
		}
	}

	return dst.Interface()
}
