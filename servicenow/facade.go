package servicenow

// TODO: Major refactoring required for the facade

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

func GetAccessToken(client models.Client) (*models.AccessToken, error) {
	var accessToken models.AccessToken

	// set up the client credentials
	clientRequest := url.Values{}
	clientRequest.Set("client_id", client.ClientId)
	clientRequest.Set("client_secret", client.ClientSecret)
	clientRequest.Set("username", client.UserName)
	clientRequest.Set("password", client.Password)
	clientRequest.Set("grant_type", client.GrantType)

	servicenow := URL + ACCESSTOKEN_PATH

	// call to service now to get the access token
	httpResponse, error := (&http.Client{}).PostForm(servicenow, clientRequest)
	if error != nil {
		log.Println("networking: could not retrieve the access token from service now", error)
		return &accessToken, errors.New("networking: could not retrieve the access token from service now")
	}

	if httpResponse.StatusCode == http.StatusOK {
		jsonToken, err := io.ReadAll(httpResponse.Body)
		if err != nil {
			log.Println("system: Error while parsing JSON token", err)
			return &accessToken, errors.New("system: Error while parsing JSON token")
		}

		err = json.Unmarshal(jsonToken, &accessToken)
		if err != nil {
			log.Println("system: Error while unmarshalling json token", err)
			return &accessToken, errors.New("system: Error while unmarshalling json token")
		}
		return &accessToken, nil
	} else {
		errorResponse, _ := io.ReadAll(httpResponse.Body)
		log.Println("business: No Token ", httpResponse.StatusCode, " ", string(errorResponse))
		return &accessToken, errors.New(
			fmt.Sprintf("business: No Token %s %s", httpResponse.StatusCode, string(errorResponse)))
	}

	return &accessToken, nil
}
func GetGroupByName(accessToken, name string) (*models.Group, error) {
	var group models.Group
	params := url.Values{}
	paramValue := "name=" + name
	params.Add("sysparm_query", paramValue)
	u, _ := url.ParseRequestURI(URL)
	u.Path = GROUPS_PATH
	u.RawQuery = params.Encode()
	urlStr := fmt.Sprintf("%v", u)
	request, error := http.NewRequest(http.MethodGet, urlStr, nil)

	if error != nil {
		log.Println("system: http request to query group could not be created", error)
		return nil, errors.New("system: http request to query group could not be created")
	}

	bearer := "Bearer " + accessToken
	request.Header.Set("Authorization", bearer)
	request.Header.Set("content-type", "application/json")
	httpResponse, err := (&http.Client{}).Do(request)

	if err != nil {
		log.Println("networking: could not retrieve group from service now", error)
		return &group, errors.New("networking: could not retrieve group from service now")
	}

	if httpResponse.StatusCode == http.StatusOK {
		jsonGroup, err := io.ReadAll(httpResponse.Body)
		if err != nil {
			log.Println("system: Error while parsing response body", err)
			return &group, errors.New("system: Error while parsing response body")
		}

		err = json.Unmarshal(jsonGroup, &group)

		if err != nil {
			log.Println("system: Error while unmarshalling response body", err)
			return &group, errors.New("system: Error while unmarshalling response body")
		}

		if len(group.Result) > 1 {
			log.Println("business: More than one group for onboarding")
			return &group, errors.New("business: More than one group for onboarding")
		}
		return &group, nil
	} else {
		errorResponse, _ := io.ReadAll(httpResponse.Body)
		log.Println("business: no group received ", httpResponse.StatusCode, " ", string(errorResponse))
		return &group, errors.New(
			fmt.Sprintf("business: no group received %s %s", httpResponse.StatusCode, string(errorResponse)))
	}

	return &group, nil
}

func GetUserByName(accessToken, name string) (*models.User, error) {
	var user models.User
	params := url.Values{}
	params.Add("user_name", name)
	u, _ := url.ParseRequestURI(URL)
	u.Path = USERS_PATH
	u.RawQuery = params.Encode()
	urlStr := fmt.Sprintf("%v", u)
	request, error := http.NewRequest(http.MethodGet, urlStr, nil)
	if error != nil {
		log.Println("system: http request to query user could not be created", error)
		return nil, errors.New("system: http request to query user could not be created")
	}

	bearer := "Bearer " + accessToken
	request.Header.Set("Authorization", bearer)
	request.Header.Set("content-type", "application/json")
	httpResponse, err := (&http.Client{}).Do(request)

	if err != nil {
		log.Println("networking: could not retrieve user from service now", error)
		return &user, errors.New("networking: could not retrieve user from service now")
	}

	if httpResponse.StatusCode == http.StatusOK {
		jsonGroup, err := io.ReadAll(httpResponse.Body)

		if err != nil {
			log.Println("system: Error while parsing response body", err)
			return &user, errors.New("system: Error while parsing response body")
		}
		//println(string(jsonGroup))
		err = json.Unmarshal(jsonGroup, &user)

		if err != nil {
			log.Println("system: Error while unmarshalling response body", err)
			return &user, errors.New("system: Error while unmarshalling response body")
		}

		if len(user.Result) > 1 {
			log.Println("business: system user name in service now needs to be unique", err)
			return &user, errors.New("business: system user name in service now needs to be unique")
		}
		return &user, nil
	} else {
		errorResponse, _ := io.ReadAll(httpResponse.Body)
		log.Println("business: no user received ", httpResponse.StatusCode, " ", string(errorResponse))
		return &user, errors.New(
			fmt.Sprintf("business: no user received %s %s", httpResponse.StatusCode, string(errorResponse)))
	}

	return &user, nil
}

func GetIncidentByNumber(accessToken, number string) (*models.Incident, error) {
	var incident models.Incident
	params := url.Values{}
	paramValue := "number=" + number
	params.Add("sysparm_query", paramValue)
	u, _ := url.ParseRequestURI(URL)
	u.Path = INCIDENTS_PATH
	u.RawQuery = params.Encode()
	urlStr := fmt.Sprintf("%v", u)
	request, error := http.NewRequest(http.MethodGet, urlStr, nil)
	if error != nil {
		log.Println("system: http request to query incident could not be created", error)
		return nil, errors.New("system: http request to query incident could not be created")
	}

	bearer := "Bearer " + accessToken
	request.Header.Set("Authorization", bearer)
	request.Header.Set("content-type", "application/json")
	httpResponse, err := (&http.Client{}).Do(request)

	if err != nil {
		log.Println("networking: could not retrieve incident from service now", error)
		return &incident, errors.New("networking: could not retrieve incident from service now")
	}

	if httpResponse.StatusCode == http.StatusOK {
		jsonGroup, err := io.ReadAll(httpResponse.Body)
		if err != nil {
			log.Println("system: Error while parsing response body", err)
			return &incident, errors.New("system: Error while parsing response body")
		}

		err = json.Unmarshal(jsonGroup, &incident)

		if err != nil {
			log.Println("system: Error while unmarshalling response body", err)
			return &incident, errors.New("system: Error while unmarshalling response body")
		}

		if len(incident.Result) > 1 {
			log.Println("business: incident number in service now is expected to be unique", err)
			return &incident, errors.New("incident number in service now is expected to be unique")
		}

		return &incident, nil
	} else {
		errorResponse, _ := io.ReadAll(httpResponse.Body)
		log.Println("business: No incident received ", httpResponse.StatusCode, " ", string(errorResponse))
		return &incident, errors.New(
			fmt.Sprintf("business: No incident received %s %s", httpResponse.StatusCode, string(errorResponse)))
	}

	return &incident, nil
}

func CreateIncident(serviceNowIncident *models.ServiceNowIncident) (*models.IncidentResponse, error) {
	var incident models.IncidentResponse
	var group *models.Group
	var user *models.User
	client := models.Client{
		ClientId:     CLIENT_ID,
		ClientSecret: CLIENT_SECRET,
		UserName:     USER_NAME,
		Password:     PASSWORD,
		GrantType:    "password",
	}
	token, _ := GetAccessToken(client)
	if strings.TrimSpace(serviceNowIncident.AssignedGroupName) != "" {
		var error = errors.New("")
		group, error = GetGroupByName(token.Token, serviceNowIncident.AssignedGroupName)
		if error != nil {
			log.Println("business: Cannot find the group for onboarding", error)
			return nil, errors.New(
				fmt.Sprintf("business: Cannot find the group for onboarding %s", error))
		}
	}

	if strings.TrimSpace(serviceNowIncident.CallerName) != "" {
		var error = errors.New("")
		user, error = GetUserByName(token.Token, serviceNowIncident.CallerName)
		if error != nil {
			log.Println("business: Cannot find the user for onboarding", error)
			return nil, errors.New(
				fmt.Sprintf("business: Cannot find the user for onboarding %s", error))
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
		log.Println("system: error while marshalling incident", error)
		return &incident, errors.New("system: error while marshalling incident")
	}

	urlStr := URL + INCIDENTS_PATH
	request, error := http.NewRequest(http.MethodPost, urlStr, bytes.NewReader(jsonIncident))
	if error != nil {
		log.Println("system: http request to create incident could not be created", error)
		return nil, errors.New("system: http request to create incident could not be created")
	}

	bearer := "Bearer " + token.Token
	request.Header.Set("Authorization", bearer)
	request.Header.Set("content-type", "application/json")
	httpResponse, err := (&http.Client{}).Do(request)

	if err != nil {
		log.Println("networking: could not create incident in service now", error)
		return &incident, errors.New("networking: could not create incident in service now")
	}
	if httpResponse.StatusCode == http.StatusCreated {
		jsonIncident, err := io.ReadAll(httpResponse.Body)
		if err != nil {
			log.Println("system: Error while parsing response body", err)
			return &incident, errors.New("system: Error while parsing response body")
		}

		err = json.Unmarshal(jsonIncident, &incident)

		if err != nil {
			log.Println("system: Error while unmarshalling response body", err)
			return &incident, errors.New("system: Error while unmarshalling response body")
		}

		return &incident, nil
	} else {
		errorResponse, _ := io.ReadAll(httpResponse.Body)
		log.Println("business: No incident could be created ", httpResponse.StatusCode, " ", string(errorResponse))
		return &incident, errors.New(
			fmt.Sprintf("business: No incident could be created %s %s", httpResponse.StatusCode, string(errorResponse)))
	}

	return &incident, nil
}

func AddAttachment(attachment *models.Attachment) error {
	var incident *models.Incident
	client := models.Client{
		ClientId:     CLIENT_ID,
		ClientSecret: CLIENT_SECRET,
		UserName:     USER_NAME,
		Password:     PASSWORD,
		GrantType:    "password",
	}
	token, _ := GetAccessToken(client)

	// Get the Incident ID from Incident number

	if strings.TrimSpace(attachment.IncidentNumber) != "" {
		var error = errors.New("")
		incident, error = GetIncidentByNumber(token.Token, attachment.IncidentNumber)
		if error != nil {
			log.Println("business: Cannot find the incident to attach the file", error)
			return errors.New(
				fmt.Sprintf("business: Cannot find the incident to attach the file %s", error))
		}

		if len(incident.Result) > 1 {
			log.Println("business: More than one incident for onboarding")
			return errors.New("business: More than one incident for onboarding")
		}

	} else {
		return errors.New("business: Incident number is required for attaching the file")
	}

	params := url.Values{}
	params.Add("table_name", "incident")
	params.Add("table_sys_id", (*incident).Result[0].ID)
	params.Add("file_name", attachment.FileName)
	u, _ := url.ParseRequestURI(URL)
	u.Path = ATTACHMENT_PATH
	u.RawQuery = params.Encode()
	urlStr := fmt.Sprintf("%v", u)

	request, error := http.NewRequest(http.MethodPost, urlStr, bytes.NewReader(attachment.Content))
	if error != nil {
		log.Println("system: http request to add attachment to incident could not created", error)
		return errors.New("system: http request to add attachment to incident could not created")
	}

	bearer := "Bearer " + token.Token
	request.Header.Set("Authorization", bearer)
	request.Header.Set("content-type", "application/json")
	httpResponse, err := (&http.Client{}).Do(request)

	if err != nil {
		log.Println("networking: could not add attachment to the incident", err)
		return errors.New("networking: could not add attachment to the incident")
	}
	if httpResponse.StatusCode == http.StatusCreated {
		_, err := io.ReadAll(httpResponse.Body)
		if err != nil {
			log.Println("system: Error while parsing the response body", err)
			return errors.New("system: Error while parsing the response body")
		}

		return nil
	} else {
		errorResponse, _ := io.ReadAll(httpResponse.Body)
		log.Println("business: no attachment created ", httpResponse.StatusCode, " ", string(errorResponse))
		return errors.New(
			fmt.Sprintf("business: no attachment created %s %s", httpResponse.StatusCode, string(errorResponse)))
	}

	return nil
}
