package query

type QueryHandler[Response any] interface {
	Handle(query Query[Response]) (Response, error)
}
