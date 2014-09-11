package binoculars

type Store struct {
	Set func(x Any) Any
	Get func() Any
}

func NewStore(set func(Any) Any, get func() Any) Store {
	return Store{
		Set: set,
		Get: get,
	}
}

func (s Store) Map(f func(Any) Any) Store {
	return s.Extend(func(x Store) Any {
		return f(x.Extract())
	})
}

func (s Store) Extend(f func(Store) Any) Store {
	return Store{
		Set: func(a Any) Any {
			return f(Store{
				Set: s.Set,
				Get: constant(a),
			})
		},
		Get: s.Get,
	}
}

func (s Store) Extract() Any {
	return s.Set(s.Get())
}
