package main

func main() {
	done := make(chan interface{})
	defer close(done)

	repeat := func(done <-chan interface{}, values ...interface{}) <-chan interface{} {
		stream := make(chan interface{})
		go func() {
			defer close(stream)
			for {
				for v := range values {
					select {
					case <-done:
						return
					case stream <- v:
					}
				}
			}
		}()
		return stream
	}

	take := func(done <-chan interface{}, count int, c <-chan interface{}) <-chan interface{} {
		valStream := make(chan interface{})
		go func() {
			defer close(valStream)

			for i := 0; i < count; i++ {
				select {
				case <-done:
					return
				case valStream <- <-c:
				}
			}
		}()
		return valStream
	}
}
