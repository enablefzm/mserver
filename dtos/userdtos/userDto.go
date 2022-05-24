package userdtos

type UserInfoDto struct {
	GUID string 			`json:"guid"`
	UID string				`json:"uid"`
	Name string 			`json:"name"`
	Sex bool 				`json:"sex"`
	IdentityCard string 	`json:"identity_card"`
	PhoneNumber string 		`json:"phone_number"`
	TelephoneNumber string 	`json:"telephone_number"`
	IsOnJob bool 			`json:"is_on_job"`
}
