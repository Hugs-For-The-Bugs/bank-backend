package middleware

type UserRepository struct {
	// Your database connection or other fields can go here
}

func New() *UserRepository {
	return &UserRepository{}
}
func (repo *UserRepository) FindBySSN(ssn string) User {
	// Your code to fetch user from database by Social Security Number
	// For the sake of this example, I'm returning a dummy user.
	return User{SocialSecurityNumber: ssn, Pass: "password"}
}

type User struct {
	SocialSecurityNumber string
	Pass                 string
}

//// Dummy UserRepository to simulate fetching a user by SSN
//func FindBySSN(ssn string) User {
//	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("mysecretpassword"), bcrypt.DefaultCost)
//	return User{SocialSecurityNumber: ssn, Pass: string(hashedPassword)}
//}
