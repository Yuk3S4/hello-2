package greet

// Identificadores en con la primer letra mayúscula se exportará y con la primer letra en minúscula NO se exportará
var emoji = "🙋‍♂️" // Variables a nivel de paquete se tienen que declarar con var

// Todas lo que se declare se pueden usar en todos lo archivos de la misma carpeta

// Retorna saludo en inglés
func English() string {
	return "Hi " + emoji
}

// Retorna saludo en italiano
func Italian() string {
	return "Ciao " + emoji
}
