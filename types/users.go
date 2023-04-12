package types

type UserAbout struct {
	Kind string `json:"kind"`
	Data struct {
		IsEmployee bool `json:"is_employee"`
		IsFriend   bool `json:"is_friend"`
		Subreddit  struct {
			DefaultSet                 bool          `json:"default_set"`
			UserIsContributor          bool          `json:"user_is_contributor"`
			BannerImg                  string        `json:"banner_img"`
			AllowedMediaInComments     []interface{} `json:"allowed_media_in_comments"`
			UserIsBanned               bool          `json:"user_is_banned"`
			FreeFormReports            bool          `json:"free_form_reports"`
			CommunityIcon              interface{}   `json:"community_icon"`
			ShowMedia                  bool          `json:"show_media"`
			IconColor                  string        `json:"icon_color"`
			UserIsMuted                interface{}   `json:"user_is_muted"`
			DisplayName                string        `json:"display_name"`
			HeaderImg                  interface{}   `json:"header_img"`
			Title                      string        `json:"title"`
			PreviousNames              []interface{} `json:"previous_names"`
			Over18                     bool          `json:"over_18"`
			IconSize                   []int         `json:"icon_size"`
			PrimaryColor               string        `json:"primary_color"`
			IconImg                    string        `json:"icon_img"`
			Description                string        `json:"description"`
			SubmitLinkLabel            string        `json:"submit_link_label"`
			HeaderSize                 interface{}   `json:"header_size"`
			RestrictPosting            bool          `json:"restrict_posting"`
			RestrictCommenting         bool          `json:"restrict_commenting"`
			Subscribers                int           `json:"subscribers"`
			SubmitTextLabel            string        `json:"submit_text_label"`
			IsDefaultIcon              bool          `json:"is_default_icon"`
			LinkFlairPosition          string        `json:"link_flair_position"`
			DisplayNamePrefixed        string        `json:"display_name_prefixed"`
			KeyColor                   string        `json:"key_color"`
			Name                       string        `json:"name"`
			IsDefaultBanner            bool          `json:"is_default_banner"`
			URL                        string        `json:"url"`
			Quarantine                 bool          `json:"quarantine"`
			BannerSize                 []int         `json:"banner_size"`
			UserIsModerator            bool          `json:"user_is_moderator"`
			AcceptFollowers            bool          `json:"accept_followers"`
			PublicDescription          string        `json:"public_description"`
			LinkFlairEnabled           bool          `json:"link_flair_enabled"`
			DisableContributorRequests bool          `json:"disable_contributor_requests"`
			SubredditType              string        `json:"subreddit_type"`
			UserIsSubscriber           bool          `json:"user_is_subscriber"`
		} `json:"subreddit"`
		SnoovatarSize     []int   `json:"snoovatar_size"`
		AwardeeKarma      int     `json:"awardee_karma"`
		ID                string  `json:"id"`
		Verified          bool    `json:"verified"`
		IsGold            bool    `json:"is_gold"`
		IsMod             bool    `json:"is_mod"`
		AwarderKarma      int     `json:"awarder_karma"`
		HasVerifiedEmail  bool    `json:"has_verified_email"`
		IconImg           string  `json:"icon_img"`
		HideFromRobots    bool    `json:"hide_from_robots"`
		LinkKarma         int     `json:"link_karma"`
		PrefShowSnoovatar bool    `json:"pref_show_snoovatar"`
		IsBlocked         bool    `json:"is_blocked"`
		TotalKarma        int     `json:"total_karma"`
		AcceptChats       bool    `json:"accept_chats"`
		Name              string  `json:"name"`
		Created           float64 `json:"created"`
		CreatedUtc        float64 `json:"created_utc"`
		SnoovatarImg      string  `json:"snoovatar_img"`
		CommentKarma      int     `json:"comment_karma"`
		AcceptFollowers   bool    `json:"accept_followers"`
		HasSubscribed     bool    `json:"has_subscribed"`
		AcceptPms         bool    `json:"accept_pms"`
	} `json:"data"`
}