package binoculars

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
