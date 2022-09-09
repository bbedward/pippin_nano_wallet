// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/appditto/pippin_nano_wallet/libs/database/ent/account"
	"github.com/appditto/pippin_nano_wallet/libs/database/ent/adhocaccount"
	"github.com/appditto/pippin_nano_wallet/libs/database/ent/wallet"
	"github.com/google/uuid"
)

// WalletCreate is the builder for creating a Wallet entity.
type WalletCreate struct {
	config
	mutation *WalletMutation
	hooks    []Hook
}

// SetSeed sets the "seed" field.
func (wc *WalletCreate) SetSeed(s string) *WalletCreate {
	wc.mutation.SetSeed(s)
	return wc
}

// SetRepresentative sets the "representative" field.
func (wc *WalletCreate) SetRepresentative(s string) *WalletCreate {
	wc.mutation.SetRepresentative(s)
	return wc
}

// SetNillableRepresentative sets the "representative" field if the given value is not nil.
func (wc *WalletCreate) SetNillableRepresentative(s *string) *WalletCreate {
	if s != nil {
		wc.SetRepresentative(*s)
	}
	return wc
}

// SetEncrypted sets the "encrypted" field.
func (wc *WalletCreate) SetEncrypted(b bool) *WalletCreate {
	wc.mutation.SetEncrypted(b)
	return wc
}

// SetNillableEncrypted sets the "encrypted" field if the given value is not nil.
func (wc *WalletCreate) SetNillableEncrypted(b *bool) *WalletCreate {
	if b != nil {
		wc.SetEncrypted(*b)
	}
	return wc
}

// SetWork sets the "work" field.
func (wc *WalletCreate) SetWork(b bool) *WalletCreate {
	wc.mutation.SetWork(b)
	return wc
}

// SetNillableWork sets the "work" field if the given value is not nil.
func (wc *WalletCreate) SetNillableWork(b *bool) *WalletCreate {
	if b != nil {
		wc.SetWork(*b)
	}
	return wc
}

// SetCreatedAt sets the "created_at" field.
func (wc *WalletCreate) SetCreatedAt(t time.Time) *WalletCreate {
	wc.mutation.SetCreatedAt(t)
	return wc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (wc *WalletCreate) SetNillableCreatedAt(t *time.Time) *WalletCreate {
	if t != nil {
		wc.SetCreatedAt(*t)
	}
	return wc
}

// SetID sets the "id" field.
func (wc *WalletCreate) SetID(u uuid.UUID) *WalletCreate {
	wc.mutation.SetID(u)
	return wc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (wc *WalletCreate) SetNillableID(u *uuid.UUID) *WalletCreate {
	if u != nil {
		wc.SetID(*u)
	}
	return wc
}

// AddAccountIDs adds the "accounts" edge to the Account entity by IDs.
func (wc *WalletCreate) AddAccountIDs(ids ...uuid.UUID) *WalletCreate {
	wc.mutation.AddAccountIDs(ids...)
	return wc
}

// AddAccounts adds the "accounts" edges to the Account entity.
func (wc *WalletCreate) AddAccounts(a ...*Account) *WalletCreate {
	ids := make([]uuid.UUID, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return wc.AddAccountIDs(ids...)
}

// AddAdhocAccountIDs adds the "adhoc_accounts" edge to the AdhocAccount entity by IDs.
func (wc *WalletCreate) AddAdhocAccountIDs(ids ...uuid.UUID) *WalletCreate {
	wc.mutation.AddAdhocAccountIDs(ids...)
	return wc
}

// AddAdhocAccounts adds the "adhoc_accounts" edges to the AdhocAccount entity.
func (wc *WalletCreate) AddAdhocAccounts(a ...*AdhocAccount) *WalletCreate {
	ids := make([]uuid.UUID, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return wc.AddAdhocAccountIDs(ids...)
}

// Mutation returns the WalletMutation object of the builder.
func (wc *WalletCreate) Mutation() *WalletMutation {
	return wc.mutation
}

// Save creates the Wallet in the database.
func (wc *WalletCreate) Save(ctx context.Context) (*Wallet, error) {
	var (
		err  error
		node *Wallet
	)
	wc.defaults()
	if len(wc.hooks) == 0 {
		if err = wc.check(); err != nil {
			return nil, err
		}
		node, err = wc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*WalletMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = wc.check(); err != nil {
				return nil, err
			}
			wc.mutation = mutation
			if node, err = wc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(wc.hooks) - 1; i >= 0; i-- {
			if wc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = wc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, wc.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Wallet)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from WalletMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (wc *WalletCreate) SaveX(ctx context.Context) *Wallet {
	v, err := wc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (wc *WalletCreate) Exec(ctx context.Context) error {
	_, err := wc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (wc *WalletCreate) ExecX(ctx context.Context) {
	if err := wc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (wc *WalletCreate) defaults() {
	if _, ok := wc.mutation.Encrypted(); !ok {
		v := wallet.DefaultEncrypted
		wc.mutation.SetEncrypted(v)
	}
	if _, ok := wc.mutation.Work(); !ok {
		v := wallet.DefaultWork
		wc.mutation.SetWork(v)
	}
	if _, ok := wc.mutation.CreatedAt(); !ok {
		v := wallet.DefaultCreatedAt()
		wc.mutation.SetCreatedAt(v)
	}
	if _, ok := wc.mutation.ID(); !ok {
		v := wallet.DefaultID()
		wc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (wc *WalletCreate) check() error {
	if _, ok := wc.mutation.Seed(); !ok {
		return &ValidationError{Name: "seed", err: errors.New(`ent: missing required field "Wallet.seed"`)}
	}
	if v, ok := wc.mutation.Seed(); ok {
		if err := wallet.SeedValidator(v); err != nil {
			return &ValidationError{Name: "seed", err: fmt.Errorf(`ent: validator failed for field "Wallet.seed": %w`, err)}
		}
	}
	if v, ok := wc.mutation.Representative(); ok {
		if err := wallet.RepresentativeValidator(v); err != nil {
			return &ValidationError{Name: "representative", err: fmt.Errorf(`ent: validator failed for field "Wallet.representative": %w`, err)}
		}
	}
	if _, ok := wc.mutation.Encrypted(); !ok {
		return &ValidationError{Name: "encrypted", err: errors.New(`ent: missing required field "Wallet.encrypted"`)}
	}
	if _, ok := wc.mutation.Work(); !ok {
		return &ValidationError{Name: "work", err: errors.New(`ent: missing required field "Wallet.work"`)}
	}
	if _, ok := wc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Wallet.created_at"`)}
	}
	return nil
}

func (wc *WalletCreate) sqlSave(ctx context.Context) (*Wallet, error) {
	_node, _spec := wc.createSpec()
	if err := sqlgraph.CreateNode(ctx, wc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*uuid.UUID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	return _node, nil
}

func (wc *WalletCreate) createSpec() (*Wallet, *sqlgraph.CreateSpec) {
	var (
		_node = &Wallet{config: wc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: wallet.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: wallet.FieldID,
			},
		}
	)
	if id, ok := wc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := wc.mutation.Seed(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: wallet.FieldSeed,
		})
		_node.Seed = value
	}
	if value, ok := wc.mutation.Representative(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: wallet.FieldRepresentative,
		})
		_node.Representative = &value
	}
	if value, ok := wc.mutation.Encrypted(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: wallet.FieldEncrypted,
		})
		_node.Encrypted = value
	}
	if value, ok := wc.mutation.Work(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: wallet.FieldWork,
		})
		_node.Work = value
	}
	if value, ok := wc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: wallet.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if nodes := wc.mutation.AccountsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   wallet.AccountsTable,
			Columns: []string{wallet.AccountsColumn},
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := wc.mutation.AdhocAccountsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   wallet.AdhocAccountsTable,
			Columns: []string{wallet.AdhocAccountsColumn},
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// WalletCreateBulk is the builder for creating many Wallet entities in bulk.
type WalletCreateBulk struct {
	config
	builders []*WalletCreate
}

// Save creates the Wallet entities in the database.
func (wcb *WalletCreateBulk) Save(ctx context.Context) ([]*Wallet, error) {
	specs := make([]*sqlgraph.CreateSpec, len(wcb.builders))
	nodes := make([]*Wallet, len(wcb.builders))
	mutators := make([]Mutator, len(wcb.builders))
	for i := range wcb.builders {
		func(i int, root context.Context) {
			builder := wcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*WalletMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, wcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, wcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, wcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (wcb *WalletCreateBulk) SaveX(ctx context.Context) []*Wallet {
	v, err := wcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (wcb *WalletCreateBulk) Exec(ctx context.Context) error {
	_, err := wcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (wcb *WalletCreateBulk) ExecX(ctx context.Context) {
	if err := wcb.Exec(ctx); err != nil {
		panic(err)
	}
}
