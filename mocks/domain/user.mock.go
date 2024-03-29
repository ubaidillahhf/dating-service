package domain_mock

import "github.com/ubaidillahhf/dating-service/app/domain"

func MakeMockUser() domain.User {
	return domain.User{
		Fullname: "a",
		Username: "a",
		Email:    "a@gmail.com",
		Password: "a",
	}
}

func MakeMockUserRegister() domain.RegisterRequest {
	return domain.RegisterRequest{
		Fullname: "a",
		Username: "a",
		Email:    "a@gmail.com",
		Password: "a",
	}
}
