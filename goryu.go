package goryu

func MapChan[Input any, Out any](in <-chan Input, mapper func(data Input) <-chan Out) <-chan <-chan Out {
	return MapUntil(nil, 0, in, mapper)
}

// Done returns a channel that closes when all channels from the input arguments are closed
// Edited from EveryDone
func Done[Sig any](done ...<-chan Sig) <-chan Sig {

	switch len(done) {
	case 0:
		return nil
	case 1:
		return done[0]
	}

	allDone := make(chan Sig)
	go func() {
		defer close(allDone)
		for _, d := range done {
			<-d
		}
	}()

	return allDone
}
