package addons

type Resource struct {
}

func LoadResource(path string) (*Resource, error) {
	return new(Resource), nil
}

func LoadResources(path string) ([]*Resource, error) {
	return make([]*Resource, 0), nil
}
