package consumer

type consumer struct {
	hooks map[string]func()
}

func NewConsumer() consumer {
	return consumer{
		hooks: make(map[string]func()),
	}
}

func (c consumer) Register(key string, callback func()) {
	c.hooks[key] = callback
}

func (c consumer) fire(key string) {
	c.hooks[key]()
}
