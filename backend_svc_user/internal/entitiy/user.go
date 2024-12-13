package entitiy

// 핵심 엔티티와 도메인 객체
type User struct {
	Email    string
	Password string
}

func (u *User) validate() error {
	return nil
}
