package vtgroups

type VtGroup string

const (
	General          = VtGroup("General")
	ProductDetection = VtGroup("Product detection")
	DefaultAccounts  = VtGroup("Default Accounts")
	ShellRemotely    = VtGroup("Gain a shell remotely")
	BruteLogins      = VtGroup("Brute force attacks")
	PortScanner      = VtGroup("Port scanners")
	ServiceDetection = VtGroup("Service detection")
)
