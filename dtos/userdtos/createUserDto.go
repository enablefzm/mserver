package userdtos

import "strings"

type CreateUpdateUserDto struct {
	UID string				`json:"uid"`
	Name string 			`json:"name"`
	PassWord string			`json:"password"`
	Sex bool				`json:"sex"`
	IdentityCard string 	`json:"identity_card"`
	PhoneNumber string  	`json:"phone_number"`
	TelephoneNumber string 	`json:"telephone_number"`
	IsOnJob bool			`json:"is_on_job"`
}

func (p *CreateUpdateUserDto) GetUid() string {
	return strings.Replace(strings.ToLower(p.UID), " ", "", -1)
}

func (p *CreateUpdateUserDto) GetPassword() string {
	return strings.Replace(p.PassWord, " ", "", -1)
}
