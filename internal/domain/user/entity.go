package user

import "golang.org/x/crypto/bcrypt"

type User struct {
	ID        string  `json:"id" gorm:"type:uuid;primaryKey"`
	FirstName string  `json:"first_name" gorm:"type:varchar(100);not null"`
	LastName  *string `json:"last_name,omitempty" gorm:"type:varchar(100)"`
	Password  string  `json:"password" gorm:"type:varchar(255);not null"`
	Email     string  `json:"email" gorm:"type:varchar(255);not null;unique"`
	CreatedAt string  `json:"created_at" gorm:"type:timestamptz;not null;default:now()"`
	UpdatedAt string  `json:"updated_at" gorm:"type:timestamptz;not null;default:now()"`
	DeletedAt *string `json:"deleted_at,omitempty" gorm:"type:timestamptz"`
}

// HashPassword hashes the provided plain-text password and sets the hashed password in the User struct.
func (u *User) HashPassword(password string) error {
	bPassword := []byte(password)

	hashedPassword, err := bcrypt.GenerateFromPassword(bPassword, bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	u.Password = string(hashedPassword)
	return nil
}

// CheckPassword compares the provided plain-text password with the hashed password stored in the User struct.
func (u *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}
