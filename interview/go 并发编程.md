## go 并发编程

```golang
func say(s string) {
	fmt.Printf("%s say\n", s)
}
func main() {
	go say("lisi")
	say("zhangsan")
}
```