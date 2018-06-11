# Freshbooks API CLI

Bored of using GUI to add time entries to Freshbooks? This should help.

*NOTE:* This is a starting point. You can help by adding more methods, organizing code and implementing functionality and/or tests.

## Using it
* export `AUTHENTICATION_TOKEN` and `FRESHBOOKS_API_URL` with values taken from `MyAccount > FreshBooks API page`. It has general address `https://<YOUR APP>.freshbooks.com/apiEnable`.
* `go run main.go` - you should be presented with list of your projects w/ tasks. If you have more than 25 projects, or more than 25 tasks, it'll be truncated. Happy coding ;)
