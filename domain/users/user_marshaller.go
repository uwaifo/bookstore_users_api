package users

import "encoding/json"

//PublicUser  . . .
type PublicUser struct {
	ID          int64  `json:"id"`
	DateCreated string `json:"date_created"`
	Status      string `json:"status"`
}

//PrivateUser . . .
type PrivateUser struct {
	ID          int64  `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	DateCreated string `json:"date_created"`
	Status      string `json:"status"`
}

//Marshall single . . .
func (user *User) Marshall(isPublic bool) interface{} {
	if isPublic {
		return PublicUser{
			ID:          user.ID,
			DateCreated: user.DateCreated,
			Status:      user.Status,
		}
	}
	userJSON, _ := json.Marshal(user)

	var privateUser PrivateUser
	json.Unmarshal(userJSON, &privateUser)
	return privateUser

}

//Marshall slices . . .
func (users Users) Marshall(isPublic bool) []interface{} {
	//this is one way of handling/marshaling multiple returned objects
	//another way is to have the above functionality embeded in the Marshal domain function
	//as a Method for hadling  slices of User
	result := make([]interface{}, len(users))
	for index, user := range users {
		result[index] = user.Marshall(isPublic)
	}

	return result

}
