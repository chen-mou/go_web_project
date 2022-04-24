package model

import "project/main/module/user/entity"

type Auth interface {
	auth(user entity.User) bool
}

//func getAuth(auth string) Auth {
//
//}

type public struct {
}

type onlySelf struct {
}

//func (public) auth(user entity.User) bool {
//
//}
