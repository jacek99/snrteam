package generics
import (
	"container/list"
	"github.com/joeshaw/gengen/generic"
)

//Converts a List to an array types
func List2Array(l *list.List, entities []generic.T) []generic.T {
	index := 0
	for e := l.Front(); e != nil; e = e.Next() {
		entities[index] = e.Value.(generic.T)
		index++
	}
	return entities
}



