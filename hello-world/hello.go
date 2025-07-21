package helloworld

func hello(name, lang string) string {
	var helloPrefix string

	if name == "" {
		name = "World"
	}

	switch lang {
	case "English":
		helloPrefix = "Hello, "
	case "Spanish":
		helloPrefix = "Hola, "
	case "French":
		helloPrefix = "Bonjour, "
	default:
		helloPrefix = "Hello, "
	}

	return helloPrefix + name
}
