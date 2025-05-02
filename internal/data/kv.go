package data

import (
	"context"
	"strconv"

	"github.com/tuihub/librarian/internal/data/internal/ent"
	"github.com/tuihub/librarian/internal/data/internal/ent/kv"
	"github.com/tuihub/librarian/internal/lib/libcodec"

	"entgo.io/ent/dialect/sql"
)

func (d *Data) kvSet(ctx context.Context, bucket, key, value string) error {
	err := d.db.KV.Create().
		SetBucket(bucket).
		SetKey(key).
		SetValue(value).
		OnConflict(
			sql.ConflictColumns(
				kv.FieldBucket,
				kv.FieldKey,
			),
			resolveWithIgnores([]string{
				kv.FieldBucket,
				kv.FieldKey,
			}),
		).
		Exec(ctx)
	return err
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
	res, err := d.db.KV.Query().
		Where(
			kv.Bucket(bucket),
			kv.Key(key),
		).
		Only(ctx)
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
		if ent.IsNotFound(err) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}
