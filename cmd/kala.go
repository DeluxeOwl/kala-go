package main

import (
	"context"
	"fmt"
	"log"

	"github.com/DeluxeOwl/kala-go/ent"
	"github.com/DeluxeOwl/kala-go/internal/models"
	"github.com/DeluxeOwl/kala-go/internal/services"
	_ "github.com/mattn/go-sqlite3"
)

func main() {

	// TODO: handler struct
	client, err := ent.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	defer client.Close()
	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	ctx := context.Background()

	h := services.Handler{
		Client: client,
	}

	tc, err := h.CreateTypeConfig(ctx,
		&models.TypeConfigReq{Name: "user"},
	)
	fmt.Println(tc, err)

	tc, err = h.CreateTypeConfig(ctx,
		&models.TypeConfigReq{
			Name: "group",
			Relations: map[string]string{
				"member": "user",
			},
		})
	fmt.Println(tc, err)

	tc, err = h.CreateTypeConfig(ctx,
		&models.TypeConfigReq{
			Name: "folder",
			Relations: map[string]string{
				"reader": "user | group#member",
			},
		})
	fmt.Println(tc, err)

	tc, err = h.CreateTypeConfig(ctx,
		&models.TypeConfigReq{
			Name: "document",
			Relations: map[string]string{
				"parent_folder": "folder",
				"writer":        "user",
				"reader":        "user",
			},
			Permissions: map[string]string{
				"read":           "reader | writer | parent_folder.reader",
				"read_and_write": "reader & writer",
				"read_only":      "reader | !writer",
			},
		})
	fmt.Println(tc, err)

	subjects := []models.SubjectReq{
		{
			TypeConfigName: "document",
			SubjectName:    "report.csv",
		},
		{
			TypeConfigName: "user",
			SubjectName:    "anna",
		},
		{
			TypeConfigName: "user",
			SubjectName:    "john",
		},
		{
			TypeConfigName: "user",
			SubjectName:    "steve",
		},
		{
			TypeConfigName: "folder",
			SubjectName:    "secret_folder",
		},
		{
			TypeConfigName: "group",
			SubjectName:    "dev",
		},
		{
			TypeConfigName: "group",
			SubjectName:    "test_group",
		},
	}

	for _, v := range subjects {
		subj, err := h.CreateSubject(ctx, &v)
		fmt.Println(subj, err)
	}

	tuples := []models.TupleReqRelation{
		{
			Subject: &models.SubjectReq{
				TypeConfigName: "user",
				SubjectName:    "anna",
			},
			Relation: "reader",
			Resource: &models.SubjectReq{
				TypeConfigName: "document",
				SubjectName:    "report.csv",
			},
		},
		{
			Subject: &models.SubjectReq{
				TypeConfigName: "user",
				SubjectName:    "anna",
			},
			Relation: "writer",
			Resource: &models.SubjectReq{
				TypeConfigName: "document",
				SubjectName:    "report.csv",
			},
		},
		{
			Subject: &models.SubjectReq{
				TypeConfigName: "folder",
				SubjectName:    "secret_folder",
			},
			Relation: "parent_folder",
			Resource: &models.SubjectReq{
				TypeConfigName: "document",
				SubjectName:    "report.csv",
			},
		},
		{
			Subject: &models.SubjectReq{
				TypeConfigName: "user",
				SubjectName:    "john",
			},
			Relation: "reader",
			Resource: &models.SubjectReq{
				TypeConfigName: "folder",
				SubjectName:    "secret_folder",
			},
		},
		{
			Subject: &models.SubjectReq{
				TypeConfigName: "user",
				SubjectName:    "john",
			},
			Relation: "member",
			Resource: &models.SubjectReq{
				TypeConfigName: "group",
				SubjectName:    "dev",
			},
		},
		{
			Subject: &models.SubjectReq{
				TypeConfigName: "group",
				SubjectName:    "dev#member",
			},
			Relation: "reader",
			Resource: &models.SubjectReq{
				TypeConfigName: "folder",
				SubjectName:    "secret_folder",
			},
		},
		{
			Subject: &models.SubjectReq{
				TypeConfigName: "group",
				SubjectName:    "test_group#member",
			},
			Relation: "reader",
			Resource: &models.SubjectReq{
				TypeConfigName: "folder",
				SubjectName:    "secret_folder",
			},
		},
		{
			Subject: &models.SubjectReq{
				TypeConfigName: "user",
				SubjectName:    "steve",
			},
			Relation: "member",
			Resource: &models.SubjectReq{
				TypeConfigName: "group",
				SubjectName:    "dev",
			},
		},
	}

	for _, v := range tuples {
		tuple, err := h.CreateTuple(ctx, &v)
		fmt.Println(tuple, err)
	}

	hasRead, err := h.CheckPermission(ctx, &models.TupleReqPermission{
		Subject: &models.SubjectReq{
			TypeConfigName: "user",
			SubjectName:    "steve",
		},
		Permission: "read",
		Resource: &models.SubjectReq{
			TypeConfigName: "document",
			SubjectName:    "report.csv",
		},
	})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("⟶\tPermission:", hasRead)

	// TEST: empty permissions
	// tc, err = h.CreateTypeConfig(ctx,
	// 	&TypeConfig{
	// 		Name: "test",
	// 		Permissions: map[string]string{
	// 			"read":           "reader | writer | parent_folder.reader",
	// 			"read_and_write": "reader & writer",
	// 			"read_only":      "reader & !writer",
	// 		},
	// 	})
	// fmt.Println(tc, err)

	// reader | writer | parent_folder.reader
	// reader & writer
	// reader | writer
	// reader | !writer
	// !writer
	// reader
	// reader & writer | !tester
	// !tester & reader | writer

	// reader |
	// !writer |
	// reader &
	// writer & reader | !
}
