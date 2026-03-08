package auth

type UserEntity struct {
	UserId   string `gorm:"column:UserId;primaryKey"`
	FullName string `gorm:"column:FullName;size:255"`
	Email    string `gorm:"column:Email;size:255"`
	Picture  string `gorm:"column:Picture;size:255"`
}

type UserCredential struct {
	ID            string `json:"id"`
	Email         string `json:"email"`
	VerifiedEmail bool   `json:"verified_email"`
	Name          string `json:"name"`
	GivenName     string `json:"given_name"`
	FamilyName    string `json:"family_name"`
	Picture       string `json:"picture"`
	Locale        string `json:"locale"`
}
