package titan

import (
	"modules/world/bus"

	"github.com/herb-go/herbplugin"
)

type PluginOpt struct {
	Titan *Titan
	Bus   *bus.Bus
}

func (o *PluginOpt) GetLocation() *herbplugin.Location {
	l := herbplugin.NewLoaction()
	l.Path = o.Titan.GetScriptPath(o.Bus.GetScriptID())
	return l
}
func (o *PluginOpt) GetParam(name string) string {
	return o.Bus.GetParam(name)
}
func (o *PluginOpt) MustAuthorizeDomain(domain string) bool {
	return o.Bus.GetTrusted().MustAuthorizeDomain(domain)
}
func (o *PluginOpt) MustAuthorizePath(path string) bool {
	return o.Bus.GetTrusted().MustAuthorizePath(path)
}
func (o *PluginOpt) MustAuthorizePermission(permission string) bool {
	return herbplugin.MustAuthorizePermission(o.Bus.GetPermissions(), permission)
}
