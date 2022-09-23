package utils

var (
	IdVerify             = Rules{"ID": []string{NotEmpty()}}
	LoginVerify          = Rules{"Username": {NotEmpty()}, "Password": {NotEmpty()}, "AuthorityId": {NotEmpty()}}
	RegisterVerify       = Rules{"Username": {NotEmpty()}, "NickName": {NotEmpty()}, "Password": {NotEmpty()}}
	PageInfoVerify       = Rules{"Page": {NotEmpty()}, "PageSize": {NotEmpty()}}
	ChangePasswordVerify = Rules{"Password": {NotEmpty()}, "NewPassword": {NotEmpty()}}
)
