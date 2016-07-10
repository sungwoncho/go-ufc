/*
Package ufc provies an API client for Ultimate Fighting Championship's
API v3.

Get an API client by using New()

func main() {
	client := ufc.New()
}

client has methods that calls UFC's API, unmarshalls JSON response,
and returns a struct representing the response.
*/
package ufc
