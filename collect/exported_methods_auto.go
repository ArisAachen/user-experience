// Code generated by "dbusutil-gen em -type DBusModule dbus.go"; DO NOT EDIT.

package collect

import (
	"pkg.deepin.io/lib/dbusutil"
)

func (v *DBusModule) GetExportedMethods() dbusutil.ExportedMethods {
	return dbusutil.ExportedMethods{
		{
			Name:   "Enable",
			Fn:     v.Enable,
			InArgs: []string{"enabled"},
		},
		{
			Name:    "IsEnabled",
			Fn:      v.IsEnabled,
			OutArgs: []string{"outArg0"},
		},
		{
			Name:   "SendAppInstallData",
			Fn:     v.SendAppInstallData,
			InArgs: []string{"msg", "path", "name", "id"},
		},
		{
			Name:   "SendAppStateData",
			Fn:     v.SendAppStateData,
			InArgs: []string{"msg", "path", "name", "id"},
		},
		{
			Name:   "SendLogonData",
			Fn:     v.SendLogonData,
			InArgs: []string{"msg"},
		},
	}
}
