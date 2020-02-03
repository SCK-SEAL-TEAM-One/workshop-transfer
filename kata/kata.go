package kata

import "time"

const yeaeOfHours = 8765.8

type Tranfer struct {
	Now func() time.Time
}
type UserInfo struct {
	UserID    int
	BirthDate time.Time
}

func (t Tranfer) CalculateAge(brithDate time.Time) int {
	age := t.Now().Sub(brithDate)
	return int(age.Hours() / yeaeOfHours)
}

func (t Tranfer) CalulateFee(userInfo UserInfo, amount float64) float64 {
	age := t.CalculateAge(userInfo.BirthDate)
	if age < 15 {
		return 20.0
	}
	return 15.0
}
