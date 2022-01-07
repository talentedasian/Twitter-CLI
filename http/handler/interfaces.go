package handler

import "twitter/http/client"

type Handler interface {
	handle(client client.Client) (string, error)
}

type URLReq interface {
	URI() (string, error)
}
