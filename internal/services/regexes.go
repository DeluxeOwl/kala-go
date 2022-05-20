package services

import "regexp"

// valid regex for relations, permissions and type names
var regexPropertyName = regexp.MustCompile(`^[a-zA-Z_]{1,64}$`)
var regexTypeName = regexp.MustCompile(`^[a-zA-Z_]{1,64}$`)

// valid regexes for relation, permission values
var regexRelValue = regexp.MustCompile(`^[a-zA-Z_]{1,64}(#[a-zA-Z_]{1,64})?( \| [a-zA-Z_]{1,64}(#[a-zA-Z_]{1,64})?)*$`)

var regexPermValue = regexp.MustCompile(`^((!)?[a-zA-Z_]{1,64}(\.[a-zA-Z_]{1,64})?)((( \| )|( & ))((!)?[a-zA-Z_]{1,64}(\.[a-zA-Z_]{1,64})?))*$`)

// delimiter for the value of a relation
var refValueDelim = " | "

// delimiter for a subrelation
var refSubrelationDelim = "#"

// delimiter for parent relation in permission
var parentRelDelim = "."

// delimiter for permissions
var permDelim = regexp.MustCompile(`[a-zA-Z_]{1,64}(\.([a-zA-Z_]{1,64}))?`)


var regexSubjName = regexp.MustCompile(`^[a-zA-Z0-9\._\/-]{1,64}$`)
