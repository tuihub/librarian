// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/tuihub/librarian/app/sephirah/internal/data/internal/ent/app"
	"github.com/tuihub/librarian/app/sephirah/internal/data/internal/ent/apppackage"
	"github.com/tuihub/librarian/app/sephirah/internal/data/internal/ent/predicate"
	"github.com/tuihub/librarian/app/sephirah/internal/data/internal/ent/user"
	"github.com/tuihub/librarian/internal/model"
)

// AppPackageUpdate is the builder for updating AppPackage entities.
type AppPackageUpdate struct {
	config
	hooks    []Hook
	mutation *AppPackageMutation
}

// Where appends a list predicates to the AppPackageUpdate builder.
func (apu *AppPackageUpdate) Where(ps ...predicate.AppPackage) *AppPackageUpdate {
	apu.mutation.Where(ps...)
	return apu
}

// SetSource sets the "source" field.
func (apu *AppPackageUpdate) SetSource(a apppackage.Source) *AppPackageUpdate {
	apu.mutation.SetSource(a)
	return apu
}

// SetSourceID sets the "source_id" field.
func (apu *AppPackageUpdate) SetSourceID(mi model.InternalID) *AppPackageUpdate {
	apu.mutation.ResetSourceID()
	apu.mutation.SetSourceID(mi)
	return apu
}

// AddSourceID adds mi to the "source_id" field.
func (apu *AppPackageUpdate) AddSourceID(mi model.InternalID) *AppPackageUpdate {
	apu.mutation.AddSourceID(mi)
	return apu
}

// SetName sets the "name" field.
func (apu *AppPackageUpdate) SetName(s string) *AppPackageUpdate {
	apu.mutation.SetName(s)
	return apu
}

// SetDescription sets the "description" field.
func (apu *AppPackageUpdate) SetDescription(s string) *AppPackageUpdate {
	apu.mutation.SetDescription(s)
	return apu
}

// SetPublic sets the "public" field.
func (apu *AppPackageUpdate) SetPublic(b bool) *AppPackageUpdate {
	apu.mutation.SetPublic(b)
	return apu
}

// SetBinaryName sets the "binary_name" field.
func (apu *AppPackageUpdate) SetBinaryName(s string) *AppPackageUpdate {
	apu.mutation.SetBinaryName(s)
	return apu
}

// SetBinarySizeBytes sets the "binary_size_bytes" field.
func (apu *AppPackageUpdate) SetBinarySizeBytes(i int64) *AppPackageUpdate {
	apu.mutation.ResetBinarySizeBytes()
	apu.mutation.SetBinarySizeBytes(i)
	return apu
}

// AddBinarySizeBytes adds i to the "binary_size_bytes" field.
func (apu *AppPackageUpdate) AddBinarySizeBytes(i int64) *AppPackageUpdate {
	apu.mutation.AddBinarySizeBytes(i)
	return apu
}

// SetBinaryPublicURL sets the "binary_public_url" field.
func (apu *AppPackageUpdate) SetBinaryPublicURL(s string) *AppPackageUpdate {
	apu.mutation.SetBinaryPublicURL(s)
	return apu
}

// SetBinarySha256 sets the "binary_sha256" field.
func (apu *AppPackageUpdate) SetBinarySha256(b []byte) *AppPackageUpdate {
	apu.mutation.SetBinarySha256(b)
	return apu
}

// SetUpdatedAt sets the "updated_at" field.
func (apu *AppPackageUpdate) SetUpdatedAt(t time.Time) *AppPackageUpdate {
	apu.mutation.SetUpdatedAt(t)
	return apu
}

// SetCreatedAt sets the "created_at" field.
func (apu *AppPackageUpdate) SetCreatedAt(t time.Time) *AppPackageUpdate {
	apu.mutation.SetCreatedAt(t)
	return apu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (apu *AppPackageUpdate) SetNillableCreatedAt(t *time.Time) *AppPackageUpdate {
	if t != nil {
		apu.SetCreatedAt(*t)
	}
	return apu
}

// SetOwnerID sets the "owner" edge to the User entity by ID.
func (apu *AppPackageUpdate) SetOwnerID(id model.InternalID) *AppPackageUpdate {
	apu.mutation.SetOwnerID(id)
	return apu
}

// SetOwner sets the "owner" edge to the User entity.
func (apu *AppPackageUpdate) SetOwner(u *User) *AppPackageUpdate {
	return apu.SetOwnerID(u.ID)
}

// SetAppID sets the "app" edge to the App entity by ID.
func (apu *AppPackageUpdate) SetAppID(id model.InternalID) *AppPackageUpdate {
	apu.mutation.SetAppID(id)
	return apu
}

// SetNillableAppID sets the "app" edge to the App entity by ID if the given value is not nil.
func (apu *AppPackageUpdate) SetNillableAppID(id *model.InternalID) *AppPackageUpdate {
	if id != nil {
		apu = apu.SetAppID(*id)
	}
	return apu
}

// SetApp sets the "app" edge to the App entity.
func (apu *AppPackageUpdate) SetApp(a *App) *AppPackageUpdate {
	return apu.SetAppID(a.ID)
}

// Mutation returns the AppPackageMutation object of the builder.
func (apu *AppPackageUpdate) Mutation() *AppPackageMutation {
	return apu.mutation
}

// ClearOwner clears the "owner" edge to the User entity.
func (apu *AppPackageUpdate) ClearOwner() *AppPackageUpdate {
	apu.mutation.ClearOwner()
	return apu
}

// ClearApp clears the "app" edge to the App entity.
func (apu *AppPackageUpdate) ClearApp() *AppPackageUpdate {
	apu.mutation.ClearApp()
	return apu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (apu *AppPackageUpdate) Save(ctx context.Context) (int, error) {
	apu.defaults()
	return withHooks(ctx, apu.sqlSave, apu.mutation, apu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (apu *AppPackageUpdate) SaveX(ctx context.Context) int {
	affected, err := apu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (apu *AppPackageUpdate) Exec(ctx context.Context) error {
	_, err := apu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (apu *AppPackageUpdate) ExecX(ctx context.Context) {
	if err := apu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (apu *AppPackageUpdate) defaults() {
	if _, ok := apu.mutation.UpdatedAt(); !ok {
		v := apppackage.UpdateDefaultUpdatedAt()
		apu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (apu *AppPackageUpdate) check() error {
	if v, ok := apu.mutation.Source(); ok {
		if err := apppackage.SourceValidator(v); err != nil {
			return &ValidationError{Name: "source", err: fmt.Errorf(`ent: validator failed for field "AppPackage.source": %w`, err)}
		}
	}
	if _, ok := apu.mutation.OwnerID(); apu.mutation.OwnerCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "AppPackage.owner"`)
	}
	return nil
}

func (apu *AppPackageUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := apu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(apppackage.Table, apppackage.Columns, sqlgraph.NewFieldSpec(apppackage.FieldID, field.TypeInt64))
	if ps := apu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := apu.mutation.Source(); ok {
		_spec.SetField(apppackage.FieldSource, field.TypeEnum, value)
	}
	if value, ok := apu.mutation.SourceID(); ok {
		_spec.SetField(apppackage.FieldSourceID, field.TypeInt64, value)
	}
	if value, ok := apu.mutation.AddedSourceID(); ok {
		_spec.AddField(apppackage.FieldSourceID, field.TypeInt64, value)
	}
	if value, ok := apu.mutation.Name(); ok {
		_spec.SetField(apppackage.FieldName, field.TypeString, value)
	}
	if value, ok := apu.mutation.Description(); ok {
		_spec.SetField(apppackage.FieldDescription, field.TypeString, value)
	}
	if value, ok := apu.mutation.Public(); ok {
		_spec.SetField(apppackage.FieldPublic, field.TypeBool, value)
	}
	if value, ok := apu.mutation.BinaryName(); ok {
		_spec.SetField(apppackage.FieldBinaryName, field.TypeString, value)
	}
	if value, ok := apu.mutation.BinarySizeBytes(); ok {
		_spec.SetField(apppackage.FieldBinarySizeBytes, field.TypeInt64, value)
	}
	if value, ok := apu.mutation.AddedBinarySizeBytes(); ok {
		_spec.AddField(apppackage.FieldBinarySizeBytes, field.TypeInt64, value)
	}
	if value, ok := apu.mutation.BinaryPublicURL(); ok {
		_spec.SetField(apppackage.FieldBinaryPublicURL, field.TypeString, value)
	}
	if value, ok := apu.mutation.BinarySha256(); ok {
		_spec.SetField(apppackage.FieldBinarySha256, field.TypeBytes, value)
	}
	if value, ok := apu.mutation.UpdatedAt(); ok {
		_spec.SetField(apppackage.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := apu.mutation.CreatedAt(); ok {
		_spec.SetField(apppackage.FieldCreatedAt, field.TypeTime, value)
	}
	if apu.mutation.OwnerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   apppackage.OwnerTable,
			Columns: []string{apppackage.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := apu.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   apppackage.OwnerTable,
			Columns: []string{apppackage.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if apu.mutation.AppCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   apppackage.AppTable,
			Columns: []string{apppackage.AppColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(app.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := apu.mutation.AppIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   apppackage.AppTable,
			Columns: []string{apppackage.AppColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(app.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, apu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{apppackage.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	apu.mutation.done = true
	return n, nil
}

// AppPackageUpdateOne is the builder for updating a single AppPackage entity.
type AppPackageUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *AppPackageMutation
}

// SetSource sets the "source" field.
func (apuo *AppPackageUpdateOne) SetSource(a apppackage.Source) *AppPackageUpdateOne {
	apuo.mutation.SetSource(a)
	return apuo
}

// SetSourceID sets the "source_id" field.
func (apuo *AppPackageUpdateOne) SetSourceID(mi model.InternalID) *AppPackageUpdateOne {
	apuo.mutation.ResetSourceID()
	apuo.mutation.SetSourceID(mi)
	return apuo
}

// AddSourceID adds mi to the "source_id" field.
func (apuo *AppPackageUpdateOne) AddSourceID(mi model.InternalID) *AppPackageUpdateOne {
	apuo.mutation.AddSourceID(mi)
	return apuo
}

// SetName sets the "name" field.
func (apuo *AppPackageUpdateOne) SetName(s string) *AppPackageUpdateOne {
	apuo.mutation.SetName(s)
	return apuo
}

// SetDescription sets the "description" field.
func (apuo *AppPackageUpdateOne) SetDescription(s string) *AppPackageUpdateOne {
	apuo.mutation.SetDescription(s)
	return apuo
}

// SetPublic sets the "public" field.
func (apuo *AppPackageUpdateOne) SetPublic(b bool) *AppPackageUpdateOne {
	apuo.mutation.SetPublic(b)
	return apuo
}

// SetBinaryName sets the "binary_name" field.
func (apuo *AppPackageUpdateOne) SetBinaryName(s string) *AppPackageUpdateOne {
	apuo.mutation.SetBinaryName(s)
	return apuo
}

// SetBinarySizeBytes sets the "binary_size_bytes" field.
func (apuo *AppPackageUpdateOne) SetBinarySizeBytes(i int64) *AppPackageUpdateOne {
	apuo.mutation.ResetBinarySizeBytes()
	apuo.mutation.SetBinarySizeBytes(i)
	return apuo
}

// AddBinarySizeBytes adds i to the "binary_size_bytes" field.
func (apuo *AppPackageUpdateOne) AddBinarySizeBytes(i int64) *AppPackageUpdateOne {
	apuo.mutation.AddBinarySizeBytes(i)
	return apuo
}

// SetBinaryPublicURL sets the "binary_public_url" field.
func (apuo *AppPackageUpdateOne) SetBinaryPublicURL(s string) *AppPackageUpdateOne {
	apuo.mutation.SetBinaryPublicURL(s)
	return apuo
}

// SetBinarySha256 sets the "binary_sha256" field.
func (apuo *AppPackageUpdateOne) SetBinarySha256(b []byte) *AppPackageUpdateOne {
	apuo.mutation.SetBinarySha256(b)
	return apuo
}

// SetUpdatedAt sets the "updated_at" field.
func (apuo *AppPackageUpdateOne) SetUpdatedAt(t time.Time) *AppPackageUpdateOne {
	apuo.mutation.SetUpdatedAt(t)
	return apuo
}

// SetCreatedAt sets the "created_at" field.
func (apuo *AppPackageUpdateOne) SetCreatedAt(t time.Time) *AppPackageUpdateOne {
	apuo.mutation.SetCreatedAt(t)
	return apuo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (apuo *AppPackageUpdateOne) SetNillableCreatedAt(t *time.Time) *AppPackageUpdateOne {
	if t != nil {
		apuo.SetCreatedAt(*t)
	}
	return apuo
}

// SetOwnerID sets the "owner" edge to the User entity by ID.
func (apuo *AppPackageUpdateOne) SetOwnerID(id model.InternalID) *AppPackageUpdateOne {
	apuo.mutation.SetOwnerID(id)
	return apuo
}

// SetOwner sets the "owner" edge to the User entity.
func (apuo *AppPackageUpdateOne) SetOwner(u *User) *AppPackageUpdateOne {
	return apuo.SetOwnerID(u.ID)
}

// SetAppID sets the "app" edge to the App entity by ID.
func (apuo *AppPackageUpdateOne) SetAppID(id model.InternalID) *AppPackageUpdateOne {
	apuo.mutation.SetAppID(id)
	return apuo
}

// SetNillableAppID sets the "app" edge to the App entity by ID if the given value is not nil.
func (apuo *AppPackageUpdateOne) SetNillableAppID(id *model.InternalID) *AppPackageUpdateOne {
	if id != nil {
		apuo = apuo.SetAppID(*id)
	}
	return apuo
}

// SetApp sets the "app" edge to the App entity.
func (apuo *AppPackageUpdateOne) SetApp(a *App) *AppPackageUpdateOne {
	return apuo.SetAppID(a.ID)
}

// Mutation returns the AppPackageMutation object of the builder.
func (apuo *AppPackageUpdateOne) Mutation() *AppPackageMutation {
	return apuo.mutation
}

// ClearOwner clears the "owner" edge to the User entity.
func (apuo *AppPackageUpdateOne) ClearOwner() *AppPackageUpdateOne {
	apuo.mutation.ClearOwner()
	return apuo
}

// ClearApp clears the "app" edge to the App entity.
func (apuo *AppPackageUpdateOne) ClearApp() *AppPackageUpdateOne {
	apuo.mutation.ClearApp()
	return apuo
}

// Where appends a list predicates to the AppPackageUpdate builder.
func (apuo *AppPackageUpdateOne) Where(ps ...predicate.AppPackage) *AppPackageUpdateOne {
	apuo.mutation.Where(ps...)
	return apuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (apuo *AppPackageUpdateOne) Select(field string, fields ...string) *AppPackageUpdateOne {
	apuo.fields = append([]string{field}, fields...)
	return apuo
}

// Save executes the query and returns the updated AppPackage entity.
func (apuo *AppPackageUpdateOne) Save(ctx context.Context) (*AppPackage, error) {
	apuo.defaults()
	return withHooks(ctx, apuo.sqlSave, apuo.mutation, apuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (apuo *AppPackageUpdateOne) SaveX(ctx context.Context) *AppPackage {
	node, err := apuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (apuo *AppPackageUpdateOne) Exec(ctx context.Context) error {
	_, err := apuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (apuo *AppPackageUpdateOne) ExecX(ctx context.Context) {
	if err := apuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (apuo *AppPackageUpdateOne) defaults() {
	if _, ok := apuo.mutation.UpdatedAt(); !ok {
		v := apppackage.UpdateDefaultUpdatedAt()
		apuo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (apuo *AppPackageUpdateOne) check() error {
	if v, ok := apuo.mutation.Source(); ok {
		if err := apppackage.SourceValidator(v); err != nil {
			return &ValidationError{Name: "source", err: fmt.Errorf(`ent: validator failed for field "AppPackage.source": %w`, err)}
		}
	}
	if _, ok := apuo.mutation.OwnerID(); apuo.mutation.OwnerCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "AppPackage.owner"`)
	}
	return nil
}

func (apuo *AppPackageUpdateOne) sqlSave(ctx context.Context) (_node *AppPackage, err error) {
	if err := apuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(apppackage.Table, apppackage.Columns, sqlgraph.NewFieldSpec(apppackage.FieldID, field.TypeInt64))
	id, ok := apuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "AppPackage.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := apuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, apppackage.FieldID)
		for _, f := range fields {
			if !apppackage.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != apppackage.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := apuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := apuo.mutation.Source(); ok {
		_spec.SetField(apppackage.FieldSource, field.TypeEnum, value)
	}
	if value, ok := apuo.mutation.SourceID(); ok {
		_spec.SetField(apppackage.FieldSourceID, field.TypeInt64, value)
	}
	if value, ok := apuo.mutation.AddedSourceID(); ok {
		_spec.AddField(apppackage.FieldSourceID, field.TypeInt64, value)
	}
	if value, ok := apuo.mutation.Name(); ok {
		_spec.SetField(apppackage.FieldName, field.TypeString, value)
	}
	if value, ok := apuo.mutation.Description(); ok {
		_spec.SetField(apppackage.FieldDescription, field.TypeString, value)
	}
	if value, ok := apuo.mutation.Public(); ok {
		_spec.SetField(apppackage.FieldPublic, field.TypeBool, value)
	}
	if value, ok := apuo.mutation.BinaryName(); ok {
		_spec.SetField(apppackage.FieldBinaryName, field.TypeString, value)
	}
	if value, ok := apuo.mutation.BinarySizeBytes(); ok {
		_spec.SetField(apppackage.FieldBinarySizeBytes, field.TypeInt64, value)
	}
	if value, ok := apuo.mutation.AddedBinarySizeBytes(); ok {
		_spec.AddField(apppackage.FieldBinarySizeBytes, field.TypeInt64, value)
	}
	if value, ok := apuo.mutation.BinaryPublicURL(); ok {
		_spec.SetField(apppackage.FieldBinaryPublicURL, field.TypeString, value)
	}
	if value, ok := apuo.mutation.BinarySha256(); ok {
		_spec.SetField(apppackage.FieldBinarySha256, field.TypeBytes, value)
	}
	if value, ok := apuo.mutation.UpdatedAt(); ok {
		_spec.SetField(apppackage.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := apuo.mutation.CreatedAt(); ok {
		_spec.SetField(apppackage.FieldCreatedAt, field.TypeTime, value)
	}
	if apuo.mutation.OwnerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   apppackage.OwnerTable,
			Columns: []string{apppackage.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := apuo.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   apppackage.OwnerTable,
			Columns: []string{apppackage.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if apuo.mutation.AppCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   apppackage.AppTable,
			Columns: []string{apppackage.AppColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(app.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := apuo.mutation.AppIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   apppackage.AppTable,
			Columns: []string{apppackage.AppColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(app.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &AppPackage{config: apuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, apuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{apppackage.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	apuo.mutation.done = true
	return _node, nil
}
