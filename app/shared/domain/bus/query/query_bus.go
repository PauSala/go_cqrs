package query

type QueryBus[Response any] interface {
	Ask(query Query[Response]) (Response, error)
}
