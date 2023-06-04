package servicenow

// TODO: host url and API paths

const (
	URL              string = "https://dev124376.service-now.com"
	ACCESSTOKEN_PATH string = "/oauth_token.do"
	GROUPS_PATH      string = "/api/now/table/sys_user_group"
	USERS_PATH       string = "/api/now/table/sys_user"

	INCIDENTS_PATH string = "/api/now/table/incident"

	ATTACHMENT_PATH = "/api/now/attachment/file"
)
