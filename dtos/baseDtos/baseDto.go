package baseDtos

import "time"

type BaseDto struct {
	Guid string 			`json:"guid"`	// domainID 由GUID创建
	CreatedAt time.Time		`json:"created_at"`
	UpdatedAt time.Time		`json:"updated_at"`
}
