package decorator

/*
	装饰器模式跟代理模式代码实现上无本质区别，区别在于该模式为类增强功能，功能间有关联，而代理模式对功能本身无改动；
*/

type IDraw interface {
	Draw() string
}

// Square 正方形
type Square struct{}

func (s Square) Draw() string {
	return "this is a square"
}

// ColorSquare 有颜色的正方形
type ColorSquare struct {
	square IDraw
	color  string
}

func NewColorSquare(square IDraw, color string) ColorSquare {
	return ColorSquare{color: color, square: square}
}

func (c ColorSquare) Draw() string {
	return c.square.Draw() + ", color is " + c.color
}
