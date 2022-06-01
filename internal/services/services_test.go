package services

import (
	"context"
	"testing"

	"github.com/DeluxeOwl/kala-go/ent/enttest"
	"github.com/DeluxeOwl/kala-go/internal/models"

	_ "github.com/mattn/go-sqlite3"
)

func TestTypeConfig(t *testing.T) {
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	defer client.Close()

	h := Handler{
		Db: client,
	}

	ctx := context.Background()

	tables := []struct {
		m             *models.TypeConfigReq
		name          string
		nrRelations   int
		nrPermissions int
	}{
		{&models.TypeConfigReq{Name: "user"}, "user", 0, 0},
		{&models.TypeConfigReq{
			Name: "group",
			Relations: map[string]string{
				"member": "user",
			},
		}, "group", 1, 0},
		{&models.TypeConfigReq{
			Name: "folder",
			Relations: map[string]string{
				"reader": "user | group#member",
			},
		}, "folder", 1, 0},
		{&models.TypeConfigReq{
			Name: "document",
			Relations: map[string]string{
				"parent_folder": "folder",
				"writer":        "user",
				"reader":        "user",
			},
			Permissions: map[string]string{
				"read":           "reader | writer | parent_folder.reader",
				"read_and_write": "reader & writer",
				"read_only":      "reader & !writer",
			},
		}, "document", 3, 3},
	}
	for _, table := range tables {
		tc, err := h.CreateTypeConfig(ctx, table.m)
		if err != nil {
			t.Fatalf("on type: %s, shouldn't error: %s", table.m.Name, err)
		}
		if tc.Name != table.m.Name {
			t.Errorf("name %s is incorrect, wanted %s", tc.Name, table.m.Name)
		}

		nrRelations, err := tc.Unwrap().QueryRelations().Count(ctx)
		if err != nil {
			t.Fatalf("on type: %s, shouldn't error: %s", table.m.Name, err)
		}
		if nrRelations != table.nrRelations {
			t.Errorf("on type: %s, number of relations %d is incorrect, wanted %d", table.m.Name, nrRelations, table.nrRelations)
		}

		nrPermissions, err := tc.QueryPermissions().Count(ctx)
		if err != nil {
			t.Fatalf("on type: %s, shouldn't error: %s", table.m.Name, err)
		}
		if nrPermissions != table.nrPermissions {
			t.Errorf("on type: %s, number of permissions %d is incorrect, wanted %d", table.m.Name, nrPermissions, table.nrPermissions)
		}

	}
}

func TestSubjectCreation(t *testing.T) {
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	defer client.Close()

	h := Handler{
		Db: client,
	}

	ctx := context.Background()

	h.CreateTypeConfig(ctx,
		&models.TypeConfigReq{Name: "user"},
	)

	h.CreateTypeConfig(ctx,
		&models.TypeConfigReq{
			Name: "group",
			Relations: map[string]string{
				"member": "user",
			},
		})

	h.CreateTypeConfig(ctx,
		&models.TypeConfigReq{
			Name: "folder",
			Relations: map[string]string{
				"reader": "user | group#member",
			},
		})
	h.CreateTypeConfig(ctx,
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
				"read_only":      "reader & !writer",
			},
		})

	tables := []struct {
		s    models.SubjectReq
		name string
	}{
		{models.SubjectReq{
			TypeConfigName: "document",
			SubjectName:    "report.csv",
		}, "report.csv"},
		{models.SubjectReq{
			TypeConfigName: "user",
			SubjectName:    "anna",
		}, "anna"},
		{models.SubjectReq{
			TypeConfigName: "user",
			SubjectName:    "john",
		}, "john"},
		{models.SubjectReq{
			TypeConfigName: "user",
			SubjectName:    "steve",
		}, "steve"},
		{models.SubjectReq{
			TypeConfigName: "folder",
			SubjectName:    "secret_folder",
		}, "secret_folder"},
		{models.SubjectReq{
			TypeConfigName: "group",
			SubjectName:    "dev",
		}, "dev"},
		{models.SubjectReq{
			TypeConfigName: "group",
			SubjectName:    "test_group",
		}, "test_group"},
	}

	for _, table := range tables {
		subj, err := h.CreateSubject(ctx, &table.s)
		if err != nil {
			t.Fatalf("at subject %s, shouldn't get error %s", table.name, err)
		}
		if subj.Name != table.name {
			t.Errorf("wanted %s, got %s", table.name, subj.Name)
		}
	}

}

func TestPermissionCheck(t *testing.T) {
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	defer client.Close()

	h := Handler{
		Db: client,
	}

	ctx := context.Background()

	h.CreateTypeConfig(ctx,
		&models.TypeConfigReq{Name: "user"},
	)

	h.CreateTypeConfig(ctx,
		&models.TypeConfigReq{
			Name: "group",
			Relations: map[string]string{
				"member": "user",
			},
		})

	h.CreateTypeConfig(ctx,
		&models.TypeConfigReq{
			Name: "folder",
			Relations: map[string]string{
				"reader": "user | group#member",
			},
		})

	h.CreateTypeConfig(ctx,
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
				"read_only":      "reader & !writer",
			},
		})

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
		h.CreateSubject(ctx, &v)
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
		h.CreateTuple(ctx, &v)
	}

	resource := &models.SubjectReq{
		TypeConfigName: "document",
		SubjectName:    "report.csv",
	}

	tables := []struct {
		subject    *models.SubjectReq
		permission string
		hasPerm    bool
	}{
		{
			&models.SubjectReq{
				TypeConfigName: "user",
				SubjectName:    "john",
			}, "read", true,
		},
		{
			&models.SubjectReq{
				TypeConfigName: "user",
				SubjectName:    "anna",
			}, "read", true,
		},
		{
			&models.SubjectReq{
				TypeConfigName: "user",
				SubjectName:    "steve",
			}, "read", true,
		},
		{
			&models.SubjectReq{
				TypeConfigName: "user",
				SubjectName:    "anna",
			}, "read_only", false,
		},
		{
			&models.SubjectReq{
				TypeConfigName: "user",
				SubjectName:    "anna",
			}, "read_and_write", true,
		},
		{
			&models.SubjectReq{
				TypeConfigName: "user",
				SubjectName:    "steve",
			}, "read_and_write", false,
		},
		{
			&models.SubjectReq{
				TypeConfigName: "user",
				SubjectName:    "john",
			}, "read_and_write", false,
		},
	}

	for i := 0; i < 100; i++ {
		for _, table := range tables {
			hasPerm, err := h.CheckPermission(ctx, &models.TupleReqPermission{
				Subject:    table.subject,
				Permission: table.permission,
				Resource:   resource,
			})
			if err != nil {
				if err != nil {
					t.Fatalf("at subject `%s` and permission `%s`, shouldn't get error %s",
						table.subject.SubjectName, table.permission, err)
				}
				if hasPerm != table.hasPerm {
					t.Errorf("at subject `%s` and permission %s, wanted `%t` got `%t`",
						table.subject.SubjectName, table.permission, table.hasPerm, hasPerm)
				}
			}
		}
	}

}
