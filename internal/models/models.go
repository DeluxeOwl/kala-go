package models

import "github.com/DeluxeOwl/kala-go/ent"

type TypeConfigReq struct {
	Name        string            `json:"type"`
	Relations   map[string]string `json:"relations"`
	Permissions map[string]string `json:"permissions"`
}

type SubjectReq struct {
	TypeConfigName string
	SubjectName    string
}

type TupleReqRelation struct {
	Subject  *SubjectReq
	Relation string
	Resource *SubjectReq
}

type RelationCheck struct {
	Subj *ent.Subject
	Rel  *ent.Relation
	Res  *ent.Subject
}

type PermissionCheck struct {
	Subj *ent.Subject
	Perm *ent.Permission
	Res  *ent.Subject
}

type TupleReqPermission struct {
	Subject    *SubjectReq
	Permission string
	Resource   *SubjectReq
}
