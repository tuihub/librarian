// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/dialect/sql/sqljson"
	"entgo.io/ent/schema/field"
	"github.com/tuihub/librarian/internal/data/internal/ent/appinfo"
	"github.com/tuihub/librarian/internal/data/internal/ent/predicate"
	"github.com/tuihub/librarian/internal/model"
)

// AppInfoUpdate is the builder for updating AppInfo entities.
type AppInfoUpdate struct {
	config
	hooks    []Hook
	mutation *AppInfoMutation
}

// Where appends a list predicates to the AppInfoUpdate builder.
func (aiu *AppInfoUpdate) Where(ps ...predicate.AppInfo) *AppInfoUpdate {
	aiu.mutation.Where(ps...)
	return aiu
}

// SetSource sets the "source" field.
func (aiu *AppInfoUpdate) SetSource(s string) *AppInfoUpdate {
	aiu.mutation.SetSource(s)
	return aiu
}

// SetNillableSource sets the "source" field if the given value is not nil.
func (aiu *AppInfoUpdate) SetNillableSource(s *string) *AppInfoUpdate {
	if s != nil {
		aiu.SetSource(*s)
	}
	return aiu
}

// SetSourceAppID sets the "source_app_id" field.
func (aiu *AppInfoUpdate) SetSourceAppID(s string) *AppInfoUpdate {
	aiu.mutation.SetSourceAppID(s)
	return aiu
}

// SetNillableSourceAppID sets the "source_app_id" field if the given value is not nil.
func (aiu *AppInfoUpdate) SetNillableSourceAppID(s *string) *AppInfoUpdate {
	if s != nil {
		aiu.SetSourceAppID(*s)
	}
	return aiu
}

// SetSourceURL sets the "source_url" field.
func (aiu *AppInfoUpdate) SetSourceURL(s string) *AppInfoUpdate {
	aiu.mutation.SetSourceURL(s)
	return aiu
}

// SetNillableSourceURL sets the "source_url" field if the given value is not nil.
func (aiu *AppInfoUpdate) SetNillableSourceURL(s *string) *AppInfoUpdate {
	if s != nil {
		aiu.SetSourceURL(*s)
	}
	return aiu
}

// ClearSourceURL clears the value of the "source_url" field.
func (aiu *AppInfoUpdate) ClearSourceURL() *AppInfoUpdate {
	aiu.mutation.ClearSourceURL()
	return aiu
}

// SetName sets the "name" field.
func (aiu *AppInfoUpdate) SetName(s string) *AppInfoUpdate {
	aiu.mutation.SetName(s)
	return aiu
}

// SetNillableName sets the "name" field if the given value is not nil.
func (aiu *AppInfoUpdate) SetNillableName(s *string) *AppInfoUpdate {
	if s != nil {
		aiu.SetName(*s)
	}
	return aiu
}

// SetType sets the "type" field.
func (aiu *AppInfoUpdate) SetType(a appinfo.Type) *AppInfoUpdate {
	aiu.mutation.SetType(a)
	return aiu
}

// SetNillableType sets the "type" field if the given value is not nil.
func (aiu *AppInfoUpdate) SetNillableType(a *appinfo.Type) *AppInfoUpdate {
	if a != nil {
		aiu.SetType(*a)
	}
	return aiu
}

// SetShortDescription sets the "short_description" field.
func (aiu *AppInfoUpdate) SetShortDescription(s string) *AppInfoUpdate {
	aiu.mutation.SetShortDescription(s)
	return aiu
}

// SetNillableShortDescription sets the "short_description" field if the given value is not nil.
func (aiu *AppInfoUpdate) SetNillableShortDescription(s *string) *AppInfoUpdate {
	if s != nil {
		aiu.SetShortDescription(*s)
	}
	return aiu
}

// ClearShortDescription clears the value of the "short_description" field.
func (aiu *AppInfoUpdate) ClearShortDescription() *AppInfoUpdate {
	aiu.mutation.ClearShortDescription()
	return aiu
}

// SetDescription sets the "description" field.
func (aiu *AppInfoUpdate) SetDescription(s string) *AppInfoUpdate {
	aiu.mutation.SetDescription(s)
	return aiu
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (aiu *AppInfoUpdate) SetNillableDescription(s *string) *AppInfoUpdate {
	if s != nil {
		aiu.SetDescription(*s)
	}
	return aiu
}

// ClearDescription clears the value of the "description" field.
func (aiu *AppInfoUpdate) ClearDescription() *AppInfoUpdate {
	aiu.mutation.ClearDescription()
	return aiu
}

// SetIconImageURL sets the "icon_image_url" field.
func (aiu *AppInfoUpdate) SetIconImageURL(s string) *AppInfoUpdate {
	aiu.mutation.SetIconImageURL(s)
	return aiu
}

// SetNillableIconImageURL sets the "icon_image_url" field if the given value is not nil.
func (aiu *AppInfoUpdate) SetNillableIconImageURL(s *string) *AppInfoUpdate {
	if s != nil {
		aiu.SetIconImageURL(*s)
	}
	return aiu
}

// ClearIconImageURL clears the value of the "icon_image_url" field.
func (aiu *AppInfoUpdate) ClearIconImageURL() *AppInfoUpdate {
	aiu.mutation.ClearIconImageURL()
	return aiu
}

// SetIconImageID sets the "icon_image_id" field.
func (aiu *AppInfoUpdate) SetIconImageID(mi model.InternalID) *AppInfoUpdate {
	aiu.mutation.ResetIconImageID()
	aiu.mutation.SetIconImageID(mi)
	return aiu
}

// SetNillableIconImageID sets the "icon_image_id" field if the given value is not nil.
func (aiu *AppInfoUpdate) SetNillableIconImageID(mi *model.InternalID) *AppInfoUpdate {
	if mi != nil {
		aiu.SetIconImageID(*mi)
	}
	return aiu
}

// AddIconImageID adds mi to the "icon_image_id" field.
func (aiu *AppInfoUpdate) AddIconImageID(mi model.InternalID) *AppInfoUpdate {
	aiu.mutation.AddIconImageID(mi)
	return aiu
}

// SetBackgroundImageURL sets the "background_image_url" field.
func (aiu *AppInfoUpdate) SetBackgroundImageURL(s string) *AppInfoUpdate {
	aiu.mutation.SetBackgroundImageURL(s)
	return aiu
}

// SetNillableBackgroundImageURL sets the "background_image_url" field if the given value is not nil.
func (aiu *AppInfoUpdate) SetNillableBackgroundImageURL(s *string) *AppInfoUpdate {
	if s != nil {
		aiu.SetBackgroundImageURL(*s)
	}
	return aiu
}

// ClearBackgroundImageURL clears the value of the "background_image_url" field.
func (aiu *AppInfoUpdate) ClearBackgroundImageURL() *AppInfoUpdate {
	aiu.mutation.ClearBackgroundImageURL()
	return aiu
}

// SetBackgroundImageID sets the "background_image_id" field.
func (aiu *AppInfoUpdate) SetBackgroundImageID(mi model.InternalID) *AppInfoUpdate {
	aiu.mutation.ResetBackgroundImageID()
	aiu.mutation.SetBackgroundImageID(mi)
	return aiu
}

// SetNillableBackgroundImageID sets the "background_image_id" field if the given value is not nil.
func (aiu *AppInfoUpdate) SetNillableBackgroundImageID(mi *model.InternalID) *AppInfoUpdate {
	if mi != nil {
		aiu.SetBackgroundImageID(*mi)
	}
	return aiu
}

// AddBackgroundImageID adds mi to the "background_image_id" field.
func (aiu *AppInfoUpdate) AddBackgroundImageID(mi model.InternalID) *AppInfoUpdate {
	aiu.mutation.AddBackgroundImageID(mi)
	return aiu
}

// SetCoverImageURL sets the "cover_image_url" field.
func (aiu *AppInfoUpdate) SetCoverImageURL(s string) *AppInfoUpdate {
	aiu.mutation.SetCoverImageURL(s)
	return aiu
}

// SetNillableCoverImageURL sets the "cover_image_url" field if the given value is not nil.
func (aiu *AppInfoUpdate) SetNillableCoverImageURL(s *string) *AppInfoUpdate {
	if s != nil {
		aiu.SetCoverImageURL(*s)
	}
	return aiu
}

// ClearCoverImageURL clears the value of the "cover_image_url" field.
func (aiu *AppInfoUpdate) ClearCoverImageURL() *AppInfoUpdate {
	aiu.mutation.ClearCoverImageURL()
	return aiu
}

// SetCoverImageID sets the "cover_image_id" field.
func (aiu *AppInfoUpdate) SetCoverImageID(mi model.InternalID) *AppInfoUpdate {
	aiu.mutation.ResetCoverImageID()
	aiu.mutation.SetCoverImageID(mi)
	return aiu
}

// SetNillableCoverImageID sets the "cover_image_id" field if the given value is not nil.
func (aiu *AppInfoUpdate) SetNillableCoverImageID(mi *model.InternalID) *AppInfoUpdate {
	if mi != nil {
		aiu.SetCoverImageID(*mi)
	}
	return aiu
}

// AddCoverImageID adds mi to the "cover_image_id" field.
func (aiu *AppInfoUpdate) AddCoverImageID(mi model.InternalID) *AppInfoUpdate {
	aiu.mutation.AddCoverImageID(mi)
	return aiu
}

// SetReleaseDate sets the "release_date" field.
func (aiu *AppInfoUpdate) SetReleaseDate(s string) *AppInfoUpdate {
	aiu.mutation.SetReleaseDate(s)
	return aiu
}

// SetNillableReleaseDate sets the "release_date" field if the given value is not nil.
func (aiu *AppInfoUpdate) SetNillableReleaseDate(s *string) *AppInfoUpdate {
	if s != nil {
		aiu.SetReleaseDate(*s)
	}
	return aiu
}

// ClearReleaseDate clears the value of the "release_date" field.
func (aiu *AppInfoUpdate) ClearReleaseDate() *AppInfoUpdate {
	aiu.mutation.ClearReleaseDate()
	return aiu
}

// SetDeveloper sets the "developer" field.
func (aiu *AppInfoUpdate) SetDeveloper(s string) *AppInfoUpdate {
	aiu.mutation.SetDeveloper(s)
	return aiu
}

// SetNillableDeveloper sets the "developer" field if the given value is not nil.
func (aiu *AppInfoUpdate) SetNillableDeveloper(s *string) *AppInfoUpdate {
	if s != nil {
		aiu.SetDeveloper(*s)
	}
	return aiu
}

// ClearDeveloper clears the value of the "developer" field.
func (aiu *AppInfoUpdate) ClearDeveloper() *AppInfoUpdate {
	aiu.mutation.ClearDeveloper()
	return aiu
}

// SetPublisher sets the "publisher" field.
func (aiu *AppInfoUpdate) SetPublisher(s string) *AppInfoUpdate {
	aiu.mutation.SetPublisher(s)
	return aiu
}

// SetNillablePublisher sets the "publisher" field if the given value is not nil.
func (aiu *AppInfoUpdate) SetNillablePublisher(s *string) *AppInfoUpdate {
	if s != nil {
		aiu.SetPublisher(*s)
	}
	return aiu
}

// ClearPublisher clears the value of the "publisher" field.
func (aiu *AppInfoUpdate) ClearPublisher() *AppInfoUpdate {
	aiu.mutation.ClearPublisher()
	return aiu
}

// SetTags sets the "tags" field.
func (aiu *AppInfoUpdate) SetTags(s []string) *AppInfoUpdate {
	aiu.mutation.SetTags(s)
	return aiu
}

// AppendTags appends s to the "tags" field.
func (aiu *AppInfoUpdate) AppendTags(s []string) *AppInfoUpdate {
	aiu.mutation.AppendTags(s)
	return aiu
}

// SetAlternativeNames sets the "alternative_names" field.
func (aiu *AppInfoUpdate) SetAlternativeNames(s []string) *AppInfoUpdate {
	aiu.mutation.SetAlternativeNames(s)
	return aiu
}

// AppendAlternativeNames appends s to the "alternative_names" field.
func (aiu *AppInfoUpdate) AppendAlternativeNames(s []string) *AppInfoUpdate {
	aiu.mutation.AppendAlternativeNames(s)
	return aiu
}

// SetRawData sets the "raw_data" field.
func (aiu *AppInfoUpdate) SetRawData(s string) *AppInfoUpdate {
	aiu.mutation.SetRawData(s)
	return aiu
}

// SetNillableRawData sets the "raw_data" field if the given value is not nil.
func (aiu *AppInfoUpdate) SetNillableRawData(s *string) *AppInfoUpdate {
	if s != nil {
		aiu.SetRawData(*s)
	}
	return aiu
}

// SetUpdatedAt sets the "updated_at" field.
func (aiu *AppInfoUpdate) SetUpdatedAt(t time.Time) *AppInfoUpdate {
	aiu.mutation.SetUpdatedAt(t)
	return aiu
}

// SetCreatedAt sets the "created_at" field.
func (aiu *AppInfoUpdate) SetCreatedAt(t time.Time) *AppInfoUpdate {
	aiu.mutation.SetCreatedAt(t)
	return aiu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (aiu *AppInfoUpdate) SetNillableCreatedAt(t *time.Time) *AppInfoUpdate {
	if t != nil {
		aiu.SetCreatedAt(*t)
	}
	return aiu
}

// Mutation returns the AppInfoMutation object of the builder.
func (aiu *AppInfoUpdate) Mutation() *AppInfoMutation {
	return aiu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (aiu *AppInfoUpdate) Save(ctx context.Context) (int, error) {
	aiu.defaults()
	return withHooks(ctx, aiu.sqlSave, aiu.mutation, aiu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (aiu *AppInfoUpdate) SaveX(ctx context.Context) int {
	affected, err := aiu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (aiu *AppInfoUpdate) Exec(ctx context.Context) error {
	_, err := aiu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (aiu *AppInfoUpdate) ExecX(ctx context.Context) {
	if err := aiu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (aiu *AppInfoUpdate) defaults() {
	if _, ok := aiu.mutation.UpdatedAt(); !ok {
		v := appinfo.UpdateDefaultUpdatedAt()
		aiu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (aiu *AppInfoUpdate) check() error {
	if v, ok := aiu.mutation.GetType(); ok {
		if err := appinfo.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf(`ent: validator failed for field "AppInfo.type": %w`, err)}
		}
	}
	return nil
}

func (aiu *AppInfoUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := aiu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(appinfo.Table, appinfo.Columns, sqlgraph.NewFieldSpec(appinfo.FieldID, field.TypeInt64))
	if ps := aiu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := aiu.mutation.Source(); ok {
		_spec.SetField(appinfo.FieldSource, field.TypeString, value)
	}
	if value, ok := aiu.mutation.SourceAppID(); ok {
		_spec.SetField(appinfo.FieldSourceAppID, field.TypeString, value)
	}
	if value, ok := aiu.mutation.SourceURL(); ok {
		_spec.SetField(appinfo.FieldSourceURL, field.TypeString, value)
	}
	if aiu.mutation.SourceURLCleared() {
		_spec.ClearField(appinfo.FieldSourceURL, field.TypeString)
	}
	if value, ok := aiu.mutation.Name(); ok {
		_spec.SetField(appinfo.FieldName, field.TypeString, value)
	}
	if value, ok := aiu.mutation.GetType(); ok {
		_spec.SetField(appinfo.FieldType, field.TypeEnum, value)
	}
	if value, ok := aiu.mutation.ShortDescription(); ok {
		_spec.SetField(appinfo.FieldShortDescription, field.TypeString, value)
	}
	if aiu.mutation.ShortDescriptionCleared() {
		_spec.ClearField(appinfo.FieldShortDescription, field.TypeString)
	}
	if value, ok := aiu.mutation.Description(); ok {
		_spec.SetField(appinfo.FieldDescription, field.TypeString, value)
	}
	if aiu.mutation.DescriptionCleared() {
		_spec.ClearField(appinfo.FieldDescription, field.TypeString)
	}
	if value, ok := aiu.mutation.IconImageURL(); ok {
		_spec.SetField(appinfo.FieldIconImageURL, field.TypeString, value)
	}
	if aiu.mutation.IconImageURLCleared() {
		_spec.ClearField(appinfo.FieldIconImageURL, field.TypeString)
	}
	if value, ok := aiu.mutation.IconImageID(); ok {
		_spec.SetField(appinfo.FieldIconImageID, field.TypeInt64, value)
	}
	if value, ok := aiu.mutation.AddedIconImageID(); ok {
		_spec.AddField(appinfo.FieldIconImageID, field.TypeInt64, value)
	}
	if value, ok := aiu.mutation.BackgroundImageURL(); ok {
		_spec.SetField(appinfo.FieldBackgroundImageURL, field.TypeString, value)
	}
	if aiu.mutation.BackgroundImageURLCleared() {
		_spec.ClearField(appinfo.FieldBackgroundImageURL, field.TypeString)
	}
	if value, ok := aiu.mutation.BackgroundImageID(); ok {
		_spec.SetField(appinfo.FieldBackgroundImageID, field.TypeInt64, value)
	}
	if value, ok := aiu.mutation.AddedBackgroundImageID(); ok {
		_spec.AddField(appinfo.FieldBackgroundImageID, field.TypeInt64, value)
	}
	if value, ok := aiu.mutation.CoverImageURL(); ok {
		_spec.SetField(appinfo.FieldCoverImageURL, field.TypeString, value)
	}
	if aiu.mutation.CoverImageURLCleared() {
		_spec.ClearField(appinfo.FieldCoverImageURL, field.TypeString)
	}
	if value, ok := aiu.mutation.CoverImageID(); ok {
		_spec.SetField(appinfo.FieldCoverImageID, field.TypeInt64, value)
	}
	if value, ok := aiu.mutation.AddedCoverImageID(); ok {
		_spec.AddField(appinfo.FieldCoverImageID, field.TypeInt64, value)
	}
	if value, ok := aiu.mutation.ReleaseDate(); ok {
		_spec.SetField(appinfo.FieldReleaseDate, field.TypeString, value)
	}
	if aiu.mutation.ReleaseDateCleared() {
		_spec.ClearField(appinfo.FieldReleaseDate, field.TypeString)
	}
	if value, ok := aiu.mutation.Developer(); ok {
		_spec.SetField(appinfo.FieldDeveloper, field.TypeString, value)
	}
	if aiu.mutation.DeveloperCleared() {
		_spec.ClearField(appinfo.FieldDeveloper, field.TypeString)
	}
	if value, ok := aiu.mutation.Publisher(); ok {
		_spec.SetField(appinfo.FieldPublisher, field.TypeString, value)
	}
	if aiu.mutation.PublisherCleared() {
		_spec.ClearField(appinfo.FieldPublisher, field.TypeString)
	}
	if value, ok := aiu.mutation.Tags(); ok {
		_spec.SetField(appinfo.FieldTags, field.TypeJSON, value)
	}
	if value, ok := aiu.mutation.AppendedTags(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, appinfo.FieldTags, value)
		})
	}
	if value, ok := aiu.mutation.AlternativeNames(); ok {
		_spec.SetField(appinfo.FieldAlternativeNames, field.TypeJSON, value)
	}
	if value, ok := aiu.mutation.AppendedAlternativeNames(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, appinfo.FieldAlternativeNames, value)
		})
	}
	if value, ok := aiu.mutation.RawData(); ok {
		_spec.SetField(appinfo.FieldRawData, field.TypeString, value)
	}
	if value, ok := aiu.mutation.UpdatedAt(); ok {
		_spec.SetField(appinfo.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := aiu.mutation.CreatedAt(); ok {
		_spec.SetField(appinfo.FieldCreatedAt, field.TypeTime, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, aiu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{appinfo.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	aiu.mutation.done = true
	return n, nil
}

// AppInfoUpdateOne is the builder for updating a single AppInfo entity.
type AppInfoUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *AppInfoMutation
}

// SetSource sets the "source" field.
func (aiuo *AppInfoUpdateOne) SetSource(s string) *AppInfoUpdateOne {
	aiuo.mutation.SetSource(s)
	return aiuo
}

// SetNillableSource sets the "source" field if the given value is not nil.
func (aiuo *AppInfoUpdateOne) SetNillableSource(s *string) *AppInfoUpdateOne {
	if s != nil {
		aiuo.SetSource(*s)
	}
	return aiuo
}

// SetSourceAppID sets the "source_app_id" field.
func (aiuo *AppInfoUpdateOne) SetSourceAppID(s string) *AppInfoUpdateOne {
	aiuo.mutation.SetSourceAppID(s)
	return aiuo
}

// SetNillableSourceAppID sets the "source_app_id" field if the given value is not nil.
func (aiuo *AppInfoUpdateOne) SetNillableSourceAppID(s *string) *AppInfoUpdateOne {
	if s != nil {
		aiuo.SetSourceAppID(*s)
	}
	return aiuo
}

// SetSourceURL sets the "source_url" field.
func (aiuo *AppInfoUpdateOne) SetSourceURL(s string) *AppInfoUpdateOne {
	aiuo.mutation.SetSourceURL(s)
	return aiuo
}

// SetNillableSourceURL sets the "source_url" field if the given value is not nil.
func (aiuo *AppInfoUpdateOne) SetNillableSourceURL(s *string) *AppInfoUpdateOne {
	if s != nil {
		aiuo.SetSourceURL(*s)
	}
	return aiuo
}

// ClearSourceURL clears the value of the "source_url" field.
func (aiuo *AppInfoUpdateOne) ClearSourceURL() *AppInfoUpdateOne {
	aiuo.mutation.ClearSourceURL()
	return aiuo
}

// SetName sets the "name" field.
func (aiuo *AppInfoUpdateOne) SetName(s string) *AppInfoUpdateOne {
	aiuo.mutation.SetName(s)
	return aiuo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (aiuo *AppInfoUpdateOne) SetNillableName(s *string) *AppInfoUpdateOne {
	if s != nil {
		aiuo.SetName(*s)
	}
	return aiuo
}

// SetType sets the "type" field.
func (aiuo *AppInfoUpdateOne) SetType(a appinfo.Type) *AppInfoUpdateOne {
	aiuo.mutation.SetType(a)
	return aiuo
}

// SetNillableType sets the "type" field if the given value is not nil.
func (aiuo *AppInfoUpdateOne) SetNillableType(a *appinfo.Type) *AppInfoUpdateOne {
	if a != nil {
		aiuo.SetType(*a)
	}
	return aiuo
}

// SetShortDescription sets the "short_description" field.
func (aiuo *AppInfoUpdateOne) SetShortDescription(s string) *AppInfoUpdateOne {
	aiuo.mutation.SetShortDescription(s)
	return aiuo
}

// SetNillableShortDescription sets the "short_description" field if the given value is not nil.
func (aiuo *AppInfoUpdateOne) SetNillableShortDescription(s *string) *AppInfoUpdateOne {
	if s != nil {
		aiuo.SetShortDescription(*s)
	}
	return aiuo
}

// ClearShortDescription clears the value of the "short_description" field.
func (aiuo *AppInfoUpdateOne) ClearShortDescription() *AppInfoUpdateOne {
	aiuo.mutation.ClearShortDescription()
	return aiuo
}

// SetDescription sets the "description" field.
func (aiuo *AppInfoUpdateOne) SetDescription(s string) *AppInfoUpdateOne {
	aiuo.mutation.SetDescription(s)
	return aiuo
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (aiuo *AppInfoUpdateOne) SetNillableDescription(s *string) *AppInfoUpdateOne {
	if s != nil {
		aiuo.SetDescription(*s)
	}
	return aiuo
}

// ClearDescription clears the value of the "description" field.
func (aiuo *AppInfoUpdateOne) ClearDescription() *AppInfoUpdateOne {
	aiuo.mutation.ClearDescription()
	return aiuo
}

// SetIconImageURL sets the "icon_image_url" field.
func (aiuo *AppInfoUpdateOne) SetIconImageURL(s string) *AppInfoUpdateOne {
	aiuo.mutation.SetIconImageURL(s)
	return aiuo
}

// SetNillableIconImageURL sets the "icon_image_url" field if the given value is not nil.
func (aiuo *AppInfoUpdateOne) SetNillableIconImageURL(s *string) *AppInfoUpdateOne {
	if s != nil {
		aiuo.SetIconImageURL(*s)
	}
	return aiuo
}

// ClearIconImageURL clears the value of the "icon_image_url" field.
func (aiuo *AppInfoUpdateOne) ClearIconImageURL() *AppInfoUpdateOne {
	aiuo.mutation.ClearIconImageURL()
	return aiuo
}

// SetIconImageID sets the "icon_image_id" field.
func (aiuo *AppInfoUpdateOne) SetIconImageID(mi model.InternalID) *AppInfoUpdateOne {
	aiuo.mutation.ResetIconImageID()
	aiuo.mutation.SetIconImageID(mi)
	return aiuo
}

// SetNillableIconImageID sets the "icon_image_id" field if the given value is not nil.
func (aiuo *AppInfoUpdateOne) SetNillableIconImageID(mi *model.InternalID) *AppInfoUpdateOne {
	if mi != nil {
		aiuo.SetIconImageID(*mi)
	}
	return aiuo
}

// AddIconImageID adds mi to the "icon_image_id" field.
func (aiuo *AppInfoUpdateOne) AddIconImageID(mi model.InternalID) *AppInfoUpdateOne {
	aiuo.mutation.AddIconImageID(mi)
	return aiuo
}

// SetBackgroundImageURL sets the "background_image_url" field.
func (aiuo *AppInfoUpdateOne) SetBackgroundImageURL(s string) *AppInfoUpdateOne {
	aiuo.mutation.SetBackgroundImageURL(s)
	return aiuo
}

// SetNillableBackgroundImageURL sets the "background_image_url" field if the given value is not nil.
func (aiuo *AppInfoUpdateOne) SetNillableBackgroundImageURL(s *string) *AppInfoUpdateOne {
	if s != nil {
		aiuo.SetBackgroundImageURL(*s)
	}
	return aiuo
}

// ClearBackgroundImageURL clears the value of the "background_image_url" field.
func (aiuo *AppInfoUpdateOne) ClearBackgroundImageURL() *AppInfoUpdateOne {
	aiuo.mutation.ClearBackgroundImageURL()
	return aiuo
}

// SetBackgroundImageID sets the "background_image_id" field.
func (aiuo *AppInfoUpdateOne) SetBackgroundImageID(mi model.InternalID) *AppInfoUpdateOne {
	aiuo.mutation.ResetBackgroundImageID()
	aiuo.mutation.SetBackgroundImageID(mi)
	return aiuo
}

// SetNillableBackgroundImageID sets the "background_image_id" field if the given value is not nil.
func (aiuo *AppInfoUpdateOne) SetNillableBackgroundImageID(mi *model.InternalID) *AppInfoUpdateOne {
	if mi != nil {
		aiuo.SetBackgroundImageID(*mi)
	}
	return aiuo
}

// AddBackgroundImageID adds mi to the "background_image_id" field.
func (aiuo *AppInfoUpdateOne) AddBackgroundImageID(mi model.InternalID) *AppInfoUpdateOne {
	aiuo.mutation.AddBackgroundImageID(mi)
	return aiuo
}

// SetCoverImageURL sets the "cover_image_url" field.
func (aiuo *AppInfoUpdateOne) SetCoverImageURL(s string) *AppInfoUpdateOne {
	aiuo.mutation.SetCoverImageURL(s)
	return aiuo
}

// SetNillableCoverImageURL sets the "cover_image_url" field if the given value is not nil.
func (aiuo *AppInfoUpdateOne) SetNillableCoverImageURL(s *string) *AppInfoUpdateOne {
	if s != nil {
		aiuo.SetCoverImageURL(*s)
	}
	return aiuo
}

// ClearCoverImageURL clears the value of the "cover_image_url" field.
func (aiuo *AppInfoUpdateOne) ClearCoverImageURL() *AppInfoUpdateOne {
	aiuo.mutation.ClearCoverImageURL()
	return aiuo
}

// SetCoverImageID sets the "cover_image_id" field.
func (aiuo *AppInfoUpdateOne) SetCoverImageID(mi model.InternalID) *AppInfoUpdateOne {
	aiuo.mutation.ResetCoverImageID()
	aiuo.mutation.SetCoverImageID(mi)
	return aiuo
}

// SetNillableCoverImageID sets the "cover_image_id" field if the given value is not nil.
func (aiuo *AppInfoUpdateOne) SetNillableCoverImageID(mi *model.InternalID) *AppInfoUpdateOne {
	if mi != nil {
		aiuo.SetCoverImageID(*mi)
	}
	return aiuo
}

// AddCoverImageID adds mi to the "cover_image_id" field.
func (aiuo *AppInfoUpdateOne) AddCoverImageID(mi model.InternalID) *AppInfoUpdateOne {
	aiuo.mutation.AddCoverImageID(mi)
	return aiuo
}

// SetReleaseDate sets the "release_date" field.
func (aiuo *AppInfoUpdateOne) SetReleaseDate(s string) *AppInfoUpdateOne {
	aiuo.mutation.SetReleaseDate(s)
	return aiuo
}

// SetNillableReleaseDate sets the "release_date" field if the given value is not nil.
func (aiuo *AppInfoUpdateOne) SetNillableReleaseDate(s *string) *AppInfoUpdateOne {
	if s != nil {
		aiuo.SetReleaseDate(*s)
	}
	return aiuo
}

// ClearReleaseDate clears the value of the "release_date" field.
func (aiuo *AppInfoUpdateOne) ClearReleaseDate() *AppInfoUpdateOne {
	aiuo.mutation.ClearReleaseDate()
	return aiuo
}

// SetDeveloper sets the "developer" field.
func (aiuo *AppInfoUpdateOne) SetDeveloper(s string) *AppInfoUpdateOne {
	aiuo.mutation.SetDeveloper(s)
	return aiuo
}

// SetNillableDeveloper sets the "developer" field if the given value is not nil.
func (aiuo *AppInfoUpdateOne) SetNillableDeveloper(s *string) *AppInfoUpdateOne {
	if s != nil {
		aiuo.SetDeveloper(*s)
	}
	return aiuo
}

// ClearDeveloper clears the value of the "developer" field.
func (aiuo *AppInfoUpdateOne) ClearDeveloper() *AppInfoUpdateOne {
	aiuo.mutation.ClearDeveloper()
	return aiuo
}

// SetPublisher sets the "publisher" field.
func (aiuo *AppInfoUpdateOne) SetPublisher(s string) *AppInfoUpdateOne {
	aiuo.mutation.SetPublisher(s)
	return aiuo
}

// SetNillablePublisher sets the "publisher" field if the given value is not nil.
func (aiuo *AppInfoUpdateOne) SetNillablePublisher(s *string) *AppInfoUpdateOne {
	if s != nil {
		aiuo.SetPublisher(*s)
	}
	return aiuo
}

// ClearPublisher clears the value of the "publisher" field.
func (aiuo *AppInfoUpdateOne) ClearPublisher() *AppInfoUpdateOne {
	aiuo.mutation.ClearPublisher()
	return aiuo
}

// SetTags sets the "tags" field.
func (aiuo *AppInfoUpdateOne) SetTags(s []string) *AppInfoUpdateOne {
	aiuo.mutation.SetTags(s)
	return aiuo
}

// AppendTags appends s to the "tags" field.
func (aiuo *AppInfoUpdateOne) AppendTags(s []string) *AppInfoUpdateOne {
	aiuo.mutation.AppendTags(s)
	return aiuo
}

// SetAlternativeNames sets the "alternative_names" field.
func (aiuo *AppInfoUpdateOne) SetAlternativeNames(s []string) *AppInfoUpdateOne {
	aiuo.mutation.SetAlternativeNames(s)
	return aiuo
}

// AppendAlternativeNames appends s to the "alternative_names" field.
func (aiuo *AppInfoUpdateOne) AppendAlternativeNames(s []string) *AppInfoUpdateOne {
	aiuo.mutation.AppendAlternativeNames(s)
	return aiuo
}

// SetRawData sets the "raw_data" field.
func (aiuo *AppInfoUpdateOne) SetRawData(s string) *AppInfoUpdateOne {
	aiuo.mutation.SetRawData(s)
	return aiuo
}

// SetNillableRawData sets the "raw_data" field if the given value is not nil.
func (aiuo *AppInfoUpdateOne) SetNillableRawData(s *string) *AppInfoUpdateOne {
	if s != nil {
		aiuo.SetRawData(*s)
	}
	return aiuo
}

// SetUpdatedAt sets the "updated_at" field.
func (aiuo *AppInfoUpdateOne) SetUpdatedAt(t time.Time) *AppInfoUpdateOne {
	aiuo.mutation.SetUpdatedAt(t)
	return aiuo
}

// SetCreatedAt sets the "created_at" field.
func (aiuo *AppInfoUpdateOne) SetCreatedAt(t time.Time) *AppInfoUpdateOne {
	aiuo.mutation.SetCreatedAt(t)
	return aiuo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (aiuo *AppInfoUpdateOne) SetNillableCreatedAt(t *time.Time) *AppInfoUpdateOne {
	if t != nil {
		aiuo.SetCreatedAt(*t)
	}
	return aiuo
}

// Mutation returns the AppInfoMutation object of the builder.
func (aiuo *AppInfoUpdateOne) Mutation() *AppInfoMutation {
	return aiuo.mutation
}

// Where appends a list predicates to the AppInfoUpdate builder.
func (aiuo *AppInfoUpdateOne) Where(ps ...predicate.AppInfo) *AppInfoUpdateOne {
	aiuo.mutation.Where(ps...)
	return aiuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (aiuo *AppInfoUpdateOne) Select(field string, fields ...string) *AppInfoUpdateOne {
	aiuo.fields = append([]string{field}, fields...)
	return aiuo
}

// Save executes the query and returns the updated AppInfo entity.
func (aiuo *AppInfoUpdateOne) Save(ctx context.Context) (*AppInfo, error) {
	aiuo.defaults()
	return withHooks(ctx, aiuo.sqlSave, aiuo.mutation, aiuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (aiuo *AppInfoUpdateOne) SaveX(ctx context.Context) *AppInfo {
	node, err := aiuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (aiuo *AppInfoUpdateOne) Exec(ctx context.Context) error {
	_, err := aiuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (aiuo *AppInfoUpdateOne) ExecX(ctx context.Context) {
	if err := aiuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (aiuo *AppInfoUpdateOne) defaults() {
	if _, ok := aiuo.mutation.UpdatedAt(); !ok {
		v := appinfo.UpdateDefaultUpdatedAt()
		aiuo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (aiuo *AppInfoUpdateOne) check() error {
	if v, ok := aiuo.mutation.GetType(); ok {
		if err := appinfo.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf(`ent: validator failed for field "AppInfo.type": %w`, err)}
		}
	}
	return nil
}

func (aiuo *AppInfoUpdateOne) sqlSave(ctx context.Context) (_node *AppInfo, err error) {
	if err := aiuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(appinfo.Table, appinfo.Columns, sqlgraph.NewFieldSpec(appinfo.FieldID, field.TypeInt64))
	id, ok := aiuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "AppInfo.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := aiuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, appinfo.FieldID)
		for _, f := range fields {
			if !appinfo.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != appinfo.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := aiuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := aiuo.mutation.Source(); ok {
		_spec.SetField(appinfo.FieldSource, field.TypeString, value)
	}
	if value, ok := aiuo.mutation.SourceAppID(); ok {
		_spec.SetField(appinfo.FieldSourceAppID, field.TypeString, value)
	}
	if value, ok := aiuo.mutation.SourceURL(); ok {
		_spec.SetField(appinfo.FieldSourceURL, field.TypeString, value)
	}
	if aiuo.mutation.SourceURLCleared() {
		_spec.ClearField(appinfo.FieldSourceURL, field.TypeString)
	}
	if value, ok := aiuo.mutation.Name(); ok {
		_spec.SetField(appinfo.FieldName, field.TypeString, value)
	}
	if value, ok := aiuo.mutation.GetType(); ok {
		_spec.SetField(appinfo.FieldType, field.TypeEnum, value)
	}
	if value, ok := aiuo.mutation.ShortDescription(); ok {
		_spec.SetField(appinfo.FieldShortDescription, field.TypeString, value)
	}
	if aiuo.mutation.ShortDescriptionCleared() {
		_spec.ClearField(appinfo.FieldShortDescription, field.TypeString)
	}
	if value, ok := aiuo.mutation.Description(); ok {
		_spec.SetField(appinfo.FieldDescription, field.TypeString, value)
	}
	if aiuo.mutation.DescriptionCleared() {
		_spec.ClearField(appinfo.FieldDescription, field.TypeString)
	}
	if value, ok := aiuo.mutation.IconImageURL(); ok {
		_spec.SetField(appinfo.FieldIconImageURL, field.TypeString, value)
	}
	if aiuo.mutation.IconImageURLCleared() {
		_spec.ClearField(appinfo.FieldIconImageURL, field.TypeString)
	}
	if value, ok := aiuo.mutation.IconImageID(); ok {
		_spec.SetField(appinfo.FieldIconImageID, field.TypeInt64, value)
	}
	if value, ok := aiuo.mutation.AddedIconImageID(); ok {
		_spec.AddField(appinfo.FieldIconImageID, field.TypeInt64, value)
	}
	if value, ok := aiuo.mutation.BackgroundImageURL(); ok {
		_spec.SetField(appinfo.FieldBackgroundImageURL, field.TypeString, value)
	}
	if aiuo.mutation.BackgroundImageURLCleared() {
		_spec.ClearField(appinfo.FieldBackgroundImageURL, field.TypeString)
	}
	if value, ok := aiuo.mutation.BackgroundImageID(); ok {
		_spec.SetField(appinfo.FieldBackgroundImageID, field.TypeInt64, value)
	}
	if value, ok := aiuo.mutation.AddedBackgroundImageID(); ok {
		_spec.AddField(appinfo.FieldBackgroundImageID, field.TypeInt64, value)
	}
	if value, ok := aiuo.mutation.CoverImageURL(); ok {
		_spec.SetField(appinfo.FieldCoverImageURL, field.TypeString, value)
	}
	if aiuo.mutation.CoverImageURLCleared() {
		_spec.ClearField(appinfo.FieldCoverImageURL, field.TypeString)
	}
	if value, ok := aiuo.mutation.CoverImageID(); ok {
		_spec.SetField(appinfo.FieldCoverImageID, field.TypeInt64, value)
	}
	if value, ok := aiuo.mutation.AddedCoverImageID(); ok {
		_spec.AddField(appinfo.FieldCoverImageID, field.TypeInt64, value)
	}
	if value, ok := aiuo.mutation.ReleaseDate(); ok {
		_spec.SetField(appinfo.FieldReleaseDate, field.TypeString, value)
	}
	if aiuo.mutation.ReleaseDateCleared() {
		_spec.ClearField(appinfo.FieldReleaseDate, field.TypeString)
	}
	if value, ok := aiuo.mutation.Developer(); ok {
		_spec.SetField(appinfo.FieldDeveloper, field.TypeString, value)
	}
	if aiuo.mutation.DeveloperCleared() {
		_spec.ClearField(appinfo.FieldDeveloper, field.TypeString)
	}
	if value, ok := aiuo.mutation.Publisher(); ok {
		_spec.SetField(appinfo.FieldPublisher, field.TypeString, value)
	}
	if aiuo.mutation.PublisherCleared() {
		_spec.ClearField(appinfo.FieldPublisher, field.TypeString)
	}
	if value, ok := aiuo.mutation.Tags(); ok {
		_spec.SetField(appinfo.FieldTags, field.TypeJSON, value)
	}
	if value, ok := aiuo.mutation.AppendedTags(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, appinfo.FieldTags, value)
		})
	}
	if value, ok := aiuo.mutation.AlternativeNames(); ok {
		_spec.SetField(appinfo.FieldAlternativeNames, field.TypeJSON, value)
	}
	if value, ok := aiuo.mutation.AppendedAlternativeNames(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, appinfo.FieldAlternativeNames, value)
		})
	}
	if value, ok := aiuo.mutation.RawData(); ok {
		_spec.SetField(appinfo.FieldRawData, field.TypeString, value)
	}
	if value, ok := aiuo.mutation.UpdatedAt(); ok {
		_spec.SetField(appinfo.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := aiuo.mutation.CreatedAt(); ok {
		_spec.SetField(appinfo.FieldCreatedAt, field.TypeTime, value)
	}
	_node = &AppInfo{config: aiuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, aiuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{appinfo.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	aiuo.mutation.done = true
	return _node, nil
}
