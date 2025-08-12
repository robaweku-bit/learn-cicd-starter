func GetAPIKey(r *http.Request) (string, error) {
    // something like:
    // - get "Authorization" header
    // - check if it starts with "ApiKey "
    // - return the key or an error
}
