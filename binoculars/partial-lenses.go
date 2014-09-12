package binoculars

import "reflect"

func IdPartialLens() PartialLens {
	return NewPartialLens(func(a Any) Option {
		return NewSome(IdLens().Run(a))
	})
}

func SlicePartialLens(index uint) PartialLens {
	return NewPartialLens(func(a Any) Option {
		num := reflect.ValueOf(a).Len()
		if index >= 0 && int(index) < num {
			x := SliceLens(index)
			return NewSome(x.Run(a))
		}
		return NewNone()
	})
}

func ObjectPartialLens(property Property) PartialLens {
	return NewPartialLens(func(a Any) Option {
		val := reflect.ValueOf(a).FieldByName(property.String())
		if val.Kind() != reflect.Invalid {
			x := ObjectLens(property)
			return NewSome(x.Run(a))
		}
		return NewNone()
	})
}

func ObjectPartialLenses(properties []Property) map[Property]PartialLens {
	res := make(map[Property]PartialLens, len(properties))
	for _, v := range properties {
		res[v] = ObjectPartialLens(v)
	}
	return res
}
