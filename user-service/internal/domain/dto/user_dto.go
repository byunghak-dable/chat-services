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
	Idx      uint
	Email    string
	Name     string
	ImageUrl string
}
