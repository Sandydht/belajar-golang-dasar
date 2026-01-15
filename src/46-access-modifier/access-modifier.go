package accessmodifier

var version = "1.0.0" // tidak bisa diakses dari luar package
var Application = "golang"

func sayGoodBye(name string) string { // tidak bisa diakses dari luar package
	return "Good bye " + name
}

func SayHello(name string) string {
	return "Hello " + name
}
