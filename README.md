# /incidents
```agsl
{
	"short_description":  "Sample Incident",
	"urgency":           "2",
	"impact":            "2",
	"description":       "Please oboard customer foo, admin: firstName: John, lastName: Doe,email: John.Doe@gmail.com, ClientId: foo, ClientSecret: xxxxxxxx,Discovery Endpoint: https://accounts.google.com/.well-known/openid-configuration",
	"contact_type":       "mpaul@redhat.com",
	"caller_id":        "survey.user",
	"assignment_group": "Onboarding"
}
```
```agsl
{
	"short_description":  "Sample Incident",
	"urgency":           "2",
	"impact":            "2",
	"description":       "Please oboard customer foo, admin: firstName: John, lastName: Doe,email: John.Doe@gmail.com, ClientId: foo, ClientSecret: xxxxxxxx,Discovery Endpoint: https://accounts.google.com/.well-known/openid-configuration",
	"contact_type":       "mpaul@redhat.com"
}
```

# /upload
```agsl
curl --location --request POST 'http://localhost:3000/upload?file_name=foo.yaml&number=INC0010036' \
--header 'Content-Type: text/yaml' \
--data-binary '@/Users/mrigankapaul/Downloads/foo.yaml'
```