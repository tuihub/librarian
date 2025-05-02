package model

// https://github.com/babelshift/SteamWebAPI2

type GetPlayerSummariesRequest struct {
	SteamID uint64 `url:"steamids,string"`
}

type GetPlayerSummariesResponse struct {
	Response struct {
		Players []PlayerSummary `json:"players"`
	} `json:"response"`
}

type PlayerSummary struct {
	// Unique Steam ID of the player.
	SteamID int64 `json:"steamid,string"`
	// Determines the visibility of the user's profile (public, private, friends)
	ProfileVisibility ProfileVisibility `json:"communityvisibilitystate"`
	// If set to 1, the user has configured his profile.
	ProfileState uint `json:"profilestate"`
	// User's current nick name (displayed in profile and friends list)
	Nickname string `json:"personaname"`
	// The URL for the player's Steam Community profile
	ProfileURL string `json:"profileurl"`
	// The normal sized uploaded avatar image for the user's Steam profile
	Avatar string `json:"avatar"`
	// The medium sized uploaded avatar image for the user's Steam profile
	AvatarMedium string `json:"avatarmedium"`
	// The full sized uploaded avatar image for the user's Steam profile
	AvatarFull string `json:"avatarfull"`
	//
	AvatarHash string `json:"avatarhash"`
	// The date at which the user last logged off Steam
	LastLogOffDate int64 `json:"lastlogoff"`
	// The selected privacy/visibility level of the player's comments section on their Steam Community profile
	CommentPermission CommentPermission `json:"commentpermission"`
	// The current status of the user on the Steam network
	UserStatus UserStatus `json:"personastate"`
	// The player's real name as entered on their Steam profile
	RealName string `json:"realname"`
	// The player's selected primary group to display on their Steam profile
	PrimaryGroupID string `json:"primaryclanid"`
	// The date at which the user created their Steam account
	AccountCreatedDate int64 `json:"timecreated"`
	//
	PersonaStateFlags int `json:"personastateflags"`
	// The player's selected country
	CountryCode string `json:"loccountrycode"`
	// The player's selected state
	StateCode string
	// The player's selected city.
	// This seems to refer to a database city id, so I'm not sure how to make use of this field.
	CityCode uint `json:"loccityid"`
	// The name of the game that a player is currently playing
	PlayingGameName string `json:"gameextrainfo"`
	// The id of the game that the player is currently playing.
	// This doesn't seem to be an appid, so I'm not sure how to make use of this field.
	PlayingGameID string `json:"gameid"`
	// The IP of the server the user is currently playing on.
	PlayingGameServerIP string `json:"gameserverip"`
}

// ProfileVisibility Indicates the selected privacy/visibility level of the player's Steam Community profile.
type ProfileVisibility int

const (
	ProfileVisibilityUnknown     ProfileVisibility = 0
	ProfileVisibilityPrivate     ProfileVisibility = 1
	ProfileVisibilityFriendsOnly ProfileVisibility = 2
	ProfileVisibilityPublic      ProfileVisibility = 3
)

// CommentPermission Indicates the selected privacy/visibility level of the player's comments section
// on their Steam Community profile.
type CommentPermission int

const (
	CommentPermissionUnknown     ProfileVisibility = 0
	CommentPermissionFriendsOnly ProfileVisibility = 1
	CommentPermissionPrivate     ProfileVisibility = 2
	CommentPermissionPublic      ProfileVisibility = 3
)

// UserStatus Indicates the current status of the user on the Steam network.
type UserStatus int

const (
	UserStatusOffline UserStatus = 0
	UserStatusOnline  UserStatus = 1
	UserStatusBusy    UserStatus = 2
	UserStatusAway    UserStatus = 3
	UserStatusSnooze  UserStatus = 4
	UserStatusUnknown UserStatus = 5
	UserStatusInGame  UserStatus = 6
)

type GetOwnedGamesRequest struct {
	SteamID uint64 `url:"steamid,string"`
	// true if we want additional details (name, icon) about each game
	IncludeAppInfo bool `url:"include_appinfo,omitempty"`
	// free games are excluded by default. If this is set, free games the user has played will be returned.
	IncludePlayedFreeGames bool `url:"include_played_free_games,omitempty"`
	// some games are in the free sub, which are excluded by default.
	IncludeFreeSub bool `url:"include_free_sub,omitempty"`
	// if set, skip unvetted store apps
	SkipUnvettedApps bool `url:"skip_unvetted_apps,omitempty"`
	// will return appinfo in this language
	Language LanguageCode `url:"language,omitempty"`
	// true if we want even more details (capsule, sortas, and capabilities) about each game.
	// include_appinfo must also be true.
	IncludeExtendedAppInfo bool `url:"include_extended_appinfo,omitempty"`
}

type GetOwnedGamesResponse struct {
	Response OwnedGames `json:"response"`
}

type OwnedGames struct {
	GameCount int `json:"game_count"`
	Games     []struct {
		AppID                    uint   `json:"appid"`
		Name                     string `json:"name"`
		PlaytimeForever          uint   `json:"playtime_forever"`
		PlaytimeWindows          uint   `json:"playtime_windows_forever"`
		PlaytimeMac              uint   `json:"playtime_mac_forever"`
		PlaytimeLinux            uint   `json:"playtime_linux_forever"`
		ImgIconURL               string `json:"img_icon_url"`
		ImgLogoURL               string `json:"img_logo_url"`
		HasCommunityVisibleStats bool   `json:"has_community_visible_stats"`
	} `json:"games"`
}

type GetAppDetailsRequest struct {
	AppIDs      []int        `url:"appids"`
	CountryCode ProductCC    `url:"cc,omitempty"`
	Language    LanguageCode `url:"l,omitempty"`
}

type AppDetailsBase struct {
	Type                string   `json:"type"`
	Name                string   `json:"name"`
	AppID               int      `json:"steam_appid"`
	IsFree              bool     `json:"is_free"`
	ControllerSupport   string   `json:"controller_support"`
	DetailedDescription string   `json:"detailed_description"`
	AboutTheGame        string   `json:"about_the_game"`
	ShortDescription    string   `json:"short_description"`
	SupportedLanguages  string   `json:"supported_languages"`
	Reviews             string   `json:"reviews"`
	HeaderImage         string   `json:"header_image"`
	Website             string   `json:"website"`
	Developers          []string `json:"developers"`
	Publishers          []string `json:"publishers"`
	ReleaseDate         struct {
		ComingSoon bool   `json:"coming_soon"`
		Date       string `json:"date"`
	} `json:"release_date"`
}

type AppDetailsBasic struct {
	Success bool            `json:"success"`
	Data    *AppDetailsBase `json:"data"`
}

type AppDetails struct {
	Success bool `json:"success"`
	Data    *struct {
		AppDetailsBase
		RequiredAge int   `json:"required_age"` // may be string in some time
		DLC         []int `json:"dlc"`
		Fullgame    struct {
			AppID int    `json:"appid,string"`
			Name  string `json:"name"`
		} `json:"fullgame"`
		PcRequirements struct {
			Minimum     string `json:"minimum"`
			Recommended string `json:"recommended"`
		} `json:"pc_requirements"`
		MacRequirements struct { // may be an empty slice some time
			Minimum     string `json:"minimum"`
			Recommended string `json:"recommended"`
		} `json:"mac_requirements"`
		LinuxRequirements struct { // may be an empty slice some time
			Minimum     string `json:"minimum"`
			Recommended string `json:"recommended"`
		} `json:"linux_requirements"`
		LegalNotice          string `json:"legal_notice"`
		ExtUserAccountNotice string `json:"ext_user_account_notice"`
		DRMNotice            string `json:"drm_notice"`
		Demos                []struct {
			AppID       int    `json:"appid,string"`
			Description string `json:"description"`
		} `json:"demos"`
		PriceOverview *struct {
			Currency         CurrencyCode `json:"currency"`
			Initial          int          `json:"initial"`
			Final            int          `json:"final"`
			DiscountPercent  int          `json:"discount_percent"`
			InitialFormatted string       `json:"initial_formatted"`
			FinalFormatted   string       `json:"final_formatted"`
			RecurringSub     interface{}  `json:"recurring_sub"` // Either "false" or a sub id int
			RecurringSubDesc string       `json:"recurring_sub_desc"`
		} `json:"price_overview"`
		Packages      []int `json:"packages"`
		PackageGroups []struct {
			Name                    string `json:"name"`
			Title                   string `json:"title"`
			Description             string `json:"description"`
			SelectionText           string `json:"selection_text"`
			SaveText                string `json:"save_text"`
			DisplayType             int    `json:"display_type"`
			IsRecurringSubscription string `json:"is_recurring_subscription"`
			Subs                    []struct {
				PackageID                int    `json:"packageid"`
				PercentSavingsText       string `json:"percent_savings_text"`
				PercentSavings           int    `json:"percent_savings"`
				OptionText               string `json:"option_text"`
				OptionDescription        string `json:"option_description"`
				CanGetFreeLicense        string `json:"can_get_free_license"`
				IsFreeLicense            bool   `json:"is_free_license"`
				PriceInCentsWithDiscount int    `json:"price_in_cents_with_discount"`
			} `json:"subs"`
		} `json:"package_groups"`
		Platforms struct {
			Windows bool `json:"windows"`
			Mac     bool `json:"mac"`
			Linux   bool `json:"linux"`
		} `json:"platforms"`
		Metacritic struct {
			Score int8   `json:"score"`
			URL   string `json:"url"`
		} `json:"metacritic"`
		Categories  AppDetailsCategory `json:"categories"`
		Genres      AppDetailsGenre    `json:"genres"`
		Screenshots []struct {
			ID            int    `json:"id"`
			PathThumbnail string `json:"path_thumbnail"`
			PathFull      string `json:"path_full"`
		} `json:"screenshots"`
		Movies []struct {
			ID        int    `json:"id"`
			Name      string `json:"name"`
			Thumbnail string `json:"thumbnail"`
			Webm      struct {
				Num480 string `json:"480"`
				Max    string `json:"max"`
			} `json:"webm"`
			Highlight bool `json:"highlight"`
		} `json:"movies"`
		Recommendations struct {
			Total int `json:"total"`
		} `json:"recommendations"`
		Achievements struct {
			Total       int `json:"total"`
			Highlighted []struct {
				Name string `json:"name"`
				Path string `json:"path"`
			} `json:"highlighted"`
		} `json:"achievements"`
		SupportInfo struct {
			URL   string `json:"url"`
			Email string `json:"email"`
		} `json:"support_info"`
		Background         string `json:"background"`
		ContentDescriptors struct {
			IDs   interface{}
			Notes interface{}
		} `json:"content_descriptors"`
	} `json:"data"`
}

type AppDetailsGenre []struct {
	ID          int    `json:"id,string"`
	Description string `json:"description"`
}

func (g AppDetailsGenre) IDs() []int {
	ids := make([]int, len(g))
	for i, v := range g {
		ids[i] = v.ID
	}
	return ids
}

func (g AppDetailsGenre) Names() []string {
	names := make([]string, len(g))
	for i, v := range g {
		names[i] = v.Description
	}
	return names
}

type AppDetailsCategory []struct {
	ID          int    `json:"id"`
	Description string `json:"description"`
}

func (c AppDetailsCategory) IDs() []int {
	ids := make([]int, len(c))
	for i, v := range c {
		ids[i] = v.ID
	}
	return ids
}

func (c AppDetailsCategory) Names() []string {
	names := make([]string, len(c))
	for i, v := range c {
		names[i] = v.Description
	}
	return names
}

type GetAppListRequest struct {
	// will return appinfo in this language
	Language LanguageCode `url:"l,omitempty"`
}

type GetAppListResponse struct {
	AppList AppList `json:"applist"`
}

type AppList struct {
	Apps []struct {
		AppID int    `json:"appid"`
		Name  string `json:"name"`
	} `json:"apps"`
}
