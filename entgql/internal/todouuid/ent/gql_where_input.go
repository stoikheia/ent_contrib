// Copyright 2019-present Facebook
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by entc, DO NOT EDIT.

package ent

import (
	"encoding/json"
	"fmt"
	"io"
	"time"

	"entgo.io/contrib/entgql/internal/todo/ent/schema/schematype"
	"entgo.io/contrib/entgql/internal/todouuid/ent/category"
	"entgo.io/contrib/entgql/internal/todouuid/ent/group"
	"entgo.io/contrib/entgql/internal/todouuid/ent/predicate"
	"entgo.io/contrib/entgql/internal/todouuid/ent/todo"
	"entgo.io/contrib/entgql/internal/todouuid/ent/user"
	"github.com/google/uuid"
)

// CategoryWhereInput represents a where input for filtering Category queries.
type CategoryWhereInput struct {
	Not *CategoryWhereInput   `json:"not,omitempty"`
	Or  []*CategoryWhereInput `json:"or,omitempty"`
	And []*CategoryWhereInput `json:"and,omitempty"`

	// "id" field predicates.
	ID      *uuid.UUID  `json:"id,omitempty"`
	IDNEQ   *uuid.UUID  `json:"idNEQ,omitempty"`
	IDIn    []uuid.UUID `json:"idIn,omitempty"`
	IDNotIn []uuid.UUID `json:"idNotIn,omitempty"`
	IDGT    *uuid.UUID  `json:"idGT,omitempty"`
	IDGTE   *uuid.UUID  `json:"idGTE,omitempty"`
	IDLT    *uuid.UUID  `json:"idLT,omitempty"`
	IDLTE   *uuid.UUID  `json:"idLTE,omitempty"`

	// "text" field predicates.
	Text             *string  `json:"text,omitempty"`
	TextNEQ          *string  `json:"textNEQ,omitempty"`
	TextIn           []string `json:"textIn,omitempty"`
	TextNotIn        []string `json:"textNotIn,omitempty"`
	TextGT           *string  `json:"textGT,omitempty"`
	TextGTE          *string  `json:"textGTE,omitempty"`
	TextLT           *string  `json:"textLT,omitempty"`
	TextLTE          *string  `json:"textLTE,omitempty"`
	TextContains     *string  `json:"textContains,omitempty"`
	TextHasPrefix    *string  `json:"textHasPrefix,omitempty"`
	TextHasSuffix    *string  `json:"textHasSuffix,omitempty"`
	TextEqualFold    *string  `json:"textEqualFold,omitempty"`
	TextContainsFold *string  `json:"textContainsFold,omitempty"`

	// "status" field predicates.
	Status      *category.Status  `json:"status,omitempty"`
	StatusNEQ   *category.Status  `json:"statusNEQ,omitempty"`
	StatusIn    []category.Status `json:"statusIn,omitempty"`
	StatusNotIn []category.Status `json:"statusNotIn,omitempty"`

	// "config" field predicates.
	Config       *schematype.CategoryConfig   `json:"config,omitempty"`
	ConfigNEQ    *schematype.CategoryConfig   `json:"configNEQ,omitempty"`
	ConfigIn     []*schematype.CategoryConfig `json:"configIn,omitempty"`
	ConfigNotIn  []*schematype.CategoryConfig `json:"configNotIn,omitempty"`
	ConfigGT     *schematype.CategoryConfig   `json:"configGT,omitempty"`
	ConfigGTE    *schematype.CategoryConfig   `json:"configGTE,omitempty"`
	ConfigLT     *schematype.CategoryConfig   `json:"configLT,omitempty"`
	ConfigLTE    *schematype.CategoryConfig   `json:"configLTE,omitempty"`
	ConfigIsNil  bool                         `json:"configIsNil,omitempty"`
	ConfigNotNil bool                         `json:"configNotNil,omitempty"`

	// "duration" field predicates.
	Duration       *time.Duration  `json:"duration,omitempty"`
	DurationNEQ    *time.Duration  `json:"durationNEQ,omitempty"`
	DurationIn     []time.Duration `json:"durationIn,omitempty"`
	DurationNotIn  []time.Duration `json:"durationNotIn,omitempty"`
	DurationGT     *time.Duration  `json:"durationGT,omitempty"`
	DurationGTE    *time.Duration  `json:"durationGTE,omitempty"`
	DurationLT     *time.Duration  `json:"durationLT,omitempty"`
	DurationLTE    *time.Duration  `json:"durationLTE,omitempty"`
	DurationIsNil  bool            `json:"durationIsNil,omitempty"`
	DurationNotNil bool            `json:"durationNotNil,omitempty"`

	// "count" field predicates.
	Count       *uint64  `json:"count,omitempty"`
	CountNEQ    *uint64  `json:"countNEQ,omitempty"`
	CountIn     []uint64 `json:"countIn,omitempty"`
	CountNotIn  []uint64 `json:"countNotIn,omitempty"`
	CountGT     *uint64  `json:"countGT,omitempty"`
	CountGTE    *uint64  `json:"countGTE,omitempty"`
	CountLT     *uint64  `json:"countLT,omitempty"`
	CountLTE    *uint64  `json:"countLTE,omitempty"`
	CountIsNil  bool     `json:"countIsNil,omitempty"`
	CountNotNil bool     `json:"countNotNil,omitempty"`

	// "todos" edge predicates.
	HasTodos     *bool             `json:"hasTodos,omitempty"`
	HasTodosWith []*TodoWhereInput `json:"hasTodosWith,omitempty"`
}

// Filter applies the CategoryWhereInput filter on the CategoryQuery builder.
func (i *CategoryWhereInput) Filter(q *CategoryQuery) (*CategoryQuery, error) {
	if i == nil {
		return q, nil
	}
	p, err := i.P()
	if err != nil {
		return nil, err
	}
	return q.Where(p), nil
}

// P returns a predicate for filtering categories.
// An error is returned if the input is empty or invalid.
func (i *CategoryWhereInput) P() (predicate.Category, error) {
	var predicates []predicate.Category
	if i.Not != nil {
		p, err := i.Not.P()
		if err != nil {
			return nil, err
		}
		predicates = append(predicates, category.Not(p))
	}
	switch n := len(i.Or); {
	case n == 1:
		p, err := i.Or[0].P()
		if err != nil {
			return nil, err
		}
		predicates = append(predicates, p)
	case n > 1:
		or := make([]predicate.Category, 0, n)
		for _, w := range i.Or {
			p, err := w.P()
			if err != nil {
				return nil, err
			}
			or = append(or, p)
		}
		predicates = append(predicates, category.Or(or...))
	}
	switch n := len(i.And); {
	case n == 1:
		p, err := i.And[0].P()
		if err != nil {
			return nil, err
		}
		predicates = append(predicates, p)
	case n > 1:
		and := make([]predicate.Category, 0, n)
		for _, w := range i.And {
			p, err := w.P()
			if err != nil {
				return nil, err
			}
			and = append(and, p)
		}
		predicates = append(predicates, category.And(and...))
	}
	if i.ID != nil {
		predicates = append(predicates, category.IDEQ(*i.ID))
	}
	if i.IDNEQ != nil {
		predicates = append(predicates, category.IDNEQ(*i.IDNEQ))
	}
	if len(i.IDIn) > 0 {
		predicates = append(predicates, category.IDIn(i.IDIn...))
	}
	if len(i.IDNotIn) > 0 {
		predicates = append(predicates, category.IDNotIn(i.IDNotIn...))
	}
	if i.IDGT != nil {
		predicates = append(predicates, category.IDGT(*i.IDGT))
	}
	if i.IDGTE != nil {
		predicates = append(predicates, category.IDGTE(*i.IDGTE))
	}
	if i.IDLT != nil {
		predicates = append(predicates, category.IDLT(*i.IDLT))
	}
	if i.IDLTE != nil {
		predicates = append(predicates, category.IDLTE(*i.IDLTE))
	}
	if i.Text != nil {
		predicates = append(predicates, category.TextEQ(*i.Text))
	}
	if i.TextNEQ != nil {
		predicates = append(predicates, category.TextNEQ(*i.TextNEQ))
	}
	if len(i.TextIn) > 0 {
		predicates = append(predicates, category.TextIn(i.TextIn...))
	}
	if len(i.TextNotIn) > 0 {
		predicates = append(predicates, category.TextNotIn(i.TextNotIn...))
	}
	if i.TextGT != nil {
		predicates = append(predicates, category.TextGT(*i.TextGT))
	}
	if i.TextGTE != nil {
		predicates = append(predicates, category.TextGTE(*i.TextGTE))
	}
	if i.TextLT != nil {
		predicates = append(predicates, category.TextLT(*i.TextLT))
	}
	if i.TextLTE != nil {
		predicates = append(predicates, category.TextLTE(*i.TextLTE))
	}
	if i.TextContains != nil {
		predicates = append(predicates, category.TextContains(*i.TextContains))
	}
	if i.TextHasPrefix != nil {
		predicates = append(predicates, category.TextHasPrefix(*i.TextHasPrefix))
	}
	if i.TextHasSuffix != nil {
		predicates = append(predicates, category.TextHasSuffix(*i.TextHasSuffix))
	}
	if i.TextEqualFold != nil {
		predicates = append(predicates, category.TextEqualFold(*i.TextEqualFold))
	}
	if i.TextContainsFold != nil {
		predicates = append(predicates, category.TextContainsFold(*i.TextContainsFold))
	}
	if i.Status != nil {
		predicates = append(predicates, category.StatusEQ(*i.Status))
	}
	if i.StatusNEQ != nil {
		predicates = append(predicates, category.StatusNEQ(*i.StatusNEQ))
	}
	if len(i.StatusIn) > 0 {
		predicates = append(predicates, category.StatusIn(i.StatusIn...))
	}
	if len(i.StatusNotIn) > 0 {
		predicates = append(predicates, category.StatusNotIn(i.StatusNotIn...))
	}
	if i.Config != nil {
		predicates = append(predicates, category.ConfigEQ(i.Config))
	}
	if i.ConfigNEQ != nil {
		predicates = append(predicates, category.ConfigNEQ(i.ConfigNEQ))
	}
	if len(i.ConfigIn) > 0 {
		predicates = append(predicates, category.ConfigIn(i.ConfigIn...))
	}
	if len(i.ConfigNotIn) > 0 {
		predicates = append(predicates, category.ConfigNotIn(i.ConfigNotIn...))
	}
	if i.ConfigGT != nil {
		predicates = append(predicates, category.ConfigGT(i.ConfigGT))
	}
	if i.ConfigGTE != nil {
		predicates = append(predicates, category.ConfigGTE(i.ConfigGTE))
	}
	if i.ConfigLT != nil {
		predicates = append(predicates, category.ConfigLT(i.ConfigLT))
	}
	if i.ConfigLTE != nil {
		predicates = append(predicates, category.ConfigLTE(i.ConfigLTE))
	}
	if i.ConfigIsNil {
		predicates = append(predicates, category.ConfigIsNil())
	}
	if i.ConfigNotNil {
		predicates = append(predicates, category.ConfigNotNil())
	}
	if i.Duration != nil {
		predicates = append(predicates, category.DurationEQ(*i.Duration))
	}
	if i.DurationNEQ != nil {
		predicates = append(predicates, category.DurationNEQ(*i.DurationNEQ))
	}
	if len(i.DurationIn) > 0 {
		predicates = append(predicates, category.DurationIn(i.DurationIn...))
	}
	if len(i.DurationNotIn) > 0 {
		predicates = append(predicates, category.DurationNotIn(i.DurationNotIn...))
	}
	if i.DurationGT != nil {
		predicates = append(predicates, category.DurationGT(*i.DurationGT))
	}
	if i.DurationGTE != nil {
		predicates = append(predicates, category.DurationGTE(*i.DurationGTE))
	}
	if i.DurationLT != nil {
		predicates = append(predicates, category.DurationLT(*i.DurationLT))
	}
	if i.DurationLTE != nil {
		predicates = append(predicates, category.DurationLTE(*i.DurationLTE))
	}
	if i.DurationIsNil {
		predicates = append(predicates, category.DurationIsNil())
	}
	if i.DurationNotNil {
		predicates = append(predicates, category.DurationNotNil())
	}
	if i.Count != nil {
		predicates = append(predicates, category.CountEQ(*i.Count))
	}
	if i.CountNEQ != nil {
		predicates = append(predicates, category.CountNEQ(*i.CountNEQ))
	}
	if len(i.CountIn) > 0 {
		predicates = append(predicates, category.CountIn(i.CountIn...))
	}
	if len(i.CountNotIn) > 0 {
		predicates = append(predicates, category.CountNotIn(i.CountNotIn...))
	}
	if i.CountGT != nil {
		predicates = append(predicates, category.CountGT(*i.CountGT))
	}
	if i.CountGTE != nil {
		predicates = append(predicates, category.CountGTE(*i.CountGTE))
	}
	if i.CountLT != nil {
		predicates = append(predicates, category.CountLT(*i.CountLT))
	}
	if i.CountLTE != nil {
		predicates = append(predicates, category.CountLTE(*i.CountLTE))
	}
	if i.CountIsNil {
		predicates = append(predicates, category.CountIsNil())
	}
	if i.CountNotNil {
		predicates = append(predicates, category.CountNotNil())
	}

	if i.HasTodos != nil {
		p := category.HasTodos()
		if !*i.HasTodos {
			p = category.Not(p)
		}
		predicates = append(predicates, p)
	}
	if len(i.HasTodosWith) > 0 {
		with := make([]predicate.Todo, 0, len(i.HasTodosWith))
		for _, w := range i.HasTodosWith {
			p, err := w.P()
			if err != nil {
				return nil, err
			}
			with = append(with, p)
		}
		predicates = append(predicates, category.HasTodosWith(with...))
	}
	switch len(predicates) {
	case 0:
		return nil, fmt.Errorf("empty predicate CategoryWhereInput")
	case 1:
		return predicates[0], nil
	default:
		return category.And(predicates...), nil
	}
}

// UnmarshalGQL implements the graphql.Unmarshaler interface
func (i *CategoryWhereInput) UnmarshalGQL(obj interface{}) error {
	raw, err := json.Marshal(obj)
	if err != nil {
		return err
	}
	return json.Unmarshal(raw, i)
}

// MarshalGQL implements the graphql.Marshaler interface
//
// This function should never be called because our object is Input.
// It exists to make GQLGen aware of objects has custom unmarshal function
func (i CategoryWhereInput) MarshalGQL(w io.Writer) {
	panic("CategoryWhereInput.MarshalGQL not implemented")
}

// GroupWhereInput represents a where input for filtering Group queries.
type GroupWhereInput struct {
	Not *GroupWhereInput   `json:"not,omitempty"`
	Or  []*GroupWhereInput `json:"or,omitempty"`
	And []*GroupWhereInput `json:"and,omitempty"`

	// "id" field predicates.
	ID      *uuid.UUID  `json:"id,omitempty"`
	IDNEQ   *uuid.UUID  `json:"idNEQ,omitempty"`
	IDIn    []uuid.UUID `json:"idIn,omitempty"`
	IDNotIn []uuid.UUID `json:"idNotIn,omitempty"`
	IDGT    *uuid.UUID  `json:"idGT,omitempty"`
	IDGTE   *uuid.UUID  `json:"idGTE,omitempty"`
	IDLT    *uuid.UUID  `json:"idLT,omitempty"`
	IDLTE   *uuid.UUID  `json:"idLTE,omitempty"`

	// "name" field predicates.
	Name             *string  `json:"name,omitempty"`
	NameNEQ          *string  `json:"nameNEQ,omitempty"`
	NameIn           []string `json:"nameIn,omitempty"`
	NameNotIn        []string `json:"nameNotIn,omitempty"`
	NameGT           *string  `json:"nameGT,omitempty"`
	NameGTE          *string  `json:"nameGTE,omitempty"`
	NameLT           *string  `json:"nameLT,omitempty"`
	NameLTE          *string  `json:"nameLTE,omitempty"`
	NameContains     *string  `json:"nameContains,omitempty"`
	NameHasPrefix    *string  `json:"nameHasPrefix,omitempty"`
	NameHasSuffix    *string  `json:"nameHasSuffix,omitempty"`
	NameEqualFold    *string  `json:"nameEqualFold,omitempty"`
	NameContainsFold *string  `json:"nameContainsFold,omitempty"`

	// "users" edge predicates.
	HasUsers     *bool             `json:"hasUsers,omitempty"`
	HasUsersWith []*UserWhereInput `json:"hasUsersWith,omitempty"`
}

// Filter applies the GroupWhereInput filter on the GroupQuery builder.
func (i *GroupWhereInput) Filter(q *GroupQuery) (*GroupQuery, error) {
	if i == nil {
		return q, nil
	}
	p, err := i.P()
	if err != nil {
		return nil, err
	}
	return q.Where(p), nil
}

// P returns a predicate for filtering groups.
// An error is returned if the input is empty or invalid.
func (i *GroupWhereInput) P() (predicate.Group, error) {
	var predicates []predicate.Group
	if i.Not != nil {
		p, err := i.Not.P()
		if err != nil {
			return nil, err
		}
		predicates = append(predicates, group.Not(p))
	}
	switch n := len(i.Or); {
	case n == 1:
		p, err := i.Or[0].P()
		if err != nil {
			return nil, err
		}
		predicates = append(predicates, p)
	case n > 1:
		or := make([]predicate.Group, 0, n)
		for _, w := range i.Or {
			p, err := w.P()
			if err != nil {
				return nil, err
			}
			or = append(or, p)
		}
		predicates = append(predicates, group.Or(or...))
	}
	switch n := len(i.And); {
	case n == 1:
		p, err := i.And[0].P()
		if err != nil {
			return nil, err
		}
		predicates = append(predicates, p)
	case n > 1:
		and := make([]predicate.Group, 0, n)
		for _, w := range i.And {
			p, err := w.P()
			if err != nil {
				return nil, err
			}
			and = append(and, p)
		}
		predicates = append(predicates, group.And(and...))
	}
	if i.ID != nil {
		predicates = append(predicates, group.IDEQ(*i.ID))
	}
	if i.IDNEQ != nil {
		predicates = append(predicates, group.IDNEQ(*i.IDNEQ))
	}
	if len(i.IDIn) > 0 {
		predicates = append(predicates, group.IDIn(i.IDIn...))
	}
	if len(i.IDNotIn) > 0 {
		predicates = append(predicates, group.IDNotIn(i.IDNotIn...))
	}
	if i.IDGT != nil {
		predicates = append(predicates, group.IDGT(*i.IDGT))
	}
	if i.IDGTE != nil {
		predicates = append(predicates, group.IDGTE(*i.IDGTE))
	}
	if i.IDLT != nil {
		predicates = append(predicates, group.IDLT(*i.IDLT))
	}
	if i.IDLTE != nil {
		predicates = append(predicates, group.IDLTE(*i.IDLTE))
	}
	if i.Name != nil {
		predicates = append(predicates, group.NameEQ(*i.Name))
	}
	if i.NameNEQ != nil {
		predicates = append(predicates, group.NameNEQ(*i.NameNEQ))
	}
	if len(i.NameIn) > 0 {
		predicates = append(predicates, group.NameIn(i.NameIn...))
	}
	if len(i.NameNotIn) > 0 {
		predicates = append(predicates, group.NameNotIn(i.NameNotIn...))
	}
	if i.NameGT != nil {
		predicates = append(predicates, group.NameGT(*i.NameGT))
	}
	if i.NameGTE != nil {
		predicates = append(predicates, group.NameGTE(*i.NameGTE))
	}
	if i.NameLT != nil {
		predicates = append(predicates, group.NameLT(*i.NameLT))
	}
	if i.NameLTE != nil {
		predicates = append(predicates, group.NameLTE(*i.NameLTE))
	}
	if i.NameContains != nil {
		predicates = append(predicates, group.NameContains(*i.NameContains))
	}
	if i.NameHasPrefix != nil {
		predicates = append(predicates, group.NameHasPrefix(*i.NameHasPrefix))
	}
	if i.NameHasSuffix != nil {
		predicates = append(predicates, group.NameHasSuffix(*i.NameHasSuffix))
	}
	if i.NameEqualFold != nil {
		predicates = append(predicates, group.NameEqualFold(*i.NameEqualFold))
	}
	if i.NameContainsFold != nil {
		predicates = append(predicates, group.NameContainsFold(*i.NameContainsFold))
	}

	if i.HasUsers != nil {
		p := group.HasUsers()
		if !*i.HasUsers {
			p = group.Not(p)
		}
		predicates = append(predicates, p)
	}
	if len(i.HasUsersWith) > 0 {
		with := make([]predicate.User, 0, len(i.HasUsersWith))
		for _, w := range i.HasUsersWith {
			p, err := w.P()
			if err != nil {
				return nil, err
			}
			with = append(with, p)
		}
		predicates = append(predicates, group.HasUsersWith(with...))
	}
	switch len(predicates) {
	case 0:
		return nil, fmt.Errorf("empty predicate GroupWhereInput")
	case 1:
		return predicates[0], nil
	default:
		return group.And(predicates...), nil
	}
}

// UnmarshalGQL implements the graphql.Unmarshaler interface
func (i *GroupWhereInput) UnmarshalGQL(obj interface{}) error {
	raw, err := json.Marshal(obj)
	if err != nil {
		return err
	}
	return json.Unmarshal(raw, i)
}

// MarshalGQL implements the graphql.Marshaler interface
//
// This function should never be called because our object is Input.
// It exists to make GQLGen aware of objects has custom unmarshal function
func (i GroupWhereInput) MarshalGQL(w io.Writer) {
	panic("GroupWhereInput.MarshalGQL not implemented")
}

// TodoWhereInput represents a where input for filtering Todo queries.
type TodoWhereInput struct {
	Not *TodoWhereInput   `json:"not,omitempty"`
	Or  []*TodoWhereInput `json:"or,omitempty"`
	And []*TodoWhereInput `json:"and,omitempty"`

	// "id" field predicates.
	ID      *uuid.UUID  `json:"id,omitempty"`
	IDNEQ   *uuid.UUID  `json:"idNEQ,omitempty"`
	IDIn    []uuid.UUID `json:"idIn,omitempty"`
	IDNotIn []uuid.UUID `json:"idNotIn,omitempty"`
	IDGT    *uuid.UUID  `json:"idGT,omitempty"`
	IDGTE   *uuid.UUID  `json:"idGTE,omitempty"`
	IDLT    *uuid.UUID  `json:"idLT,omitempty"`
	IDLTE   *uuid.UUID  `json:"idLTE,omitempty"`

	// "created_at" field predicates.
	CreatedAt      *time.Time  `json:"createdAt,omitempty"`
	CreatedAtNEQ   *time.Time  `json:"createdAtNEQ,omitempty"`
	CreatedAtIn    []time.Time `json:"createdAtIn,omitempty"`
	CreatedAtNotIn []time.Time `json:"createdAtNotIn,omitempty"`
	CreatedAtGT    *time.Time  `json:"createdAtGT,omitempty"`
	CreatedAtGTE   *time.Time  `json:"createdAtGTE,omitempty"`
	CreatedAtLT    *time.Time  `json:"createdAtLT,omitempty"`
	CreatedAtLTE   *time.Time  `json:"createdAtLTE,omitempty"`

	// "status" field predicates.
	Status      *todo.Status  `json:"status,omitempty"`
	StatusNEQ   *todo.Status  `json:"statusNEQ,omitempty"`
	StatusIn    []todo.Status `json:"statusIn,omitempty"`
	StatusNotIn []todo.Status `json:"statusNotIn,omitempty"`

	// "priority" field predicates.
	Priority      *int  `json:"priority,omitempty"`
	PriorityNEQ   *int  `json:"priorityNEQ,omitempty"`
	PriorityIn    []int `json:"priorityIn,omitempty"`
	PriorityNotIn []int `json:"priorityNotIn,omitempty"`
	PriorityGT    *int  `json:"priorityGT,omitempty"`
	PriorityGTE   *int  `json:"priorityGTE,omitempty"`
	PriorityLT    *int  `json:"priorityLT,omitempty"`
	PriorityLTE   *int  `json:"priorityLTE,omitempty"`

	// "text" field predicates.
	Text             *string  `json:"text,omitempty"`
	TextNEQ          *string  `json:"textNEQ,omitempty"`
	TextIn           []string `json:"textIn,omitempty"`
	TextNotIn        []string `json:"textNotIn,omitempty"`
	TextGT           *string  `json:"textGT,omitempty"`
	TextGTE          *string  `json:"textGTE,omitempty"`
	TextLT           *string  `json:"textLT,omitempty"`
	TextLTE          *string  `json:"textLTE,omitempty"`
	TextContains     *string  `json:"textContains,omitempty"`
	TextHasPrefix    *string  `json:"textHasPrefix,omitempty"`
	TextHasSuffix    *string  `json:"textHasSuffix,omitempty"`
	TextEqualFold    *string  `json:"textEqualFold,omitempty"`
	TextContainsFold *string  `json:"textContainsFold,omitempty"`

	// "parent" edge predicates.
	HasParent     *bool             `json:"hasParent,omitempty"`
	HasParentWith []*TodoWhereInput `json:"hasParentWith,omitempty"`

	// "children" edge predicates.
	HasChildren     *bool             `json:"hasChildren,omitempty"`
	HasChildrenWith []*TodoWhereInput `json:"hasChildrenWith,omitempty"`

	// "category" edge predicates.
	HasCategory     *bool                 `json:"hasCategory,omitempty"`
	HasCategoryWith []*CategoryWhereInput `json:"hasCategoryWith,omitempty"`
}

// Filter applies the TodoWhereInput filter on the TodoQuery builder.
func (i *TodoWhereInput) Filter(q *TodoQuery) (*TodoQuery, error) {
	if i == nil {
		return q, nil
	}
	p, err := i.P()
	if err != nil {
		return nil, err
	}
	return q.Where(p), nil
}

// P returns a predicate for filtering todos.
// An error is returned if the input is empty or invalid.
func (i *TodoWhereInput) P() (predicate.Todo, error) {
	var predicates []predicate.Todo
	if i.Not != nil {
		p, err := i.Not.P()
		if err != nil {
			return nil, err
		}
		predicates = append(predicates, todo.Not(p))
	}
	switch n := len(i.Or); {
	case n == 1:
		p, err := i.Or[0].P()
		if err != nil {
			return nil, err
		}
		predicates = append(predicates, p)
	case n > 1:
		or := make([]predicate.Todo, 0, n)
		for _, w := range i.Or {
			p, err := w.P()
			if err != nil {
				return nil, err
			}
			or = append(or, p)
		}
		predicates = append(predicates, todo.Or(or...))
	}
	switch n := len(i.And); {
	case n == 1:
		p, err := i.And[0].P()
		if err != nil {
			return nil, err
		}
		predicates = append(predicates, p)
	case n > 1:
		and := make([]predicate.Todo, 0, n)
		for _, w := range i.And {
			p, err := w.P()
			if err != nil {
				return nil, err
			}
			and = append(and, p)
		}
		predicates = append(predicates, todo.And(and...))
	}
	if i.ID != nil {
		predicates = append(predicates, todo.IDEQ(*i.ID))
	}
	if i.IDNEQ != nil {
		predicates = append(predicates, todo.IDNEQ(*i.IDNEQ))
	}
	if len(i.IDIn) > 0 {
		predicates = append(predicates, todo.IDIn(i.IDIn...))
	}
	if len(i.IDNotIn) > 0 {
		predicates = append(predicates, todo.IDNotIn(i.IDNotIn...))
	}
	if i.IDGT != nil {
		predicates = append(predicates, todo.IDGT(*i.IDGT))
	}
	if i.IDGTE != nil {
		predicates = append(predicates, todo.IDGTE(*i.IDGTE))
	}
	if i.IDLT != nil {
		predicates = append(predicates, todo.IDLT(*i.IDLT))
	}
	if i.IDLTE != nil {
		predicates = append(predicates, todo.IDLTE(*i.IDLTE))
	}
	if i.CreatedAt != nil {
		predicates = append(predicates, todo.CreatedAtEQ(*i.CreatedAt))
	}
	if i.CreatedAtNEQ != nil {
		predicates = append(predicates, todo.CreatedAtNEQ(*i.CreatedAtNEQ))
	}
	if len(i.CreatedAtIn) > 0 {
		predicates = append(predicates, todo.CreatedAtIn(i.CreatedAtIn...))
	}
	if len(i.CreatedAtNotIn) > 0 {
		predicates = append(predicates, todo.CreatedAtNotIn(i.CreatedAtNotIn...))
	}
	if i.CreatedAtGT != nil {
		predicates = append(predicates, todo.CreatedAtGT(*i.CreatedAtGT))
	}
	if i.CreatedAtGTE != nil {
		predicates = append(predicates, todo.CreatedAtGTE(*i.CreatedAtGTE))
	}
	if i.CreatedAtLT != nil {
		predicates = append(predicates, todo.CreatedAtLT(*i.CreatedAtLT))
	}
	if i.CreatedAtLTE != nil {
		predicates = append(predicates, todo.CreatedAtLTE(*i.CreatedAtLTE))
	}
	if i.Status != nil {
		predicates = append(predicates, todo.StatusEQ(*i.Status))
	}
	if i.StatusNEQ != nil {
		predicates = append(predicates, todo.StatusNEQ(*i.StatusNEQ))
	}
	if len(i.StatusIn) > 0 {
		predicates = append(predicates, todo.StatusIn(i.StatusIn...))
	}
	if len(i.StatusNotIn) > 0 {
		predicates = append(predicates, todo.StatusNotIn(i.StatusNotIn...))
	}
	if i.Priority != nil {
		predicates = append(predicates, todo.PriorityEQ(*i.Priority))
	}
	if i.PriorityNEQ != nil {
		predicates = append(predicates, todo.PriorityNEQ(*i.PriorityNEQ))
	}
	if len(i.PriorityIn) > 0 {
		predicates = append(predicates, todo.PriorityIn(i.PriorityIn...))
	}
	if len(i.PriorityNotIn) > 0 {
		predicates = append(predicates, todo.PriorityNotIn(i.PriorityNotIn...))
	}
	if i.PriorityGT != nil {
		predicates = append(predicates, todo.PriorityGT(*i.PriorityGT))
	}
	if i.PriorityGTE != nil {
		predicates = append(predicates, todo.PriorityGTE(*i.PriorityGTE))
	}
	if i.PriorityLT != nil {
		predicates = append(predicates, todo.PriorityLT(*i.PriorityLT))
	}
	if i.PriorityLTE != nil {
		predicates = append(predicates, todo.PriorityLTE(*i.PriorityLTE))
	}
	if i.Text != nil {
		predicates = append(predicates, todo.TextEQ(*i.Text))
	}
	if i.TextNEQ != nil {
		predicates = append(predicates, todo.TextNEQ(*i.TextNEQ))
	}
	if len(i.TextIn) > 0 {
		predicates = append(predicates, todo.TextIn(i.TextIn...))
	}
	if len(i.TextNotIn) > 0 {
		predicates = append(predicates, todo.TextNotIn(i.TextNotIn...))
	}
	if i.TextGT != nil {
		predicates = append(predicates, todo.TextGT(*i.TextGT))
	}
	if i.TextGTE != nil {
		predicates = append(predicates, todo.TextGTE(*i.TextGTE))
	}
	if i.TextLT != nil {
		predicates = append(predicates, todo.TextLT(*i.TextLT))
	}
	if i.TextLTE != nil {
		predicates = append(predicates, todo.TextLTE(*i.TextLTE))
	}
	if i.TextContains != nil {
		predicates = append(predicates, todo.TextContains(*i.TextContains))
	}
	if i.TextHasPrefix != nil {
		predicates = append(predicates, todo.TextHasPrefix(*i.TextHasPrefix))
	}
	if i.TextHasSuffix != nil {
		predicates = append(predicates, todo.TextHasSuffix(*i.TextHasSuffix))
	}
	if i.TextEqualFold != nil {
		predicates = append(predicates, todo.TextEqualFold(*i.TextEqualFold))
	}
	if i.TextContainsFold != nil {
		predicates = append(predicates, todo.TextContainsFold(*i.TextContainsFold))
	}

	if i.HasParent != nil {
		p := todo.HasParent()
		if !*i.HasParent {
			p = todo.Not(p)
		}
		predicates = append(predicates, p)
	}
	if len(i.HasParentWith) > 0 {
		with := make([]predicate.Todo, 0, len(i.HasParentWith))
		for _, w := range i.HasParentWith {
			p, err := w.P()
			if err != nil {
				return nil, err
			}
			with = append(with, p)
		}
		predicates = append(predicates, todo.HasParentWith(with...))
	}
	if i.HasChildren != nil {
		p := todo.HasChildren()
		if !*i.HasChildren {
			p = todo.Not(p)
		}
		predicates = append(predicates, p)
	}
	if len(i.HasChildrenWith) > 0 {
		with := make([]predicate.Todo, 0, len(i.HasChildrenWith))
		for _, w := range i.HasChildrenWith {
			p, err := w.P()
			if err != nil {
				return nil, err
			}
			with = append(with, p)
		}
		predicates = append(predicates, todo.HasChildrenWith(with...))
	}
	if i.HasCategory != nil {
		p := todo.HasCategory()
		if !*i.HasCategory {
			p = todo.Not(p)
		}
		predicates = append(predicates, p)
	}
	if len(i.HasCategoryWith) > 0 {
		with := make([]predicate.Category, 0, len(i.HasCategoryWith))
		for _, w := range i.HasCategoryWith {
			p, err := w.P()
			if err != nil {
				return nil, err
			}
			with = append(with, p)
		}
		predicates = append(predicates, todo.HasCategoryWith(with...))
	}
	switch len(predicates) {
	case 0:
		return nil, fmt.Errorf("empty predicate TodoWhereInput")
	case 1:
		return predicates[0], nil
	default:
		return todo.And(predicates...), nil
	}
}

// UnmarshalGQL implements the graphql.Unmarshaler interface
func (i *TodoWhereInput) UnmarshalGQL(obj interface{}) error {
	raw, err := json.Marshal(obj)
	if err != nil {
		return err
	}
	return json.Unmarshal(raw, i)
}

// MarshalGQL implements the graphql.Marshaler interface
//
// This function should never be called because our object is Input.
// It exists to make GQLGen aware of objects has custom unmarshal function
func (i TodoWhereInput) MarshalGQL(w io.Writer) {
	panic("TodoWhereInput.MarshalGQL not implemented")
}

// UserWhereInput represents a where input for filtering User queries.
type UserWhereInput struct {
	Not *UserWhereInput   `json:"not,omitempty"`
	Or  []*UserWhereInput `json:"or,omitempty"`
	And []*UserWhereInput `json:"and,omitempty"`

	// "id" field predicates.
	ID      *uuid.UUID  `json:"id,omitempty"`
	IDNEQ   *uuid.UUID  `json:"idNEQ,omitempty"`
	IDIn    []uuid.UUID `json:"idIn,omitempty"`
	IDNotIn []uuid.UUID `json:"idNotIn,omitempty"`
	IDGT    *uuid.UUID  `json:"idGT,omitempty"`
	IDGTE   *uuid.UUID  `json:"idGTE,omitempty"`
	IDLT    *uuid.UUID  `json:"idLT,omitempty"`
	IDLTE   *uuid.UUID  `json:"idLTE,omitempty"`

	// "name" field predicates.
	Name             *string  `json:"name,omitempty"`
	NameNEQ          *string  `json:"nameNEQ,omitempty"`
	NameIn           []string `json:"nameIn,omitempty"`
	NameNotIn        []string `json:"nameNotIn,omitempty"`
	NameGT           *string  `json:"nameGT,omitempty"`
	NameGTE          *string  `json:"nameGTE,omitempty"`
	NameLT           *string  `json:"nameLT,omitempty"`
	NameLTE          *string  `json:"nameLTE,omitempty"`
	NameContains     *string  `json:"nameContains,omitempty"`
	NameHasPrefix    *string  `json:"nameHasPrefix,omitempty"`
	NameHasSuffix    *string  `json:"nameHasSuffix,omitempty"`
	NameEqualFold    *string  `json:"nameEqualFold,omitempty"`
	NameContainsFold *string  `json:"nameContainsFold,omitempty"`

	// "groups" edge predicates.
	HasGroups     *bool              `json:"hasGroups,omitempty"`
	HasGroupsWith []*GroupWhereInput `json:"hasGroupsWith,omitempty"`
}

// Filter applies the UserWhereInput filter on the UserQuery builder.
func (i *UserWhereInput) Filter(q *UserQuery) (*UserQuery, error) {
	if i == nil {
		return q, nil
	}
	p, err := i.P()
	if err != nil {
		return nil, err
	}
	return q.Where(p), nil
}

// P returns a predicate for filtering users.
// An error is returned if the input is empty or invalid.
func (i *UserWhereInput) P() (predicate.User, error) {
	var predicates []predicate.User
	if i.Not != nil {
		p, err := i.Not.P()
		if err != nil {
			return nil, err
		}
		predicates = append(predicates, user.Not(p))
	}
	switch n := len(i.Or); {
	case n == 1:
		p, err := i.Or[0].P()
		if err != nil {
			return nil, err
		}
		predicates = append(predicates, p)
	case n > 1:
		or := make([]predicate.User, 0, n)
		for _, w := range i.Or {
			p, err := w.P()
			if err != nil {
				return nil, err
			}
			or = append(or, p)
		}
		predicates = append(predicates, user.Or(or...))
	}
	switch n := len(i.And); {
	case n == 1:
		p, err := i.And[0].P()
		if err != nil {
			return nil, err
		}
		predicates = append(predicates, p)
	case n > 1:
		and := make([]predicate.User, 0, n)
		for _, w := range i.And {
			p, err := w.P()
			if err != nil {
				return nil, err
			}
			and = append(and, p)
		}
		predicates = append(predicates, user.And(and...))
	}
	if i.ID != nil {
		predicates = append(predicates, user.IDEQ(*i.ID))
	}
	if i.IDNEQ != nil {
		predicates = append(predicates, user.IDNEQ(*i.IDNEQ))
	}
	if len(i.IDIn) > 0 {
		predicates = append(predicates, user.IDIn(i.IDIn...))
	}
	if len(i.IDNotIn) > 0 {
		predicates = append(predicates, user.IDNotIn(i.IDNotIn...))
	}
	if i.IDGT != nil {
		predicates = append(predicates, user.IDGT(*i.IDGT))
	}
	if i.IDGTE != nil {
		predicates = append(predicates, user.IDGTE(*i.IDGTE))
	}
	if i.IDLT != nil {
		predicates = append(predicates, user.IDLT(*i.IDLT))
	}
	if i.IDLTE != nil {
		predicates = append(predicates, user.IDLTE(*i.IDLTE))
	}
	if i.Name != nil {
		predicates = append(predicates, user.NameEQ(*i.Name))
	}
	if i.NameNEQ != nil {
		predicates = append(predicates, user.NameNEQ(*i.NameNEQ))
	}
	if len(i.NameIn) > 0 {
		predicates = append(predicates, user.NameIn(i.NameIn...))
	}
	if len(i.NameNotIn) > 0 {
		predicates = append(predicates, user.NameNotIn(i.NameNotIn...))
	}
	if i.NameGT != nil {
		predicates = append(predicates, user.NameGT(*i.NameGT))
	}
	if i.NameGTE != nil {
		predicates = append(predicates, user.NameGTE(*i.NameGTE))
	}
	if i.NameLT != nil {
		predicates = append(predicates, user.NameLT(*i.NameLT))
	}
	if i.NameLTE != nil {
		predicates = append(predicates, user.NameLTE(*i.NameLTE))
	}
	if i.NameContains != nil {
		predicates = append(predicates, user.NameContains(*i.NameContains))
	}
	if i.NameHasPrefix != nil {
		predicates = append(predicates, user.NameHasPrefix(*i.NameHasPrefix))
	}
	if i.NameHasSuffix != nil {
		predicates = append(predicates, user.NameHasSuffix(*i.NameHasSuffix))
	}
	if i.NameEqualFold != nil {
		predicates = append(predicates, user.NameEqualFold(*i.NameEqualFold))
	}
	if i.NameContainsFold != nil {
		predicates = append(predicates, user.NameContainsFold(*i.NameContainsFold))
	}

	if i.HasGroups != nil {
		p := user.HasGroups()
		if !*i.HasGroups {
			p = user.Not(p)
		}
		predicates = append(predicates, p)
	}
	if len(i.HasGroupsWith) > 0 {
		with := make([]predicate.Group, 0, len(i.HasGroupsWith))
		for _, w := range i.HasGroupsWith {
			p, err := w.P()
			if err != nil {
				return nil, err
			}
			with = append(with, p)
		}
		predicates = append(predicates, user.HasGroupsWith(with...))
	}
	switch len(predicates) {
	case 0:
		return nil, fmt.Errorf("empty predicate UserWhereInput")
	case 1:
		return predicates[0], nil
	default:
		return user.And(predicates...), nil
	}
}

// UnmarshalGQL implements the graphql.Unmarshaler interface
func (i *UserWhereInput) UnmarshalGQL(obj interface{}) error {
	raw, err := json.Marshal(obj)
	if err != nil {
		return err
	}
	return json.Unmarshal(raw, i)
}

// MarshalGQL implements the graphql.Marshaler interface
//
// This function should never be called because our object is Input.
// It exists to make GQLGen aware of objects has custom unmarshal function
func (i UserWhereInput) MarshalGQL(w io.Writer) {
	panic("UserWhereInput.MarshalGQL not implemented")
}
