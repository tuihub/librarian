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
	"github.com/tuihub/librarian/internal/data/internal/ent/predicate"
	"github.com/tuihub/librarian/internal/data/internal/ent/sentinelappbinary"
	"github.com/tuihub/librarian/internal/data/internal/ent/sentinelappbinaryfile"
	"github.com/tuihub/librarian/internal/model"
)

// SentinelAppBinaryFileUpdate is the builder for updating SentinelAppBinaryFile entities.
type SentinelAppBinaryFileUpdate struct {
	config
	hooks    []Hook
	mutation *SentinelAppBinaryFileMutation
}

// Where appends a list predicates to the SentinelAppBinaryFileUpdate builder.
func (sabfu *SentinelAppBinaryFileUpdate) Where(ps ...predicate.SentinelAppBinaryFile) *SentinelAppBinaryFileUpdate {
	sabfu.mutation.Where(ps...)
	return sabfu
}

// SetSentinelAppBinaryID sets the "sentinel_app_binary_id" field.
func (sabfu *SentinelAppBinaryFileUpdate) SetSentinelAppBinaryID(mi model.InternalID) *SentinelAppBinaryFileUpdate {
	sabfu.mutation.SetSentinelAppBinaryID(mi)
	return sabfu
}

// SetNillableSentinelAppBinaryID sets the "sentinel_app_binary_id" field if the given value is not nil.
func (sabfu *SentinelAppBinaryFileUpdate) SetNillableSentinelAppBinaryID(mi *model.InternalID) *SentinelAppBinaryFileUpdate {
	if mi != nil {
		sabfu.SetSentinelAppBinaryID(*mi)
	}
	return sabfu
}

// SetName sets the "name" field.
func (sabfu *SentinelAppBinaryFileUpdate) SetName(s string) *SentinelAppBinaryFileUpdate {
	sabfu.mutation.SetName(s)
	return sabfu
}

// SetNillableName sets the "name" field if the given value is not nil.
func (sabfu *SentinelAppBinaryFileUpdate) SetNillableName(s *string) *SentinelAppBinaryFileUpdate {
	if s != nil {
		sabfu.SetName(*s)
	}
	return sabfu
}

// SetSizeBytes sets the "size_bytes" field.
func (sabfu *SentinelAppBinaryFileUpdate) SetSizeBytes(i int64) *SentinelAppBinaryFileUpdate {
	sabfu.mutation.ResetSizeBytes()
	sabfu.mutation.SetSizeBytes(i)
	return sabfu
}

// SetNillableSizeBytes sets the "size_bytes" field if the given value is not nil.
func (sabfu *SentinelAppBinaryFileUpdate) SetNillableSizeBytes(i *int64) *SentinelAppBinaryFileUpdate {
	if i != nil {
		sabfu.SetSizeBytes(*i)
	}
	return sabfu
}

// AddSizeBytes adds i to the "size_bytes" field.
func (sabfu *SentinelAppBinaryFileUpdate) AddSizeBytes(i int64) *SentinelAppBinaryFileUpdate {
	sabfu.mutation.AddSizeBytes(i)
	return sabfu
}

// SetSha256 sets the "sha256" field.
func (sabfu *SentinelAppBinaryFileUpdate) SetSha256(b []byte) *SentinelAppBinaryFileUpdate {
	sabfu.mutation.SetSha256(b)
	return sabfu
}

// SetServerFilePath sets the "server_file_path" field.
func (sabfu *SentinelAppBinaryFileUpdate) SetServerFilePath(s string) *SentinelAppBinaryFileUpdate {
	sabfu.mutation.SetServerFilePath(s)
	return sabfu
}

// SetNillableServerFilePath sets the "server_file_path" field if the given value is not nil.
func (sabfu *SentinelAppBinaryFileUpdate) SetNillableServerFilePath(s *string) *SentinelAppBinaryFileUpdate {
	if s != nil {
		sabfu.SetServerFilePath(*s)
	}
	return sabfu
}

// SetChunksInfo sets the "chunks_info" field.
func (sabfu *SentinelAppBinaryFileUpdate) SetChunksInfo(s string) *SentinelAppBinaryFileUpdate {
	sabfu.mutation.SetChunksInfo(s)
	return sabfu
}

// SetNillableChunksInfo sets the "chunks_info" field if the given value is not nil.
func (sabfu *SentinelAppBinaryFileUpdate) SetNillableChunksInfo(s *string) *SentinelAppBinaryFileUpdate {
	if s != nil {
		sabfu.SetChunksInfo(*s)
	}
	return sabfu
}

// ClearChunksInfo clears the value of the "chunks_info" field.
func (sabfu *SentinelAppBinaryFileUpdate) ClearChunksInfo() *SentinelAppBinaryFileUpdate {
	sabfu.mutation.ClearChunksInfo()
	return sabfu
}

// SetUpdatedAt sets the "updated_at" field.
func (sabfu *SentinelAppBinaryFileUpdate) SetUpdatedAt(t time.Time) *SentinelAppBinaryFileUpdate {
	sabfu.mutation.SetUpdatedAt(t)
	return sabfu
}

// SetCreatedAt sets the "created_at" field.
func (sabfu *SentinelAppBinaryFileUpdate) SetCreatedAt(t time.Time) *SentinelAppBinaryFileUpdate {
	sabfu.mutation.SetCreatedAt(t)
	return sabfu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (sabfu *SentinelAppBinaryFileUpdate) SetNillableCreatedAt(t *time.Time) *SentinelAppBinaryFileUpdate {
	if t != nil {
		sabfu.SetCreatedAt(*t)
	}
	return sabfu
}

// SetSentinelAppBinary sets the "sentinel_app_binary" edge to the SentinelAppBinary entity.
func (sabfu *SentinelAppBinaryFileUpdate) SetSentinelAppBinary(s *SentinelAppBinary) *SentinelAppBinaryFileUpdate {
	return sabfu.SetSentinelAppBinaryID(s.ID)
}

// Mutation returns the SentinelAppBinaryFileMutation object of the builder.
func (sabfu *SentinelAppBinaryFileUpdate) Mutation() *SentinelAppBinaryFileMutation {
	return sabfu.mutation
}

// ClearSentinelAppBinary clears the "sentinel_app_binary" edge to the SentinelAppBinary entity.
func (sabfu *SentinelAppBinaryFileUpdate) ClearSentinelAppBinary() *SentinelAppBinaryFileUpdate {
	sabfu.mutation.ClearSentinelAppBinary()
	return sabfu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (sabfu *SentinelAppBinaryFileUpdate) Save(ctx context.Context) (int, error) {
	sabfu.defaults()
	return withHooks(ctx, sabfu.sqlSave, sabfu.mutation, sabfu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (sabfu *SentinelAppBinaryFileUpdate) SaveX(ctx context.Context) int {
	affected, err := sabfu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (sabfu *SentinelAppBinaryFileUpdate) Exec(ctx context.Context) error {
	_, err := sabfu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sabfu *SentinelAppBinaryFileUpdate) ExecX(ctx context.Context) {
	if err := sabfu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (sabfu *SentinelAppBinaryFileUpdate) defaults() {
	if _, ok := sabfu.mutation.UpdatedAt(); !ok {
		v := sentinelappbinaryfile.UpdateDefaultUpdatedAt()
		sabfu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (sabfu *SentinelAppBinaryFileUpdate) check() error {
	if sabfu.mutation.SentinelAppBinaryCleared() && len(sabfu.mutation.SentinelAppBinaryIDs()) > 0 {
		return errors.New(`ent: clearing a required unique edge "SentinelAppBinaryFile.sentinel_app_binary"`)
	}
	return nil
}

func (sabfu *SentinelAppBinaryFileUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := sabfu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(sentinelappbinaryfile.Table, sentinelappbinaryfile.Columns, sqlgraph.NewFieldSpec(sentinelappbinaryfile.FieldID, field.TypeInt64))
	if ps := sabfu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := sabfu.mutation.Name(); ok {
		_spec.SetField(sentinelappbinaryfile.FieldName, field.TypeString, value)
	}
	if value, ok := sabfu.mutation.SizeBytes(); ok {
		_spec.SetField(sentinelappbinaryfile.FieldSizeBytes, field.TypeInt64, value)
	}
	if value, ok := sabfu.mutation.AddedSizeBytes(); ok {
		_spec.AddField(sentinelappbinaryfile.FieldSizeBytes, field.TypeInt64, value)
	}
	if value, ok := sabfu.mutation.Sha256(); ok {
		_spec.SetField(sentinelappbinaryfile.FieldSha256, field.TypeBytes, value)
	}
	if value, ok := sabfu.mutation.ServerFilePath(); ok {
		_spec.SetField(sentinelappbinaryfile.FieldServerFilePath, field.TypeString, value)
	}
	if value, ok := sabfu.mutation.ChunksInfo(); ok {
		_spec.SetField(sentinelappbinaryfile.FieldChunksInfo, field.TypeString, value)
	}
	if sabfu.mutation.ChunksInfoCleared() {
		_spec.ClearField(sentinelappbinaryfile.FieldChunksInfo, field.TypeString)
	}
	if value, ok := sabfu.mutation.UpdatedAt(); ok {
		_spec.SetField(sentinelappbinaryfile.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := sabfu.mutation.CreatedAt(); ok {
		_spec.SetField(sentinelappbinaryfile.FieldCreatedAt, field.TypeTime, value)
	}
	if sabfu.mutation.SentinelAppBinaryCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   sentinelappbinaryfile.SentinelAppBinaryTable,
			Columns: []string{sentinelappbinaryfile.SentinelAppBinaryColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(sentinelappbinary.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := sabfu.mutation.SentinelAppBinaryIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   sentinelappbinaryfile.SentinelAppBinaryTable,
			Columns: []string{sentinelappbinaryfile.SentinelAppBinaryColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(sentinelappbinary.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, sabfu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{sentinelappbinaryfile.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	sabfu.mutation.done = true
	return n, nil
}

// SentinelAppBinaryFileUpdateOne is the builder for updating a single SentinelAppBinaryFile entity.
type SentinelAppBinaryFileUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *SentinelAppBinaryFileMutation
}

// SetSentinelAppBinaryID sets the "sentinel_app_binary_id" field.
func (sabfuo *SentinelAppBinaryFileUpdateOne) SetSentinelAppBinaryID(mi model.InternalID) *SentinelAppBinaryFileUpdateOne {
	sabfuo.mutation.SetSentinelAppBinaryID(mi)
	return sabfuo
}

// SetNillableSentinelAppBinaryID sets the "sentinel_app_binary_id" field if the given value is not nil.
func (sabfuo *SentinelAppBinaryFileUpdateOne) SetNillableSentinelAppBinaryID(mi *model.InternalID) *SentinelAppBinaryFileUpdateOne {
	if mi != nil {
		sabfuo.SetSentinelAppBinaryID(*mi)
	}
	return sabfuo
}

// SetName sets the "name" field.
func (sabfuo *SentinelAppBinaryFileUpdateOne) SetName(s string) *SentinelAppBinaryFileUpdateOne {
	sabfuo.mutation.SetName(s)
	return sabfuo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (sabfuo *SentinelAppBinaryFileUpdateOne) SetNillableName(s *string) *SentinelAppBinaryFileUpdateOne {
	if s != nil {
		sabfuo.SetName(*s)
	}
	return sabfuo
}

// SetSizeBytes sets the "size_bytes" field.
func (sabfuo *SentinelAppBinaryFileUpdateOne) SetSizeBytes(i int64) *SentinelAppBinaryFileUpdateOne {
	sabfuo.mutation.ResetSizeBytes()
	sabfuo.mutation.SetSizeBytes(i)
	return sabfuo
}

// SetNillableSizeBytes sets the "size_bytes" field if the given value is not nil.
func (sabfuo *SentinelAppBinaryFileUpdateOne) SetNillableSizeBytes(i *int64) *SentinelAppBinaryFileUpdateOne {
	if i != nil {
		sabfuo.SetSizeBytes(*i)
	}
	return sabfuo
}

// AddSizeBytes adds i to the "size_bytes" field.
func (sabfuo *SentinelAppBinaryFileUpdateOne) AddSizeBytes(i int64) *SentinelAppBinaryFileUpdateOne {
	sabfuo.mutation.AddSizeBytes(i)
	return sabfuo
}

// SetSha256 sets the "sha256" field.
func (sabfuo *SentinelAppBinaryFileUpdateOne) SetSha256(b []byte) *SentinelAppBinaryFileUpdateOne {
	sabfuo.mutation.SetSha256(b)
	return sabfuo
}

// SetServerFilePath sets the "server_file_path" field.
func (sabfuo *SentinelAppBinaryFileUpdateOne) SetServerFilePath(s string) *SentinelAppBinaryFileUpdateOne {
	sabfuo.mutation.SetServerFilePath(s)
	return sabfuo
}

// SetNillableServerFilePath sets the "server_file_path" field if the given value is not nil.
func (sabfuo *SentinelAppBinaryFileUpdateOne) SetNillableServerFilePath(s *string) *SentinelAppBinaryFileUpdateOne {
	if s != nil {
		sabfuo.SetServerFilePath(*s)
	}
	return sabfuo
}

// SetChunksInfo sets the "chunks_info" field.
func (sabfuo *SentinelAppBinaryFileUpdateOne) SetChunksInfo(s string) *SentinelAppBinaryFileUpdateOne {
	sabfuo.mutation.SetChunksInfo(s)
	return sabfuo
}

// SetNillableChunksInfo sets the "chunks_info" field if the given value is not nil.
func (sabfuo *SentinelAppBinaryFileUpdateOne) SetNillableChunksInfo(s *string) *SentinelAppBinaryFileUpdateOne {
	if s != nil {
		sabfuo.SetChunksInfo(*s)
	}
	return sabfuo
}

// ClearChunksInfo clears the value of the "chunks_info" field.
func (sabfuo *SentinelAppBinaryFileUpdateOne) ClearChunksInfo() *SentinelAppBinaryFileUpdateOne {
	sabfuo.mutation.ClearChunksInfo()
	return sabfuo
}

// SetUpdatedAt sets the "updated_at" field.
func (sabfuo *SentinelAppBinaryFileUpdateOne) SetUpdatedAt(t time.Time) *SentinelAppBinaryFileUpdateOne {
	sabfuo.mutation.SetUpdatedAt(t)
	return sabfuo
}

// SetCreatedAt sets the "created_at" field.
func (sabfuo *SentinelAppBinaryFileUpdateOne) SetCreatedAt(t time.Time) *SentinelAppBinaryFileUpdateOne {
	sabfuo.mutation.SetCreatedAt(t)
	return sabfuo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (sabfuo *SentinelAppBinaryFileUpdateOne) SetNillableCreatedAt(t *time.Time) *SentinelAppBinaryFileUpdateOne {
	if t != nil {
		sabfuo.SetCreatedAt(*t)
	}
	return sabfuo
}

// SetSentinelAppBinary sets the "sentinel_app_binary" edge to the SentinelAppBinary entity.
func (sabfuo *SentinelAppBinaryFileUpdateOne) SetSentinelAppBinary(s *SentinelAppBinary) *SentinelAppBinaryFileUpdateOne {
	return sabfuo.SetSentinelAppBinaryID(s.ID)
}

// Mutation returns the SentinelAppBinaryFileMutation object of the builder.
func (sabfuo *SentinelAppBinaryFileUpdateOne) Mutation() *SentinelAppBinaryFileMutation {
	return sabfuo.mutation
}

// ClearSentinelAppBinary clears the "sentinel_app_binary" edge to the SentinelAppBinary entity.
func (sabfuo *SentinelAppBinaryFileUpdateOne) ClearSentinelAppBinary() *SentinelAppBinaryFileUpdateOne {
	sabfuo.mutation.ClearSentinelAppBinary()
	return sabfuo
}

// Where appends a list predicates to the SentinelAppBinaryFileUpdate builder.
func (sabfuo *SentinelAppBinaryFileUpdateOne) Where(ps ...predicate.SentinelAppBinaryFile) *SentinelAppBinaryFileUpdateOne {
	sabfuo.mutation.Where(ps...)
	return sabfuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (sabfuo *SentinelAppBinaryFileUpdateOne) Select(field string, fields ...string) *SentinelAppBinaryFileUpdateOne {
	sabfuo.fields = append([]string{field}, fields...)
	return sabfuo
}

// Save executes the query and returns the updated SentinelAppBinaryFile entity.
func (sabfuo *SentinelAppBinaryFileUpdateOne) Save(ctx context.Context) (*SentinelAppBinaryFile, error) {
	sabfuo.defaults()
	return withHooks(ctx, sabfuo.sqlSave, sabfuo.mutation, sabfuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (sabfuo *SentinelAppBinaryFileUpdateOne) SaveX(ctx context.Context) *SentinelAppBinaryFile {
	node, err := sabfuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (sabfuo *SentinelAppBinaryFileUpdateOne) Exec(ctx context.Context) error {
	_, err := sabfuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sabfuo *SentinelAppBinaryFileUpdateOne) ExecX(ctx context.Context) {
	if err := sabfuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (sabfuo *SentinelAppBinaryFileUpdateOne) defaults() {
	if _, ok := sabfuo.mutation.UpdatedAt(); !ok {
		v := sentinelappbinaryfile.UpdateDefaultUpdatedAt()
		sabfuo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (sabfuo *SentinelAppBinaryFileUpdateOne) check() error {
	if sabfuo.mutation.SentinelAppBinaryCleared() && len(sabfuo.mutation.SentinelAppBinaryIDs()) > 0 {
		return errors.New(`ent: clearing a required unique edge "SentinelAppBinaryFile.sentinel_app_binary"`)
	}
	return nil
}

func (sabfuo *SentinelAppBinaryFileUpdateOne) sqlSave(ctx context.Context) (_node *SentinelAppBinaryFile, err error) {
	if err := sabfuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(sentinelappbinaryfile.Table, sentinelappbinaryfile.Columns, sqlgraph.NewFieldSpec(sentinelappbinaryfile.FieldID, field.TypeInt64))
	id, ok := sabfuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "SentinelAppBinaryFile.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := sabfuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, sentinelappbinaryfile.FieldID)
		for _, f := range fields {
			if !sentinelappbinaryfile.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != sentinelappbinaryfile.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := sabfuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := sabfuo.mutation.Name(); ok {
		_spec.SetField(sentinelappbinaryfile.FieldName, field.TypeString, value)
	}
	if value, ok := sabfuo.mutation.SizeBytes(); ok {
		_spec.SetField(sentinelappbinaryfile.FieldSizeBytes, field.TypeInt64, value)
	}
	if value, ok := sabfuo.mutation.AddedSizeBytes(); ok {
		_spec.AddField(sentinelappbinaryfile.FieldSizeBytes, field.TypeInt64, value)
	}
	if value, ok := sabfuo.mutation.Sha256(); ok {
		_spec.SetField(sentinelappbinaryfile.FieldSha256, field.TypeBytes, value)
	}
	if value, ok := sabfuo.mutation.ServerFilePath(); ok {
		_spec.SetField(sentinelappbinaryfile.FieldServerFilePath, field.TypeString, value)
	}
	if value, ok := sabfuo.mutation.ChunksInfo(); ok {
		_spec.SetField(sentinelappbinaryfile.FieldChunksInfo, field.TypeString, value)
	}
	if sabfuo.mutation.ChunksInfoCleared() {
		_spec.ClearField(sentinelappbinaryfile.FieldChunksInfo, field.TypeString)
	}
	if value, ok := sabfuo.mutation.UpdatedAt(); ok {
		_spec.SetField(sentinelappbinaryfile.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := sabfuo.mutation.CreatedAt(); ok {
		_spec.SetField(sentinelappbinaryfile.FieldCreatedAt, field.TypeTime, value)
	}
	if sabfuo.mutation.SentinelAppBinaryCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   sentinelappbinaryfile.SentinelAppBinaryTable,
			Columns: []string{sentinelappbinaryfile.SentinelAppBinaryColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(sentinelappbinary.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := sabfuo.mutation.SentinelAppBinaryIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   sentinelappbinaryfile.SentinelAppBinaryTable,
			Columns: []string{sentinelappbinaryfile.SentinelAppBinaryColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(sentinelappbinary.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &SentinelAppBinaryFile{config: sabfuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, sabfuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{sentinelappbinaryfile.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	sabfuo.mutation.done = true
	return _node, nil
}
