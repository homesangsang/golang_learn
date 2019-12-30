package mock

type Retriever struct {
	Contents string
}

// Retriever并没有实现上层main.go中的Retriever接口，只要含有Get方法，即可
func (r Retriever) Get(url string) string {
	return "https://" + r.Contents
}

func (r Retriever) Get1(url string, queryParam string) string {
	return "https://" + r.Contents + "?" + queryParam
}
