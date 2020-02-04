package sqlmapapi

import "strconv"

type arguments []string

type option func(*arguments)

// Host allows an alternate hostname to be used by the SQLMap client
// or server.  The default is 127.0.0.1.
func Host(host string) option { //nolint:golint
	return func(args *arguments) {
		*args = append(*args, "--host="+host)
	}
}

// Port allows an alternate port number to be used by the SQLMap client
// or server.  The default is 8775.
func Port(port int16) option { //nolint:golint
	return func(args *arguments) {
		*args = append(*args, "--port="+strconv.FormatInt(int64(port), 16))
	}
}

// Username specifies (with Password) the credentials to be used between
// the SQLMap client and server using HTTP Basic Authentication.  The
// default is no authentication.
func Username(username string) option { //nolint:golint
	return func(args *arguments) {
		*args = append(*args, "--username="+username)
	}
}

// Password specifies (with Username) the credentials to be used between
// the SQLMap client and server using HTTP Basic Authentication.  The
// default is no authentication.
func Password(password string) option { //nolint:golint
	return func(args *arguments) {
		*args = append(*args, "--password="+password)
	}
}
