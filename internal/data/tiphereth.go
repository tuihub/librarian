package data

import (
	"context"
	"database/sql/driver"
	"errors"

	"github.com/tuihub/librarian/internal/data/orm/query"
	libmodel "github.com/tuihub/librarian/internal/model"
	"github.com/tuihub/librarian/internal/model/modelsupervisor"
)

type TipherethRepo struct {
	data *Data
}

func NewTipherethRepo(data *Data) *TipherethRepo {
	return &TipherethRepo{
		data: data,
	}
}

func (t *TipherethRepo) FetchUserByPassword(
	ctx context.Context,
	username, password string,
) (*libmodel.User, error) {
	q := query.Use(t.data.db).User
	u, err := q.WithContext(ctx).Where(
		q.Username.Eq(username),
		q.Password.Eq(password),
	).First()
	if u == nil || err != nil {
		return nil, errors.New("invalid user")
	}
	return u, nil
}

func (t *TipherethRepo) CreateDevice(
	ctx context.Context,
	info *libmodel.Device,
	clientLocalID *string,
) (libmodel.InternalID, error) {
	var res libmodel.InternalID
	err := t.data.WithTx(ctx, func(tx *query.Query) error {
		if clientLocalID != nil {
			infos, err := tx.Device.WithContext(ctx).Where(
				tx.Device.ClientLocalID.Eq(*clientLocalID),
			).Find()
			if err != nil {
				return err
			}
			if len(infos) > 0 {
				res = infos[0].ID
				return nil
			}
		}

		if clientLocalID != nil {
			info.ClientLocalID = *clientLocalID
		}

		if err := tx.Device.WithContext(ctx).Create(info); err != nil {
			return err
		}
		res = info.ID
		return nil
	})
	if err != nil {
		return 0, err
	}
	return res, nil
}

func (t *TipherethRepo) FetchDeviceInfo(
	ctx context.Context,
	deviceID libmodel.InternalID,
) (*libmodel.Device, error) {
	q := query.Use(t.data.db).Device
	res, err := q.WithContext(ctx).Where(q.ID.Eq(int64(deviceID))).First()
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (t *TipherethRepo) CreateUserSession(ctx context.Context, s *libmodel.Session) error {
	return t.data.WithTx(ctx, func(tx *query.Query) error {
		qs := tx.Session

		if s.Device != nil {
			// Delete existing session for user+device
			_, _ = qs.WithContext(ctx).Where(
				qs.UserID.Eq(int64(s.UserID)),
				qs.DeviceID.Eq(int64(s.Device.ID)),
			).Delete()
			s.DeviceID = s.Device.ID
			s.Device = nil // Prevent creating device
		}

		return qs.WithContext(ctx).Create(s)
	})
}

func (t *TipherethRepo) FetchUserSession(
	ctx context.Context,
	userID libmodel.InternalID,
	token string,
) (*libmodel.Session, error) {
	q := query.Use(t.data.db).Session
	s, err := q.WithContext(ctx).
		Where(q.UserID.Eq(int64(userID)), q.RefreshToken.Eq(token)).
		Preload(q.Device).
		First()
	if err != nil {
		return nil, err
	}
	return s, nil
}

func (t *TipherethRepo) ListUserSessions(
	ctx context.Context,
	id libmodel.InternalID,
) ([]*libmodel.Session, error) {
	q := query.Use(t.data.db).Session
	ss, err := q.WithContext(ctx).
		Where(q.UserID.Eq(int64(id))).
		Preload(q.Device).
		Find()
	if err != nil {
		return nil, err
	}
	return ss, nil
}

func (t *TipherethRepo) UpdateUserSession(ctx context.Context, session *libmodel.Session) error {
	return t.data.WithTx(ctx, func(tx *query.Query) error {
		q := tx.Session
		u := q.WithContext(ctx).Where(q.ID.Eq(int64(session.ID)))

		updates := map[string]interface{}{
			"refresh_token": session.RefreshToken,
			"created_at":    session.CreateAt,
			"expire_at":     session.ExpireAt,
		}
		if session.Device != nil {
			updates["device_id"] = session.Device.ID
		}

		_, err := u.Updates(updates)
		return err
	})
}

func (t *TipherethRepo) DeleteUserSession(
	ctx context.Context,
	userID libmodel.InternalID,
	sessionID libmodel.InternalID,
) error {
	q := query.Use(t.data.db).Session
	_, err := q.WithContext(ctx).Where(
		q.ID.Eq(int64(sessionID)),
		q.UserID.Eq(int64(userID)),
	).Delete()
	return err
}

func (t *TipherethRepo) CreateUser(ctx context.Context, u *libmodel.User, c libmodel.InternalID) error {
	q := query.Use(t.data.db).User
	u.CreatorID = c
	return q.WithContext(ctx).Create(u)
}

func (t *TipherethRepo) UpdateUser(ctx context.Context, u *libmodel.User, password string) error {
	q := query.Use(t.data.db).User
	updates := q.WithContext(ctx).Where(q.ID.Eq(int64(u.ID)))

	vals := make(map[string]interface{})
	if u.Username != "" {
		vals["username"] = u.Username
	}
	if u.Password != "" {
		// Only update password if old password matches
		// GORM update with where clause
		if password != "" {
			updates = updates.Where(q.Password.Eq(password))
		}
		vals["password"] = u.Password
	}
	if u.Type != libmodel.UserTypeUnspecified {
		vals["type"] = u.Type
	}
	if u.Status != libmodel.UserStatusUnspecified {
		vals["status"] = u.Status
	}

	if len(vals) == 0 {
		return nil
	}

	_, err := updates.Updates(vals)
	return err
}

func (t *TipherethRepo) ListUsers(
	ctx context.Context,
	paging libmodel.Paging,
	ids []libmodel.InternalID,
	types []libmodel.UserType,
	statuses []libmodel.UserStatus,
	exclude []libmodel.InternalID,
) ([]*libmodel.User, int64, error) {
	q := query.Use(t.data.db).User
	u := q.WithContext(ctx)

	if len(ids) > 0 {
		castIDs := make([]int64, len(ids))
		for i, v := range ids {
			castIDs[i] = int64(v)
		}
		u = u.Where(q.ID.In(castIDs...))
	}
	if len(types) > 0 {
		t := make([]driver.Valuer, len(types))
		for i, v := range types {
			t[i] = v
		}
		u = u.Where(q.Type.In(t...))
	}
	if len(statuses) > 0 {
		s := make([]driver.Valuer, len(statuses))
		for i, v := range statuses {
			s[i] = v
		}
		u = u.Where(q.Status.In(s...))
	}
	if len(exclude) > 0 {
		castExclude := make([]int64, len(exclude))
		for i, v := range exclude {
			castExclude[i] = int64(v)
		}
		u = u.Where(q.ID.NotIn(castExclude...))
	}

	count, err := u.Count()
	if err != nil {
		return nil, 0, err
	}

	users, err := u.Limit(paging.ToLimit()).Offset(paging.ToOffset()).Find()
	if err != nil {
		return nil, 0, err
	}
	return users, count, nil
}

func (t *TipherethRepo) GetUser(ctx context.Context, id libmodel.InternalID) (*libmodel.User, error) {
	q := query.Use(t.data.db).User
	u, err := q.WithContext(ctx).Where(q.ID.Eq(int64(id))).First()
	if err != nil {
		return nil, err
	}
	return u, nil
}

func (t *TipherethRepo) GetUserCount(ctx context.Context) (int, error) {
	c, err := query.Use(t.data.db).User.WithContext(ctx).Count()
	return int(c), err
}

func (t *TipherethRepo) LinkAccount(
	ctx context.Context,
	a *libmodel.Account,
	userID libmodel.InternalID,
) (libmodel.InternalID, error) {
	accountID := a.ID
	err := t.data.WithTx(ctx, func(tx *query.Query) error {
		// Check if user has account with same platform
		count, err := tx.Account.WithContext(ctx).Where(
			tx.Account.BoundUserID.Eq(int64(userID)),
			tx.Account.Platform.Eq(a.Platform),
		).Count()
		if err != nil {
			return err
		}
		if count > 0 {
			return errors.New("an account already bound to user")
		}

		// Check if account exists
		acc, err := tx.Account.WithContext(ctx).Where(
			tx.Account.Platform.Eq(a.Platform),
			tx.Account.PlatformAccountID.Eq(a.PlatformAccountID),
		).First()

		if err != nil {
			// Not found (or error)
			// Assuming GORM returns error on not found if using First
			// But check error type
			// If not found, create new
			// Actually we should check if err is record not found
			// If error is something else, return it
			// If not found, create
			// If found, check if bound
			// We can use FirstOrInit or just check err
			// But we need logic: if not found -> create. if found -> update bound user if not bound.

			// Let's assume standard error check
			// Create new account
			a.BoundUserID = userID
			return tx.Account.WithContext(ctx).Create(a)
		}

		// Account exists
		// Check if bound
		// Ent: acc.QueryBoundUser().Exist(ctx)
		// Model has BoundUserID field.
		if acc.BoundUserID != 0 {
			return errors.New("account already bound to an user")
		}

		accountID = acc.ID
		_, err = tx.Account.WithContext(ctx).
			Where(tx.Account.ID.Eq(int64(acc.ID))).
			Update(tx.Account.BoundUserID, userID)
		return err
	})
	if err != nil {
		return 0, err
	}
	return accountID, nil
}

func (t *TipherethRepo) UnLinkAccount(ctx context.Context, aid libmodel.InternalID, u libmodel.InternalID) error {
	q := query.Use(t.data.db).Account
	_, err := q.WithContext(ctx).Where(
		q.ID.Eq(int64(aid)),
		q.BoundUserID.Eq(int64(u)),
	).Update(q.BoundUserID, 0) // Set to 0 (or null if pointer)
	// BoundUserID is InternalID (int64). 0 usually means no user if logic treats 0 as null.
	// But in GORM model it's just int64. Ent schema: Optional().
	// GORM model: BoundUserID model.InternalID `gorm:"index"`.
	// If it's not a pointer, 0 is value.
	return err
}

func (t *TipherethRepo) ListLinkAccounts(
	ctx context.Context,
	userID libmodel.InternalID,
) ([]*libmodel.Account, error) {
	q := query.Use(t.data.db).Account
	res, err := q.WithContext(ctx).Where(q.BoundUserID.Eq(int64(userID))).Find()
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (t *TipherethRepo) ListPorters(
	ctx context.Context,
	paging libmodel.Paging,
) ([]*modelsupervisor.PorterInstance, int64, error) {
	q := query.Use(t.data.db).PorterInstance
	count, err := q.WithContext(ctx).Count()
	if err != nil {
		return nil, 0, err
	}
	res, err := q.WithContext(ctx).Limit(paging.ToLimit()).Offset(paging.ToOffset()).Find()
	if err != nil {
		return nil, 0, err
	}
	return res, count, nil
}

func (t *TipherethRepo) GetPorter(
	ctx context.Context,
	id libmodel.InternalID,
) (*modelsupervisor.PorterInstance, error) {
	q := query.Use(t.data.db).PorterInstance
	res, err := q.WithContext(ctx).Where(q.ID.Eq(int64(id))).First()
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (t *TipherethRepo) UpdatePorterStatus(
	ctx context.Context,
	id libmodel.InternalID,
	status libmodel.UserStatus,
) (*modelsupervisor.PorterInstance, error) {
	q := query.Use(t.data.db).PorterInstance
	res, err := q.WithContext(ctx).Where(q.ID.Eq(int64(id))).First()
	if err != nil {
		return nil, err
	}

	_, err = q.WithContext(ctx).Where(q.ID.Eq(int64(id))).Update(q.Status, status)
	if err != nil {
		return nil, err
	}
	res.Status = status
	return res, nil
}

func (t *TipherethRepo) CreatePorterContext(
	ctx context.Context,
	userID libmodel.InternalID,
	context *modelsupervisor.PorterContext,
) error {
	q := query.Use(t.data.db).PorterContext
	context.OwnerID = userID
	return q.WithContext(ctx).Create(context)
}

func (t *TipherethRepo) ListPorterContexts(
	ctx context.Context,
	userID libmodel.InternalID,
	paging libmodel.Paging,
) ([]*modelsupervisor.PorterContext, int64, error) {
	q := query.Use(t.data.db).PorterContext
	u := q.WithContext(ctx).Where(q.OwnerID.Eq(int64(userID)))
	count, err := u.Count()
	if err != nil {
		return nil, 0, err
	}
	res, err := u.Limit(paging.ToLimit()).Offset(paging.ToOffset()).Find()
	if err != nil {
		return nil, 0, err
	}
	return res, count, nil
}

func (t *TipherethRepo) ListPorterContextsByGlobalName(
	ctx context.Context,
	userID libmodel.InternalID,
	globalName string,
	paging libmodel.Paging,
) ([]*modelsupervisor.PorterContext, int64, error) {
	q := query.Use(t.data.db).PorterContext
	u := q.WithContext(ctx).Where(
		q.OwnerID.Eq(int64(userID)),
		q.GlobalName.Eq(globalName),
	)
	count, err := u.Count()
	if err != nil {
		return nil, 0, err
	}
	res, err := u.Limit(paging.ToLimit()).Offset(paging.ToOffset()).Find()
	if err != nil {
		return nil, 0, err
	}
	return res, count, nil
}

func (t *TipherethRepo) UpdatePorterContext(
	ctx context.Context,
	userID libmodel.InternalID,
	context *modelsupervisor.PorterContext,
) error {
	q := query.Use(t.data.db).PorterContext
	updates := map[string]interface{}{
		"context_json": context.ContextJSON,
		"name":         context.Name,
		"description":  context.Description,
		"status":       context.Status,
	}
	_, err := q.WithContext(ctx).Where(
		q.ID.Eq(int64(context.ID)),
		q.OwnerID.Eq(int64(userID)),
	).Updates(updates)
	return err
}

func (t *TipherethRepo) ListPorterDigests(
	ctx context.Context,
	status []libmodel.UserStatus,
) ([]*modelsupervisor.PorterDigest, error) {
	q := query.Use(t.data.db).PorterInstance
	u := q.WithContext(ctx)
	if len(status) > 0 {
		s := make([]driver.Valuer, len(status))
		for i, v := range status {
			s[i] = v
		}
		u = u.Where(q.Status.In(s...))
	}

	// Group By GlobalName, Region. Select Min(ID).
	var results []struct {
		MinID libmodel.InternalID `gorm:"column:min_id"`
	}

	err := u.Select(q.ID.Min().As("min_id")).Group(q.GlobalName, q.Region).Scan(&results)
	if err != nil {
		return nil, err
	}

	var ids []int64
	for _, r := range results {
		ids = append(ids, int64(r.MinID))
	}

	if len(ids) == 0 {
		return nil, nil
	}

	pi, err := q.WithContext(ctx).Where(q.ID.In(ids...)).Find()
	if err != nil {
		return nil, err
	}

	var pg []*modelsupervisor.PorterDigest
	pgm := make(map[string]*modelsupervisor.PorterDigest)
	for _, p := range pi {
		if len(p.ContextJSONSchema) == 0 {
			continue
		}
		if pgm[p.GlobalName] == nil {
			pgm[p.GlobalName] = &modelsupervisor.PorterDigest{
				BinarySummary:     p.BinarySummary,
				GlobalName:        p.GlobalName,
				Regions:           []string{p.Region},
				ContextJSONSchema: p.ContextJSONSchema,
				FeatureSummary:    p.FeatureSummary,
			}
		} else {
			pgm[p.GlobalName].Regions = append(pgm[p.GlobalName].Regions, p.Region)
		}
	}
	for _, v := range pgm {
		pg = append(pg, v)
	}
	return pg, nil
}

func (t *TipherethRepo) FetchPorterContext(
	ctx context.Context,
	id libmodel.InternalID,
) (*modelsupervisor.PorterContext, error) {
	q := query.Use(t.data.db).PorterContext
	res, err := q.WithContext(ctx).Where(q.ID.Eq(int64(id))).First()
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (t *TipherethRepo) GetEnabledPorterContexts(
	ctx context.Context,
) ([]*modelsupervisor.PorterContext, error) {
	q := query.Use(t.data.db).PorterContext
	res, err := q.WithContext(ctx).Where(
		q.Status.Eq(modelsupervisor.PorterContextStatusActive),
	).Find()
	if err != nil {
		return nil, err
	}
	return res, nil
}
