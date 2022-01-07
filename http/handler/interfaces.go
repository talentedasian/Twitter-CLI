package handler

type Handler interface {
	Handle() (string, error)
}

type URLReq interface {
	URI() (string, error)
}
