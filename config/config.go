package config

// Add config variables to here
// will read config variables into here
// from a file
type configStruct struct {
	ListenPort   string // port to listen on
	LogFile      string // file to log requests
	Quiet        bool   // be quiet?
	Verb         bool   // for debugging
	ProxyVerbose bool   // setting for proxy Verbosity
	LogRequests  bool   //Log requests? not yet implemented
}

func ReadConfigFile(filename string) configStruct {
	return configStruct{}
}

func DefaultConfig() *configStruct {
	return &configStruct{
		ListenPort:   "9000",
		LogFile:      "/etc/adproxy/log/proxylog.log",
		Quiet:        true,
		Verb:         false,
		ProxyVerbose: false}
}
