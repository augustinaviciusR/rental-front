package model

import "fmt"

type InquiryForm struct {
	Name        string `form:"name"`
	Email       string `form:"email"`
	Description string `form:"description"`
}

func (i *InquiryForm) String() string {
	return fmt.Sprintf("name: %s, email: %s, description: %s", i.Name, i.Email, i.Description)
}
