package services

import (
	"context"
	"errors"
	"fmt"
	"log"
	"strings"

	"github.com/DeluxeOwl/kala-go/ent"
	"github.com/DeluxeOwl/kala-go/ent/permission"
	"github.com/DeluxeOwl/kala-go/ent/relation"
	"github.com/DeluxeOwl/kala-go/ent/subject"
	"github.com/DeluxeOwl/kala-go/ent/tuple"
	"github.com/DeluxeOwl/kala-go/ent/typeconfig"
	"github.com/DeluxeOwl/kala-go/internal/models"
	"golang.org/x/exp/maps"
	"golang.org/x/exp/slices"
)

// TODO: make the errors as variables?
func (h *Handler) CreateTypeConfig(ctx context.Context, tcInput *models.TypeConfigReq) (*ent.TypeConfig, error) {

	if !regexTypeName.MatchString(tcInput.Name) {
		return nil, fmt.Errorf("malformed type name input: '%s'", tcInput.Name)
	}

	tc, err := h.Client.TypeConfig.Create().
		SetName(tcInput.Name).
		Save(ctx)

	if err != nil {
		return nil, fmt.Errorf("failed creating type config: %w", err)
	}

	if tcInput.Relations == nil && tcInput.Permissions != nil {
		return nil, errors.New("failed creating type config: relations cannot be empty while permissions exist")
	}
	// if no relations, we created just the type with the name
	if tcInput.Relations == nil {
		return tc, nil
	}

	relSlice := make([]*ent.RelationCreate, len(tcInput.Relations))
	cnt := 0

	log.Printf("for type: %s\n", tcInput.Name)
	for relName, relValue := range tcInput.Relations {

		log.Printf("-> validating '%s: %s'\n", relName, relValue)

		if !regexPropertyName.MatchString(relName) {
			return nil, fmt.Errorf("malformed relation name input: '%s'", relName)
		}

		if !regexRelValue.MatchString(relValue) {
			return nil, fmt.Errorf("malformed relation reference input: '%s'", relValue)
		}

		referencedTypeIDsSlice := make([]int, 0)

		// type with relation, ex: user | group#member
		for _, referencedType := range strings.Split(relValue, refValueDelim) {

			if !strings.Contains(referencedType, refSubrelationDelim) {

				id, err := h.Client.TypeConfig.
					Query().
					Where(typeconfig.Name(referencedType)).
					OnlyID(ctx)

				if err != nil {
					return nil, fmt.Errorf("referenced type does not exist: '%s'", referencedType)
				}

				log.Printf("---> found id %d for %s\n", id, referencedType)

				referencedTypeIDsSlice = append(referencedTypeIDsSlice, id)

			} else {
				// checking composed relation
				s := strings.Split(referencedType, refSubrelationDelim)

				refTypeName := s[0]
				refTypeRelation := s[1]

				id, err := h.Client.TypeConfig.
					Query().
					Where(typeconfig.Name(refTypeName)).
					OnlyID(ctx)

				if err != nil {
					return nil, fmt.Errorf("referenced type does not exist: '%s'", refTypeName)
				}

				hasRelation, err := h.Client.TypeConfig.
					Query().
					Where(typeconfig.IDEQ(id)).
					QueryRelations().
					Where(relation.NameEQ(refTypeRelation)).
					Exist(ctx)

				if !hasRelation || err != nil {
					return nil, fmt.Errorf("referenced relation does not exist: '%s#%s'", refTypeName, refTypeRelation)
				}

				log.Printf("---> found id %d for %s\n", id, refTypeName)
				referencedTypeIDsSlice = append(referencedTypeIDsSlice, id)

			}
		}

		relSlice[cnt] = h.Client.Relation.
			Create().
			SetName(relName).
			SetValue(relValue).
			AddRelTypeconfigIDs(referencedTypeIDsSlice...)

		cnt++
	}

	// save relations and permissions in db
	relBulk, err := h.Client.Relation.CreateBulk(relSlice...).Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed when creating relations: %w", err)
	}

	tc, err = tc.Update().AddRelations(relBulk...).Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed when adding relations: %w", err)
	}

	// validate permissions

	permSlice := make([]*ent.PermissionCreate, len(tcInput.Permissions))
	cnt = 0

	inputRelations := maps.Keys(tcInput.Relations)
	fmt.Println("relations:", inputRelations)

	getIdForRelationName := func(name string) (int, error) {
		for _, rel := range relBulk {
			if rel.Name == name {
				return rel.ID, nil
			}
		}

		return -1, fmt.Errorf("adding relation: couldn't get id for relation '%s'", name)
	}

	for permName, permValue := range tcInput.Permissions {

		if !regexPropertyName.MatchString(permName) {
			return nil, fmt.Errorf("malformed permission name input: '%s'", permName)
		}

		if !regexPermValue.MatchString(permValue) {
			return nil, fmt.Errorf("malformed permission reference input: '%s'", permValue)
		}

		referencedTypeIDsSlice := make([]int, 0)

		fmt.Printf("-> validating '%s: %s'\n", permName, permValue)
		for _, referencedRelation := range permDelim.Split(permValue, -1) {
			if referencedRelation == "" {
				continue
			}
			// check direct relations
			if !strings.Contains(referencedRelation, parentRelDelim) {
				if !slices.Contains(inputRelations, referencedRelation) {
					return nil, fmt.Errorf("referenced relation '%s' in permission '%s' not found", referencedRelation, permName)
				}

				relID, err := getIdForRelationName(referencedRelation)
				if err != nil {
					return nil, err
				}

				referencedTypeIDsSlice = append(referencedTypeIDsSlice, relID)

				fmt.Printf("--> found relation '%s' with id %d \n", referencedRelation, relID)
			} else {
				// check parent relations
				s := strings.Split(referencedRelation, parentRelDelim)

				refParent := s[0]
				refParentRel := s[1]

				if !slices.Contains(inputRelations, refParent) {
					return nil, fmt.Errorf("referenced relation '%s' in permission '%s' not found", referencedRelation, permName)
				}
				// error if composed relation: user | group#member
				referencedType := tcInput.Relations[refParent]

				if strings.Contains(referencedType, refValueDelim) {
					return nil, fmt.Errorf("referenced relation '%s' can't contain a composed value: '%s'", refParent, referencedType)
				}

				hasRelation, err := h.Client.TypeConfig.
					Query().
					Where(typeconfig.NameEQ(referencedType)).
					QueryRelations().
					Where(relation.NameEQ(refParentRel)).
					Exist(ctx)

				if !hasRelation || err != nil {
					return nil, fmt.Errorf("referenced type '%s' in relation '%s' doesn't have a '%s' relation",
						referencedType,
						refParent,
						refParentRel)
				}

				relID, err := getIdForRelationName(refParent)
				if err != nil {
					return nil, err
				}

				referencedTypeIDsSlice = append(referencedTypeIDsSlice, relID)

				fmt.Printf("--> found parent relation '%s' with relation '%s' with id %d\n", refParent, refParentRel, relID)
			}

		}

		permSlice[cnt] = h.Client.Permission.
			Create().
			SetName(permName).
			SetValue(permValue).
			AddRelationIDs(referencedTypeIDsSlice...)

		cnt++

	}

	// save relations and permissions in db
	permBulk, err := h.Client.Permission.CreateBulk(permSlice...).Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed when creating permissions: %w", err)
	}

	tc, err = tc.Update().AddPermissions(permBulk...).Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed when adding permissions: %w", err)
	}

	return tc, nil
}

func (h *Handler) CreateSubject(ctx context.Context, s *models.SubjectReq) (*ent.Subject, error) {

	if !regexSubjName.MatchString(s.SubjectName) {
		return nil, fmt.Errorf("malformed subj name input: '%s'", s.SubjectName)
	}

	exists, err := h.Client.TypeConfig.
		Query().
		Where(typeconfig.NameEQ(s.TypeConfigName)).
		Exist(ctx)

	if err != nil {
		return nil, fmt.Errorf("failed querying type config: %w", err)
	}
	if !exists {
		return nil, fmt.Errorf("failed when creating subject, type '%s' does not exist", s.TypeConfigName)
	}

	subj, err := h.Client.Subject.Create().SetName(s.SubjectName).Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed creating subject: %w", err)
	}

	_, err = h.Client.TypeConfig.Update().
		Where(typeconfig.NameEQ(s.TypeConfigName)).
		AddSubjects(subj).
		Save(ctx)

	if err != nil {
		return nil, fmt.Errorf("error when updating typeconfig: %w", err)
	}

	return subj, nil
}

// TODO handle group:* viewer on document
// create a default subject called '*' when creating the type?
func (h *Handler) CreateTuple(ctx context.Context, tr *models.TupleReqRelation) (*ent.Tuple, error) {

	var subjectName string
	var subjectRelation string = ""

	if !strings.Contains(tr.Subject.SubjectName, "#") {
		subjectName = tr.Subject.SubjectName
	} else {
		subjectName = strings.Split(tr.Subject.SubjectName, "#")[0]
		subjectRelation = strings.Split(tr.Subject.SubjectName, "#")[1]
	}

	// handling cases like group:dev#member and group:dev#*
	if subjectRelation != "" {
		hasRelation, err := h.Client.TypeConfig.
			Query().
			Where(typeconfig.NameEQ(tr.Subject.TypeConfigName)).
			QueryRelations().
			Where(relation.NameEQ(subjectRelation)).
			Exist(ctx)
		if !hasRelation || err != nil {
			return nil, fmt.Errorf("error when checking subject relation: %w", err)
		}
	}

	subj, err := h.Client.Subject.
		Query().
		Where(
			subject.And(
				subject.NameEQ(subjectName),
				subject.HasTypeWith(
					typeconfig.NameEQ(tr.Subject.TypeConfigName)))).
		Only(ctx)

	if err != nil {
		return nil, fmt.Errorf("error when getting subject: %w", err)
	}

	res, err := h.Client.Subject.
		Query().
		Where(
			subject.And(
				subject.NameEQ(tr.Resource.SubjectName),
				subject.HasTypeWith(
					typeconfig.NameEQ(tr.Resource.TypeConfigName)))).
		Only(ctx)

	if err != nil {
		return nil, fmt.Errorf("error when getting resource: %w", err)
	}

	rel, err := h.Client.Relation.
		Query().
		Where(
			relation.And(
				relation.HasTypeconfigWith(
					typeconfig.NameEQ(tr.Resource.TypeConfigName)),
				relation.NameEQ(tr.Relation))).
		Only(ctx)

	if err != nil {
		return nil, fmt.Errorf("error when getting relation for resource: %w", err)
	}

	tupleCreate := h.Client.Tuple.
		Create().
		SetSubject(subj).
		SetRelation(rel).
		SetResource(res)

	if subjectRelation != "" {
		tupleCreate.SetSubjectRel(subjectRelation)
	}

	tuple, err := tupleCreate.Save(ctx)

	if err != nil {
		return nil, fmt.Errorf("error when creating tuple: %w", err)
	}

	return tuple, nil
}

// TODO: add some depth, maybe add path taken (ident in a json and print graph?), goroutines
func (h *Handler) CheckRelation(ctx context.Context, rc *models.RelationCheck, depth int) bool {

	fmt.Println(strings.Repeat("\t", depth+1)+"• Checking relation:", rc)

	subjTypeString := rc.Subj.QueryType().Select(typeconfig.FieldName).StringX(ctx)

	// check if direct
	if !strings.Contains(rc.Rel.Value, refValueDelim) {

		if rc.Rel.Value != subjTypeString {
			fmt.Println("\t\t Type doesnt match with subject")
			return false
		}

		tupleExists, err := h.Client.Tuple.
			Query().
			Where(
				tuple.And(
					tuple.SubjectID(rc.Subj.ID),
					tuple.RelationID(rc.Rel.ID),
					tuple.ResourceID(rc.Res.ID),
				),
			).
			Exist(ctx)

		fmt.Printf("\t\tDirect relationship? %t\n", tupleExists)

		if err != nil {
			fmt.Println(err)
		}

		return tupleExists

	} else {
		for _, referencedType := range strings.Split(rc.Rel.Value, refValueDelim) {

			// if direct type
			if referencedType == subjTypeString {
				tupleExists, err := h.Client.Tuple.
					Query().
					Where(
						tuple.And(
							tuple.SubjectID(rc.Subj.ID),
							tuple.RelationID(rc.Rel.ID),
							tuple.ResourceID(rc.Res.ID),
						),
					).
					Exist(ctx)

				fmt.Printf("\t\tDirect relationship? %t\n", tupleExists)

				if err != nil {
					fmt.Println(err)
				}

				if tupleExists {
					return true
				}

			} else {
				// checking composed relation
				// TODO: repeated, own function
				s := strings.Split(referencedType, refSubrelationDelim)

				refTypeName := s[0]
				refTypeRelation := s[1]

				// Get all refTypeName and get the relation refTypeName (e.g. group#member) with relation rc.Rel.Name
				fmt.Printf("%sGetting all `%s:<name>#%s` with relation `%s` on `%s`\n",
					strings.Repeat("\t", depth+2),
					refTypeName,
					refTypeRelation,
					rc.Rel.Name,
					rc.Res.Name)
				fmt.Printf("%sand checking if subject `%s` has relation `%s` on each group\n",
					strings.Repeat("\t", depth+2),
					rc.Subj.Name,
					refTypeRelation)

				refTypeRelationObject, err := h.Client.TypeConfig.
					Query().
					Where(typeconfig.NameEQ(refTypeName)).
					QueryRelations().
					Where(relation.NameEQ(refTypeRelation)).
					Only(ctx)

				if err != nil {
					fmt.Println(err)
					return false
				}

				subjects, err := h.Client.Tuple.
					Query().
					Where(tuple.And(
						tuple.SubjectRelEQ(refTypeRelation),
						tuple.RelationID(rc.Rel.ID),
						tuple.ResourceID(rc.Res.ID),
					)).
					QuerySubject().
					All(ctx)

				if err != nil {
					fmt.Println(err)
					return false
				}

				relExists := false
				for _, s := range subjects {

					relExists = relExists || h.CheckRelation(ctx, &models.RelationCheck{
						Subj: rc.Subj,
						Rel:  refTypeRelationObject,
						Res:  s,
					}, depth+1)

					if relExists {
						return relExists
					}
				}
			}
		}
	}

	return false
}

func (h *Handler) Subject(ctx context.Context, tfName string, name string) (*ent.Subject, error) {
	subj, err := h.Client.Subject.
		Query().
		Where(
			subject.And(
				subject.NameEQ(name),
				subject.HasTypeWith(typeconfig.NameEQ(tfName)))).
		Only(ctx)
	return subj, err
}

func (h *Handler) CheckPermission(ctx context.Context, tr *models.TupleReqPermission) (bool, error) {
	fmt.Printf("====> Does `%s:%s` have `%s` permission on `%s:%s`?\n",
		tr.Subject.TypeConfigName,
		tr.Subject.SubjectName,
		tr.Permission,
		tr.Resource.TypeConfigName,
		tr.Resource.SubjectName)

	subj, err := h.Subject(ctx, tr.Subject.TypeConfigName, tr.Subject.SubjectName)

	if err != nil {
		return false, fmt.Errorf("error when querying subject: %w", err)
	}

	fmt.Println("Checking if subject exists:", subj)

	res, err := h.Subject(ctx, tr.Resource.TypeConfigName, tr.Resource.SubjectName)

	if err != nil {
		return false, fmt.Errorf("error when querying resource: %w", err)
	}

	fmt.Println("Checking if resource exists:", res)

	permQuery := res.
		QueryType().
		QueryPermissions().
		Where(permission.NameEQ(tr.Permission))

	perm, err := permQuery.
		Only(ctx)

	if err != nil {
		return false, fmt.Errorf("error when querying permission: %w", err)
	}

	fmt.Println("Getting relations for perm:", perm)

	// Getting the result from the check relations
	hasPerm := false

	// TODO: refactor this in its own function, same one as in CreateTypeConfig
	// or you should probably parse it
	// probably shouldnt return on all errors
	for _, referencedRelation := range permDelim.Split(perm.Value, -1) {
		if referencedRelation == "" {
			continue
		}

		// check direct relations and indirect relations
		if !strings.Contains(referencedRelation, parentRelDelim) {
			rel, err := permQuery.
				QueryRelations().
				Where(relation.NameEQ(referencedRelation)).
				Only(ctx)

			if err != nil {
				return false, fmt.Errorf("error when querying relation: %w", err)
			}

			hasPerm = hasPerm || h.CheckRelation(ctx, &models.RelationCheck{
				Subj: subj,
				Rel:  rel,
				Res:  res,
			}, 0)

			if hasPerm {
				return hasPerm, nil
			}

		} else {
			s := strings.Split(referencedRelation, parentRelDelim)

			refParent := s[0]
			refParentRel := s[1]

			// Get the relation, for this type of query
			// it can only be one anyway
			r, err := perm.QueryRelations().
				Where(relation.NameEQ(refParent)).
				QueryRelTypeconfigs().
				QueryRelations().
				Where(relation.NameEQ(refParentRel)).
				Only(ctx)

			if err != nil {
				return false, fmt.Errorf("error when querying relation: %w", err)
			}

			// Get all referenced subjects
			subjects, err := perm.QueryRelations().
				Where(relation.NameEQ(refParent)).
				QueryTuples().
				QuerySubject().
				All(ctx)

			// if it doesn't have subjects, just move on
			if err != nil {
				fmt.Println(err)
				continue
			}

			for _, s := range subjects {
				hasPerm = hasPerm || h.CheckRelation(ctx, &models.RelationCheck{
					Subj: subj,
					Rel:  r,
					Res:  s,
				}, 0)

				if hasPerm {
					return hasPerm, nil
				}
			}
		}

	}

	return hasPerm, nil
}
