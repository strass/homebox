// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/sysadminsmedia/homebox/backend/internal/data/ent/authroles"
	"github.com/sysadminsmedia/homebox/backend/internal/data/ent/authtokens"
	"github.com/sysadminsmedia/homebox/backend/internal/data/ent/predicate"
)

// AuthRolesUpdate is the builder for updating AuthRoles entities.
type AuthRolesUpdate struct {
	config
	hooks    []Hook
	mutation *AuthRolesMutation
}

// Where appends a list predicates to the AuthRolesUpdate builder.
func (aru *AuthRolesUpdate) Where(ps ...predicate.AuthRoles) *AuthRolesUpdate {
	aru.mutation.Where(ps...)
	return aru
}

// SetRole sets the "role" field.
func (aru *AuthRolesUpdate) SetRole(a authroles.Role) *AuthRolesUpdate {
	aru.mutation.SetRole(a)
	return aru
}

// SetNillableRole sets the "role" field if the given value is not nil.
func (aru *AuthRolesUpdate) SetNillableRole(a *authroles.Role) *AuthRolesUpdate {
	if a != nil {
		aru.SetRole(*a)
	}
	return aru
}

// SetTokenID sets the "token" edge to the AuthTokens entity by ID.
func (aru *AuthRolesUpdate) SetTokenID(id uuid.UUID) *AuthRolesUpdate {
	aru.mutation.SetTokenID(id)
	return aru
}

// SetNillableTokenID sets the "token" edge to the AuthTokens entity by ID if the given value is not nil.
func (aru *AuthRolesUpdate) SetNillableTokenID(id *uuid.UUID) *AuthRolesUpdate {
	if id != nil {
		aru = aru.SetTokenID(*id)
	}
	return aru
}

// SetToken sets the "token" edge to the AuthTokens entity.
func (aru *AuthRolesUpdate) SetToken(a *AuthTokens) *AuthRolesUpdate {
	return aru.SetTokenID(a.ID)
}

// Mutation returns the AuthRolesMutation object of the builder.
func (aru *AuthRolesUpdate) Mutation() *AuthRolesMutation {
	return aru.mutation
}

// ClearToken clears the "token" edge to the AuthTokens entity.
func (aru *AuthRolesUpdate) ClearToken() *AuthRolesUpdate {
	aru.mutation.ClearToken()
	return aru
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (aru *AuthRolesUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, aru.sqlSave, aru.mutation, aru.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (aru *AuthRolesUpdate) SaveX(ctx context.Context) int {
	affected, err := aru.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (aru *AuthRolesUpdate) Exec(ctx context.Context) error {
	_, err := aru.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (aru *AuthRolesUpdate) ExecX(ctx context.Context) {
	if err := aru.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (aru *AuthRolesUpdate) check() error {
	if v, ok := aru.mutation.Role(); ok {
		if err := authroles.RoleValidator(v); err != nil {
			return &ValidationError{Name: "role", err: fmt.Errorf(`ent: validator failed for field "AuthRoles.role": %w`, err)}
		}
	}
	return nil
}

func (aru *AuthRolesUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := aru.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(authroles.Table, authroles.Columns, sqlgraph.NewFieldSpec(authroles.FieldID, field.TypeInt))
	if ps := aru.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := aru.mutation.Role(); ok {
		_spec.SetField(authroles.FieldRole, field.TypeEnum, value)
	}
	if aru.mutation.TokenCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   authroles.TokenTable,
			Columns: []string{authroles.TokenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(authtokens.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := aru.mutation.TokenIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   authroles.TokenTable,
			Columns: []string{authroles.TokenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(authtokens.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, aru.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{authroles.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	aru.mutation.done = true
	return n, nil
}

// AuthRolesUpdateOne is the builder for updating a single AuthRoles entity.
type AuthRolesUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *AuthRolesMutation
}

// SetRole sets the "role" field.
func (aruo *AuthRolesUpdateOne) SetRole(a authroles.Role) *AuthRolesUpdateOne {
	aruo.mutation.SetRole(a)
	return aruo
}

// SetNillableRole sets the "role" field if the given value is not nil.
func (aruo *AuthRolesUpdateOne) SetNillableRole(a *authroles.Role) *AuthRolesUpdateOne {
	if a != nil {
		aruo.SetRole(*a)
	}
	return aruo
}

// SetTokenID sets the "token" edge to the AuthTokens entity by ID.
func (aruo *AuthRolesUpdateOne) SetTokenID(id uuid.UUID) *AuthRolesUpdateOne {
	aruo.mutation.SetTokenID(id)
	return aruo
}

// SetNillableTokenID sets the "token" edge to the AuthTokens entity by ID if the given value is not nil.
func (aruo *AuthRolesUpdateOne) SetNillableTokenID(id *uuid.UUID) *AuthRolesUpdateOne {
	if id != nil {
		aruo = aruo.SetTokenID(*id)
	}
	return aruo
}

// SetToken sets the "token" edge to the AuthTokens entity.
func (aruo *AuthRolesUpdateOne) SetToken(a *AuthTokens) *AuthRolesUpdateOne {
	return aruo.SetTokenID(a.ID)
}

// Mutation returns the AuthRolesMutation object of the builder.
func (aruo *AuthRolesUpdateOne) Mutation() *AuthRolesMutation {
	return aruo.mutation
}

// ClearToken clears the "token" edge to the AuthTokens entity.
func (aruo *AuthRolesUpdateOne) ClearToken() *AuthRolesUpdateOne {
	aruo.mutation.ClearToken()
	return aruo
}

// Where appends a list predicates to the AuthRolesUpdate builder.
func (aruo *AuthRolesUpdateOne) Where(ps ...predicate.AuthRoles) *AuthRolesUpdateOne {
	aruo.mutation.Where(ps...)
	return aruo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (aruo *AuthRolesUpdateOne) Select(field string, fields ...string) *AuthRolesUpdateOne {
	aruo.fields = append([]string{field}, fields...)
	return aruo
}

// Save executes the query and returns the updated AuthRoles entity.
func (aruo *AuthRolesUpdateOne) Save(ctx context.Context) (*AuthRoles, error) {
	return withHooks(ctx, aruo.sqlSave, aruo.mutation, aruo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (aruo *AuthRolesUpdateOne) SaveX(ctx context.Context) *AuthRoles {
	node, err := aruo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (aruo *AuthRolesUpdateOne) Exec(ctx context.Context) error {
	_, err := aruo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (aruo *AuthRolesUpdateOne) ExecX(ctx context.Context) {
	if err := aruo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (aruo *AuthRolesUpdateOne) check() error {
	if v, ok := aruo.mutation.Role(); ok {
		if err := authroles.RoleValidator(v); err != nil {
			return &ValidationError{Name: "role", err: fmt.Errorf(`ent: validator failed for field "AuthRoles.role": %w`, err)}
		}
	}
	return nil
}

func (aruo *AuthRolesUpdateOne) sqlSave(ctx context.Context) (_node *AuthRoles, err error) {
	if err := aruo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(authroles.Table, authroles.Columns, sqlgraph.NewFieldSpec(authroles.FieldID, field.TypeInt))
	id, ok := aruo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "AuthRoles.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := aruo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, authroles.FieldID)
		for _, f := range fields {
			if !authroles.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != authroles.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := aruo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := aruo.mutation.Role(); ok {
		_spec.SetField(authroles.FieldRole, field.TypeEnum, value)
	}
	if aruo.mutation.TokenCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   authroles.TokenTable,
			Columns: []string{authroles.TokenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(authtokens.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := aruo.mutation.TokenIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   authroles.TokenTable,
			Columns: []string{authroles.TokenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(authtokens.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &AuthRoles{config: aruo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, aruo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{authroles.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	aruo.mutation.done = true
	return _node, nil
}
