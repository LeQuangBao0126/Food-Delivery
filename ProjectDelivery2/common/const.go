package common


const (
	DbTypeRestaurant = 1
	DbTypeFood = 2
	DbTypeCategory  =3
	DbTypeUser  = 4
)

type Requester interface{
	GetUserId() int
	GetEmail() string
	GetRole() string
}

