package apirequest

type ApiRequestHandler interface {
	//Find(word string) string
	Get(url string) ([]byte, error)
}
