package main

import "sync"

func main() {
	// syncパッケージの使用
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		// 重い処理
	}()
	// 別の重い処理
	wg.Wait()

}
