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

func ObjectLens(property Property) Lens {
	return NewLens(func(a Any) Store {
		src := reflect.ValueOf(a)
		dst := reflect.New(src.Type()).Elem()
		dst.Set(src)

		val := dst.FieldByName(property.String())

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

func ObjectLenses(properties []Property) map[Property]Lens {
	res := make(map[Property]Lens, len(properties))
	for _, v := range properties {
		res[v] = ObjectLens(v)
	}
	return res
}
