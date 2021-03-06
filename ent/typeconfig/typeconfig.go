// Code generated by entc, DO NOT EDIT.

package typeconfig

const (
	// Label holds the string label denoting the typeconfig type in the database.
	Label = "type_config"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// EdgeRelations holds the string denoting the relations edge name in mutations.
	EdgeRelations = "relations"
	// EdgePermissions holds the string denoting the permissions edge name in mutations.
	EdgePermissions = "permissions"
	// EdgeSubjects holds the string denoting the subjects edge name in mutations.
	EdgeSubjects = "subjects"
	// EdgeRelTypeconfigs holds the string denoting the rel_typeconfigs edge name in mutations.
	EdgeRelTypeconfigs = "rel_typeconfigs"
	// Table holds the table name of the typeconfig in the database.
	Table = "type_configs"
	// RelationsTable is the table that holds the relations relation/edge.
	RelationsTable = "relations"
	// RelationsInverseTable is the table name for the Relation entity.
	// It exists in this package in order to avoid circular dependency with the "relation" package.
	RelationsInverseTable = "relations"
	// RelationsColumn is the table column denoting the relations relation/edge.
	RelationsColumn = "type_config_relations"
	// PermissionsTable is the table that holds the permissions relation/edge.
	PermissionsTable = "permissions"
	// PermissionsInverseTable is the table name for the Permission entity.
	// It exists in this package in order to avoid circular dependency with the "permission" package.
	PermissionsInverseTable = "permissions"
	// PermissionsColumn is the table column denoting the permissions relation/edge.
	PermissionsColumn = "type_config_permissions"
	// SubjectsTable is the table that holds the subjects relation/edge.
	SubjectsTable = "subjects"
	// SubjectsInverseTable is the table name for the Subject entity.
	// It exists in this package in order to avoid circular dependency with the "subject" package.
	SubjectsInverseTable = "subjects"
	// SubjectsColumn is the table column denoting the subjects relation/edge.
	SubjectsColumn = "type_config_subjects"
	// RelTypeconfigsTable is the table that holds the rel_typeconfigs relation/edge. The primary key declared below.
	RelTypeconfigsTable = "relation_rel_typeconfigs"
	// RelTypeconfigsInverseTable is the table name for the Relation entity.
	// It exists in this package in order to avoid circular dependency with the "relation" package.
	RelTypeconfigsInverseTable = "relations"
)

// Columns holds all SQL columns for typeconfig fields.
var Columns = []string{
	FieldID,
	FieldName,
}

var (
	// RelTypeconfigsPrimaryKey and RelTypeconfigsColumn2 are the table columns denoting the
	// primary key for the rel_typeconfigs relation (M2M).
	RelTypeconfigsPrimaryKey = []string{"relation_id", "type_config_id"}
)

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}
