package goryu

func MapChan[Input any, Out any](in <-chan Input, mapper func(data Input) <-chan Out) <-chan <-chan Out {
	return MapUntil(nil, 0, in, mapper)
}
