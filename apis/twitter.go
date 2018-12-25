// Copyright 2018 by caixw, All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package apis

// https://developer.twitter.com/en/docs/api-reference-index

import "net/http"

func init() {
	twitter.init()
}

// https://developer.github.com/v3/
var twitter = &Collection{
	Name: "Twitter API",
	Desc: "采集自 Twitter 的 API，以静态路由为主。",
	APIS: []*API{
		//Authentication
		{Method: http.MethodGet, Brace: "/oauth/authenticate"},
		{Method: http.MethodGet, Brace: "/oauth/authorize"},
		{Method: http.MethodPost, Brace: "/oauth/access_token"},
		{Method: http.MethodPost, Brace: "/oauth/invalidate_token"},
		{Method: http.MethodPost, Brace: "/oauth/request_token"},
		{Method: http.MethodPost, Brace: "/oauth2/invalidate_token"},
		{Method: http.MethodPost, Brace: "/oauth2/token"},

		// Create,Brace and manage lists
		{Method: http.MethodGet, Brace: "/lists/list"},
		{Method: http.MethodGet, Brace: "/lists/members"},
		{Method: http.MethodGet, Brace: "/lists/members/show"},
		{Method: http.MethodGet, Brace: "/lists/memberships"},
		{Method: http.MethodGet, Brace: "/lists/ownerships"},
		{Method: http.MethodGet, Brace: "/lists/show"},
		{Method: http.MethodGet, Brace: "/lists/statuses"},
		{Method: http.MethodGet, Brace: "/lists/subscribers"},
		{Method: http.MethodGet, Brace: "/lists/subscribers/show"},
		{Method: http.MethodGet, Brace: "/lists/subscriptions"},
		{Method: http.MethodPost, Brace: "/lists/create"},
		{Method: http.MethodPost, Brace: "/lists/destroy"},
		{Method: http.MethodPost, Brace: "/lists/members/create"},
		{Method: http.MethodPost, Brace: "/lists/members/create_all"},
		{Method: http.MethodPost, Brace: "/lists/members/destroy"},
		{Method: http.MethodPost, Brace: "/lists/members/destroy_all"},
		{Method: http.MethodPost, Brace: "/lists/subscribers/create"},
		{Method: http.MethodPost, Brace: "/lists/subscribers/destroy"},
		{Method: http.MethodPost, Brace: "/lists/update"},

		// Follow, search, and get users
		{Method: http.MethodGet, Brace: "/followers/ids"},
		{Method: http.MethodGet, Brace: "/followers/list"},
		{Method: http.MethodGet, Brace: "/friends/ids"},
		{Method: http.MethodGet, Brace: "/friends/list"},
		{Method: http.MethodGet, Brace: "/friendships/incoming"},
		{Method: http.MethodGet, Brace: "/friendships/lookup"},
		{Method: http.MethodGet, Brace: "/friendships/no_retweets/ids"},
		{Method: http.MethodGet, Brace: "/friendships/outgoing"},
		{Method: http.MethodGet, Brace: "/friendships/show"},
		{Method: http.MethodGet, Brace: "/users/lookup"},
		{Method: http.MethodGet, Brace: "/users/search"},
		{Method: http.MethodGet, Brace: "/users/show"},
		{Method: http.MethodGet, Brace: "/users/suggestions"},
		{Method: http.MethodGet, Brace: "/users/suggestions/{slug}"},
		{Method: http.MethodGet, Brace: "/users/suggestions/{slug}/members"},
		{Method: http.MethodPost, Brace: "/friendships/create"},
		{Method: http.MethodPost, Brace: "/friendships/destroy"},
		{Method: http.MethodPost, Brace: "/friendships/update"},

		// {Method:http.MethodManage,Brace:"account settings and profile"},
		{Method: http.MethodGet, Brace: "/account/settings"},
		{Method: http.MethodGet, Brace: "/account/verify_credentials"},
		{Method: http.MethodGet, Brace: "/saved_searches/list"},
		{Method: http.MethodGet, Brace: "/saved_searches/show/{id}"},
		{Method: http.MethodGet, Brace: "/users/profile_banner"},
		{Method: http.MethodPost, Brace: "/account/remove_profile_banner"},
		{Method: http.MethodPost, Brace: "/account/settings"},
		{Method: http.MethodPost, Brace: "/account/update_profile"},
		{Method: http.MethodPost, Brace: "/account/update_profile_background_image"},
		{Method: http.MethodPost, Brace: "/account/update_profile_banner"},
		{Method: http.MethodPost, Brace: "/account/update_profile_image"},
		{Method: http.MethodPost, Brace: "/saved_searches/create"},
		{Method: http.MethodPost, Brace: "/saved_searches/destroy/{id}"},

		// {Method:http.MethodMute,,Brace:"block and report users"},
		{Method: http.MethodGet, Brace: "/blocks/ids"},
		{Method: http.MethodGet, Brace: "/blocks/list"},
		{Method: http.MethodGet, Brace: "/mutes/users/ids"},
		{Method: http.MethodGet, Brace: "/mutes/users/list"},
		{Method: http.MethodPost, Brace: "/blocks/create"},
		{Method: http.MethodPost, Brace: "/blocks/destroy"},
		{Method: http.MethodPost, Brace: "/mutes/users/create"},
		{Method: http.MethodPost, Brace: "/mutes/users/destroy"},
		{Method: http.MethodPost, Brace: "/users/report_spam"},

		// {Method:http.MethodCurate,Brace:"a collection of Tweets"},
		{Method: http.MethodGet, Brace: "/collections/entries"},
		{Method: http.MethodGet, Brace: "/collections/list"},
		{Method: http.MethodGet, Brace: "/collections/show"},
		{Method: http.MethodPost, Brace: "/collections/create"},
		{Method: http.MethodPost, Brace: "/collections/destroy"},
		{Method: http.MethodPost, Brace: "/collections/entries/add"},
		{Method: http.MethodPost, Brace: "/collections/entries/curate"},
		{Method: http.MethodPost, Brace: "/collections/entries/move"},
		{Method: http.MethodPost, Brace: "/collections/entries/remove"},
		{Method: http.MethodPost, Brace: "/collections/update"},

		// {Method: http.MethodReplay, Brace: "/API"},
		{Method: http.MethodPost, Brace: "/statuses/filter"},

		// {Method:http.MethodGet,Brace:"Tweet timelines"},
		{Method: http.MethodGet, Brace: "/statuses/home_timeline"},
		{Method: http.MethodGet, Brace: "/statuses/mentions_timeline"},
		{Method: http.MethodGet, Brace: "/statuses/user_timeline"},

		// {Method:http.MethodPost,,Brace:"retrieve and engage with Tweets"},
		{Method: http.MethodGet, Brace: "/favorites/list"},
		{Method: http.MethodGet, Brace: "/statuses/lookup"},
		{Method: http.MethodGet, Brace: "/statuses/oembed"},
		{Method: http.MethodGet, Brace: "/statuses/retweeters/ids"},
		{Method: http.MethodGet, Brace: "/statuses/retweets/{id}"},
		{Method: http.MethodGet, Brace: "/statuses/retweets_of_me"},
		{Method: http.MethodGet, Brace: "/statuses/show/{id}"},
		{Method: http.MethodPost, Brace: "/favorites/create"},
		{Method: http.MethodPost, Brace: "/favorites/destroy"},
		{Method: http.MethodPost, Brace: "/statuses/destroy/{id}"},
		{Method: http.MethodPost, Brace: "/statuses/retweet/{id}"},
		{Method: http.MethodPost, Brace: "/statuses/unretweet/{id}"},
		{Method: http.MethodPost, Brace: "/statuses/update"},
		{Method: http.MethodPost, Brace: "/statuses/update_with_media"},

		// {Method: http.MethodDecahose, Brace: "/stream"},
		{Method: http.MethodGet, Brace: "/statuses/sample"},

		// {Method:http.MethodTweet,Brace:"compliance"},
		{Method: http.MethodGet, Brace: "/compliance/firehose"},

		// {Method:http.MethodCustom,Brace:"profiles"},
		{Method: http.MethodDelete, Brace: "/custom_profiles/destroy.json"},
		{Method: http.MethodGet, Brace: "/custom_profiles/{id}"},
		{Method: http.MethodGet, Brace: "/custom_profiles/list"},
		{Method: http.MethodPost, Brace: "/custom_profiles/new.json"},

		// Customer feedback cards
		{Method: http.MethodGet, Brace: "/feedback/events.json"},
		{Method: http.MethodGet, Brace: "/feedback/show/{id}.json"},
		{Method: http.MethodPost, Brace: "/feedback/create.json"},

		// Sending,Brace and receiving events
		{Method: http.MethodDelete, Brace: "/direct_messages/events/destroy"},
		{Method: http.MethodGet, Brace: "/direct_messages/events/list"},
		{Method: http.MethodGet, Brace: "/direct_messages/events/show"},
		{Method: http.MethodPost, Brace: "/direct_messages/events/new"},

		// {Method:http.MethodTyping,Brace:"indicator and read receipts"},
		{Method: http.MethodPost, Brace: "/direct_messages/indicate_typing"},
		{Method: http.MethodPost, Brace: "/direct_messages/mark_read"},

		// {Method:http.MethodWelcome,Brace:"Messages"},
		{Method: http.MethodDelete, Brace: "/direct_messages/welcome_messages/destroy"},
		{Method: http.MethodDelete, Brace: "/direct_messages/welcome_messages/rules/destroy"},
		{Method: http.MethodPut, Brace: "/direct_messages/welcome_messages/update"},
		{Method: http.MethodGet, Brace: "/direct_messages/welcome_messages/list"},
		{Method: http.MethodGet, Brace: "/direct_messages/welcome_messages/rules/list"},
		{Method: http.MethodGet, Brace: "/direct_messages/welcome_messages/rules/show"},
		{Method: http.MethodGet, Brace: "/direct_messages/welcome_messages/show"},
		{Method: http.MethodPost, Brace: "/direct_messages/welcome_messages/new"},
		{Method: http.MethodPost, Brace: "/direct_messages/welcome_messages/rules/new"},

		// {Method:http.MethodUpload,Brace:"media"},
		{Method: http.MethodGet, Brace: "/media/upload"},
		{Method: http.MethodPost, Brace: "/media/metadata/create"},
		{Method: http.MethodPost, Brace: "/media/subtitles/create"},
		{Method: http.MethodPost, Brace: "/media/subtitles/delete"},
		{Method: http.MethodPost, Brace: "/media/upload"},

		// {Method:http.MethodGet,Brace:"locations with trending topics"},
		{Method: http.MethodGet, Brace: "/trends/available"},
		{Method: http.MethodGet, Brace: "/trends/closest"},

		// {Method:http.MethodGet,Brace:"information about a place"},
		{Method: http.MethodGet, Brace: "/geo/id/{place_id}"},

		// {Method:http.MethodGet,Brace:"places near a location"},
		{Method: http.MethodGet, Brace: "/geo/reverse_geocode"},
		{Method: http.MethodGet, Brace: "/geo/search"},
	},
}
