package utils

type Bar struct {
}

func NewBar(count int, MyStrStart, MyStrEnd string) *Bar {
	return &Bar{}
}

func (b *Bar) Grow(num int, MyStrVal string) {
}

func (b *Bar) Done() {
}
