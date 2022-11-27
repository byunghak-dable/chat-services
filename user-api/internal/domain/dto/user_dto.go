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
	Idx      uint32
	Email    string
	Name     string
	ImageUrl string
}
