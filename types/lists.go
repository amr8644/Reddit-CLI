package types

type Lists struct {
	Kind string `json:"kind"`
	Data struct {
		After     string      `json:"after"`
		Dist      int         `json:"dist"`
		Modhash   interface{} `json:"modhash"`
		GeoFilter string      `json:"geo_filter"`
		Children  []struct {
			Kind string `json:"kind"`
			Data struct {
				ApprovedAtUtc         interface{}   `json:"approved_at_utc"`
				Subreddit             string        `json:"subreddit"`
				Selftext              string        `json:"selftext"`
				AuthorFullname        string        `json:"author_fullname"`
				Saved                 bool          `json:"saved"`
				ModReasonTitle        interface{}   `json:"mod_reason_title"`
				Gilded                int           `json:"gilded"`
				Clicked               bool          `json:"clicked"`
				IsGallery             bool          `json:"is_gallery"`
				Title                 string        `json:"title"`
				LinkFlairRichtext     []interface{} `json:"link_flair_richtext"`
				SubredditNamePrefixed string        `json:"subreddit_name_prefixed"`
				Hidden                bool          `json:"hidden"`
				Pwls                  int           `json:"pwls"`
				LinkFlairCSSClass     string        `json:"link_flair_css_class"`
				Downs                 int           `json:"downs"`
				ThumbnailHeight       int           `json:"thumbnail_height"`
				TopAwardedType        interface{}   `json:"top_awarded_type"`
				HideScore             bool          `json:"hide_score"`
				MediaMetadata         struct {
					ThreeAat6Ojydhta1 struct {
						Status string `json:"status"`
						E      string `json:"e"`
						M      string `json:"m"`
						P      []struct {
							Y int    `json:"y"`
							X int    `json:"x"`
							U string `json:"u"`
						} `json:"p"`
						S struct {
							Y int    `json:"y"`
							X int    `json:"x"`
							U string `json:"u"`
						} `json:"s"`
						ID string `json:"id"`
					} `json:"3aat6ojydhta1"`
					R9B4Fnjydhta1 struct {
						Status string `json:"status"`
						E      string `json:"e"`
						M      string `json:"m"`
						P      []struct {
							Y int    `json:"y"`
							X int    `json:"x"`
							U string `json:"u"`
						} `json:"p"`
						S struct {
							Y int    `json:"y"`
							X int    `json:"x"`
							U string `json:"u"`
						} `json:"s"`
						ID string `json:"id"`
					} `json:"r9b4fnjydhta1"`
				} `json:"media_metadata"`
				Name                       string      `json:"name"`
				Quarantine                 bool        `json:"quarantine"`
				LinkFlairTextColor         string      `json:"link_flair_text_color"`
				UpvoteRatio                float64     `json:"upvote_ratio"`
				AuthorFlairBackgroundColor interface{} `json:"author_flair_background_color"`
				Ups                        int         `json:"ups"`
				Domain                     string      `json:"domain"`
				MediaEmbed                 struct {
				} `json:"media_embed"`
				ThumbnailWidth        int           `json:"thumbnail_width"`
				AuthorFlairTemplateID interface{}   `json:"author_flair_template_id"`
				IsOriginalContent     bool          `json:"is_original_content"`
				UserReports           []interface{} `json:"user_reports"`
				SecureMedia           interface{}   `json:"secure_media"`
				IsRedditMediaDomain   bool          `json:"is_reddit_media_domain"`
				IsMeta                bool          `json:"is_meta"`
				Category              interface{}   `json:"category"`
				SecureMediaEmbed      struct {
				} `json:"secure_media_embed"`
				GalleryData struct {
					Items []struct {
						MediaID string `json:"media_id"`
						ID      int    `json:"id"`
					} `json:"items"`
				} `json:"gallery_data"`
				LinkFlairText       string        `json:"link_flair_text"`
				CanModPost          bool          `json:"can_mod_post"`
				Score               int           `json:"score"`
				ApprovedBy          interface{}   `json:"approved_by"`
				IsCreatedFromAdsUI  bool          `json:"is_created_from_ads_ui"`
				AuthorPremium       bool          `json:"author_premium"`
				Thumbnail           string        `json:"thumbnail"`
				Edited              bool          `json:"edited"`
				AuthorFlairCSSClass interface{}   `json:"author_flair_css_class"`
				AuthorFlairRichtext []interface{} `json:"author_flair_richtext"`
				Gildings            struct {
				} `json:"gildings"`
				ContentCategories        interface{}   `json:"content_categories"`
				IsSelf                   bool          `json:"is_self"`
				SubredditType            string        `json:"subreddit_type"`
				Created                  int           `json:"created"`
				LinkFlairType            string        `json:"link_flair_type"`
				Wls                      int           `json:"wls"`
				RemovedByCategory        interface{}   `json:"removed_by_category"`
				BannedBy                 interface{}   `json:"banned_by"`
				AuthorFlairType          string        `json:"author_flair_type"`
				TotalAwardsReceived      int           `json:"total_awards_received"`
				AllowLiveComments        bool          `json:"allow_live_comments"`
				SelftextHTML             string        `json:"selftext_html"`
				Likes                    interface{}   `json:"likes"`
				SuggestedSort            interface{}   `json:"suggested_sort"`
				BannedAtUtc              interface{}   `json:"banned_at_utc"`
				URLOverriddenByDest      string        `json:"url_overridden_by_dest"`
				ViewCount                interface{}   `json:"view_count"`
				Archived                 bool          `json:"archived"`
				NoFollow                 bool          `json:"no_follow"`
				IsCrosspostable          bool          `json:"is_crosspostable"`
				Pinned                   bool          `json:"pinned"`
				Over18                   bool          `json:"over_18"`
				AllAwardings             []interface{} `json:"all_awardings"`
				Awarders                 []interface{} `json:"awarders"`
				MediaOnly                bool          `json:"media_only"`
				LinkFlairTemplateID      string        `json:"link_flair_template_id"`
				CanGild                  bool          `json:"can_gild"`
				Spoiler                  bool          `json:"spoiler"`
				Locked                   bool          `json:"locked"`
				AuthorFlairText          interface{}   `json:"author_flair_text"`
				TreatmentTags            []interface{} `json:"treatment_tags"`
				Visited                  bool          `json:"visited"`
				RemovedBy                interface{}   `json:"removed_by"`
				ModNote                  interface{}   `json:"mod_note"`
				Distinguished            interface{}   `json:"distinguished"`
				SubredditID              string        `json:"subreddit_id"`
				AuthorIsBlocked          bool          `json:"author_is_blocked"`
				ModReasonBy              interface{}   `json:"mod_reason_by"`
				NumReports               interface{}   `json:"num_reports"`
				RemovalReason            interface{}   `json:"removal_reason"`
				LinkFlairBackgroundColor string        `json:"link_flair_background_color"`
				ID                       string        `json:"id"`
				IsRobotIndexable         bool          `json:"is_robot_indexable"`
				ReportReasons            interface{}   `json:"report_reasons"`
				Author                   string        `json:"author"`
				DiscussionType           interface{}   `json:"discussion_type"`
				NumComments              int           `json:"num_comments"`
				SendReplies              bool          `json:"send_replies"`
				WhitelistStatus          string        `json:"whitelist_status"`
				ContestMode              bool          `json:"contest_mode"`
				ModReports               []interface{} `json:"mod_reports"`
				AuthorPatreonFlair       bool          `json:"author_patreon_flair"`
				AuthorFlairTextColor     interface{}   `json:"author_flair_text_color"`
				Permalink                string        `json:"permalink"`
				ParentWhitelistStatus    string        `json:"parent_whitelist_status"`
				Stickied                 bool          `json:"stickied"`
				URL                      string        `json:"url"`
				SubredditSubscribers     int           `json:"subreddit_subscribers"`
				CreatedUtc               int           `json:"created_utc"`
				NumCrossposts            int           `json:"num_crossposts"`
				Media                    interface{}   `json:"media"`
				IsVideo                  bool          `json:"is_video"`
			} `json:"data"`
		} `json:"children"`
		Before interface{} `json:"before"`
	} `json:"data"`
}