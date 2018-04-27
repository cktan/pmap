package pmap

type Func func(int) error

func Pmap(n int, m int, fn Func) error {

	done := make(chan error)
	fin := make(chan error)
	gate := make(chan int)
	go func() {
		// seed
		for i := 0; i < m; i++ {
			gate <- 0
		}
	}()

	go func() {
		// harvest
                var err error
		for i := 0; i < n && err == nil; i++ {
			err <-fin
		}
		done <- err
	}()

	for i := 0; i < n; i++ {
		<-gate
		go func() {
			fin <- fn(i)
			gate <- 0
		}()
	}

	return <-done
}
