package dto

type RegisterReqDto struct {
	email    string
	name     string
	imageUrl string
	token    string
}

type SigninReqDto struct {
	token string
}

type SigninResDto struct {
	idx      uint32
	email    string
	name     string
	imageUrl string
}
