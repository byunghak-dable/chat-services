package dto

type RegisterReqDto struct {
	email    string
	name     string
	imageUrl string
	token    string
}

type GoogleSigninReqDto struct {
	token string
}

type GoogleSigninResDto struct {
	idx      uint32
	email    string
	name     string
	imageUrl string
}
