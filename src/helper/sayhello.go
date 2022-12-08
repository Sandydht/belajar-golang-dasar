package helper

var version = "1.0.0"      // Tidak bisa diakses dari luar
var Application = "golang" // Bisa diakses dari luar

func sayGoodBye(name string) string { // Tidak bisa diakses dari luar
	return "Good bye " + name
}

func SayHello(name string) string { // Bisa diakses dari luar
	return "Hello " + name
}
