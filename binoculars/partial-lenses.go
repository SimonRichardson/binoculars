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

func ObjectPartialLens(property string) PartialLens {
	return NewPartialLens(func(a Any) Option {
		val := reflect.ValueOf(a).FieldByName(property)
		if val.Kind() != reflect.Invalid {
			x := ObjectLens(property)
			return NewSome(x.Run(a))
		}
		return NewNone()
	})
}

func ObjectPartialLenses(properties []string) []PartialLens {
	num := len(properties)
	res := make([]PartialLens, num, num)
	for k, v := range properties {
		res[k] = ObjectPartialLens(v)
	}
	return res
}
