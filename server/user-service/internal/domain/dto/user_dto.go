package dto

type RegisterReqDto struct {
	Email    string
	Name     string
	ImageUrl string
	Token    string
}

type GoogleSigninReqDto struct {
	Token string
}

type GoogleSigninResDto struct {
	Email    string
	Name     string
	ImageUrl string
	Idx      uint
}
