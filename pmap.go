package pmap

type Func func(int) bool

func Pmap(n int, m int, fn Func) bool {

	done := make(chan bool)
	fin := make(chan bool)
	gate := make(chan int)
	go func() {
		// seed
		for i := 0; i < m; i++ {
			gate <- 0
		}
	}()

	go func() {
		// harvest
		res := true
		for i := 0; i < n; i++ {
			res = res && <-fin
		}
		done <- res
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
