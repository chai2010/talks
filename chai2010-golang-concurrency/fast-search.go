// +build ignore

package main

func main() {
    ch := make(chan string, 32)

    go func() {
        ch <- searchByBing("golang")
    }()
    go func() {
        ch <- searchByGoogle("golang")
    }()
    go func() {
        ch <- searchByBaidu("golang")
    }()

    fmt.Println(<-ch)
}

func searchByBing(key string) string {
	return ""
}

func searchByGoogle(key string) string {
	return ""
}

func searchByBaidu(key string) string {
	return ""
}
