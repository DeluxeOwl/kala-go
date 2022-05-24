package services

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/pkg/errors"

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

func (h *Handler) WithTx(ctx context.Context, fn func(tx *ent.Tx) error) error {
	tx, err := h.Db.Tx(ctx)
	if err != nil {
		return err
	}
	defer func() {
		if v := recover(); v != nil {
			tx.Rollback()
			panic(v)
		}
	}()
	if err := fn(tx); err != nil {
		if rerr := tx.Rollback(); rerr != nil {
			err = errors.Wrapf(err, "rolling back transaction: %v", rerr)
		}
		return err
	}
	if err := tx.Commit(); err != nil {

		return errors.Wrapf(err, "committing transaction: %v", err)
	}
	return nil
}

func (h *Handler) CreateTypeConfig(ctx context.Context, tcInput *models.TypeConfigReq) (*ent.TypeConfig, error) {
	var tc *ent.TypeConfig

	if err := h.WithTx(ctx, func(tx *ent.Tx) error {
		newH := Handler{
			Db: tx.Client(),
		}
		var err error
		tc, err = newH.DoCreateTypeConfig(ctx, tcInput)

		return err
	}); err != nil {
		return nil, err
	}

	return tc, nil
}
func (h *Handler) CreateSubject(ctx context.Context, s *models.SubjectReq) (*ent.Subject, error) {
	var subj *ent.Subject

	if err := h.WithTx(ctx, func(tx *ent.Tx) error {
		newH := Handler{
			Db: tx.Client(),
		}
		var err error
		subj, err = newH.DoCreateSubject(ctx, s)

		return err
	}); err != nil {
		return nil, err
	}

	return subj, nil
}
func (h *Handler) CreateTuple(ctx context.Context, tr *models.TupleReqRelation) (*ent.Tuple, error) {
	var tuple *ent.Tuple

	if err := h.WithTx(ctx, func(tx *ent.Tx) error {
		newH := Handler{
			Db: tx.Client(),
		}
		var err error
		tuple, err = newH.DoCreateTuple(ctx, tr)

		return err
	}); err != nil {
		return nil, err
	}

	return tuple, nil
}

// error: cannot start a transaction within a transaction
// func (h *Handler) CheckRelation(ctx context.Context, rc *models.RelationCheck, depth int) bool {
// 	hasRel := false

// 	if err := h.WithTx(ctx, func(tx *ent.Tx) error {
// 		newH := Handler{
// 			Client: tx.Client(),
// 		}

// 		var err error
// 		hasRel = newH.DoCheckRelation(ctx, rc, depth)

// 		return err
// 	}); err != nil {
// 		fmt.Println(err)
// 		return false
// 	}

// 	return hasRel
// }

func (h *Handler) CheckPermission(ctx context.Context, tr *models.TupleReqPermission) (bool, error) {
	hasRel := false

	if err := h.WithTx(ctx, func(tx *ent.Tx) error {
		newH := Handler{
			Db: tx.Client(),
		}
		var err error
		hasRel, err = newH.DoCheckPermission(ctx, tr)

		return err
	}); err != nil {
		return false, err
	}

	return hasRel, nil
}

// TODO: make the errors as variables?
func (h *Handler) DoCreateTypeConfig(ctx context.Context, tcInput *models.TypeConfigReq) (*ent.TypeConfig, error) {

	if !regexTypeName.MatchString(tcInput.Name) {
		return nil, fmt.Errorf("malformed type name input: `%s`", tcInput.Name)
	}

	tc, err := h.Db.TypeConfig.Create().
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
			return nil, fmt.Errorf("malformed relation name input: `%s`", relName)
		}

		if !regexRelValue.MatchString(relValue) {
			return nil, fmt.Errorf("malformed relation reference input: `%s`", relValue)
		}

		referencedTypeIDsSlice := make([]int, 0)

		// type with relation, ex: user | group#member
		for _, referencedType := range strings.Split(relValue, refValueDelim) {

			if !strings.Contains(referencedType, refSubrelationDelim) {

				id, err := h.Db.TypeConfig.
					Query().
					Where(typeconfig.Name(referencedType)).
					OnlyID(ctx)

				if err != nil {
					return nil, fmt.Errorf("in type `%s`,\n referenced type does not exist: `%s`", tcInput.Name, referencedType)
				}

				log.Printf("---> found id %d for %s\n", id, referencedType)

				referencedTypeIDsSlice = append(referencedTypeIDsSlice, id)

			} else {
				// checking composed relation
				s := strings.Split(referencedType, refSubrelationDelim)

				refTypeName := s[0]
				refTypeRelation := s[1]

				id, err := h.Db.TypeConfig.
					Query().
					Where(typeconfig.Name(refTypeName)).
					OnlyID(ctx)

				if err != nil {
					return nil, fmt.Errorf("in type `%s`,\n referenced type does not exist: `%s`", tcInput.Name, refTypeName)
				}

				hasRelation, err := h.Db.TypeConfig.
					Query().
					Where(typeconfig.IDEQ(id)).
					QueryRelations().
					Where(relation.NameEQ(refTypeRelation)).
					Exist(ctx)

				if !hasRelation || err != nil {
					return nil, fmt.Errorf("in type `%s`,\n referenced relation does not exist: `%s#%s`", tcInput.Name, refTypeName, refTypeRelation)
				}

				log.Printf("---> found id %d for %s\n", id, refTypeName)
				referencedTypeIDsSlice = append(referencedTypeIDsSlice, id)

			}
		}

		relSlice[cnt] = h.Db.Relation.
			Create().
			SetName(relName).
			SetValue(relValue).
			AddRelTypeconfigIDs(referencedTypeIDsSlice...)

		cnt++
	}

	// save relations and permissions in db
	relBulk, err := h.Db.Relation.CreateBulk(relSlice...).Save(ctx)
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

		return -1, fmt.Errorf("adding relation: couldn't get id for relation `%s`", name)
	}

	for permName, permValue := range tcInput.Permissions {

		if !regexPropertyName.MatchString(permName) {
			return nil, fmt.Errorf("in type `%s`,\n malformed permission name input: `%s`", tcInput.Name, permName)
		}

		if !regexPermValue.MatchString(permValue) {
			return nil, fmt.Errorf("in type `%s`,\n malformed permission reference input: `%s`", tcInput.Name, permValue)
		}

		referencedTypeIDsSlice := make([]int, 0)

		fmt.Printf("-> validating '%s: %s'\n", permName, permValue)
		for _, referencedRelation := range permDelim.FindAllString(permValue, -1) {
			// check direct relations
			if !strings.Contains(referencedRelation, parentRelDelim) {
				if !slices.Contains(inputRelations, referencedRelation) {
					return nil, fmt.Errorf("in type `%s`,\n referenced relation `%s` in permission `%s` not found",
						tcInput.Name,
						referencedRelation,
						permName)
				}

				relID, err := getIdForRelationName(referencedRelation)
				if err != nil {
					return nil, err
				}

				referencedTypeIDsSlice = append(referencedTypeIDsSlice, relID)

				fmt.Printf("--> found relation `%s` with id %d \n", referencedRelation, relID)
			} else {
				// check parent relations
				s := strings.Split(referencedRelation, parentRelDelim)

				refParent := s[0]
				refParentRel := s[1]

				if !slices.Contains(inputRelations, refParent) {
					return nil, fmt.Errorf("in type `%s`,\n referenced relation `%s` in permission `%s` not found", tcInput.Name, referencedRelation, permName)
				}
				// error if composed relation: user | group#member
				referencedType := tcInput.Relations[refParent]

				if strings.Contains(referencedType, refValueDelim) {
					return nil, fmt.Errorf("in type `%s`,\n referenced relation `%s` can't contain a composed value: `%s`", tcInput.Name, refParent, referencedType)
				}

				hasRelation, err := h.Db.TypeConfig.
					Query().
					Where(typeconfig.NameEQ(referencedType)).
					QueryRelations().
					Where(relation.NameEQ(refParentRel)).
					Exist(ctx)

				if !hasRelation || err != nil {
					return nil, fmt.Errorf("in type `%s`,\n referenced type `%s` in relation `%s` doesn't have a `%s` relation",
						tcInput.Name,
						referencedType,
						refParent,
						refParentRel)
				}

				relID, err := getIdForRelationName(refParent)
				if err != nil {
					return nil, err
				}

				referencedTypeIDsSlice = append(referencedTypeIDsSlice, relID)

				fmt.Printf("--> found parent relation `%s` with relation `%s` with id %d\n", refParent, refParentRel, relID)
			}

		}

		permSlice[cnt] = h.Db.Permission.
			Create().
			SetName(permName).
			SetValue(permValue).
			AddRelationIDs(referencedTypeIDsSlice...)

		cnt++

	}

	// save relations and permissions in db
	permBulk, err := h.Db.Permission.CreateBulk(permSlice...).Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed when creating permissions: %w", err)
	}

	tc, err = tc.Update().AddPermissions(permBulk...).Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed when adding permissions: %w", err)
	}

	return tc, nil
}

func (h *Handler) DoCreateSubject(ctx context.Context, s *models.SubjectReq) (*ent.Subject, error) {

	if !regexSubjName.MatchString(s.SubjectName) {
		return nil, fmt.Errorf("malformed subj name input: `%s`", s.SubjectName)
	}

	exists, err := h.Db.TypeConfig.
		Query().
		Where(typeconfig.NameEQ(s.TypeConfigName)).
		Exist(ctx)

	if err != nil {
		return nil, fmt.Errorf("failed querying type config: %w", err)
	}
	if !exists {
		return nil, fmt.Errorf("failed when creating subject, type `%s` does not exist", s.TypeConfigName)
	}

	subj, err := h.Db.Subject.Create().SetName(s.SubjectName).Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed creating subject: %w", err)
	}

	_, err = h.Db.TypeConfig.Update().
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
func (h *Handler) DoCreateTuple(ctx context.Context, tr *models.TupleReqRelation) (*ent.Tuple, error) {

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
		hasRelation, err := h.Db.TypeConfig.
			Query().
			Where(typeconfig.NameEQ(tr.Subject.TypeConfigName)).
			QueryRelations().
			Where(relation.NameEQ(subjectRelation)).
			Exist(ctx)
		if !hasRelation || err != nil {
			return nil, fmt.Errorf("error when checking subject relation: %w", err)
		}
	}

	subj, err := h.Db.Subject.
		Query().
		Where(
			subject.And(
				subject.NameEQ(subjectName),
				subject.HasTypeWith(
					typeconfig.NameEQ(tr.Subject.TypeConfigName)))).
		Only(ctx)

	if err != nil {
		return nil, fmt.Errorf("error when getting subject `%s:%s`: %w", tr.Subject.TypeConfigName, subjectName, err)
	}

	res, err := h.Db.Subject.
		Query().
		Where(
			subject.And(
				subject.NameEQ(tr.Resource.SubjectName),
				subject.HasTypeWith(
					typeconfig.NameEQ(tr.Resource.TypeConfigName)))).
		Only(ctx)

	if err != nil {
		return nil, fmt.Errorf("error when getting resource `%s:%s`: %w",
			tr.Resource.TypeConfigName,
			tr.Resource.SubjectName,
			err)
	}

	rel, err := h.Db.Relation.
		Query().
		Where(
			relation.And(
				relation.HasTypeconfigWith(
					typeconfig.NameEQ(tr.Resource.TypeConfigName)),
				relation.NameEQ(tr.Relation))).
		Only(ctx)

	if err != nil {
		return nil, fmt.Errorf("error when getting relation `%s` for resource `%s:%s: %w",
			tr.Relation,
			tr.Resource.TypeConfigName,
			tr.Resource.SubjectName,
			err)
	}

	tupleCreate := h.Db.Tuple.
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

	if ctx.Err() != nil {
		return false
	}

	subjTypeString, err := rc.Subj.QueryType().Select(typeconfig.FieldName).String(ctx)
	if err != nil {
		fmt.Printf("getting type string in check relation: %s\n", err)
		return false
	}

	// check if direct
	if !strings.Contains(rc.Rel.Value, refValueDelim) {

		if rc.Rel.Value != subjTypeString {
			fmt.Println("\t\t Type doesnt match with subject")
			return false
		}

		if ctx.Err() != nil {
			return false
		}

		tupleExists, err := h.Db.Tuple.
			Query().
			Where(
				tuple.And(
					tuple.SubjectID(rc.Subj.ID),
					tuple.RelationID(rc.Rel.ID),
					tuple.ResourceID(rc.Res.ID),
				),
			).
			Exist(ctx)

		fmt.Printf("%s• Checking relation: %v\n", strings.Repeat("\t", depth+1), rc)
		fmt.Printf("\t\tDirect relationship? %t\n", tupleExists)

		if err != nil {
			fmt.Printf("check if tuple exists in check relation: %s\n", err)
		}

		return tupleExists

	} else {
		for _, referencedType := range strings.Split(rc.Rel.Value, refValueDelim) {

			if ctx.Err() != nil {
				return false
			}
			// if direct type
			if referencedType == subjTypeString {
				tupleExists, err := h.Db.Tuple.
					Query().
					Where(
						tuple.And(
							tuple.SubjectID(rc.Subj.ID),
							tuple.RelationID(rc.Rel.ID),
							tuple.ResourceID(rc.Res.ID),
						),
					).
					Exist(ctx)

				fmt.Printf("%s• Checking relation: %v\n", strings.Repeat("\t", depth+1), rc)
				fmt.Printf("\t\tDirect relationship? %t\n", tupleExists)

				if err != nil {
					fmt.Printf("check if tuple exists in check relation: %s\n", err)
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

				if ctx.Err() != nil {
					return false
				}

				refTypeRelationObject, err := h.Db.TypeConfig.
					Query().
					Where(typeconfig.NameEQ(refTypeName)).
					QueryRelations().
					Where(relation.NameEQ(refTypeRelation)).
					Only(ctx)

				if err != nil {
					fmt.Printf("getting parent relation in check relation: %s\n", err)
					return false
				}

				if ctx.Err() != nil {
					return false
				}

				subjects, err := h.Db.Tuple.
					Query().
					Where(tuple.And(
						tuple.SubjectRelEQ(refTypeRelation),
						tuple.RelationID(rc.Rel.ID),
						tuple.ResourceID(rc.Res.ID),
					)).
					QuerySubject().
					All(ctx)

				if err != nil {
					fmt.Printf("getting subjects in check relation: %s\n", err)
					return false
				}

				if ctx.Err() != nil {
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

func (h *Handler) DoCheckPermission(ctx context.Context, tr *models.TupleReqPermission) (bool, error) {
	fmt.Printf("====> Does `%s:%s` have `%s` permission on `%s:%s`?\n",
		tr.Subject.TypeConfigName,
		tr.Subject.SubjectName,
		tr.Permission,
		tr.Resource.TypeConfigName,
		tr.Resource.SubjectName)

	subj, err := h.subject(ctx, tr.Subject.TypeConfigName, tr.Subject.SubjectName)

	if err != nil {
		return false, fmt.Errorf("error when querying subject: %w", err)
	}

	fmt.Println("Checking if subject exists:", subj)

	res, err := h.subject(ctx, tr.Resource.TypeConfigName, tr.Resource.SubjectName)

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

	hasPerm, err := h.ParsePermissionAndEvaluate(perm.Value, &models.PermissionCheck{
		Subj: subj,
		Perm: perm,
		Res:  res,
	})
	if err != nil {
		return false, fmt.Errorf("error when parsing and evaluation permissions: %w", err)
	}
	return hasPerm, nil
}

func (h *Handler) subject(ctx context.Context, tfName string, name string) (*ent.Subject, error) {
	subj, err := h.Db.Subject.
		Query().
		Where(
			subject.And(
				subject.NameEQ(name),
				subject.HasTypeWith(typeconfig.NameEQ(tfName)))).
		Only(ctx)
	return subj, err
}

func (h *Handler) DeleteEverything(ctx context.Context) {
	_, err := h.Db.Tuple.Delete().Exec(ctx)
	if err != nil {
		fmt.Printf("error when deleting tuples: %s\n", err)
	}
	_, err = h.Db.Subject.Delete().Exec(ctx)
	if err != nil {
		fmt.Printf("error when deleting subjects: %s\n", err)
	}
	_, err = h.Db.Permission.Delete().Exec(ctx)
	if err != nil {
		fmt.Printf("error when deleting permissions: %s\n", err)
	}
	_, err = h.Db.Relation.Delete().Exec(ctx)
	if err != nil {
		fmt.Printf("error when deleting relations: %s\n", err)
	}
	_, err = h.Db.TypeConfig.Delete().Exec(ctx)
	if err != nil {
		fmt.Printf("error when deleting typeconfigs: %s\n", err)
	}
}
func (h *Handler) DeleteSubjects(ctx context.Context) {
	_, err := h.Db.Subject.Delete().Exec(ctx)
	if err != nil {
		fmt.Printf("error when deleting subjects: %s\n", err)
	}
}
func (h *Handler) DeleteTuples(ctx context.Context) {
	_, err := h.Db.Tuple.Delete().Exec(ctx)
	if err != nil {
		fmt.Printf("error when deleting subjects: %s\n", err)
	}
}
