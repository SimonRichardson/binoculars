package binoculars

type Option interface {
	Of(Any) Option
	Chain(func(Any) Option) Option
	Map(func(Any) Any) Option
	Fold(func(Any) Option, func() Option) Option
	GetOrElse(func() Any) Any
}

type Some struct {
	x Any
}

func NewSome(x Any) Some {
	return Some{
		x: x,
	}
}

func (x Some) Of(v Any) Option {
	return NewSome(v)
}

func (x Some) Chain(f func(v Any) Option) Option {
	return f(x.x)
}

func (x Some) Map(f func(v Any) Any) Option {
	return x.Chain(func(v Any) Option {
		return x.Of(f(v))
	})
}

func (x Some) Fold(f func(v Any) Option, g func() Option) Option {
	return f(x.x)
}

func (x Some) GetOrElse(v func() Any) Any {
	return x.x
}

type None struct{}

func NewNone() None {
	return None{}
}

func (x None) Of(v Any) Option {
	return NewSome(v)
}

func (x None) Chain(f func(v Any) Option) Option {
	return x
}

func (x None) Map(f func(v Any) Any) Option {
	return x
}

func (x None) Fold(f func(v Any) Option, g func() Option) Option {
	return g()
}

func (x None) GetOrElse(v func() Any) Any {
	return v()
}
