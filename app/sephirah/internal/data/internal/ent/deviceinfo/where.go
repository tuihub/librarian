// Code generated by ent, DO NOT EDIT.

package deviceinfo

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/tuihub/librarian/app/sephirah/internal/data/internal/ent/predicate"
	"github.com/tuihub/librarian/internal/model"
)

// ID filters vertices based on their ID field.
func ID(id model.InternalID) predicate.DeviceInfo {
	return predicate.DeviceInfo(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id model.InternalID) predicate.DeviceInfo {
	return predicate.DeviceInfo(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id model.InternalID) predicate.DeviceInfo {
	return predicate.DeviceInfo(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...model.InternalID) predicate.DeviceInfo {
	return predicate.DeviceInfo(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...model.InternalID) predicate.DeviceInfo {
	return predicate.DeviceInfo(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id model.InternalID) predicate.DeviceInfo {
	return predicate.DeviceInfo(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id model.InternalID) predicate.DeviceInfo {
	return predicate.DeviceInfo(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id model.InternalID) predicate.DeviceInfo {
	return predicate.DeviceInfo(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id model.InternalID) predicate.DeviceInfo {
	return predicate.DeviceInfo(sql.FieldLTE(FieldID, id))
}

// DeviceName applies equality check predicate on the "device_name" field. It's identical to DeviceNameEQ.
func DeviceName(v string) predicate.DeviceInfo {
	return predicate.DeviceInfo(sql.FieldEQ(FieldDeviceName, v))
}

// SystemVersion applies equality check predicate on the "system_version" field. It's identical to SystemVersionEQ.
func SystemVersion(v string) predicate.DeviceInfo {
	return predicate.DeviceInfo(sql.FieldEQ(FieldSystemVersion, v))
}

// ClientName applies equality check predicate on the "client_name" field. It's identical to ClientNameEQ.
func ClientName(v string) predicate.DeviceInfo {
	return predicate.DeviceInfo(sql.FieldEQ(FieldClientName, v))
}

// ClientSourceCodeAddress applies equality check predicate on the "client_source_code_address" field. It's identical to ClientSourceCodeAddressEQ.
func ClientSourceCodeAddress(v string) predicate.DeviceInfo {
	return predicate.DeviceInfo(sql.FieldEQ(FieldClientSourceCodeAddress, v))
}

// ClientVersion applies equality check predicate on the "client_version" field. It's identical to ClientVersionEQ.
func ClientVersion(v string) predicate.DeviceInfo {
	return predicate.DeviceInfo(sql.FieldEQ(FieldClientVersion, v))
}

// ClientLocalID applies equality check predicate on the "client_local_id" field. It's identical to ClientLocalIDEQ.
func ClientLocalID(v string) predicate.DeviceInfo {
	return predicate.DeviceInfo(sql.FieldEQ(FieldClientLocalID, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.DeviceInfo {
	return predicate.DeviceInfo(sql.FieldEQ(FieldUpdatedAt, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.DeviceInfo {
	return predicate.DeviceInfo(sql.FieldEQ(FieldCreatedAt, v))
}

// DeviceNameEQ applies the EQ predicate on the "device_name" field.
func DeviceNameEQ(v string) predicate.DeviceInfo {
	return predicate.DeviceInfo(sql.FieldEQ(FieldDeviceName, v))
}

// DeviceNameNEQ applies the NEQ predicate on the "device_name" field.
func DeviceNameNEQ(v string) predicate.DeviceInfo {
	return predicate.DeviceInfo(sql.FieldNEQ(FieldDeviceName, v))
}

// DeviceNameIn applies the In predicate on the "device_name" field.
func DeviceNameIn(vs ...string) predicate.DeviceInfo {
	return predicate.DeviceInfo(sql.FieldIn(FieldDeviceName, vs...))
}

// DeviceNameNotIn applies the NotIn predicate on the "device_name" field.
func DeviceNameNotIn(vs ...string) predicate.DeviceInfo {
	return predicate.DeviceInfo(sql.FieldNotIn(FieldDeviceName, vs...))
}

// DeviceNameGT applies the GT predicate on the "device_name" field.
func DeviceNameGT(v string) predicate.DeviceInfo {
	return predicate.DeviceInfo(sql.FieldGT(FieldDeviceName, v))
}

// DeviceNameGTE applies the GTE predicate on the "device_name" field.
func DeviceNameGTE(v string) predicate.DeviceInfo {
	return predicate.DeviceInfo(sql.FieldGTE(FieldDeviceName, v))
}

// DeviceNameLT applies the LT predicate on the "device_name" field.
func DeviceNameLT(v string) predicate.DeviceInfo {
	return predicate.DeviceInfo(sql.FieldLT(FieldDeviceName, v))
}

// DeviceNameLTE applies the LTE predicate on the "device_name" field.
func DeviceNameLTE(v string) predicate.DeviceInfo {
	return predicate.DeviceInfo(sql.FieldLTE(FieldDeviceName, v))
}

// DeviceNameContains applies the Contains predicate on the "device_name" field.
func DeviceNameContains(v string) predicate.DeviceInfo {
	return predicate.DeviceInfo(sql.FieldContains(FieldDeviceName, v))
}

// DeviceNameHasPrefix applies the HasPrefix predicate on the "device_name" field.
func DeviceNameHasPrefix(v string) predicate.DeviceInfo {
	return predicate.DeviceInfo(sql.FieldHasPrefix(FieldDeviceName, v))
}

// DeviceNameHasSuffix applies the HasSuffix predicate on the "device_name" field.
func DeviceNameHasSuffix(v string) predicate.DeviceInfo {
	return predicate.DeviceInfo(sql.FieldHasSuffix(FieldDeviceName, v))
}

// DeviceNameEqualFold applies the EqualFold predicate on the "device_name" field.
func DeviceNameEqualFold(v string) predicate.DeviceInfo {
	return predicate.DeviceInfo(sql.FieldEqualFold(FieldDeviceName, v))
}

// DeviceNameContainsFold applies the ContainsFold predicate on the "device_name" field.
func DeviceNameContainsFold(v string) predicate.DeviceInfo {
	return predicate.DeviceInfo(sql.FieldContainsFold(FieldDeviceName, v))
}

// SystemTypeEQ applies the EQ predicate on the "system_type" field.
func SystemTypeEQ(v SystemType) predicate.DeviceInfo {
	return predicate.DeviceInfo(sql.FieldEQ(FieldSystemType, v))
}

// SystemTypeNEQ applies the NEQ predicate on the "system_type" field.
func SystemTypeNEQ(v SystemType) predicate.DeviceInfo {
	return predicate.DeviceInfo(sql.FieldNEQ(FieldSystemType, v))
}

// SystemTypeIn applies the In predicate on the "system_type" field.
func SystemTypeIn(vs ...SystemType) predicate.DeviceInfo {
	return predicate.DeviceInfo(sql.FieldIn(FieldSystemType, vs...))
}

// SystemTypeNotIn applies the NotIn predicate on the "system_type" field.
func SystemTypeNotIn(vs ...SystemType) predicate.DeviceInfo {
	return predicate.DeviceInfo(sql.FieldNotIn(FieldSystemType, vs...))
}

// SystemVersionEQ applies the EQ predicate on the "system_version" field.
func SystemVersionEQ(v string) predicate.DeviceInfo {
	return predicate.DeviceInfo(sql.FieldEQ(FieldSystemVersion, v))
}

// SystemVersionNEQ applies the NEQ predicate on the "system_version" field.
func SystemVersionNEQ(v string) predicate.DeviceInfo {
	return predicate.DeviceInfo(sql.FieldNEQ(FieldSystemVersion, v))
}

// SystemVersionIn applies the In predicate on the "system_version" field.
func SystemVersionIn(vs ...string) predicate.DeviceInfo {
	return predicate.DeviceInfo(sql.FieldIn(FieldSystemVersion, vs...))
}

// SystemVersionNotIn applies the NotIn predicate on the "system_version" field.
func SystemVersionNotIn(vs ...string) predicate.DeviceInfo {
	return predicate.DeviceInfo(sql.FieldNotIn(FieldSystemVersion, vs...))
}

// SystemVersionGT applies the GT predicate on the "system_version" field.
func SystemVersionGT(v string) predicate.DeviceInfo {
	return predicate.DeviceInfo(sql.FieldGT(FieldSystemVersion, v))
}

// SystemVersionGTE applies the GTE predicate on the "system_version" field.
func SystemVersionGTE(v string) predicate.DeviceInfo {
	return predicate.DeviceInfo(sql.FieldGTE(FieldSystemVersion, v))
}

// SystemVersionLT applies the LT predicate on the "system_version" field.
func SystemVersionLT(v string) predicate.DeviceInfo {
	return predicate.DeviceInfo(sql.FieldLT(FieldSystemVersion, v))
}

// SystemVersionLTE applies the LTE predicate on the "system_version" field.
func SystemVersionLTE(v string) predicate.DeviceInfo {
	return predicate.DeviceInfo(sql.FieldLTE(FieldSystemVersion, v))
}

// SystemVersionContains applies the Contains predicate on the "system_version" field.
func SystemVersionContains(v string) predicate.DeviceInfo {
	return predicate.DeviceInfo(sql.FieldContains(FieldSystemVersion, v))
}

// SystemVersionHasPrefix applies the HasPrefix predicate on the "system_version" field.
func SystemVersionHasPrefix(v string) predicate.DeviceInfo {
	return predicate.DeviceInfo(sql.FieldHasPrefix(FieldSystemVersion, v))
}

// SystemVersionHasSuffix applies the HasSuffix predicate on the "system_version" field.
func SystemVersionHasSuffix(v string) predicate.DeviceInfo {
	return predicate.DeviceInfo(sql.FieldHasSuffix(FieldSystemVersion, v))
}

// SystemVersionEqualFold applies the EqualFold predicate on the "system_version" field.
func SystemVersionEqualFold(v string) predicate.DeviceInfo {
	return predicate.DeviceInfo(sql.FieldEqualFold(FieldSystemVersion, v))
}

// SystemVersionContainsFold applies the ContainsFold predicate on the "system_version" field.
func SystemVersionContainsFold(v string) predicate.DeviceInfo {
	return predicate.DeviceInfo(sql.FieldContainsFold(FieldSystemVersion, v))
}

// ClientNameEQ applies the EQ predicate on the "client_name" field.
func ClientNameEQ(v string) predicate.DeviceInfo {
	return predicate.DeviceInfo(sql.FieldEQ(FieldClientName, v))
}

// ClientNameNEQ applies the NEQ predicate on the "client_name" field.
func ClientNameNEQ(v string) predicate.DeviceInfo {
	return predicate.DeviceInfo(sql.FieldNEQ(FieldClientName, v))
}

// ClientNameIn applies the In predicate on the "client_name" field.
func ClientNameIn(vs ...string) predicate.DeviceInfo {
	return predicate.DeviceInfo(sql.FieldIn(FieldClientName, vs...))
}

// ClientNameNotIn applies the NotIn predicate on the "client_name" field.
func ClientNameNotIn(vs ...string) predicate.DeviceInfo {
	return predicate.DeviceInfo(sql.FieldNotIn(FieldClientName, vs...))
}

// ClientNameGT applies the GT predicate on the "client_name" field.
func ClientNameGT(v string) predicate.DeviceInfo {
	return predicate.DeviceInfo(sql.FieldGT(FieldClientName, v))
}

// ClientNameGTE applies the GTE predicate on the "client_name" field.
func ClientNameGTE(v string) predicate.DeviceInfo {
	return predicate.DeviceInfo(sql.FieldGTE(FieldClientName, v))
}

// ClientNameLT applies the LT predicate on the "client_name" field.
func ClientNameLT(v string) predicate.DeviceInfo {
	return predicate.DeviceInfo(sql.FieldLT(FieldClientName, v))
}

// ClientNameLTE applies the LTE predicate on the "client_name" field.
func ClientNameLTE(v string) predicate.DeviceInfo {
	return predicate.DeviceInfo(sql.FieldLTE(FieldClientName, v))
}

// ClientNameContains applies the Contains predicate on the "client_name" field.
func ClientNameContains(v string) predicate.DeviceInfo {
	return predicate.DeviceInfo(sql.FieldContains(FieldClientName, v))
}

// ClientNameHasPrefix applies the HasPrefix predicate on the "client_name" field.
func ClientNameHasPrefix(v string) predicate.DeviceInfo {
	return predicate.DeviceInfo(sql.FieldHasPrefix(FieldClientName, v))
}

// ClientNameHasSuffix applies the HasSuffix predicate on the "client_name" field.
func ClientNameHasSuffix(v string) predicate.DeviceInfo {
	return predicate.DeviceInfo(sql.FieldHasSuffix(FieldClientName, v))
}

// ClientNameEqualFold applies the EqualFold predicate on the "client_name" field.
func ClientNameEqualFold(v string) predicate.DeviceInfo {
	return predicate.DeviceInfo(sql.FieldEqualFold(FieldClientName, v))
}

// ClientNameContainsFold applies the ContainsFold predicate on the "client_name" field.
func ClientNameContainsFold(v string) predicate.DeviceInfo {
	return predicate.DeviceInfo(sql.FieldContainsFold(FieldClientName, v))
}

// ClientSourceCodeAddressEQ applies the EQ predicate on the "client_source_code_address" field.
func ClientSourceCodeAddressEQ(v string) predicate.DeviceInfo {
	return predicate.DeviceInfo(sql.FieldEQ(FieldClientSourceCodeAddress, v))
}

// ClientSourceCodeAddressNEQ applies the NEQ predicate on the "client_source_code_address" field.
func ClientSourceCodeAddressNEQ(v string) predicate.DeviceInfo {
	return predicate.DeviceInfo(sql.FieldNEQ(FieldClientSourceCodeAddress, v))
}

// ClientSourceCodeAddressIn applies the In predicate on the "client_source_code_address" field.
func ClientSourceCodeAddressIn(vs ...string) predicate.DeviceInfo {
	return predicate.DeviceInfo(sql.FieldIn(FieldClientSourceCodeAddress, vs...))
}

// ClientSourceCodeAddressNotIn applies the NotIn predicate on the "client_source_code_address" field.
func ClientSourceCodeAddressNotIn(vs ...string) predicate.DeviceInfo {
	return predicate.DeviceInfo(sql.FieldNotIn(FieldClientSourceCodeAddress, vs...))
}

// ClientSourceCodeAddressGT applies the GT predicate on the "client_source_code_address" field.
func ClientSourceCodeAddressGT(v string) predicate.DeviceInfo {
	return predicate.DeviceInfo(sql.FieldGT(FieldClientSourceCodeAddress, v))
}

// ClientSourceCodeAddressGTE applies the GTE predicate on the "client_source_code_address" field.
func ClientSourceCodeAddressGTE(v string) predicate.DeviceInfo {
	return predicate.DeviceInfo(sql.FieldGTE(FieldClientSourceCodeAddress, v))
}

// ClientSourceCodeAddressLT applies the LT predicate on the "client_source_code_address" field.
func ClientSourceCodeAddressLT(v string) predicate.DeviceInfo {
	return predicate.DeviceInfo(sql.FieldLT(FieldClientSourceCodeAddress, v))
}

// ClientSourceCodeAddressLTE applies the LTE predicate on the "client_source_code_address" field.
func ClientSourceCodeAddressLTE(v string) predicate.DeviceInfo {
	return predicate.DeviceInfo(sql.FieldLTE(FieldClientSourceCodeAddress, v))
}

// ClientSourceCodeAddressContains applies the Contains predicate on the "client_source_code_address" field.
func ClientSourceCodeAddressContains(v string) predicate.DeviceInfo {
	return predicate.DeviceInfo(sql.FieldContains(FieldClientSourceCodeAddress, v))
}

// ClientSourceCodeAddressHasPrefix applies the HasPrefix predicate on the "client_source_code_address" field.
func ClientSourceCodeAddressHasPrefix(v string) predicate.DeviceInfo {
	return predicate.DeviceInfo(sql.FieldHasPrefix(FieldClientSourceCodeAddress, v))
}

// ClientSourceCodeAddressHasSuffix applies the HasSuffix predicate on the "client_source_code_address" field.
func ClientSourceCodeAddressHasSuffix(v string) predicate.DeviceInfo {
	return predicate.DeviceInfo(sql.FieldHasSuffix(FieldClientSourceCodeAddress, v))
}

// ClientSourceCodeAddressEqualFold applies the EqualFold predicate on the "client_source_code_address" field.
func ClientSourceCodeAddressEqualFold(v string) predicate.DeviceInfo {
	return predicate.DeviceInfo(sql.FieldEqualFold(FieldClientSourceCodeAddress, v))
}

// ClientSourceCodeAddressContainsFold applies the ContainsFold predicate on the "client_source_code_address" field.
func ClientSourceCodeAddressContainsFold(v string) predicate.DeviceInfo {
	return predicate.DeviceInfo(sql.FieldContainsFold(FieldClientSourceCodeAddress, v))
}

// ClientVersionEQ applies the EQ predicate on the "client_version" field.
func ClientVersionEQ(v string) predicate.DeviceInfo {
	return predicate.DeviceInfo(sql.FieldEQ(FieldClientVersion, v))
}

// ClientVersionNEQ applies the NEQ predicate on the "client_version" field.
func ClientVersionNEQ(v string) predicate.DeviceInfo {
	return predicate.DeviceInfo(sql.FieldNEQ(FieldClientVersion, v))
}

// ClientVersionIn applies the In predicate on the "client_version" field.
func ClientVersionIn(vs ...string) predicate.DeviceInfo {
	return predicate.DeviceInfo(sql.FieldIn(FieldClientVersion, vs...))
}

// ClientVersionNotIn applies the NotIn predicate on the "client_version" field.
func ClientVersionNotIn(vs ...string) predicate.DeviceInfo {
	return predicate.DeviceInfo(sql.FieldNotIn(FieldClientVersion, vs...))
}

// ClientVersionGT applies the GT predicate on the "client_version" field.
func ClientVersionGT(v string) predicate.DeviceInfo {
	return predicate.DeviceInfo(sql.FieldGT(FieldClientVersion, v))
}

// ClientVersionGTE applies the GTE predicate on the "client_version" field.
func ClientVersionGTE(v string) predicate.DeviceInfo {
	return predicate.DeviceInfo(sql.FieldGTE(FieldClientVersion, v))
}

// ClientVersionLT applies the LT predicate on the "client_version" field.
func ClientVersionLT(v string) predicate.DeviceInfo {
	return predicate.DeviceInfo(sql.FieldLT(FieldClientVersion, v))
}

// ClientVersionLTE applies the LTE predicate on the "client_version" field.
func ClientVersionLTE(v string) predicate.DeviceInfo {
	return predicate.DeviceInfo(sql.FieldLTE(FieldClientVersion, v))
}

// ClientVersionContains applies the Contains predicate on the "client_version" field.
func ClientVersionContains(v string) predicate.DeviceInfo {
	return predicate.DeviceInfo(sql.FieldContains(FieldClientVersion, v))
}

// ClientVersionHasPrefix applies the HasPrefix predicate on the "client_version" field.
func ClientVersionHasPrefix(v string) predicate.DeviceInfo {
	return predicate.DeviceInfo(sql.FieldHasPrefix(FieldClientVersion, v))
}

// ClientVersionHasSuffix applies the HasSuffix predicate on the "client_version" field.
func ClientVersionHasSuffix(v string) predicate.DeviceInfo {
	return predicate.DeviceInfo(sql.FieldHasSuffix(FieldClientVersion, v))
}

// ClientVersionEqualFold applies the EqualFold predicate on the "client_version" field.
func ClientVersionEqualFold(v string) predicate.DeviceInfo {
	return predicate.DeviceInfo(sql.FieldEqualFold(FieldClientVersion, v))
}

// ClientVersionContainsFold applies the ContainsFold predicate on the "client_version" field.
func ClientVersionContainsFold(v string) predicate.DeviceInfo {
	return predicate.DeviceInfo(sql.FieldContainsFold(FieldClientVersion, v))
}

// ClientLocalIDEQ applies the EQ predicate on the "client_local_id" field.
func ClientLocalIDEQ(v string) predicate.DeviceInfo {
	return predicate.DeviceInfo(sql.FieldEQ(FieldClientLocalID, v))
}

// ClientLocalIDNEQ applies the NEQ predicate on the "client_local_id" field.
func ClientLocalIDNEQ(v string) predicate.DeviceInfo {
	return predicate.DeviceInfo(sql.FieldNEQ(FieldClientLocalID, v))
}

// ClientLocalIDIn applies the In predicate on the "client_local_id" field.
func ClientLocalIDIn(vs ...string) predicate.DeviceInfo {
	return predicate.DeviceInfo(sql.FieldIn(FieldClientLocalID, vs...))
}

// ClientLocalIDNotIn applies the NotIn predicate on the "client_local_id" field.
func ClientLocalIDNotIn(vs ...string) predicate.DeviceInfo {
	return predicate.DeviceInfo(sql.FieldNotIn(FieldClientLocalID, vs...))
}

// ClientLocalIDGT applies the GT predicate on the "client_local_id" field.
func ClientLocalIDGT(v string) predicate.DeviceInfo {
	return predicate.DeviceInfo(sql.FieldGT(FieldClientLocalID, v))
}

// ClientLocalIDGTE applies the GTE predicate on the "client_local_id" field.
func ClientLocalIDGTE(v string) predicate.DeviceInfo {
	return predicate.DeviceInfo(sql.FieldGTE(FieldClientLocalID, v))
}

// ClientLocalIDLT applies the LT predicate on the "client_local_id" field.
func ClientLocalIDLT(v string) predicate.DeviceInfo {
	return predicate.DeviceInfo(sql.FieldLT(FieldClientLocalID, v))
}

// ClientLocalIDLTE applies the LTE predicate on the "client_local_id" field.
func ClientLocalIDLTE(v string) predicate.DeviceInfo {
	return predicate.DeviceInfo(sql.FieldLTE(FieldClientLocalID, v))
}

// ClientLocalIDContains applies the Contains predicate on the "client_local_id" field.
func ClientLocalIDContains(v string) predicate.DeviceInfo {
	return predicate.DeviceInfo(sql.FieldContains(FieldClientLocalID, v))
}

// ClientLocalIDHasPrefix applies the HasPrefix predicate on the "client_local_id" field.
func ClientLocalIDHasPrefix(v string) predicate.DeviceInfo {
	return predicate.DeviceInfo(sql.FieldHasPrefix(FieldClientLocalID, v))
}

// ClientLocalIDHasSuffix applies the HasSuffix predicate on the "client_local_id" field.
func ClientLocalIDHasSuffix(v string) predicate.DeviceInfo {
	return predicate.DeviceInfo(sql.FieldHasSuffix(FieldClientLocalID, v))
}

// ClientLocalIDIsNil applies the IsNil predicate on the "client_local_id" field.
func ClientLocalIDIsNil() predicate.DeviceInfo {
	return predicate.DeviceInfo(sql.FieldIsNull(FieldClientLocalID))
}

// ClientLocalIDNotNil applies the NotNil predicate on the "client_local_id" field.
func ClientLocalIDNotNil() predicate.DeviceInfo {
	return predicate.DeviceInfo(sql.FieldNotNull(FieldClientLocalID))
}

// ClientLocalIDEqualFold applies the EqualFold predicate on the "client_local_id" field.
func ClientLocalIDEqualFold(v string) predicate.DeviceInfo {
	return predicate.DeviceInfo(sql.FieldEqualFold(FieldClientLocalID, v))
}

// ClientLocalIDContainsFold applies the ContainsFold predicate on the "client_local_id" field.
func ClientLocalIDContainsFold(v string) predicate.DeviceInfo {
	return predicate.DeviceInfo(sql.FieldContainsFold(FieldClientLocalID, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.DeviceInfo {
	return predicate.DeviceInfo(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.DeviceInfo {
	return predicate.DeviceInfo(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.DeviceInfo {
	return predicate.DeviceInfo(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.DeviceInfo {
	return predicate.DeviceInfo(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.DeviceInfo {
	return predicate.DeviceInfo(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.DeviceInfo {
	return predicate.DeviceInfo(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.DeviceInfo {
	return predicate.DeviceInfo(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.DeviceInfo {
	return predicate.DeviceInfo(sql.FieldLTE(FieldUpdatedAt, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.DeviceInfo {
	return predicate.DeviceInfo(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.DeviceInfo {
	return predicate.DeviceInfo(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.DeviceInfo {
	return predicate.DeviceInfo(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.DeviceInfo {
	return predicate.DeviceInfo(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.DeviceInfo {
	return predicate.DeviceInfo(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.DeviceInfo {
	return predicate.DeviceInfo(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.DeviceInfo {
	return predicate.DeviceInfo(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.DeviceInfo {
	return predicate.DeviceInfo(sql.FieldLTE(FieldCreatedAt, v))
}

// HasUser applies the HasEdge predicate on the "user" edge.
func HasUser() predicate.DeviceInfo {
	return predicate.DeviceInfo(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, UserTable, UserPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUserWith applies the HasEdge predicate on the "user" edge with a given conditions (other predicates).
func HasUserWith(preds ...predicate.User) predicate.DeviceInfo {
	return predicate.DeviceInfo(func(s *sql.Selector) {
		step := newUserStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasUserSession applies the HasEdge predicate on the "user_session" edge.
func HasUserSession() predicate.DeviceInfo {
	return predicate.DeviceInfo(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, UserSessionTable, UserSessionColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUserSessionWith applies the HasEdge predicate on the "user_session" edge with a given conditions (other predicates).
func HasUserSessionWith(preds ...predicate.UserSession) predicate.DeviceInfo {
	return predicate.DeviceInfo(func(s *sql.Selector) {
		step := newUserSessionStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasUserDevice applies the HasEdge predicate on the "user_device" edge.
func HasUserDevice() predicate.DeviceInfo {
	return predicate.DeviceInfo(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, UserDeviceTable, UserDeviceColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUserDeviceWith applies the HasEdge predicate on the "user_device" edge with a given conditions (other predicates).
func HasUserDeviceWith(preds ...predicate.UserDevice) predicate.DeviceInfo {
	return predicate.DeviceInfo(func(s *sql.Selector) {
		step := newUserDeviceStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.DeviceInfo) predicate.DeviceInfo {
	return predicate.DeviceInfo(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.DeviceInfo) predicate.DeviceInfo {
	return predicate.DeviceInfo(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.DeviceInfo) predicate.DeviceInfo {
	return predicate.DeviceInfo(sql.NotPredicates(p))
}
