package arithmestic

// 두 숫자의 연산 계산 제공 패키지(2)

// x, y 제곱의 합을 리턴
func (o *Numbers) MultiPlus() int {
	return (o.X + o.X) + (o.Y * o.Y)
}

// x, y 제곱의 합을 리턴
func (o *Numbers) MultiMinus() int {
	return (o.X - o.X) + (o.Y - o.Y)
}
