package model
import (
"github.com/google/uuid"
)

func CU() string {
	uuid, err := uuid.NewUUID()
	if err != nil {
	return "" 
	}
return uuid.String()
}
