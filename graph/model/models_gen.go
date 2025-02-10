// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Link struct {
	ID      string `json:"id"`
	Title   string `json:"title"`
	Address string `json:"address"`
	User    *User  `json:"user"`
	Score   int32  `json:"score"`
}

type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Mutation struct {
}

type NewLink struct {
	Title   string `json:"title"`
	Address string `json:"address"`
}

type NewUser struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Query struct {
}

type RefreshTokenInput struct {
	Token string `json:"token"`
}

type Subscription struct {
}

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type VoteInput struct {
	LinkID string `json:"linkId"`
	Vote   int32  `json:"vote"`
}
