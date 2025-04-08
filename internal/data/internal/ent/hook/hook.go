// Code generated by ent, DO NOT EDIT.

package hook

import (
	"context"
	"fmt"

	"github.com/tuihub/librarian/internal/data/internal/ent"
)

// The AccountFunc type is an adapter to allow the use of ordinary
// function as Account mutator.
type AccountFunc func(context.Context, *ent.AccountMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f AccountFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.AccountMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.AccountMutation", m)
}

// The AppFunc type is an adapter to allow the use of ordinary
// function as App mutator.
type AppFunc func(context.Context, *ent.AppMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f AppFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.AppMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.AppMutation", m)
}

// The AppAppCategoryFunc type is an adapter to allow the use of ordinary
// function as AppAppCategory mutator.
type AppAppCategoryFunc func(context.Context, *ent.AppAppCategoryMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f AppAppCategoryFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.AppAppCategoryMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.AppAppCategoryMutation", m)
}

// The AppCategoryFunc type is an adapter to allow the use of ordinary
// function as AppCategory mutator.
type AppCategoryFunc func(context.Context, *ent.AppCategoryMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f AppCategoryFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.AppCategoryMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.AppCategoryMutation", m)
}

// The AppInfoFunc type is an adapter to allow the use of ordinary
// function as AppInfo mutator.
type AppInfoFunc func(context.Context, *ent.AppInfoMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f AppInfoFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.AppInfoMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.AppInfoMutation", m)
}

// The AppRunTimeFunc type is an adapter to allow the use of ordinary
// function as AppRunTime mutator.
type AppRunTimeFunc func(context.Context, *ent.AppRunTimeMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f AppRunTimeFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.AppRunTimeMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.AppRunTimeMutation", m)
}

// The DeviceFunc type is an adapter to allow the use of ordinary
// function as Device mutator.
type DeviceFunc func(context.Context, *ent.DeviceMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f DeviceFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.DeviceMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.DeviceMutation", m)
}

// The FeedFunc type is an adapter to allow the use of ordinary
// function as Feed mutator.
type FeedFunc func(context.Context, *ent.FeedMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f FeedFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.FeedMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.FeedMutation", m)
}

// The FeedActionSetFunc type is an adapter to allow the use of ordinary
// function as FeedActionSet mutator.
type FeedActionSetFunc func(context.Context, *ent.FeedActionSetMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f FeedActionSetFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.FeedActionSetMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.FeedActionSetMutation", m)
}

// The FeedConfigFunc type is an adapter to allow the use of ordinary
// function as FeedConfig mutator.
type FeedConfigFunc func(context.Context, *ent.FeedConfigMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f FeedConfigFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.FeedConfigMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.FeedConfigMutation", m)
}

// The FeedConfigActionFunc type is an adapter to allow the use of ordinary
// function as FeedConfigAction mutator.
type FeedConfigActionFunc func(context.Context, *ent.FeedConfigActionMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f FeedConfigActionFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.FeedConfigActionMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.FeedConfigActionMutation", m)
}

// The FeedItemFunc type is an adapter to allow the use of ordinary
// function as FeedItem mutator.
type FeedItemFunc func(context.Context, *ent.FeedItemMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f FeedItemFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.FeedItemMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.FeedItemMutation", m)
}

// The FeedItemCollectionFunc type is an adapter to allow the use of ordinary
// function as FeedItemCollection mutator.
type FeedItemCollectionFunc func(context.Context, *ent.FeedItemCollectionMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f FeedItemCollectionFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.FeedItemCollectionMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.FeedItemCollectionMutation", m)
}

// The FileFunc type is an adapter to allow the use of ordinary
// function as File mutator.
type FileFunc func(context.Context, *ent.FileMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f FileFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.FileMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.FileMutation", m)
}

// The ImageFunc type is an adapter to allow the use of ordinary
// function as Image mutator.
type ImageFunc func(context.Context, *ent.ImageMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f ImageFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.ImageMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.ImageMutation", m)
}

// The NotifyFlowFunc type is an adapter to allow the use of ordinary
// function as NotifyFlow mutator.
type NotifyFlowFunc func(context.Context, *ent.NotifyFlowMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f NotifyFlowFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.NotifyFlowMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.NotifyFlowMutation", m)
}

// The NotifyFlowSourceFunc type is an adapter to allow the use of ordinary
// function as NotifyFlowSource mutator.
type NotifyFlowSourceFunc func(context.Context, *ent.NotifyFlowSourceMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f NotifyFlowSourceFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.NotifyFlowSourceMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.NotifyFlowSourceMutation", m)
}

// The NotifyFlowTargetFunc type is an adapter to allow the use of ordinary
// function as NotifyFlowTarget mutator.
type NotifyFlowTargetFunc func(context.Context, *ent.NotifyFlowTargetMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f NotifyFlowTargetFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.NotifyFlowTargetMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.NotifyFlowTargetMutation", m)
}

// The NotifySourceFunc type is an adapter to allow the use of ordinary
// function as NotifySource mutator.
type NotifySourceFunc func(context.Context, *ent.NotifySourceMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f NotifySourceFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.NotifySourceMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.NotifySourceMutation", m)
}

// The NotifyTargetFunc type is an adapter to allow the use of ordinary
// function as NotifyTarget mutator.
type NotifyTargetFunc func(context.Context, *ent.NotifyTargetMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f NotifyTargetFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.NotifyTargetMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.NotifyTargetMutation", m)
}

// The PorterContextFunc type is an adapter to allow the use of ordinary
// function as PorterContext mutator.
type PorterContextFunc func(context.Context, *ent.PorterContextMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f PorterContextFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.PorterContextMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.PorterContextMutation", m)
}

// The PorterInstanceFunc type is an adapter to allow the use of ordinary
// function as PorterInstance mutator.
type PorterInstanceFunc func(context.Context, *ent.PorterInstanceMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f PorterInstanceFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.PorterInstanceMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.PorterInstanceMutation", m)
}

// The SentinelAppBinaryFunc type is an adapter to allow the use of ordinary
// function as SentinelAppBinary mutator.
type SentinelAppBinaryFunc func(context.Context, *ent.SentinelAppBinaryMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f SentinelAppBinaryFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.SentinelAppBinaryMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.SentinelAppBinaryMutation", m)
}

// The SentinelAppBinaryFileFunc type is an adapter to allow the use of ordinary
// function as SentinelAppBinaryFile mutator.
type SentinelAppBinaryFileFunc func(context.Context, *ent.SentinelAppBinaryFileMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f SentinelAppBinaryFileFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.SentinelAppBinaryFileMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.SentinelAppBinaryFileMutation", m)
}

// The SentinelInfoFunc type is an adapter to allow the use of ordinary
// function as SentinelInfo mutator.
type SentinelInfoFunc func(context.Context, *ent.SentinelInfoMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f SentinelInfoFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.SentinelInfoMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.SentinelInfoMutation", m)
}

// The SentinelLibraryFunc type is an adapter to allow the use of ordinary
// function as SentinelLibrary mutator.
type SentinelLibraryFunc func(context.Context, *ent.SentinelLibraryMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f SentinelLibraryFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.SentinelLibraryMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.SentinelLibraryMutation", m)
}

// The SessionFunc type is an adapter to allow the use of ordinary
// function as Session mutator.
type SessionFunc func(context.Context, *ent.SessionMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f SessionFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.SessionMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.SessionMutation", m)
}

// The StoreAppFunc type is an adapter to allow the use of ordinary
// function as StoreApp mutator.
type StoreAppFunc func(context.Context, *ent.StoreAppMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f StoreAppFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.StoreAppMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.StoreAppMutation", m)
}

// The StoreAppBinaryFunc type is an adapter to allow the use of ordinary
// function as StoreAppBinary mutator.
type StoreAppBinaryFunc func(context.Context, *ent.StoreAppBinaryMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f StoreAppBinaryFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.StoreAppBinaryMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.StoreAppBinaryMutation", m)
}

// The SystemNotificationFunc type is an adapter to allow the use of ordinary
// function as SystemNotification mutator.
type SystemNotificationFunc func(context.Context, *ent.SystemNotificationMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f SystemNotificationFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.SystemNotificationMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.SystemNotificationMutation", m)
}

// The TagFunc type is an adapter to allow the use of ordinary
// function as Tag mutator.
type TagFunc func(context.Context, *ent.TagMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f TagFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.TagMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.TagMutation", m)
}

// The UserFunc type is an adapter to allow the use of ordinary
// function as User mutator.
type UserFunc func(context.Context, *ent.UserMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f UserFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.UserMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.UserMutation", m)
}

// Condition is a hook condition function.
type Condition func(context.Context, ent.Mutation) bool

// And groups conditions with the AND operator.
func And(first, second Condition, rest ...Condition) Condition {
	return func(ctx context.Context, m ent.Mutation) bool {
		if !first(ctx, m) || !second(ctx, m) {
			return false
		}
		for _, cond := range rest {
			if !cond(ctx, m) {
				return false
			}
		}
		return true
	}
}

// Or groups conditions with the OR operator.
func Or(first, second Condition, rest ...Condition) Condition {
	return func(ctx context.Context, m ent.Mutation) bool {
		if first(ctx, m) || second(ctx, m) {
			return true
		}
		for _, cond := range rest {
			if cond(ctx, m) {
				return true
			}
		}
		return false
	}
}

// Not negates a given condition.
func Not(cond Condition) Condition {
	return func(ctx context.Context, m ent.Mutation) bool {
		return !cond(ctx, m)
	}
}

// HasOp is a condition testing mutation operation.
func HasOp(op ent.Op) Condition {
	return func(_ context.Context, m ent.Mutation) bool {
		return m.Op().Is(op)
	}
}

// HasAddedFields is a condition validating `.AddedField` on fields.
func HasAddedFields(field string, fields ...string) Condition {
	return func(_ context.Context, m ent.Mutation) bool {
		if _, exists := m.AddedField(field); !exists {
			return false
		}
		for _, field := range fields {
			if _, exists := m.AddedField(field); !exists {
				return false
			}
		}
		return true
	}
}

// HasClearedFields is a condition validating `.FieldCleared` on fields.
func HasClearedFields(field string, fields ...string) Condition {
	return func(_ context.Context, m ent.Mutation) bool {
		if exists := m.FieldCleared(field); !exists {
			return false
		}
		for _, field := range fields {
			if exists := m.FieldCleared(field); !exists {
				return false
			}
		}
		return true
	}
}

// HasFields is a condition validating `.Field` on fields.
func HasFields(field string, fields ...string) Condition {
	return func(_ context.Context, m ent.Mutation) bool {
		if _, exists := m.Field(field); !exists {
			return false
		}
		for _, field := range fields {
			if _, exists := m.Field(field); !exists {
				return false
			}
		}
		return true
	}
}

// If executes the given hook under condition.
//
//	hook.If(ComputeAverage, And(HasFields(...), HasAddedFields(...)))
func If(hk ent.Hook, cond Condition) ent.Hook {
	return func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			if cond(ctx, m) {
				return hk(next).Mutate(ctx, m)
			}
			return next.Mutate(ctx, m)
		})
	}
}

// On executes the given hook only for the given operation.
//
//	hook.On(Log, ent.Delete|ent.Create)
func On(hk ent.Hook, op ent.Op) ent.Hook {
	return If(hk, HasOp(op))
}

// Unless skips the given hook only for the given operation.
//
//	hook.Unless(Log, ent.Update|ent.UpdateOne)
func Unless(hk ent.Hook, op ent.Op) ent.Hook {
	return If(hk, Not(HasOp(op)))
}

// FixedError is a hook returning a fixed error.
func FixedError(err error) ent.Hook {
	return func(ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(context.Context, ent.Mutation) (ent.Value, error) {
			return nil, err
		})
	}
}

// Reject returns a hook that rejects all operations that match op.
//
//	func (T) Hooks() []ent.Hook {
//		return []ent.Hook{
//			Reject(ent.Delete|ent.Update),
//		}
//	}
func Reject(op ent.Op) ent.Hook {
	hk := FixedError(fmt.Errorf("%s operation is not allowed", op))
	return On(hk, op)
}

// Chain acts as a list of hooks and is effectively immutable.
// Once created, it will always hold the same set of hooks in the same order.
type Chain struct {
	hooks []ent.Hook
}

// NewChain creates a new chain of hooks.
func NewChain(hooks ...ent.Hook) Chain {
	return Chain{append([]ent.Hook(nil), hooks...)}
}

// Hook chains the list of hooks and returns the final hook.
func (c Chain) Hook() ent.Hook {
	return func(mutator ent.Mutator) ent.Mutator {
		for i := len(c.hooks) - 1; i >= 0; i-- {
			mutator = c.hooks[i](mutator)
		}
		return mutator
	}
}

// Append extends a chain, adding the specified hook
// as the last ones in the mutation flow.
func (c Chain) Append(hooks ...ent.Hook) Chain {
	newHooks := make([]ent.Hook, 0, len(c.hooks)+len(hooks))
	newHooks = append(newHooks, c.hooks...)
	newHooks = append(newHooks, hooks...)
	return Chain{newHooks}
}

// Extend extends a chain, adding the specified chain
// as the last ones in the mutation flow.
func (c Chain) Extend(chain Chain) Chain {
	return c.Append(chain.hooks...)
}
