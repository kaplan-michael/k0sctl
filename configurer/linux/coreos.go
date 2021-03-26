package linux

import (
	"github.com/k0sproject/rig"
	"github.com/k0sproject/rig/os"
	"github.com/k0sproject/rig/os/registry"
	"strings"
)


// coreos provides OS support for Fedora & RHEL core os systems
type Coreos struct {
	os.Linux
	BaseLinux

}

func init() {
	registry.RegisterOSModule(
		func(os rig.OSVersion) bool {
			return os.ID == "fedora" && strings.Contains(os.Name, "CoreOS") || os.ID == "rhel" && strings.Contains(os.Name, "CoreOS")
		},
		func() interface{} {
			return &Coreos{}
		},
	)
}

func (l Coreos) InstallPackage(h os.Host, pkg ...string) error {
	return h.Execf("sudo rpm-ostree install %s --reboot", strings.Join(pkg, " "))
}
// InstallKubectl installs kubectl using the rpm-ostree. not reliable now as we need to figure out reconnection to the hosts.  
//func (l Coreos) InstallKubectl(h os.Host) error {
//	return l.InstallPackage(h, "kubectl")
//}