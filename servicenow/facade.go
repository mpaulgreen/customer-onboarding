package servicenow

import (
	"bytes"
	"customer-onboarding/models"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
)

const (
	SERVICENOW      string = "https://dev124376.service-now.com"
	ACCESSTOKENPATH string = "/oauth_token.do"
	GROUPSPATH      string = "/api/now/table/sys_user_group"
	USERSPATH       string = "/api/now/table/sys_user"

	INCIDENTSPATH string = "/api/now/table/incident"
)

func getHttpClient() http.Client {
	return http.Client{}
}
func GetAccessToken(client models.Client) (*models.AccessToken, error) {
	var accessToken models.AccessToken

	clientRequest := url.Values{}
	clientRequest.Set("client_id", client.ClientId)
	clientRequest.Set("client_secret", client.ClientSecret)
	clientRequest.Set("username", client.UserName)
	clientRequest.Set("password", client.Password)
	clientRequest.Set("grant_type", client.GrantType)

	servicenow := SERVICENOW + ACCESSTOKENPATH
	httpResponse, error := (&http.Client{}).PostForm(servicenow, clientRequest)
	if error != nil {
		log.Println("Error in calling service now to get the token", error)
		return &accessToken, error
	}
	if httpResponse.StatusCode == http.StatusOK {
		jsonToken, err := io.ReadAll(httpResponse.Body)
		if err != nil {
			log.Println(err)
			return &accessToken, err
		}

		err = json.Unmarshal(jsonToken, &accessToken)
		if err != nil {
			log.Println("Error in unmarshalling json token", err)
			return &accessToken, err
		}
		return &accessToken, nil
	} else {
		errorResponse, _ := io.ReadAll(httpResponse.Body)
		log.Println("Bad Token", errorResponse)
	}

	return &accessToken, nil
}
func GetGroupByName(accessToken, name string) (*models.Group, error) {

	var group models.Group
	//client := models.Client{
	//	ClientId:     "63a59937dccfa5108277140ed7efa4ca",
	//	ClientSecret: "8&E0^7``mm",
	//	UserName:     "fedramp",
	//	Password:     "Redhat123",
	//	GrantType:    "password",
	//}
	//token, _ := GetAccessToken(client)
	////b, _ := json.Marshal(token)
	////println(string(b))
	//
	// TODO: The above code needs to be removed

	params := url.Values{}
	paramValue := "name=" + name
	params.Add("sysparm_query", paramValue)
	u, _ := url.ParseRequestURI(SERVICENOW)
	groupspath := GROUPSPATH
	u.Path = groupspath
	u.RawQuery = params.Encode()
	urlStr := fmt.Sprintf("%v", u)
	request, error := http.NewRequest(http.MethodGet, urlStr, nil)
	if error != nil {
		println(" Error in creating request")
		log.Println("Error in creating request", error)
		return nil, error
	}

	bearer := "Bearer " + accessToken
	request.Header.Set("Authorization", bearer)
	request.Header.Set("content-type", "application/json")
	httpResponse, err := (&http.Client{}).Do(request)

	if err != nil {
		println(" Error in calling service now to get the group")
		log.Println("Error in calling service now to get the group", error)
		return &group, error
	}
	if httpResponse.StatusCode == http.StatusOK {
		jsonGroup, err := io.ReadAll(httpResponse.Body)
		//println(" jsonGroupBody", string(jsonGroup))
		if err != nil {
			log.Println(err)
			return &group, err
		}
		//println(string(jsonGroup))
		err = json.Unmarshal(jsonGroup, &group)

		if err != nil {
			//println("Error in unmarshalling json token", err)
			log.Println("Error in unmarshalling json token", err)
			return &group, err
		}

		return &group, nil
	} else {
		errorResponse, _ := io.ReadAll(httpResponse.Body)
		log.Println("Bad Group", errorResponse)
	}

	return &group, nil
}

func GetUserByName(accessToken, name string) (*models.User, error) {
	var user models.User
	//client := models.Client{
	//	ClientId:     "63a59937dccfa5108277140ed7efa4ca",
	//	ClientSecret: "8&E0^7``mm",
	//	UserName:     "fedramp",
	//	Password:     "Redhat123",
	//	GrantType:    "password",
	//}
	//token, _ := GetAccessToken(client)
	//b, _ := json.Marshal(token)
	//println(string(b))

	// TODO: The above code needs to be removed

	params := url.Values{}
	params.Add("user_name", name)
	u, _ := url.ParseRequestURI(SERVICENOW)
	userspath := USERSPATH
	u.Path = userspath
	u.RawQuery = params.Encode()
	urlStr := fmt.Sprintf("%v", u)
	request, error := http.NewRequest(http.MethodGet, urlStr, nil)
	if error != nil {
		log.Println("Error in creating request", error)
		return nil, error
	}

	bearer := "Bearer " + accessToken
	request.Header.Set("Authorization", bearer)
	request.Header.Set("content-type", "application/json")
	httpResponse, err := (&http.Client{}).Do(request)

	if err != nil {
		println(" Error in calling service now to get the user")
		log.Println("Error in calling service now to get the user", error)
		return &user, error
	}
	if httpResponse.StatusCode == http.StatusOK {
		jsonGroup, err := io.ReadAll(httpResponse.Body)
		//println(" jsonGroupBody", string(jsonGroup))
		if err != nil {
			log.Println(err)
			return &user, err
		}
		//println(string(jsonGroup))
		err = json.Unmarshal(jsonGroup, &user)

		if err != nil {
			//println("Error in unmarshalling json token", err)
			log.Println("Error in unmarshalling json token", err)
			return &user, err
		}

		return &user, nil
	} else {
		errorResponse, _ := io.ReadAll(httpResponse.Body)
		log.Println("Bad User", errorResponse)
	}

	return &user, nil
}

func GetIncidentByNumber(number string) (*models.Incident, error) {
	var incident models.Incident
	client := models.Client{
		ClientId:     "63a59937dccfa5108277140ed7efa4ca",
		ClientSecret: "8&E0^7``mm",
		UserName:     "fedramp",
		Password:     "Redhat123",
		GrantType:    "password",
	}
	token, _ := GetAccessToken(client)
	//b, _ := json.Marshal(token)
	//println(string(b))

	// TODO: The above code needs to be removed

	params := url.Values{}
	paramValue := "number=" + number
	params.Add("sysparm_query", paramValue)
	u, _ := url.ParseRequestURI(SERVICENOW)
	groupspath := INCIDENTSPATH
	u.Path = groupspath
	u.RawQuery = params.Encode()
	urlStr := fmt.Sprintf("%v", u)
	request, error := http.NewRequest(http.MethodGet, urlStr, nil)
	if error != nil {
		println(" Error in creating request")
		log.Println("Error in creating request", error)
		return nil, error
	}

	bearer := "Bearer " + token.Token
	request.Header.Set("Authorization", bearer)
	request.Header.Set("content-type", "application/json")
	httpResponse, err := (&http.Client{}).Do(request)

	if err != nil {
		println(" Error in calling service now to get the group")
		log.Println("Error in calling service now to get the group", error)
		return &incident, error
	}
	if httpResponse.StatusCode == http.StatusOK {
		jsonGroup, err := io.ReadAll(httpResponse.Body)
		//println(" jsonGroupBody", string(jsonGroup))
		if err != nil {
			log.Println(err)
			return &incident, err
		}
		//println(string(jsonGroup))
		err = json.Unmarshal(jsonGroup, &incident)

		if err != nil {
			//println("Error in unmarshalling json token", err)
			log.Println("Error in unmarshalling json token", err)
			return &incident, err
		}

		return &incident, nil
	} else {
		errorResponse, _ := io.ReadAll(httpResponse.Body)
		log.Println("Bad Incident", errorResponse)
	}

	return &incident, nil
}

func CreateIncident(serviceNowIncident *models.ServiceNowIncident) (*models.IncidentResponse, error) {
	var incident models.IncidentResponse
	var group *models.Group
	var user *models.User
	client := models.Client{
		ClientId:     "63a59937dccfa5108277140ed7efa4ca",
		ClientSecret: "8&E0^7``mm",
		UserName:     "fedramp",
		Password:     "Redhat123",
		GrantType:    "password",
	}
	token, _ := GetAccessToken(client)
	//b, _ := json.Marshal(token)
	//println(string(b))

	if strings.TrimSpace(serviceNowIncident.AssignedGroupName) != "" {
		var error = errors.New("")
		group, error = GetGroupByName(token.Token, serviceNowIncident.AssignedGroupName)
		if error != nil {
			log.Println("Error in Getting Group from Service Now", error)
			return nil, error
		}
	}

	if strings.TrimSpace(serviceNowIncident.CallerName) != "" {
		var error = errors.New("")
		user, error = GetUserByName(token.Token, serviceNowIncident.CallerName)
		if error != nil {
			log.Println("Error in Getting User from Service Now", error)
			return nil, error
		}
	}

	if group != nil {
		serviceNowIncident.AssignedGroupName = (*group).Result[0].ID // TODO
	}

	if user != nil {
		serviceNowIncident.CallerName = (*user).Result[0].ID // TODO
	}

	jsonIncident, error := json.Marshal(serviceNowIncident)
	if error != nil {
		fmt.Println("Error in Marshalling service now incident", error)
	}

	urlStr := SERVICENOW + INCIDENTSPATH
	request, error := http.NewRequest(http.MethodPost, urlStr, bytes.NewReader(jsonIncident))
	if error != nil {
		println(" Error in creating request")
		log.Println("Error in creating request", error)
		return nil, error
	}

	bearer := "Bearer " + token.Token
	request.Header.Set("Authorization", bearer)
	request.Header.Set("content-type", "application/json")
	httpResponse, err := (&http.Client{}).Do(request)

	if err != nil {
		println(" Error in calling service now to get the group")
		log.Println("Error in calling service now to get the group", error)
		return &incident, error
	}
	if httpResponse.StatusCode == http.StatusCreated {
		jsonIncident, err := io.ReadAll(httpResponse.Body)
		if err != nil {
			log.Println(err)
			return &incident, err
		}
		println(string(jsonIncident))

		err = json.Unmarshal(jsonIncident, &incident)

		if err != nil {
			println("Error in unmarshalling json incident", err)
			log.Println("Error in unmarshalling json incident", err)
			return &incident, err
		}

		return &incident, nil
	} else {
		errorResponse, _ := io.ReadAll(httpResponse.Body)
		log.Println("Bad Incident Creation", string(errorResponse))
	}

	return &incident, nil
}

func AddAttachment(attachment *models.Attachment) (models.Attachment, error) {
	// TODO: Implementation pending
	// REST call to service now
	return models.Attachment{}, nil
}
