package main

const helloPrefix = "Hello, "
const thaiHelloPrefix = "สวัสดี, "

func main() {
}

// Hello for TDD
func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}
	if language == "thai" {
		return thaiHelloPrefix + name
	}
	return helloPrefix + name
}
