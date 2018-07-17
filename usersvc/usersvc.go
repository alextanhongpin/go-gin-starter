package usersvc

// New initialize a new usersvc
func New() *Service {
	evt := NewEvent()
	rep := NewRepository()
	return NewService(rep, evt)
}
