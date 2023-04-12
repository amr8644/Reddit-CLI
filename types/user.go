package types

type User struct {
	Kind string `json:"kind"`
	Data struct {
		IsEmployee        bool    `json:"is_employee"`
		IconImg           string  `json:"icon_img"`
		PrefShowSnoovatar bool    `json:"pref_show_snoovatar"`
		Name              string  `json:"name"`
		IsFriend          bool    `json:"is_friend"`
		Created           float64 `json:"created"`
		HasSubscribed     bool    `json:"has_subscribed"`
		HideFromRobots    bool    `json:"hide_from_robots"`
		CreatedUtc        float64 `json:"created_utc"`
		LinkKarma         int     `json:"link_karma"`
		CommentKarma      int     `json:"comment_karma"`
		IsGold            bool    `json:"is_gold"`
		IsMod             bool    `json:"is_mod"`
		Verified          bool    `json:"verified"`
		Subreddit         struct {
			DefaultSet                 bool        `json:"default_set"`
			UserIsContributor          bool        `json:"user_is_contributor"`
			BannerImg                  string      `json:"banner_img"`
			DisableContributorRequests bool        `json:"disable_contributor_requests"`
			UserIsBanned               bool        `json:"user_is_banned"`
			FreeFormReports            bool        `json:"free_form_reports"`
			CommunityIcon              string      `json:"community_icon"`
			ShowMedia                  bool        `json:"show_media"`
			IconColor                  string      `json:"icon_color"`
			UserIsMuted                bool        `json:"user_is_muted"`
			DisplayName                string      `json:"display_name"`
			HeaderImg                  interface{} `json:"header_img"`
			Title                      string      `json:"title"`
			Over18                     bool        `json:"over_18"`
			IconSize                   []int       `json:"icon_size"`
			PrimaryColor               string      `json:"primary_color"`
			IconImg                    string      `json:"icon_img"`
			Description                string      `json:"description"`
			HeaderSize                 interface{} `json:"header_size"`
			RestrictPosting            bool        `json:"restrict_posting"`
			RestrictCommenting         bool        `json:"restrict_commenting"`
			Subscribers                int         `json:"subscribers"`
			IsDefaultIcon              bool        `json:"is_default_icon"`
			LinkFlairPosition          string      `json:"link_flair_position"`
			DisplayNamePrefixed        string      `json:"display_name_prefixed"`
			KeyColor                   string      `json:"key_color"`
			Name                       string      `json:"name"`
			IsDefaultBanner            bool        `json:"is_default_banner"`
			URL                        string      `json:"url"`
			BannerSize                 []int       `json:"banner_size"`
			UserIsModerator            bool        `json:"user_is_moderator"`
			PublicDescription          string      `json:"public_description"`
			LinkFlairEnabled           bool        `json:"link_flair_enabled"`
			SubredditType              string      `json:"subreddit_type"`
			UserIsSubscriber           bool        `json:"user_is_subscriber"`
		} `json:"subreddit"`
		HasVerifiedEmail bool   `json:"has_verified_email"`
		ID               string `json:"id"`
	} `json:"data"`
}
