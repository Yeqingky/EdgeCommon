package firewallconfigs

type HTTPFirewallPost307Action struct {
	IsPrior bool `yaml:"isPrior" json:"isPrior"`

	Life  int32         `yaml:"life" json:"life"`
	Scope FirewallScope `yaml:"scope" json:"scope"`
}
