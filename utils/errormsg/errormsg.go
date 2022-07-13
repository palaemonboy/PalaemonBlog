package errormsg

const (
	SUCCESS = 200
	ERROR   = 500

	//code 1000 UserErrors
	ErrorUserNameUsed     = 1001
	ErrorPasswordWrong    = 1002
	ErrorUserNotExist     = 1003
	ErrorTokenExist       = 1004
	ErrorTokenRuntime     = 1005
	ErrorTokenWrong       = 1006
	ErrorTokenTypeWrong   = 1007
	ErrorUserNoPermission = 1008

	// code = 2000 Article Errors
	ErrorArticleNotExist = 2001

	// code = 3000 Category Errors
	ErrorCategoryUsed     = 3001
	ErrorCategoryNotExist = 3002
)

var codeMsg = map[int]string{
	SUCCESS:               "OK",
	ERROR:                 "FAILED",
	ErrorUserNameUsed:     "Username already exists!",
	ErrorPasswordWrong:    "Password error",
	ErrorUserNotExist:     "User does not exist",
	ErrorTokenExist:       "TOKEN does not exist, please log in again",
	ErrorTokenRuntime:     "TOKEN has expired, please log in again",
	ErrorTokenWrong:       "TOKEN is incorrect, please log in again",
	ErrorTokenTypeWrong:   "TOKEN format is wrong, please log in again",
	ErrorUserNoPermission: "The user does not have permission",

	ErrorArticleNotExist: "Article does not exist",

	ErrorCategoryUsed:     "This category already exists",
	ErrorCategoryNotExist: "This category does not exist",
}

func GetErrMsg(code int) string {
	return codeMsg[code]
}
