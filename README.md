# Freshbooks API CLI

Bored of using GUI to add time entries to Freshbooks? This should help.

*NOTE:* This is a starting point. You can help by adding more methods, organizing code and implementing functionality and/or tests.

## Using it
### Setup
* export `AUTHENTICATION_TOKEN` and `FRESHBOOKS_API_URL` with values taken from `MyAccount > FreshBooks API page`. It has general address "https://<YOUR APP>.freshbooks.com/apiEnable".


### Listing projects
* `go run cmd/freshbooks-list-projects/main.go` - you should be presented with list of your projects w/ tasks. If you have more than 25 projects, or more than 25 tasks, it'll be truncated. Happy coding ;)

### Creating time entries
* Create CSV file with time entries named `entries.csv`, e.g:
```
project_name,task_name,hours,notes,date
Labs,General,1.5,"[Feature] Foo",2018-06-11
SecretProject,Learning,1.5,"[Feature] Bar",2018-06-12
WorldDomination,Research,4.5,"[Feature] And this CLI",2018-06-13
```

* `go run cmd/freshbooks-create-entries/main.go` - it'll create entries from file.
