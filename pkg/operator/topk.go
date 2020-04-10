package operator

import (
	"github.com/dgryski/go-topk"
)

var stream *topk.Stream
func init() {
	stream = topk.New(10)
}


// Add a element to find topk
func AddToTopK(name string, count int) {
	stream.Insert(name, count)
}


// Returns the top elements
func GetTopK() []topk.Element {
	return stream.Keys()

}