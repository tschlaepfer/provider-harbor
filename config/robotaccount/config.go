package robotaccount

import "github.com/crossplane/upjet/pkg/config"

func Configure(p *config.Provider) {
	p.AddResourceConfigurator("harbor_robot_account", func(r *config.Resource) {
		r.ShortGroup = "robotaccount"
		r.Kind = "RobotAccount"
	})
}
