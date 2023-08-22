package consts

const (
	// 局部钩子
	PreResolve          = "preResolve"
	MutatingPreResolve  = "mutatingPreResolve"
	MockResolve         = "mockResolve"
	CustomResolve       = "customResolve"
	PostResolve         = "postResolve"
	MutatingPostResolve = "mutatingPostResolve"

	// 认证钩子
	PostAuthentication         = "postAuthentication"
	MutatingPostAuthentication = "mutatingPostAuthentication"
	RevalidateAuthentication   = "revalidate"
	PostLogout                 = "postLogout"

	// 全局钩子
	HttpTransportBeforeRequest = "beforeRequest"
	HttpTransportOnRequest     = "onRequest"
	HttpTransportOnResponse    = "onResponse"
)
