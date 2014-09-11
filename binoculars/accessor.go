package binoculars

type Accessor interface {
	Get(Any) Any
	Set(Any, Any) Any
}
