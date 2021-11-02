// Code generated by entc, DO NOT EDIT.

package vehicle

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/open-farms/inventory/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id uint64) predicate.Vehicle {
	return predicate.Vehicle(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uint64) predicate.Vehicle {
	return predicate.Vehicle(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uint64) predicate.Vehicle {
	return predicate.Vehicle(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uint64) predicate.Vehicle {
	return predicate.Vehicle(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uint64) predicate.Vehicle {
	return predicate.Vehicle(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uint64) predicate.Vehicle {
	return predicate.Vehicle(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uint64) predicate.Vehicle {
	return predicate.Vehicle(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uint64) predicate.Vehicle {
	return predicate.Vehicle(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uint64) predicate.Vehicle {
	return predicate.Vehicle(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Make applies equality check predicate on the "make" field. It's identical to MakeEQ.
func Make(v string) predicate.Vehicle {
	return predicate.Vehicle(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldMake), v))
	})
}

// Model applies equality check predicate on the "model" field. It's identical to ModelEQ.
func Model(v string) predicate.Vehicle {
	return predicate.Vehicle(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldModel), v))
	})
}

// Year applies equality check predicate on the "year" field. It's identical to YearEQ.
func Year(v string) predicate.Vehicle {
	return predicate.Vehicle(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldYear), v))
	})
}

// Active applies equality check predicate on the "active" field. It's identical to ActiveEQ.
func Active(v bool) predicate.Vehicle {
	return predicate.Vehicle(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldActive), v))
	})
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Vehicle {
	return predicate.Vehicle(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Vehicle {
	return predicate.Vehicle(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// MakeEQ applies the EQ predicate on the "make" field.
func MakeEQ(v string) predicate.Vehicle {
	return predicate.Vehicle(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldMake), v))
	})
}

// MakeNEQ applies the NEQ predicate on the "make" field.
func MakeNEQ(v string) predicate.Vehicle {
	return predicate.Vehicle(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldMake), v))
	})
}

// MakeIn applies the In predicate on the "make" field.
func MakeIn(vs ...string) predicate.Vehicle {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Vehicle(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldMake), v...))
	})
}

// MakeNotIn applies the NotIn predicate on the "make" field.
func MakeNotIn(vs ...string) predicate.Vehicle {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Vehicle(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldMake), v...))
	})
}

// MakeGT applies the GT predicate on the "make" field.
func MakeGT(v string) predicate.Vehicle {
	return predicate.Vehicle(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldMake), v))
	})
}

// MakeGTE applies the GTE predicate on the "make" field.
func MakeGTE(v string) predicate.Vehicle {
	return predicate.Vehicle(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldMake), v))
	})
}

// MakeLT applies the LT predicate on the "make" field.
func MakeLT(v string) predicate.Vehicle {
	return predicate.Vehicle(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldMake), v))
	})
}

// MakeLTE applies the LTE predicate on the "make" field.
func MakeLTE(v string) predicate.Vehicle {
	return predicate.Vehicle(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldMake), v))
	})
}

// MakeContains applies the Contains predicate on the "make" field.
func MakeContains(v string) predicate.Vehicle {
	return predicate.Vehicle(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldMake), v))
	})
}

// MakeHasPrefix applies the HasPrefix predicate on the "make" field.
func MakeHasPrefix(v string) predicate.Vehicle {
	return predicate.Vehicle(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldMake), v))
	})
}

// MakeHasSuffix applies the HasSuffix predicate on the "make" field.
func MakeHasSuffix(v string) predicate.Vehicle {
	return predicate.Vehicle(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldMake), v))
	})
}

// MakeEqualFold applies the EqualFold predicate on the "make" field.
func MakeEqualFold(v string) predicate.Vehicle {
	return predicate.Vehicle(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldMake), v))
	})
}

// MakeContainsFold applies the ContainsFold predicate on the "make" field.
func MakeContainsFold(v string) predicate.Vehicle {
	return predicate.Vehicle(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldMake), v))
	})
}

// ModelEQ applies the EQ predicate on the "model" field.
func ModelEQ(v string) predicate.Vehicle {
	return predicate.Vehicle(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldModel), v))
	})
}

// ModelNEQ applies the NEQ predicate on the "model" field.
func ModelNEQ(v string) predicate.Vehicle {
	return predicate.Vehicle(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldModel), v))
	})
}

// ModelIn applies the In predicate on the "model" field.
func ModelIn(vs ...string) predicate.Vehicle {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Vehicle(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldModel), v...))
	})
}

// ModelNotIn applies the NotIn predicate on the "model" field.
func ModelNotIn(vs ...string) predicate.Vehicle {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Vehicle(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldModel), v...))
	})
}

// ModelGT applies the GT predicate on the "model" field.
func ModelGT(v string) predicate.Vehicle {
	return predicate.Vehicle(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldModel), v))
	})
}

// ModelGTE applies the GTE predicate on the "model" field.
func ModelGTE(v string) predicate.Vehicle {
	return predicate.Vehicle(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldModel), v))
	})
}

// ModelLT applies the LT predicate on the "model" field.
func ModelLT(v string) predicate.Vehicle {
	return predicate.Vehicle(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldModel), v))
	})
}

// ModelLTE applies the LTE predicate on the "model" field.
func ModelLTE(v string) predicate.Vehicle {
	return predicate.Vehicle(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldModel), v))
	})
}

// ModelContains applies the Contains predicate on the "model" field.
func ModelContains(v string) predicate.Vehicle {
	return predicate.Vehicle(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldModel), v))
	})
}

// ModelHasPrefix applies the HasPrefix predicate on the "model" field.
func ModelHasPrefix(v string) predicate.Vehicle {
	return predicate.Vehicle(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldModel), v))
	})
}

// ModelHasSuffix applies the HasSuffix predicate on the "model" field.
func ModelHasSuffix(v string) predicate.Vehicle {
	return predicate.Vehicle(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldModel), v))
	})
}

// ModelEqualFold applies the EqualFold predicate on the "model" field.
func ModelEqualFold(v string) predicate.Vehicle {
	return predicate.Vehicle(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldModel), v))
	})
}

// ModelContainsFold applies the ContainsFold predicate on the "model" field.
func ModelContainsFold(v string) predicate.Vehicle {
	return predicate.Vehicle(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldModel), v))
	})
}

// YearEQ applies the EQ predicate on the "year" field.
func YearEQ(v string) predicate.Vehicle {
	return predicate.Vehicle(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldYear), v))
	})
}

// YearNEQ applies the NEQ predicate on the "year" field.
func YearNEQ(v string) predicate.Vehicle {
	return predicate.Vehicle(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldYear), v))
	})
}

// YearIn applies the In predicate on the "year" field.
func YearIn(vs ...string) predicate.Vehicle {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Vehicle(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldYear), v...))
	})
}

// YearNotIn applies the NotIn predicate on the "year" field.
func YearNotIn(vs ...string) predicate.Vehicle {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Vehicle(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldYear), v...))
	})
}

// YearGT applies the GT predicate on the "year" field.
func YearGT(v string) predicate.Vehicle {
	return predicate.Vehicle(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldYear), v))
	})
}

// YearGTE applies the GTE predicate on the "year" field.
func YearGTE(v string) predicate.Vehicle {
	return predicate.Vehicle(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldYear), v))
	})
}

// YearLT applies the LT predicate on the "year" field.
func YearLT(v string) predicate.Vehicle {
	return predicate.Vehicle(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldYear), v))
	})
}

// YearLTE applies the LTE predicate on the "year" field.
func YearLTE(v string) predicate.Vehicle {
	return predicate.Vehicle(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldYear), v))
	})
}

// YearContains applies the Contains predicate on the "year" field.
func YearContains(v string) predicate.Vehicle {
	return predicate.Vehicle(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldYear), v))
	})
}

// YearHasPrefix applies the HasPrefix predicate on the "year" field.
func YearHasPrefix(v string) predicate.Vehicle {
	return predicate.Vehicle(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldYear), v))
	})
}

// YearHasSuffix applies the HasSuffix predicate on the "year" field.
func YearHasSuffix(v string) predicate.Vehicle {
	return predicate.Vehicle(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldYear), v))
	})
}

// YearEqualFold applies the EqualFold predicate on the "year" field.
func YearEqualFold(v string) predicate.Vehicle {
	return predicate.Vehicle(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldYear), v))
	})
}

// YearContainsFold applies the ContainsFold predicate on the "year" field.
func YearContainsFold(v string) predicate.Vehicle {
	return predicate.Vehicle(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldYear), v))
	})
}

// ActiveEQ applies the EQ predicate on the "active" field.
func ActiveEQ(v bool) predicate.Vehicle {
	return predicate.Vehicle(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldActive), v))
	})
}

// ActiveNEQ applies the NEQ predicate on the "active" field.
func ActiveNEQ(v bool) predicate.Vehicle {
	return predicate.Vehicle(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldActive), v))
	})
}

// ConditionEQ applies the EQ predicate on the "condition" field.
func ConditionEQ(v Condition) predicate.Vehicle {
	return predicate.Vehicle(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCondition), v))
	})
}

// ConditionNEQ applies the NEQ predicate on the "condition" field.
func ConditionNEQ(v Condition) predicate.Vehicle {
	return predicate.Vehicle(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCondition), v))
	})
}

// ConditionIn applies the In predicate on the "condition" field.
func ConditionIn(vs ...Condition) predicate.Vehicle {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Vehicle(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldCondition), v...))
	})
}

// ConditionNotIn applies the NotIn predicate on the "condition" field.
func ConditionNotIn(vs ...Condition) predicate.Vehicle {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Vehicle(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldCondition), v...))
	})
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Vehicle {
	return predicate.Vehicle(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Vehicle {
	return predicate.Vehicle(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Vehicle {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Vehicle(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Vehicle {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Vehicle(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Vehicle {
	return predicate.Vehicle(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Vehicle {
	return predicate.Vehicle(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Vehicle {
	return predicate.Vehicle(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Vehicle {
	return predicate.Vehicle(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Vehicle {
	return predicate.Vehicle(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Vehicle {
	return predicate.Vehicle(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Vehicle {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Vehicle(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Vehicle {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Vehicle(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Vehicle {
	return predicate.Vehicle(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Vehicle {
	return predicate.Vehicle(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Vehicle {
	return predicate.Vehicle(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Vehicle {
	return predicate.Vehicle(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdatedAt), v))
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Vehicle) predicate.Vehicle {
	return predicate.Vehicle(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Vehicle) predicate.Vehicle {
	return predicate.Vehicle(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Vehicle) predicate.Vehicle {
	return predicate.Vehicle(func(s *sql.Selector) {
		p(s.Not())
	})
}
