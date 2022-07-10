package Models

type Student struct {
	Id       uint   `json:"id"`
	Name     string `json:"name"`
	LastName string `json:"last_name"`
	Dob      string `json:"dob"`
	Subject  string `json:"subject"`
	Address  string `json:"address"`
	Marks    int    `json:"marks"`
}

func (b *Student) TableName() string {
	return "student"
}
