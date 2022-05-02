// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// PermissionsColumns holds the columns for the "permissions" table.
	PermissionsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
		{Name: "value", Type: field.TypeString},
		{Name: "type_config_permissions", Type: field.TypeInt, Nullable: true},
	}
	// PermissionsTable holds the schema information for the "permissions" table.
	PermissionsTable = &schema.Table{
		Name:       "permissions",
		Columns:    PermissionsColumns,
		PrimaryKey: []*schema.Column{PermissionsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "permissions_type_configs_permissions",
				Columns:    []*schema.Column{PermissionsColumns[3]},
				RefColumns: []*schema.Column{TypeConfigsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// RelationsColumns holds the columns for the "relations" table.
	RelationsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
		{Name: "value", Type: field.TypeString},
		{Name: "type_config_relations", Type: field.TypeInt, Nullable: true},
	}
	// RelationsTable holds the schema information for the "relations" table.
	RelationsTable = &schema.Table{
		Name:       "relations",
		Columns:    RelationsColumns,
		PrimaryKey: []*schema.Column{RelationsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "relations_type_configs_relations",
				Columns:    []*schema.Column{RelationsColumns[3]},
				RefColumns: []*schema.Column{TypeConfigsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// SubjectsColumns holds the columns for the "subjects" table.
	SubjectsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString, Unique: true},
		{Name: "type_config_subjects", Type: field.TypeInt, Nullable: true},
	}
	// SubjectsTable holds the schema information for the "subjects" table.
	SubjectsTable = &schema.Table{
		Name:       "subjects",
		Columns:    SubjectsColumns,
		PrimaryKey: []*schema.Column{SubjectsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "subjects_type_configs_subjects",
				Columns:    []*schema.Column{SubjectsColumns[2]},
				RefColumns: []*schema.Column{TypeConfigsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// TuplesColumns holds the columns for the "tuples" table.
	TuplesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "subject_id", Type: field.TypeInt},
		{Name: "relation_id", Type: field.TypeInt},
		{Name: "resource_id", Type: field.TypeInt},
	}
	// TuplesTable holds the schema information for the "tuples" table.
	TuplesTable = &schema.Table{
		Name:       "tuples",
		Columns:    TuplesColumns,
		PrimaryKey: []*schema.Column{TuplesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "tuples_subjects_subject",
				Columns:    []*schema.Column{TuplesColumns[1]},
				RefColumns: []*schema.Column{SubjectsColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "tuples_relations_relation",
				Columns:    []*schema.Column{TuplesColumns[2]},
				RefColumns: []*schema.Column{RelationsColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "tuples_subjects_resource",
				Columns:    []*schema.Column{TuplesColumns[3]},
				RefColumns: []*schema.Column{SubjectsColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// TypeConfigsColumns holds the columns for the "type_configs" table.
	TypeConfigsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString, Unique: true},
	}
	// TypeConfigsTable holds the schema information for the "type_configs" table.
	TypeConfigsTable = &schema.Table{
		Name:       "type_configs",
		Columns:    TypeConfigsColumns,
		PrimaryKey: []*schema.Column{TypeConfigsColumns[0]},
	}
	// PermissionRelationsColumns holds the columns for the "permission_relations" table.
	PermissionRelationsColumns = []*schema.Column{
		{Name: "permission_id", Type: field.TypeInt},
		{Name: "relation_id", Type: field.TypeInt},
	}
	// PermissionRelationsTable holds the schema information for the "permission_relations" table.
	PermissionRelationsTable = &schema.Table{
		Name:       "permission_relations",
		Columns:    PermissionRelationsColumns,
		PrimaryKey: []*schema.Column{PermissionRelationsColumns[0], PermissionRelationsColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "permission_relations_permission_id",
				Columns:    []*schema.Column{PermissionRelationsColumns[0]},
				RefColumns: []*schema.Column{PermissionsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "permission_relations_relation_id",
				Columns:    []*schema.Column{PermissionRelationsColumns[1]},
				RefColumns: []*schema.Column{RelationsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// RelationRelTypeconfigsColumns holds the columns for the "relation_rel_typeconfigs" table.
	RelationRelTypeconfigsColumns = []*schema.Column{
		{Name: "relation_id", Type: field.TypeInt},
		{Name: "type_config_id", Type: field.TypeInt},
	}
	// RelationRelTypeconfigsTable holds the schema information for the "relation_rel_typeconfigs" table.
	RelationRelTypeconfigsTable = &schema.Table{
		Name:       "relation_rel_typeconfigs",
		Columns:    RelationRelTypeconfigsColumns,
		PrimaryKey: []*schema.Column{RelationRelTypeconfigsColumns[0], RelationRelTypeconfigsColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "relation_rel_typeconfigs_relation_id",
				Columns:    []*schema.Column{RelationRelTypeconfigsColumns[0]},
				RefColumns: []*schema.Column{RelationsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "relation_rel_typeconfigs_type_config_id",
				Columns:    []*schema.Column{RelationRelTypeconfigsColumns[1]},
				RefColumns: []*schema.Column{TypeConfigsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		PermissionsTable,
		RelationsTable,
		SubjectsTable,
		TuplesTable,
		TypeConfigsTable,
		PermissionRelationsTable,
		RelationRelTypeconfigsTable,
	}
)

func init() {
	PermissionsTable.ForeignKeys[0].RefTable = TypeConfigsTable
	RelationsTable.ForeignKeys[0].RefTable = TypeConfigsTable
	SubjectsTable.ForeignKeys[0].RefTable = TypeConfigsTable
	TuplesTable.ForeignKeys[0].RefTable = SubjectsTable
	TuplesTable.ForeignKeys[1].RefTable = RelationsTable
	TuplesTable.ForeignKeys[2].RefTable = SubjectsTable
	PermissionRelationsTable.ForeignKeys[0].RefTable = PermissionsTable
	PermissionRelationsTable.ForeignKeys[1].RefTable = RelationsTable
	RelationRelTypeconfigsTable.ForeignKeys[0].RefTable = RelationsTable
	RelationRelTypeconfigsTable.ForeignKeys[1].RefTable = TypeConfigsTable
}
