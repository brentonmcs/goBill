package models

//RegisterService json model when a service tells Bill about itself
type RegisterService struct {
	BaseURI     string
	ServiceName string
}
