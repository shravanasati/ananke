package html2md

type counter struct {
	current int
	step    int
}

func newCounter(start, step int) *counter {
	return &counter{
		current: start,
		step:    step,
	}
}

func (c *counter) next() int {
	c.current += c.step
	return c.current - c.step
}
