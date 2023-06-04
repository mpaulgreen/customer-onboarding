package main

import (
	"customer-onboarding/models"
	"customer-onboarding/servicenow"
	"encoding/json"
)

func main() {
	//router.SetUpRoutes()
	//err := http.ListenAndServe(":3000", nil)
	//if err != nil {
	//	log.Fatal(err)
	//}

	//client := models.Client{
	//	ClientId:     "63a59937dccfa5108277140ed7efa4ca",
	//	ClientSecret: "8&E0^7``mm",
	//	UserName:     "fedramp",
	//	Password:     "Redhat123",
	//	GrantType:    "password",
	//}
	//token, _ := servicenow.GetAccessToken(client)
	//b, _ := json.Marshal(token)
	//println("From Main", string(b))

	//name := "Onboarding"
	//group, _ := servicenow.GetGroupByName(name)
	//fmt.Printf("v ==== %v \n", group)
	//b, _ := json.Marshal(group)
	//println("From Main", string(b))

	//name := "survey.user"
	//group, _ := servicenow.GetUserByName(name)
	//fmt.Printf("v ==== %v \n", group)
	//b, _ := json.Marshal(group)
	//println("From Main", string(b))

	//number := "INC0010017"
	//group, _ := servicenow.GetIncidentByNumber(number)
	//fmt.Printf("v ==== %v \n", group)
	//b, _ := json.Marshal(group)
	//println("From Main", string(b))

	incident, _ := servicenow.CreateIncident(&models.ServiceNowIncident{
		ShortDescription:  "Sample Incident",
		Urgency:           "2",
		Impact:            "2",
		Description:       "Please oboard customer foo, admin: firstName: John, lastName: Doe,email: John.Doe@gmail.com, ClientId: foo, ClientSecret: xxxxxxxx,Discovery Endpoint: https://accounts.google.com/.well-known/openid-configuration",
		ContactType:       "mpaul@redhat.com",
		CallerName:        "survey.user",
		AssignedGroupName: "Onboarding",
	})

	b, _ := json.Marshal(incident)
	println("From Main", string(b))
}
