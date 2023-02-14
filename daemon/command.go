package daemon

type Command struct {
	Name  string `json:"name"`
	Usage string `json:"usage"`
	Flags []Flag `json:"flags"`
	// Action func(args ...string) (response string, err error)
	//
	IsHidden bool
}

type Flag struct {
	Name     string   `json:"name"`
	Type     string   `json:"type"`
	Aliases  []string `json:"aliases"`
	Usage    string   `json:"usage"`
	EnvVars  []string `json:"env_vars"`
	Required bool     `json:"required"`
	Value    any      `json:"value"`
}

type Action func(args ...string) (response string, err error)
