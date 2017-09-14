package helper

type ServiceProvider struct {
	Schema        map[string]*Schema
	ResourceMap   map[string]*Resource
	Doc           *Doc
	ConfigureFunc ConfigureFunc
}
