package services

import "testing"

func TestPropertyNames(t *testing.T) {
	tables := []struct {
		propertyName string
		isValid      bool
	}{
		{"reader", true},
		{"writer", true},
		{"Reader", true},
		{"READER", true},
		{"", false},
		{"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa", false},
		{"testing_property", true},
		{"____", true},
		{"test____PROP", true},
		{"test!____PROP", false},
		{"test____PROP$", false},
	}

	for _, table := range tables {
		isMatch := regexPropertyName.MatchString(table.propertyName)
		if isMatch != table.isValid {
			t.Errorf("property name %s, wanted %t, got %t", table.propertyName, table.isValid, isMatch)
		}
	}
}
func TestTypeName(t *testing.T) {
	tables := []struct {
		propertyName string
		isValid      bool
	}{
		{"document", true},
		{"group", true},
		{"", false},
		{"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa", false},
		{"report_list_test", true},
		{"____", true},
		{"test____PROP", true},
		{"test!____PROP", false},
		{"test____PROP$", false},
	}

	for _, table := range tables {
		isMatch := regexTypeName.MatchString(table.propertyName)
		if isMatch != table.isValid {
			t.Errorf("type name %s, wanted %t, got %t", table.propertyName, table.isValid, isMatch)
		}
	}
}

func TestPermValue(t *testing.T) {
	tables := []struct {
		permValue string
		isValid   bool
	}{
		{"reader | writer | parent_folder.reader", true},
		{"reader & writer", true},
		{"", false},
		{"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa", false},
		{"report_list_test", true},
		{"____", true},
		{"test____PROP", true},
		{"test!____PROP", false},
		{"test____PROP#", false},
		{"reader | writer", true},
		{"reader | !writer", true},
		{"!writer", true},
		{"reader", true},
		{"reader & writer | !tester", true},
		{"!tester & reader | writer", true},
		{"(reader | writer) & !parent_folder.tester", false},
		{"reader |", false},
		{"!writer |", false},
		{"reader &", false},
		{"writer & reader | !", false},
	}

	for _, table := range tables {
		isMatch := regexPermValue.MatchString(table.permValue)

		if isMatch != table.isValid {
			t.Errorf("perm value %s, wanted %t, got %t", table.permValue, table.isValid, isMatch)
		}
	}
}
func TestRelValue(t *testing.T) {
	tables := []struct {
		relValue string
		isValid  bool
	}{
		{"group", true},
		{"user", true},
		{"", false},
		{"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa", false},
		{"report_list_test", true},
		{"____", true},
		{"test____PROP", true},
		{"test!____PROP", false},
		{"test____PROP#", false},
		{"group#member", true},
		{"user | group#member", true},
		{"user   | group#member", false},
		{"user | ", false},
		{"user |", false},
		{"|", false},
		{"user | group | dude | test_rule", true},
		{"group#member | group#member | group#member | group#member", true},
		{"group#member | group#member | group#member | group#", false},
		{"group#member | group#member | group# | user", false},
	}

	for _, table := range tables {
		isMatch := regexRelValue.MatchString(table.relValue)
		if isMatch != table.isValid {
			t.Errorf("rel value %s, wanted %t, got %t", table.relValue, table.isValid, isMatch)
		}
	}
}
func TestSubjName(t *testing.T) {
	tables := []struct {
		subjName string
		isValid  bool
	}{
		{"group", true},
		{"user", true},
		{"", false},
		{"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa", false},
		{"report_list_test", true},
		{"____", true},
		{"test____PROP", true},
		{"test!____PROP", false},
		{"test____PROP#", false},
		{"asd9asd90-12dd21", true},
		{"org/jwt", true},
		{"kebab-case", true},
	}

	for _, table := range tables {
		isMatch := regexSubjName.MatchString(table.subjName)
		if isMatch != table.isValid {
			t.Errorf("rel value %s, wanted %t, got %t", table.subjName, table.isValid, isMatch)
		}
	}
}
