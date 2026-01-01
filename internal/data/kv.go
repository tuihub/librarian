package data

import (
	"context"
	"strconv"

	"github.com/tuihub/librarian/internal/data/internal/gormschema"
	"github.com/tuihub/librarian/internal/lib/libcodec"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func (d *Data) kvSet(ctx context.Context, bucket, key, value string) error {
	kv := gormschema.KV{
		Bucket: bucket,
		Key:    key,
		Value:  value,
	}
	return d.db.WithContext(ctx).Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "bucket"}, {Name: "key"}},
		DoUpdates: clause.AssignmentColumns([]string{"value", "updated_at"}),
	}).Create(&kv).Error
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
	var kv gormschema.KV
	err := d.db.WithContext(ctx).
		Where("bucket = ? AND key = ?", bucket, key).
		First(&kv).Error
	if err != nil {
		return "", err
	}
	return kv.Value, nil
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
		if err == gorm.ErrRecordNotFound {
			return false, nil
		}
		return false, err
	}
	return true, nil
}
