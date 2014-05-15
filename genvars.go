package genvars

import (
	"errors"
	"os"
)

//const

type EnvManager struct {
	// The prefix to use for the env vars
	prefix string
	// The variable to detect current environment
	envVar string
	// The current enviroment we're in (ie, prod, dev, etc..)
	currEnv string
	// Value to indicate developer mode
	devTagVal string
	// Value to indicate production mode
	prodTagVal string
}

type ManagerOptions struct {
	// The variable to detect current enviroment
	EnviromentTag string
	// Value to indicate developer modes
	DevTagValue string
	// Value to indicate production mode
	ProdTagValue string
}

// Create a new app enviroment manager
// Optional ManagerOptions struct
func NewManager(appName string, opts ...ManagerOptions) (*EnvManager, error) {
	m := &EnvManager{
		prefix:     appName,
		envVar:     "APP_ENV",
		devTagVal:  "DEVELOPMENT",
		prodTagVal: "PRODUCTION",
	}

	// Apply the given options
	if len(opts) > 0 {
		applyIfNotNull(&m.envVar, opts[0].EnviromentTag)
		applyIfNotNull(&m.devTagVal, opts[0].DevTagValue)
		applyIfNotNull(&m.prodTagVal, opts[0].ProdTagValue)
	}

	// clean dirty data
	prefLen := len(m.prefix)
	if m.prefix[prefLen-1] != 95 { // check for _
		m.prefix += "_"
	}

	// detect enviroment
	env := getVar(m.prefix, m.envVar)
	if env == m.devTagVal {
		m.currEnv = m.devTagVal
	} else if env == m.prodTagVal {
		m.currEnv = m.prodTagVal
	} else {
		return nil, errors.New("Invalid value for Enviroment tag")
	}

	return m, nil
}

func (self *EnvManager) Getenv(name string) string {
	if self.IsDevelopment() {
		return getVar(self.prefix, name)
	} else {
		return os.Getenv(name)
	}
}

func (self *EnvManager) IsProduction() bool {
	return self.currEnv == self.prodTagVal
}

func (self *EnvManager) IsDevelopment() bool {
	return self.currEnv == self.devTagVal
}

func applyIfNotNull(value *string, opt string) {
	if opt != "" {
		*value = opt
	}
}

func getVar(tag, name string) string {
	key := tag + name
	return os.Getenv(key)
}
