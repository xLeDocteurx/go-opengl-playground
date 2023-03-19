package types

type JsonMap struct {
    width int
    height int
	cells []int
}

func NewJsonMap(width int, height int, cells []int) JsonMap {
	return JsonMap{width: width, height: height, cells: cells}
}
