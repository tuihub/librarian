// Package gormschema provides type converters between GORM models and biz models.
package gormschema

import (
	"github.com/tuihub/librarian/internal/model"
	"github.com/tuihub/librarian/internal/model/modelchesed"
	"github.com/tuihub/librarian/internal/model/modelfeed"
	"github.com/tuihub/librarian/internal/model/modelgebura"
	"github.com/tuihub/librarian/internal/model/modelnetzach"
	"github.com/tuihub/librarian/internal/model/modelsupervisor"
	"github.com/tuihub/librarian/internal/model/modelyesod"
)

// ============================================
// User/Account converters
// ============================================

func ToBizUser(u *User) *model.User {
	if u == nil {
		return nil
	}
	return &model.User{
		ID:       u.ID,
		Username: u.Username,
		Password: "", // Never return password
		Type:     toBizUserType(u.Type),
		Status:   toBizUserStatus(u.Status),
	}
}

func ToBizUserList(users []*User) []*model.User {
	res := make([]*model.User, len(users))
	for i, u := range users {
		res[i] = ToBizUser(u)
	}
	return res
}

func toBizUserType(t string) model.UserType {
	switch t {
	case "admin":
		return model.UserTypeAdmin
	case "normal":
		return model.UserTypeNormal
	default:
		return model.UserTypeUnspecified
	}
}

func ToSchemaUserType(t model.UserType) string {
	switch t {
	case model.UserTypeAdmin:
		return "admin"
	case model.UserTypeNormal:
		return "normal"
	default:
		return ""
	}
}

func toBizUserStatus(s string) model.UserStatus {
	switch s {
	case "active":
		return model.UserStatusActive
	case "blocked":
		return model.UserStatusBlocked
	default:
		return model.UserStatusUnspecified
	}
}

func ToSchemaUserStatus(s model.UserStatus) string {
	switch s {
	case model.UserStatusActive:
		return "active"
	case model.UserStatusBlocked:
		return "blocked"
	default:
		return ""
	}
}

func ToBizDevice(d *Device) *model.Device {
	if d == nil {
		return nil
	}
	return &model.Device{
		ID:                      d.ID,
		DeviceName:              d.DeviceName,
		SystemType:              toBizSystemType(d.SystemType),
		SystemVersion:           d.SystemVersion,
		ClientName:              d.ClientName,
		ClientSourceCodeAddress: d.ClientSourceCodeAddress,
		ClientVersion:           d.ClientVersion,
	}
}

func ToBizDeviceList(devices []*Device) []*model.Device {
	res := make([]*model.Device, len(devices))
	for i, d := range devices {
		res[i] = ToBizDevice(d)
	}
	return res
}

func toBizSystemType(s string) model.SystemType {
	switch s {
	case "ios":
		return model.SystemTypeIOS
	case "android":
		return model.SystemTypeAndroid
	case "web":
		return model.SystemTypeWeb
	case "windows":
		return model.SystemTypeWindows
	case "macos":
		return model.SystemTypeMacOS
	case "linux":
		return model.SystemTypeLinux
	default:
		return model.SystemTypeUnspecified
	}
}

func ToSchemaSystemType(s model.SystemType) string {
	switch s {
	case model.SystemTypeIOS:
		return "ios"
	case model.SystemTypeAndroid:
		return "android"
	case model.SystemTypeWeb:
		return "web"
	case model.SystemTypeWindows:
		return "windows"
	case model.SystemTypeMacOS:
		return "macos"
	case model.SystemTypeLinux:
		return "linux"
	default:
		return "unknown"
	}
}

func ToBizSession(s *Session) *model.Session {
	if s == nil {
		return nil
	}
	return &model.Session{
		ID:           s.ID,
		UserID:       s.UserID,
		RefreshToken: s.RefreshToken,
		Device:       nil, // Must be set separately
		CreateAt:     s.CreatedAt,
		ExpireAt:     s.ExpireAt,
	}
}

func ToBizSessionList(sessions []*Session) []*model.Session {
	res := make([]*model.Session, len(sessions))
	for i, s := range sessions {
		res[i] = ToBizSession(s)
	}
	return res
}

func ToBizAccount(a *Account) *model.Account {
	if a == nil {
		return nil
	}
	return &model.Account{
		ID:                a.ID,
		Platform:          a.Platform,
		PlatformAccountID: a.PlatformAccountID,
		Name:              a.Name,
		ProfileURL:        a.ProfileURL,
		AvatarURL:         a.AvatarURL,
		LatestUpdateTime:  a.UpdatedAt,
	}
}

func ToBizAccountList(accounts []*Account) []*model.Account {
	res := make([]*model.Account, len(accounts))
	for i, a := range accounts {
		res[i] = ToBizAccount(a)
	}
	return res
}

// ============================================
// Porter converters
// ============================================

func ToBizPorter(p *PorterInstance) *modelsupervisor.PorterInstance {
	if p == nil {
		return nil
	}
	return &modelsupervisor.PorterInstance{
		ID: p.ID,
		BinarySummary: &modelsupervisor.PorterBinarySummary{
			Name:              p.Name,
			Version:           p.Version,
			Description:       p.Description,
			SourceCodeAddress: p.SourceCodeAddress,
			BuildVersion:      p.BuildVersion,
			BuildDate:         p.BuildDate,
		},
		GlobalName:              p.GlobalName,
		Address:                 p.Address,
		Region:                  p.Region,
		FeatureSummary:          (*modelsupervisor.PorterFeatureSummary)(p.FeatureSummary),
		Status:                  toBizUserStatus(p.Status),
		ContextJSONSchema:       p.ContextJSONSchema,
		ConnectionStatus:        toBizPorterConnectionStatus(p.ConnectionStatus),
		ConnectionStatusMessage: p.ConnectionStatusMessage,
	}
}

func ToBizPorterList(porters []*PorterInstance) []*modelsupervisor.PorterInstance {
	res := make([]*modelsupervisor.PorterInstance, len(porters))
	for i, p := range porters {
		res[i] = ToBizPorter(p)
	}
	return res
}

func toBizPorterConnectionStatus(s string) modelsupervisor.PorterConnectionStatus {
	switch s {
	case "queueing":
		return modelsupervisor.PorterConnectionStatusQueueing
	case "connected":
		return modelsupervisor.PorterConnectionStatusConnected
	case "disconnected":
		return modelsupervisor.PorterConnectionStatusDisconnected
	case "active":
		return modelsupervisor.PorterConnectionStatusActive
	case "activation_failed":
		return modelsupervisor.PorterConnectionStatusActivationFailed
	case "downgraded":
		return modelsupervisor.PorterConnectionStatusDowngraded
	default:
		return modelsupervisor.PorterConnectionStatusUnspecified
	}
}

func ToSchemaPorterConnectionStatus(s modelsupervisor.PorterConnectionStatus) string {
	switch s {
	case modelsupervisor.PorterConnectionStatusQueueing:
		return "queueing"
	case modelsupervisor.PorterConnectionStatusConnected:
		return "connected"
	case modelsupervisor.PorterConnectionStatusDisconnected:
		return "disconnected"
	case modelsupervisor.PorterConnectionStatusActive:
		return "active"
	case modelsupervisor.PorterConnectionStatusActivationFailed:
		return "activation_failed"
	case modelsupervisor.PorterConnectionStatusDowngraded:
		return "downgraded"
	default:
		return "unspecified"
	}
}

func ToBizPorterContext(pc *PorterContext) *modelsupervisor.PorterContext {
	if pc == nil {
		return nil
	}
	return &modelsupervisor.PorterContext{
		ID:                  pc.ID,
		GlobalName:          pc.GlobalName,
		Region:              pc.Region,
		ContextJSON:         pc.ContextJSON,
		Name:                pc.Name,
		Description:         pc.Description,
		Status:              toBizPorterContextStatus(pc.Status),
		HandleStatus:        toBizPorterContextHandleStatus(pc.HandleStatus),
		HandleStatusMessage: pc.HandleStatusMessage,
	}
}

func ToBizPorterContextList(contexts []*PorterContext) []*modelsupervisor.PorterContext {
	res := make([]*modelsupervisor.PorterContext, len(contexts))
	for i, pc := range contexts {
		res[i] = ToBizPorterContext(pc)
	}
	return res
}

func toBizPorterContextStatus(s string) modelsupervisor.PorterContextStatus {
	switch s {
	case "active":
		return modelsupervisor.PorterContextStatusActive
	case "disabled":
		return modelsupervisor.PorterContextStatusDisabled
	default:
		return modelsupervisor.PorterContextStatusUnspecified
	}
}

func ToSchemaPorterContextStatus(s modelsupervisor.PorterContextStatus) string {
	switch s {
	case modelsupervisor.PorterContextStatusActive:
		return "active"
	case modelsupervisor.PorterContextStatusDisabled:
		return "disabled"
	default:
		return ""
	}
}

func toBizPorterContextHandleStatus(s string) modelsupervisor.PorterContextHandleStatus {
	switch s {
	case "active":
		return modelsupervisor.PorterContextHandleStatusActive
	case "downgraded":
		return modelsupervisor.PorterContextHandleStatusDowngraded
	case "queueing":
		return modelsupervisor.PorterContextHandleStatusQueueing
	case "blocked":
		return modelsupervisor.PorterContextHandleStatusBlocked
	default:
		return modelsupervisor.PorterContextHandleStatusUnspecified
	}
}

func ToSchemaPorterContextHandleStatus(s modelsupervisor.PorterContextHandleStatus) string {
	switch s {
	case modelsupervisor.PorterContextHandleStatusActive:
		return "active"
	case modelsupervisor.PorterContextHandleStatusDowngraded:
		return "downgraded"
	case modelsupervisor.PorterContextHandleStatusQueueing:
		return "queueing"
	case modelsupervisor.PorterContextHandleStatusBlocked:
		return "blocked"
	default:
		return "unspecified"
	}
}

// ============================================
// App converters
// ============================================

func ToBizAppInfo(a *AppInfo) *modelgebura.AppInfo {
	if a == nil {
		return nil
	}
	return &modelgebura.AppInfo{
		ID:                 a.ID,
		Source:             a.Source,
		SourceAppID:        a.SourceAppID,
		SourceURL:          a.SourceURL,
		Name:               a.Name,
		Type:               toBizAppType(a.Type),
		ShortDescription:   a.ShortDescription,
		Description:        a.Description,
		IconImageURL:       a.IconImageURL,
		IconImageID:        a.IconImageID,
		BackgroundImageURL: a.BackgroundImageURL,
		BackgroundImageID:  a.BackgroundImageID,
		CoverImageURL:      a.CoverImageURL,
		CoverImageID:       a.CoverImageID,
		ReleaseDate:        a.ReleaseDate,
		Developer:          a.Developer,
		Publisher:          a.Publisher,
		Tags:               []string(a.Tags),
		AlternativeNames:   []string(a.AlternativeNames),
		RawData:            a.RawData,
		UpdatedAt:          a.UpdatedAt,
	}
}

func ToBizAppInfoList(infos []*AppInfo) []*modelgebura.AppInfo {
	res := make([]*modelgebura.AppInfo, len(infos))
	for i, a := range infos {
		res[i] = ToBizAppInfo(a)
	}
	return res
}

func toBizAppType(t string) modelgebura.AppType {
	switch t {
	case "game":
		return modelgebura.AppTypeGame
	default:
		return modelgebura.AppTypeUnspecified
	}
}

func ToSchemaAppType(t modelgebura.AppType) string {
	switch t {
	case modelgebura.AppTypeGame:
		return "game"
	default:
		return "unknown"
	}
}

func ToBizApp(a *App) *modelgebura.App {
	if a == nil {
		return nil
	}
	return &modelgebura.App{
		ID:                 a.ID,
		VersionNumber:      a.VersionNumber,
		VersionDate:        a.VersionDate,
		CreatorDeviceID:    a.CreatorDeviceID,
		AppSources:         map[string]string(a.AppSources),
		Public:             a.Public,
		BoundStoreAppID:    a.BoundStoreAppID,
		StopStoreManage:    a.StopStoreManage,
		Name:               a.Name,
		Type:               toBizAppType(a.Type),
		ShortDescription:   a.ShortDescription,
		Description:        a.Description,
		IconImageURL:       a.IconImageURL,
		IconImageID:        a.IconImageID,
		BackgroundImageURL: a.BackgroundImageURL,
		BackgroundImageID:  a.BackgroundImageID,
		CoverImageURL:      a.CoverImageURL,
		CoverImageID:       a.CoverImageID,
		ReleaseDate:        a.ReleaseDate,
		Developer:          a.Developer,
		Publisher:          a.Publisher,
		Tags:               []string(a.Tags),
		AlternativeNames:   []string(a.AlternativeNames),
	}
}

func ToBizAppList(apps []*App) []*modelgebura.App {
	res := make([]*modelgebura.App, len(apps))
	for i, a := range apps {
		res[i] = ToBizApp(a)
	}
	return res
}

func ToBizAppRunTime(a *AppRunTime) *modelgebura.AppRunTime {
	if a == nil {
		return nil
	}
	return &modelgebura.AppRunTime{
		ID:       a.ID,
		AppID:    a.AppID,
		DeviceID: a.DeviceID,
		RunTime: &model.TimeRange{
			StartTime: a.StartTime,
			Duration:  a.Duration,
		},
	}
}

func ToBizAppRunTimeList(runtimes []*AppRunTime) []*modelgebura.AppRunTime {
	res := make([]*modelgebura.AppRunTime, len(runtimes))
	for i, a := range runtimes {
		res[i] = ToBizAppRunTime(a)
	}
	return res
}

func ToBizAppCategory(ac *AppCategory) *modelgebura.AppCategory {
	if ac == nil {
		return nil
	}
	return &modelgebura.AppCategory{
		ID:            ac.ID,
		VersionNumber: ac.VersionNumber,
		VersionDate:   ac.VersionDate,
		Name:          ac.Name,
		AppIDs:        nil, // Must be set separately
	}
}

// ============================================
// Sentinel converters
// ============================================

func ToBizSentinel(s *Sentinel) *modelgebura.Sentinel {
	if s == nil {
		return nil
	}
	return &modelgebura.Sentinel{
		ID:                   s.ID,
		Name:                 s.Name,
		Description:          s.Description,
		URL:                  s.URL,
		AlternativeUrls:      []string(s.AlternativeUrls),
		GetTokenPath:         s.GetTokenPath,
		DownloadFileBasePath: s.DownloadFileBasePath,
		Libraries:            nil, // Must be set separately
	}
}

func ToBizSentinelList(sentinels []*Sentinel) []*modelgebura.Sentinel {
	res := make([]*modelgebura.Sentinel, len(sentinels))
	for i, s := range sentinels {
		res[i] = ToBizSentinel(s)
	}
	return res
}

func ToBizSentinelSession(ss *SentinelSession) *modelgebura.SentinelSession {
	if ss == nil {
		return nil
	}
	return &modelgebura.SentinelSession{
		ID:              ss.ID,
		SentinelID:      ss.SentinelID,
		RefreshToken:    ss.RefreshToken,
		Status:          toBizSentinelSessionStatus(ss.Status),
		CreatorID:       ss.CreatorID,
		ExpireAt:        ss.ExpireAt,
		LastUsedAt:      ss.LastUsedAt,
		LastRefreshedAt: ss.LastRefreshedAt,
		RefreshCount:    ss.RefreshCount,
	}
}

func ToBizSentinelSessionList(sessions []*SentinelSession) []*modelgebura.SentinelSession {
	res := make([]*modelgebura.SentinelSession, len(sessions))
	for i, ss := range sessions {
		res[i] = ToBizSentinelSession(ss)
	}
	return res
}

func toBizSentinelSessionStatus(s string) modelgebura.SentinelSessionStatus {
	switch s {
	case "active":
		return modelgebura.SentinelSessionStatusActive
	case "suspend":
		return modelgebura.SentinelSessionStatusSuspend
	default:
		return modelgebura.SentinelSessionStatusUnspecified
	}
}

func ToSchemaSentinelSessionStatus(s modelgebura.SentinelSessionStatus) string {
	switch s {
	case modelgebura.SentinelSessionStatusActive:
		return "active"
	case modelgebura.SentinelSessionStatusSuspend:
		return "suspend"
	default:
		return ""
	}
}

func ToBizStoreApp(s *StoreApp) *modelgebura.StoreApp {
	if s == nil {
		return nil
	}
	return &modelgebura.StoreApp{
		ID:          s.ID,
		Name:        s.Name,
		Description: s.Description,
	}
}

func ToBizStoreAppList(apps []*StoreApp) []*modelgebura.StoreApp {
	res := make([]*modelgebura.StoreApp, len(apps))
	for i, a := range apps {
		res[i] = ToBizStoreApp(a)
	}
	return res
}

func ToBizStoreAppBinary(sab *SentinelAppBinary) *modelgebura.StoreAppBinary {
	if sab == nil {
		return nil
	}
	return &modelgebura.StoreAppBinary{
		ID:        sab.ID,
		AppID:     sab.StoreAppID,
		UnionID:   sab.UnionID,
		SizeBytes: sab.SizeBytes,
		NeedToken: sab.NeedToken,
		Name:      sab.Name,
		Version:   sab.Version,
		Developer: sab.Developer,
		Publisher: sab.Publisher,
	}
}

func ToBizStoreAppBinaryList(binaries []*SentinelAppBinary) []*modelgebura.StoreAppBinary {
	res := make([]*modelgebura.StoreAppBinary, len(binaries))
	for i, b := range binaries {
		res[i] = ToBizStoreAppBinary(b)
	}
	return res
}

// ============================================
// Feed converters
// ============================================

func ToBizFeed(f *Feed) *modelfeed.Feed {
	if f == nil {
		return nil
	}
	return &modelfeed.Feed{
		ID:          f.ID,
		Title:       f.Title,
		Description: f.Description,
		Link:        f.Link,
		Authors:     []*modelfeed.Person(f.Authors),
		Language:    f.Language,
		Image:       (*modelfeed.Image)(f.Image),
		Items:       nil, // Not included
		FeedType:    "",  // Not stored
		FeedVersion: "",  // Not stored
	}
}

func ToBizFeedItem(fi *FeedItem) *modelfeed.Item {
	if fi == nil {
		return nil
	}
	return &modelfeed.Item{
		ID:                fi.ID,
		Title:             fi.Title,
		Description:       fi.Description,
		Content:           fi.Content,
		Link:              fi.Link,
		Updated:           fi.Updated,
		UpdatedParsed:     fi.UpdatedParsed,
		Published:         fi.Published,
		PublishedParsed:   &fi.PublishedParsed,
		Authors:           []*modelfeed.Person(fi.Authors),
		GUID:              fi.GUID,
		Image:             (*modelfeed.Image)(fi.Image),
		Enclosures:        []*modelfeed.Enclosure(fi.Enclosures),
		PublishPlatform:   fi.PublishPlatform,
		ReadCount:         fi.ReadCount,
		DigestDescription: fi.DigestDescription,
		DigestImages:      []*modelfeed.Image(fi.DigestImages),
	}
}

func ToBizFeedItemList(items []*FeedItem) []*modelfeed.Item {
	res := make([]*modelfeed.Item, len(items))
	for i, fi := range items {
		res[i] = ToBizFeedItem(fi)
	}
	return res
}

func ToBizFeedConfig(fc *FeedConfig) *modelyesod.FeedConfig {
	if fc == nil {
		return nil
	}
	var source *model.FeatureRequest
	if fc.Source != nil {
		s := model.FeatureRequest(*fc.Source)
		source = &s
	}
	return &modelyesod.FeedConfig{
		ID:                fc.ID,
		Name:              fc.Name,
		Description:       fc.Description,
		Source:            source,
		ActionSets:        nil, // Must be set separately
		Category:          fc.Category,
		Status:            toBizFeedConfigStatus(fc.Status),
		PullInterval:      fc.PullInterval,
		LatestPullTime:    fc.LatestPullAt,
		LatestPullStatus:  toBizFeedConfigPullStatus(fc.LatestPullStatus),
		LatestPullMessage: fc.LatestPullMessage,
		HideItems:         fc.HideItems,
	}
}

func ToBizFeedConfigList(configs []*FeedConfig) []*modelyesod.FeedConfig {
	res := make([]*modelyesod.FeedConfig, len(configs))
	for i, fc := range configs {
		res[i] = ToBizFeedConfig(fc)
	}
	return res
}

func toBizFeedConfigStatus(s string) modelyesod.FeedConfigStatus {
	switch s {
	case "active":
		return modelyesod.FeedConfigStatusActive
	case "suspend":
		return modelyesod.FeedConfigStatusSuspend
	default:
		return modelyesod.FeedConfigStatusUnspecified
	}
}

func ToSchemaFeedConfigStatus(s modelyesod.FeedConfigStatus) string {
	switch s {
	case modelyesod.FeedConfigStatusActive:
		return "active"
	case modelyesod.FeedConfigStatusSuspend:
		return "suspend"
	default:
		return ""
	}
}

func toBizFeedConfigPullStatus(s string) modelyesod.FeedConfigPullStatus {
	switch s {
	case "processing":
		return modelyesod.FeedConfigPullStatusProcessing
	case "success":
		return modelyesod.FeedConfigPullStatusSuccess
	case "failed":
		return modelyesod.FeedConfigPullStatusFailed
	default:
		return modelyesod.FeedConfigPullStatusUnspecified
	}
}

func ToSchemaFeedConfigPullStatus(s modelyesod.FeedConfigPullStatus) string {
	switch s {
	case modelyesod.FeedConfigPullStatusProcessing:
		return "processing"
	case modelyesod.FeedConfigPullStatusSuccess:
		return "success"
	case modelyesod.FeedConfigPullStatusFailed:
		return "failed"
	default:
		return ""
	}
}

func ToBizFeedActionSet(fas *FeedActionSet) *modelyesod.FeedActionSet {
	if fas == nil {
		return nil
	}
	return &modelyesod.FeedActionSet{
		ID:          fas.ID,
		Name:        fas.Name,
		Description: fas.Description,
		Actions:     []*model.FeatureRequest(fas.Actions),
	}
}

func ToBizFeedActionSetList(sets []*FeedActionSet) []*modelyesod.FeedActionSet {
	res := make([]*modelyesod.FeedActionSet, len(sets))
	for i, fas := range sets {
		res[i] = ToBizFeedActionSet(fas)
	}
	return res
}

func ToBizFeedItemCollection(fic *FeedItemCollection) *modelyesod.FeedItemCollection {
	if fic == nil {
		return nil
	}
	return &modelyesod.FeedItemCollection{
		ID:          fic.ID,
		Name:        fic.Name,
		Description: fic.Description,
		Category:    fic.Category,
	}
}

func ToBizFeedItemCollectionList(collections []*FeedItemCollection) []*modelyesod.FeedItemCollection {
	res := make([]*modelyesod.FeedItemCollection, len(collections))
	for i, fic := range collections {
		res[i] = ToBizFeedItemCollection(fic)
	}
	return res
}

// ============================================
// Notify converters
// ============================================

func ToBizNotifyTarget(nt *NotifyTarget) *modelnetzach.NotifyTarget {
	if nt == nil {
		return nil
	}
	var dest *model.FeatureRequest
	if nt.Destination != nil {
		d := model.FeatureRequest(*nt.Destination)
		dest = &d
	}
	return &modelnetzach.NotifyTarget{
		ID:          nt.ID,
		Name:        nt.Name,
		Description: nt.Description,
		Destination: dest,
		Status:      toBizNotifyTargetStatus(nt.Status),
	}
}

func ToBizNotifyTargetList(targets []*NotifyTarget) []*modelnetzach.NotifyTarget {
	res := make([]*modelnetzach.NotifyTarget, len(targets))
	for i, nt := range targets {
		res[i] = ToBizNotifyTarget(nt)
	}
	return res
}

func toBizNotifyTargetStatus(s string) modelnetzach.NotifyTargetStatus {
	switch s {
	case "active":
		return modelnetzach.NotifyTargetStatusActive
	case "suspend":
		return modelnetzach.NotifyTargetStatusSuspend
	default:
		return modelnetzach.NotifyTargetStatusUnspecified
	}
}

func ToSchemaNotifyTargetStatus(s modelnetzach.NotifyTargetStatus) string {
	switch s {
	case modelnetzach.NotifyTargetStatusActive:
		return "active"
	case modelnetzach.NotifyTargetStatusSuspend:
		return "suspend"
	default:
		return ""
	}
}

func ToBizNotifyFlow(nf *NotifyFlow) *modelnetzach.NotifyFlow {
	if nf == nil {
		return nil
	}
	return &modelnetzach.NotifyFlow{
		ID:          nf.ID,
		Name:        nf.Name,
		Description: nf.Description,
		Sources:     nil, // Must be set separately
		Targets:     nil, // Must be set separately
		Status:      toBizNotifyFlowStatus(nf.Status),
	}
}

func toBizNotifyFlowStatus(s string) modelnetzach.NotifyFlowStatus {
	switch s {
	case "active":
		return modelnetzach.NotifyFlowStatusActive
	case "suspend":
		return modelnetzach.NotifyFlowStatusSuspend
	default:
		return modelnetzach.NotifyFlowStatusUnspecified
	}
}

func ToSchemaNotifyFlowStatus(s modelnetzach.NotifyFlowStatus) string {
	switch s {
	case modelnetzach.NotifyFlowStatusActive:
		return "active"
	case modelnetzach.NotifyFlowStatusSuspend:
		return "suspend"
	default:
		return ""
	}
}

func ToBizSystemNotification(sn *SystemNotification) *modelnetzach.SystemNotification {
	if sn == nil {
		return nil
	}
	return &modelnetzach.SystemNotification{
		ID:         sn.ID,
		Type:       toBizSystemNotificationType(sn.Type),
		Level:      toBizSystemNotificationLevel(sn.Level),
		Status:     toBizSystemNotificationStatus(sn.Status),
		Title:      sn.Title,
		Content:    sn.Content,
		CreateTime: sn.CreatedAt,
		UpdateTime: sn.UpdatedAt,
	}
}

func ToBizSystemNotificationList(notifications []*SystemNotification) []*modelnetzach.SystemNotification {
	res := make([]*modelnetzach.SystemNotification, len(notifications))
	for i, sn := range notifications {
		res[i] = ToBizSystemNotification(sn)
	}
	return res
}

func toBizSystemNotificationType(s string) modelnetzach.SystemNotificationType {
	switch s {
	case "system":
		return modelnetzach.SystemNotificationTypeSystem
	case "user":
		return modelnetzach.SystemNotificationTypeUser
	default:
		return modelnetzach.SystemNotificationTypeUnspecified
	}
}

func ToSchemaSystemNotificationType(s modelnetzach.SystemNotificationType) string {
	switch s {
	case modelnetzach.SystemNotificationTypeSystem:
		return "system"
	case modelnetzach.SystemNotificationTypeUser:
		return "user"
	default:
		return ""
	}
}

func toBizSystemNotificationLevel(s string) modelnetzach.SystemNotificationLevel {
	switch s {
	case "info":
		return modelnetzach.SystemNotificationLevelInfo
	case "warn":
		return modelnetzach.SystemNotificationLevelWarning
	case "error":
		return modelnetzach.SystemNotificationLevelError
	case "ongoing":
		return modelnetzach.SystemNotificationLevelOngoing
	default:
		return modelnetzach.SystemNotificationLevelUnspecified
	}
}

func ToSchemaSystemNotificationLevel(s modelnetzach.SystemNotificationLevel) string {
	switch s {
	case modelnetzach.SystemNotificationLevelInfo:
		return "info"
	case modelnetzach.SystemNotificationLevelWarning:
		return "warn"
	case modelnetzach.SystemNotificationLevelError:
		return "error"
	case modelnetzach.SystemNotificationLevelOngoing:
		return "ongoing"
	default:
		return ""
	}
}

func toBizSystemNotificationStatus(s string) modelnetzach.SystemNotificationStatus {
	switch s {
	case "unread":
		return modelnetzach.SystemNotificationStatusUnread
	case "read":
		return modelnetzach.SystemNotificationStatusRead
	case "dismissed":
		return modelnetzach.SystemNotificationStatusDismissed
	default:
		return modelnetzach.SystemNotificationStatusUnspecified
	}
}

func ToSchemaSystemNotificationStatus(s modelnetzach.SystemNotificationStatus) string {
	switch s {
	case modelnetzach.SystemNotificationStatusUnread:
		return "unread"
	case modelnetzach.SystemNotificationStatusRead:
		return "read"
	case modelnetzach.SystemNotificationStatusDismissed:
		return "dismissed"
	default:
		return ""
	}
}

// ============================================
// Image converters
// ============================================

func ToBizImage(img *Image) *modelchesed.Image {
	if img == nil {
		return nil
	}
	return &modelchesed.Image{
		ID:          img.ID,
		Name:        img.Name,
		Description: img.Description,
		Status:      toBizImageStatus(img.Status),
	}
}

func ToBizImageList(images []*Image) []*modelchesed.Image {
	res := make([]*modelchesed.Image, len(images))
	for i, img := range images {
		res[i] = ToBizImage(img)
	}
	return res
}

func toBizImageStatus(s string) modelchesed.ImageStatus {
	switch s {
	case "uploaded":
		return modelchesed.ImageStatusUploaded
	case "scanned":
		return modelchesed.ImageStatusScanned
	default:
		return modelchesed.ImageStatusUnspecified
	}
}

func ToSchemaImageStatus(s modelchesed.ImageStatus) string {
	switch s {
	case modelchesed.ImageStatusUploaded:
		return "uploaded"
	case modelchesed.ImageStatusScanned:
		return "scanned"
	default:
		return ""
	}
}
