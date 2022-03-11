package hadiscovery

////////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
////////////////////////////////////////////////////////////////////////////////
type store struct {
	Light struct {
		Availability map[string]string
		State        map[string]string
	}
}

func initStore() store {
	s := store{}
	s.Light.Availability = make(map[string]string)
	s.Light.State = make(map[string]string)
	return s
}
