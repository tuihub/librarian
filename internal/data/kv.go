package data

import (
	"context"
	"strconv"

	"github.com/tuihub/librarian/internal/data/orm/model"
	"github.com/tuihub/librarian/internal/data/orm/query"
	"github.com/tuihub/librarian/internal/lib/libcodec"

	"gorm.io/gorm/clause"
)

func (d *Data) kvSet(ctx context.Context, bucket, key, value string) error {
	k := query.KV
	return k.WithContext(ctx).Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "bucket"}, {Name: "key"}},
		DoUpdates: clause.AssignmentColumns([]string{"value"}),
	}).Create(&model.KV{
		Bucket: bucket,
		Key:    key,
		Value:  value,
	})
}

func (d *Data) kvSetJSON(ctx context.Context, bucket, key string, value any) error {
	v, err := libcodec.Marshal(libcodec.JSON, value)
	if err != nil {
		return err
	}
	return d.kvSet(ctx, bucket, key, string(v))
}

func (d *Data) kvSetInt64(ctx context.Context, bucket, key string, value int64) error { //nolint:unused // no need
	return d.kvSet(ctx, bucket, key, strconv.FormatInt(value, 10))
}

func (d *Data) kvGet(ctx context.Context, bucket, key string) (string, error) {
	k := query.KV
	res, err := k.WithContext(ctx).
		Where(k.Bucket.Eq(bucket), k.Key.Eq(key)).
		First()
	if err != nil {
		return "", err
	}
	return res.Value, nil
}

func (d *Data) kvGetJSON(ctx context.Context, bucket, key string, value any) error {
	res, err := d.kvGet(ctx, bucket, key)
	if err != nil {
		return err
	}
	return libcodec.Unmarshal(libcodec.JSON, []byte(res), value)
}

func (d *Data) kvGetInt64(ctx context.Context, bucket, key string) (int64, error) { //nolint:unused // no need
	res, err := d.kvGet(ctx, bucket, key)
	if err != nil {
		return 0, err
	}
	v, err := strconv.ParseInt(res, 10, 64)
	if err != nil {
		return 0, err
	}
	return v, nil
}

func (d *Data) kvExists(ctx context.Context, bucket, key string) (bool, error) {
	_, err := d.kvGet(ctx, bucket, key)
	if err != nil {
		if ErrorIsNotFound(err) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}
