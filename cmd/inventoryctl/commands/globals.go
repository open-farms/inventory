package commands

type Globals struct {
	Config string `help:"Path to config file" short:"c" default:"./config/config.yaml" type:"path"`
}
