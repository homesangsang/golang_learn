package main

import defer1 "golang_learn/errorhandling/defer/fun"

func main() {
	//defer1.TryDefer()
	//defer1.TryDefer1()

	defer1.WriteFileErr("fib.txt")
}
