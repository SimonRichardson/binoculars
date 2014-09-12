package binoculars

import "reflect"

func IdLens() Lens {
	return NewLens(func(a Any) Store {
		return NewStore(identity(), constant(a))
	})
}

func AccessorLens(accessor Accessor) Lens {
	return NewLens(func(a Any) Store {
		return NewStore(
			func(b Any) Any {
				return accessor.Set(a, b)
			},
			func() Any {
				return accessor.Get(a)
			},
		)
	})
}

func SliceLens(index uint) Lens {
	return NewLens(func(a Any) Store {
		src := reflect.ValueOf(a)
		dst := reflect.New(src.Type()).Elem()
		dst.Set(src)

		val := dst.Index(int(index))

		return NewStore(
			func(x Any) Any {
				val.Set(reflect.ValueOf(x))
				return dst.Interface()
			},
			func() Any {
				return val.Interface()
			},
		)
	})
}

func ObjectLens(property string) Lens {
	return NewLens(func(a Any) Store {
		src := reflect.ValueOf(a)
		dst := reflect.New(src.Type()).Elem()
		dst.Set(src)

		val := dst.FieldByName(property)

		return NewStore(
			func(x Any) Any {
				val.Set(reflect.ValueOf(x))
				return dst.Interface()
			},
			func() Any {
				return val.Interface()
			},
		)
	})
}

func ObjectLenses(properties []string) []Lens {
	num := len(properties)
	res := make([]Lens, num, num)
	for k, v := range properties {
		res[k] = ObjectLens(v)
	}
	return res
}
