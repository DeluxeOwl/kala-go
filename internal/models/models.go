package models

import "github.com/DeluxeOwl/kala-go/ent"

type TypeConfigReq struct {
	Name        string            `json:"type"`
	Relations   map[string]string `json:"relations"`
	Permissions map[string]string `json:"permissions"`
}

type SubjectReq struct {
	TypeConfigName string `json:"type"`
	SubjectName    string `json:"name"`
}

type TupleReqRelation struct {
	Subject  *SubjectReq `json:"subject"`
	Relation string      `json:"relation"`
	Resource *SubjectReq `json:"resource"`
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
	Subject    *SubjectReq `json:"subject"`
	Permission string      `json:"permission"`
	Resource   *SubjectReq `json:"resource"`
}
