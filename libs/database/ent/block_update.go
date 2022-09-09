// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/appditto/pippin_nano_wallet/libs/database/ent/account"
	"github.com/appditto/pippin_nano_wallet/libs/database/ent/adhocaccount"
	"github.com/appditto/pippin_nano_wallet/libs/database/ent/block"
	"github.com/appditto/pippin_nano_wallet/libs/database/ent/predicate"
	"github.com/google/uuid"
)

// BlockUpdate is the builder for updating Block entities.
type BlockUpdate struct {
	config
	hooks    []Hook
	mutation *BlockMutation
}

// Where appends a list predicates to the BlockUpdate builder.
func (bu *BlockUpdate) Where(ps ...predicate.Block) *BlockUpdate {
	bu.mutation.Where(ps...)
	return bu
}

// SetAccountID sets the "account_id" field.
func (bu *BlockUpdate) SetAccountID(u uuid.UUID) *BlockUpdate {
	bu.mutation.SetAccountID(u)
	return bu
}

// SetAdhocAccountID sets the "adhoc_account_id" field.
func (bu *BlockUpdate) SetAdhocAccountID(u uuid.UUID) *BlockUpdate {
	bu.mutation.SetAdhocAccountID(u)
	return bu
}

// SetSubtype sets the "subtype" field.
func (bu *BlockUpdate) SetSubtype(s string) *BlockUpdate {
	bu.mutation.SetSubtype(s)
	return bu
}

// SetAccount sets the "account" edge to the Account entity.
func (bu *BlockUpdate) SetAccount(a *Account) *BlockUpdate {
	return bu.SetAccountID(a.ID)
}

// SetAdhocAccount sets the "adhoc_account" edge to the AdhocAccount entity.
func (bu *BlockUpdate) SetAdhocAccount(a *AdhocAccount) *BlockUpdate {
	return bu.SetAdhocAccountID(a.ID)
}

// Mutation returns the BlockMutation object of the builder.
func (bu *BlockUpdate) Mutation() *BlockMutation {
	return bu.mutation
}

// ClearAccount clears the "account" edge to the Account entity.
func (bu *BlockUpdate) ClearAccount() *BlockUpdate {
	bu.mutation.ClearAccount()
	return bu
}

// ClearAdhocAccount clears the "adhoc_account" edge to the AdhocAccount entity.
func (bu *BlockUpdate) ClearAdhocAccount() *BlockUpdate {
	bu.mutation.ClearAdhocAccount()
	return bu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (bu *BlockUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(bu.hooks) == 0 {
		if err = bu.check(); err != nil {
			return 0, err
		}
		affected, err = bu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*BlockMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = bu.check(); err != nil {
				return 0, err
			}
			bu.mutation = mutation
			affected, err = bu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(bu.hooks) - 1; i >= 0; i-- {
			if bu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = bu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, bu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (bu *BlockUpdate) SaveX(ctx context.Context) int {
	affected, err := bu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (bu *BlockUpdate) Exec(ctx context.Context) error {
	_, err := bu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bu *BlockUpdate) ExecX(ctx context.Context) {
	if err := bu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (bu *BlockUpdate) check() error {
	if v, ok := bu.mutation.Subtype(); ok {
		if err := block.SubtypeValidator(v); err != nil {
			return &ValidationError{Name: "subtype", err: fmt.Errorf(`ent: validator failed for field "Block.subtype": %w`, err)}
		}
	}
	if _, ok := bu.mutation.AccountID(); bu.mutation.AccountCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Block.account"`)
	}
	if _, ok := bu.mutation.AdhocAccountID(); bu.mutation.AdhocAccountCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Block.adhoc_account"`)
	}
	return nil
}

func (bu *BlockUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   block.Table,
			Columns: block.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: block.FieldID,
			},
		},
	}
	if ps := bu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if bu.mutation.SendIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: block.FieldSendID,
		})
	}
	if value, ok := bu.mutation.Subtype(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: block.FieldSubtype,
		})
	}
	if bu.mutation.AccountCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   block.AccountTable,
			Columns: []string{block.AccountColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: account.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := bu.mutation.AccountIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   block.AccountTable,
			Columns: []string{block.AccountColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: account.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if bu.mutation.AdhocAccountCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   block.AdhocAccountTable,
			Columns: []string{block.AdhocAccountColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: adhocaccount.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := bu.mutation.AdhocAccountIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   block.AdhocAccountTable,
			Columns: []string{block.AdhocAccountColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: adhocaccount.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, bu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{block.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// BlockUpdateOne is the builder for updating a single Block entity.
type BlockUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *BlockMutation
}

// SetAccountID sets the "account_id" field.
func (buo *BlockUpdateOne) SetAccountID(u uuid.UUID) *BlockUpdateOne {
	buo.mutation.SetAccountID(u)
	return buo
}

// SetAdhocAccountID sets the "adhoc_account_id" field.
func (buo *BlockUpdateOne) SetAdhocAccountID(u uuid.UUID) *BlockUpdateOne {
	buo.mutation.SetAdhocAccountID(u)
	return buo
}

// SetSubtype sets the "subtype" field.
func (buo *BlockUpdateOne) SetSubtype(s string) *BlockUpdateOne {
	buo.mutation.SetSubtype(s)
	return buo
}

// SetAccount sets the "account" edge to the Account entity.
func (buo *BlockUpdateOne) SetAccount(a *Account) *BlockUpdateOne {
	return buo.SetAccountID(a.ID)
}

// SetAdhocAccount sets the "adhoc_account" edge to the AdhocAccount entity.
func (buo *BlockUpdateOne) SetAdhocAccount(a *AdhocAccount) *BlockUpdateOne {
	return buo.SetAdhocAccountID(a.ID)
}

// Mutation returns the BlockMutation object of the builder.
func (buo *BlockUpdateOne) Mutation() *BlockMutation {
	return buo.mutation
}

// ClearAccount clears the "account" edge to the Account entity.
func (buo *BlockUpdateOne) ClearAccount() *BlockUpdateOne {
	buo.mutation.ClearAccount()
	return buo
}

// ClearAdhocAccount clears the "adhoc_account" edge to the AdhocAccount entity.
func (buo *BlockUpdateOne) ClearAdhocAccount() *BlockUpdateOne {
	buo.mutation.ClearAdhocAccount()
	return buo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (buo *BlockUpdateOne) Select(field string, fields ...string) *BlockUpdateOne {
	buo.fields = append([]string{field}, fields...)
	return buo
}

// Save executes the query and returns the updated Block entity.
func (buo *BlockUpdateOne) Save(ctx context.Context) (*Block, error) {
	var (
		err  error
		node *Block
	)
	if len(buo.hooks) == 0 {
		if err = buo.check(); err != nil {
			return nil, err
		}
		node, err = buo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*BlockMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = buo.check(); err != nil {
				return nil, err
			}
			buo.mutation = mutation
			node, err = buo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(buo.hooks) - 1; i >= 0; i-- {
			if buo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = buo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, buo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Block)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from BlockMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (buo *BlockUpdateOne) SaveX(ctx context.Context) *Block {
	node, err := buo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (buo *BlockUpdateOne) Exec(ctx context.Context) error {
	_, err := buo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (buo *BlockUpdateOne) ExecX(ctx context.Context) {
	if err := buo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (buo *BlockUpdateOne) check() error {
	if v, ok := buo.mutation.Subtype(); ok {
		if err := block.SubtypeValidator(v); err != nil {
			return &ValidationError{Name: "subtype", err: fmt.Errorf(`ent: validator failed for field "Block.subtype": %w`, err)}
		}
	}
	if _, ok := buo.mutation.AccountID(); buo.mutation.AccountCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Block.account"`)
	}
	if _, ok := buo.mutation.AdhocAccountID(); buo.mutation.AdhocAccountCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Block.adhoc_account"`)
	}
	return nil
}

func (buo *BlockUpdateOne) sqlSave(ctx context.Context) (_node *Block, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   block.Table,
			Columns: block.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: block.FieldID,
			},
		},
	}
	id, ok := buo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Block.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := buo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, block.FieldID)
		for _, f := range fields {
			if !block.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != block.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := buo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if buo.mutation.SendIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: block.FieldSendID,
		})
	}
	if value, ok := buo.mutation.Subtype(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: block.FieldSubtype,
		})
	}
	if buo.mutation.AccountCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   block.AccountTable,
			Columns: []string{block.AccountColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: account.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := buo.mutation.AccountIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   block.AccountTable,
			Columns: []string{block.AccountColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: account.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if buo.mutation.AdhocAccountCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   block.AdhocAccountTable,
			Columns: []string{block.AdhocAccountColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: adhocaccount.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := buo.mutation.AdhocAccountIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   block.AdhocAccountTable,
			Columns: []string{block.AdhocAccountColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: adhocaccount.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Block{config: buo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, buo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{block.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
