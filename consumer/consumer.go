package consumer

type listeners map[string]func()

func (l listeners) Register(key string, callback func()) {
	l[key] = callback
}

func (l listeners) fire(key string) {
	l[key]()
}
