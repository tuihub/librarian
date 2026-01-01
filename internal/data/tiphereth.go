package data

import (
	"context"
	"errors"

	"github.com/tuihub/librarian/internal/data/internal/gormschema"
	"github.com/tuihub/librarian/internal/model"
	"github.com/tuihub/librarian/internal/model/modelsupervisor"

	"gorm.io/gorm"
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
) (*model.User, error) {
	var u gormschema.User
	err := t.data.db.WithContext(ctx).
		Where("username = ? AND password = ?", username, password).
		First(&u).Error
	if err != nil {
		return nil, errors.New("invalid user")
	}
	res := gormschema.ToBizUser(&u)
	res.Password = u.Password
	return res, nil
}

func (t *TipherethRepo) CreateDevice(
	ctx context.Context,
	info *model.Device,
	clientLocalID *string,
) (model.InternalID, error) {
	var res model.InternalID
	err := t.data.WithTx(ctx, func(tx *gorm.DB) error {
		if clientLocalID != nil {
			var existing []gormschema.Device
			if err := tx.Where("client_local_id = ?", *clientLocalID).Find(&existing).Error; err != nil {
				return err
			}
			if len(existing) > 0 {
				res = existing[0].ID
				return nil
			}
		}
		device := gormschema.Device{
			ID:                      info.ID,
			DeviceName:              info.DeviceName,
			SystemType:              gormschema.ToSchemaSystemType(info.SystemType),
			SystemVersion:           info.SystemVersion,
			ClientName:              info.ClientName,
			ClientSourceCodeAddress: info.ClientSourceCodeAddress,
			ClientVersion:           info.ClientVersion,
			ClientLocalID:           clientLocalID,
		}
		if err := tx.Create(&device).Error; err != nil {
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
	deviceID model.InternalID,
) (*model.Device, error) {
	var device gormschema.Device
	if err := t.data.db.WithContext(ctx).First(&device, deviceID).Error; err != nil {
		return nil, err
	}
	return gormschema.ToBizDevice(&device), nil
}

func (t *TipherethRepo) CreateUserSession(ctx context.Context, s *model.Session) error {
	return t.data.WithTx(ctx, func(tx *gorm.DB) error {
		if s.Device != nil {
			// Delete existing session with same user and device
			tx.Where("user_id = ? AND device_id = ?", s.UserID, s.Device.ID).
				Delete(&gormschema.Session{})
		}
		session := gormschema.Session{
			ID:           s.ID,
			UserID:       s.UserID,
			RefreshToken: s.RefreshToken,
			ExpireAt:     s.ExpireAt,
		}
		if s.Device != nil {
			session.DeviceID = &s.Device.ID
		}
		return tx.Create(&session).Error
	})
}

func (t *TipherethRepo) FetchUserSession(
	ctx context.Context,
	userID model.InternalID,
	token string,
) (*model.Session, error) {
	var session gormschema.Session
	err := t.data.db.WithContext(ctx).
		Where("user_id = ? AND refresh_token = ?", userID, token).
		First(&session).Error
	if err != nil {
		return nil, err
	}
	res := gormschema.ToBizSession(&session)
	if session.DeviceID != nil {
		var device gormschema.Device
		if err := t.data.db.WithContext(ctx).First(&device, *session.DeviceID).Error; err == nil {
			res.Device = gormschema.ToBizDevice(&device)
		}
	}
	return res, nil
}

func (t *TipherethRepo) ListUserSessions(
	ctx context.Context,
	id model.InternalID,
) ([]*model.Session, error) {
	var sessions []gormschema.Session
	err := t.data.db.WithContext(ctx).
		Where("user_id = ?", id).
		Find(&sessions).Error
	if err != nil {
		return nil, err
	}
	res := make([]*model.Session, len(sessions))
	for i, s := range sessions {
		res[i] = gormschema.ToBizSession(&s)
		if s.DeviceID != nil {
			var device gormschema.Device
			if err := t.data.db.WithContext(ctx).First(&device, *s.DeviceID).Error; err == nil {
				res[i].Device = gormschema.ToBizDevice(&device)
			}
		}
	}
	return res, nil
}

func (t *TipherethRepo) UpdateUserSession(ctx context.Context, session *model.Session) error {
	return t.data.WithTx(ctx, func(tx *gorm.DB) error {
		updates := map[string]any{
			"refresh_token": session.RefreshToken,
			"expire_at":     session.ExpireAt,
		}
		if session.Device != nil {
			updates["device_id"] = session.Device.ID
		}
		return tx.Model(&gormschema.Session{}).
			Where("id = ?", session.ID).
			Updates(updates).Error
	})
}

func (t *TipherethRepo) DeleteUserSession(
	ctx context.Context,
	userID model.InternalID,
	sessionID model.InternalID,
) error {
	return t.data.db.WithContext(ctx).
		Where("id = ? AND user_id = ?", sessionID, userID).
		Delete(&gormschema.Session{}).Error
}

func (t *TipherethRepo) CreateUser(ctx context.Context, u *model.User, c model.InternalID) error {
	user := gormschema.User{
		ID:        u.ID,
		Username:  u.Username,
		Password:  u.Password,
		Status:    gormschema.ToSchemaUserStatus(u.Status),
		Type:      gormschema.ToSchemaUserType(u.Type),
		CreatorID: c,
	}
	return t.data.db.WithContext(ctx).Create(&user).Error
}

func (t *TipherethRepo) UpdateUser(ctx context.Context, u *model.User, password string) error {
	query := t.data.db.WithContext(ctx).Model(&gormschema.User{}).Where("id = ?", u.ID)
	updates := make(map[string]any)

	if u.Username != "" {
		updates["username"] = u.Username
	}
	if u.Password != "" {
		query = query.Where("password = ?", password)
		updates["password"] = u.Password
	}
	if u.Type != model.UserTypeUnspecified {
		updates["type"] = gormschema.ToSchemaUserType(u.Type)
	}
	if u.Status != model.UserStatusUnspecified {
		updates["status"] = gormschema.ToSchemaUserStatus(u.Status)
	}

	if len(updates) == 0 {
		return nil
	}
	return query.Updates(updates).Error
}

func (t *TipherethRepo) ListUsers(
	ctx context.Context,
	paging model.Paging,
	ids []model.InternalID,
	types []model.UserType,
	statuses []model.UserStatus,
	exclude []model.InternalID,
) ([]*model.User, int64, error) {
	query := t.data.db.WithContext(ctx).Model(&gormschema.User{})

	if len(ids) > 0 {
		query = query.Where("id IN ?", ids)
	}
	if len(types) > 0 {
		typeStrs := make([]string, len(types))
		for i, tp := range types {
			typeStrs[i] = gormschema.ToSchemaUserType(tp)
		}
		query = query.Where("type IN ?", typeStrs)
	}
	if len(statuses) > 0 {
		statusStrs := make([]string, len(statuses))
		for i, st := range statuses {
			statusStrs[i] = gormschema.ToSchemaUserStatus(st)
		}
		query = query.Where("status IN ?", statusStrs)
	}
	if len(exclude) > 0 {
		query = query.Where("id NOT IN ?", exclude)
	}

	var count int64
	if err := query.Count(&count).Error; err != nil {
		return nil, 0, err
	}

	var users []gormschema.User
	if err := query.Limit(paging.ToLimit()).Offset(paging.ToOffset()).Find(&users).Error; err != nil {
		return nil, 0, err
	}

	res := make([]*model.User, len(users))
	for i := range users {
		res[i] = gormschema.ToBizUser(&users[i])
	}
	return res, count, nil
}

func (t *TipherethRepo) GetUser(ctx context.Context, id model.InternalID) (*model.User, error) {
	var u gormschema.User
	if err := t.data.db.WithContext(ctx).First(&u, id).Error; err != nil {
		return nil, err
	}
	return gormschema.ToBizUser(&u), nil
}

func (t *TipherethRepo) GetUserCount(ctx context.Context) (int, error) {
	var count int64
	if err := t.data.db.WithContext(ctx).Model(&gormschema.User{}).Count(&count).Error; err != nil {
		return 0, err
	}
	return int(count), nil
}

func (t *TipherethRepo) LinkAccount(
	ctx context.Context,
	a *model.Account,
	userID model.InternalID,
) (model.InternalID, error) {
	accountID := a.ID
	err := t.data.WithTx(ctx, func(tx *gorm.DB) error {
		// Check if user already has account on this platform
		var existingCount int64
		if err := tx.Model(&gormschema.Account{}).
			Where("platform = ? AND bound_user_id = ?", a.Platform, userID).
			Count(&existingCount).Error; err != nil {
			return err
		}
		if existingCount > 0 {
			return errors.New("an account already bound to user")
		}

		// Check if account already exists
		var acc gormschema.Account
		err := tx.Where("platform = ? AND platform_account_id = ?", a.Platform, a.PlatformAccountID).
			First(&acc).Error

		if err == gorm.ErrRecordNotFound {
			// Create new account
			newAcc := gormschema.Account{
				ID:                a.ID,
				Platform:          a.Platform,
				PlatformAccountID: a.PlatformAccountID,
				BoundUserID:       &userID,
				Name:              a.Name,
				ProfileURL:        a.ProfileURL,
				AvatarURL:         a.AvatarURL,
			}
			return tx.Create(&newAcc).Error
		}
		if err != nil {
			return err
		}

		// Account exists, check if already bound
		if acc.BoundUserID != nil {
			return errors.New("account already bound to an user")
		}

		accountID = acc.ID
		return tx.Model(&gormschema.Account{}).
			Where("id = ?", acc.ID).
			Update("bound_user_id", userID).Error
	})
	if err != nil {
		return 0, err
	}
	return accountID, nil
}

func (t *TipherethRepo) UnLinkAccount(ctx context.Context, aid model.InternalID, u model.InternalID) error {
	return t.data.db.WithContext(ctx).
		Model(&gormschema.Account{}).
		Where("id = ? AND bound_user_id = ?", aid, u).
		Update("bound_user_id", nil).Error
}

func (t *TipherethRepo) ListLinkAccounts(
	ctx context.Context,
	userID model.InternalID,
) ([]*model.Account, error) {
	var accounts []gormschema.Account
	err := t.data.db.WithContext(ctx).
		Where("bound_user_id = ?", userID).
		Find(&accounts).Error
	if err != nil {
		return nil, err
	}
	res := make([]*model.Account, len(accounts))
	for i := range accounts {
		res[i] = gormschema.ToBizAccount(&accounts[i])
	}
	return res, nil
}

func (t *TipherethRepo) ListPorters(
	ctx context.Context,
	paging model.Paging,
) ([]*modelsupervisor.PorterInstance, int64, error) {
	query := t.data.db.WithContext(ctx).Model(&gormschema.PorterInstance{})

	var count int64
	if err := query.Count(&count).Error; err != nil {
		return nil, 0, err
	}

	var porters []gormschema.PorterInstance
	if err := query.Limit(paging.ToLimit()).Offset(paging.ToOffset()).Find(&porters).Error; err != nil {
		return nil, 0, err
	}

	res := make([]*modelsupervisor.PorterInstance, len(porters))
	for i := range porters {
		res[i] = gormschema.ToBizPorter(&porters[i])
	}
	return res, count, nil
}

func (t *TipherethRepo) GetPorter(
	ctx context.Context,
	id model.InternalID,
) (*modelsupervisor.PorterInstance, error) {
	var p gormschema.PorterInstance
	if err := t.data.db.WithContext(ctx).First(&p, id).Error; err != nil {
		return nil, err
	}
	return gormschema.ToBizPorter(&p), nil
}

func (t *TipherethRepo) UpdatePorterStatus(
	ctx context.Context,
	id model.InternalID,
	status model.UserStatus,
) (*modelsupervisor.PorterInstance, error) {
	if err := t.data.db.WithContext(ctx).
		Model(&gormschema.PorterInstance{}).
		Where("id = ?", id).
		Update("status", gormschema.ToSchemaUserStatus(status)).Error; err != nil {
		return nil, err
	}

	var pi gormschema.PorterInstance
	if err := t.data.db.WithContext(ctx).First(&pi, id).Error; err != nil {
		return nil, err
	}
	return gormschema.ToBizPorter(&pi), nil
}

func (t *TipherethRepo) CreatePorterContext(
	ctx context.Context,
	userID model.InternalID,
	context *modelsupervisor.PorterContext,
) error {
	pc := gormschema.PorterContext{
		ID:           context.ID,
		OwnerID:      userID,
		GlobalName:   context.GlobalName,
		Region:       context.Region,
		ContextJSON:  context.ContextJSON,
		Name:         context.Name,
		Description:  context.Description,
		Status:       gormschema.ToSchemaPorterContextStatus(context.Status),
		HandleStatus: gormschema.ToSchemaPorterContextHandleStatus(context.HandleStatus),
	}
	return t.data.db.WithContext(ctx).Create(&pc).Error
}

func (t *TipherethRepo) ListPorterContexts(
	ctx context.Context,
	userID model.InternalID,
	paging model.Paging,
) ([]*modelsupervisor.PorterContext, int64, error) {
	query := t.data.db.WithContext(ctx).Model(&gormschema.PorterContext{}).
		Where("owner_id = ?", userID)

	var count int64
	if err := query.Count(&count).Error; err != nil {
		return nil, 0, err
	}

	var contexts []gormschema.PorterContext
	if err := query.Limit(paging.ToLimit()).Offset(paging.ToOffset()).Find(&contexts).Error; err != nil {
		return nil, 0, err
	}

	res := make([]*modelsupervisor.PorterContext, len(contexts))
	for i := range contexts {
		res[i] = gormschema.ToBizPorterContext(&contexts[i])
	}
	return res, count, nil
}

func (t *TipherethRepo) ListPorterContextsByGlobalName(
	ctx context.Context,
	userID model.InternalID,
	globalName string,
	paging model.Paging,
) ([]*modelsupervisor.PorterContext, int64, error) {
	query := t.data.db.WithContext(ctx).Model(&gormschema.PorterContext{}).
		Where("owner_id = ? AND global_name = ?", userID, globalName)

	var count int64
	if err := query.Count(&count).Error; err != nil {
		return nil, 0, err
	}

	var contexts []gormschema.PorterContext
	if err := query.Limit(paging.ToLimit()).Offset(paging.ToOffset()).Find(&contexts).Error; err != nil {
		return nil, 0, err
	}

	res := make([]*modelsupervisor.PorterContext, len(contexts))
	for i := range contexts {
		res[i] = gormschema.ToBizPorterContext(&contexts[i])
	}
	return res, count, nil
}

func (t *TipherethRepo) UpdatePorterContext(
	ctx context.Context,
	userID model.InternalID,
	context *modelsupervisor.PorterContext,
) error {
	updates := map[string]any{
		"context_json": context.ContextJSON,
		"name":         context.Name,
		"description":  context.Description,
		"status":       gormschema.ToSchemaPorterContextStatus(context.Status),
	}
	return t.data.db.WithContext(ctx).
		Model(&gormschema.PorterContext{}).
		Where("id = ? AND owner_id = ?", context.ID, userID).
		Updates(updates).Error
}

func (t *TipherethRepo) ListPorterDigests(
	ctx context.Context,
	status []model.UserStatus,
) ([]*modelsupervisor.PorterDigest, error) {
	query := t.data.db.WithContext(ctx).Model(&gormschema.PorterInstance{})

	if len(status) > 0 {
		statusStrs := make([]string, len(status))
		for i, s := range status {
			statusStrs[i] = gormschema.ToSchemaUserStatus(s)
		}
		query = query.Where("status IN ?", statusStrs)
	}

	// Group by global_name and region, get min ID
	type result struct {
		GlobalName string
		Region     string
		MinID      model.InternalID
	}
	var results []result
	if err := query.Select("global_name, region, MIN(id) as min_id").
		Group("global_name, region").
		Find(&results).Error; err != nil {
		return nil, err
	}

	if len(results) == 0 {
		return []*modelsupervisor.PorterDigest{}, nil
	}

	ids := make([]model.InternalID, len(results))
	for i, r := range results {
		ids[i] = r.MinID
	}

	var porters []gormschema.PorterInstance
	if err := t.data.db.WithContext(ctx).Where("id IN ?", ids).Find(&porters).Error; err != nil {
		return nil, err
	}

	pgm := make(map[string]*modelsupervisor.PorterDigest)
	for i := range porters {
		p := &porters[i]
		if len(p.ContextJSONSchema) == 0 {
			continue
		}
		if pgm[p.GlobalName] == nil {
			bizPorter := gormschema.ToBizPorter(p)
			pgm[p.GlobalName] = &modelsupervisor.PorterDigest{
				BinarySummary:     bizPorter.BinarySummary,
				GlobalName:        p.GlobalName,
				Regions:           []string{p.Region},
				ContextJSONSchema: p.ContextJSONSchema,
				FeatureSummary:    (*modelsupervisor.PorterFeatureSummary)(p.FeatureSummary),
			}
		} else {
			pgm[p.GlobalName].Regions = append(pgm[p.GlobalName].Regions, p.Region)
		}
	}

	var pg []*modelsupervisor.PorterDigest
	for _, v := range pgm {
		pg = append(pg, v)
	}
	return pg, nil
}

func (t *TipherethRepo) FetchPorterContext(
	ctx context.Context,
	id model.InternalID,
) (*modelsupervisor.PorterContext, error) {
	var pc gormschema.PorterContext
	if err := t.data.db.WithContext(ctx).First(&pc, id).Error; err != nil {
		return nil, err
	}
	return gormschema.ToBizPorterContext(&pc), nil
}

func (t *TipherethRepo) GetEnabledPorterContexts(
	ctx context.Context,
) ([]*modelsupervisor.PorterContext, error) {
	var contexts []gormschema.PorterContext
	if err := t.data.db.WithContext(ctx).
		Where("status = ?", "active").
		Find(&contexts).Error; err != nil {
		return nil, err
	}
	res := make([]*modelsupervisor.PorterContext, len(contexts))
	for i := range contexts {
		res[i] = gormschema.ToBizPorterContext(&contexts[i])
	}
	return res, nil
}
