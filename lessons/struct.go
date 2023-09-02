package main

func main() {
	user1 := &User{
		ID:        1,
		FirstName: "Test",
		LastName:  "DEMİR",
		UserName:  "TestD",
		Age:       1,
		Pay: &Payment{
			Salary: 10000,
			Bonus:  50,
		},
	}
}

// ? Kullanıcı yapısı
type User struct {
	ID        int
	FirstName string
	LastName  string
	UserName  string
	Age       int
	Pay       *Payment
}

// ? kullanıcı Yapıcı methodu
func NewUser() *User {
	u := new(User)
	u.Pay = NewPayment()
	return u
}

// ? Ödeme yapısı
type Payment struct {
	Salary float32
	Bonus  float32
}

// ? Ödeme Yapıcı methodu
func NewPayment() *Payment {
	p := new(Payment)

	return p
}

func (u *User) getUserName() string {
	return u.UserName
}
