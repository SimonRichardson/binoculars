package binoculars

type PartialLens struct {
	Run func(Any) Option
}

func NewPartialLens(run func(Any) Option) PartialLens {
	return PartialLens{
		Run: run,
	}
}

func (p PartialLens) Compose(x PartialLens) PartialLens {
	return NewPartialLens(func(target Any) Option {
		return x.Run(target).Chain(func(y Any) Option {
			a := y.(Store)
			return p.Run(a.Get()).Map(func(z Any) Any {
				b := z.(Store)
				return NewStore(
					compose(a.Set)(b.Set),
					b.Get,
				)
			})
		})
	})
}

func (p PartialLens) AndThen(x PartialLens) PartialLens {
	return x.Compose(p)
}
