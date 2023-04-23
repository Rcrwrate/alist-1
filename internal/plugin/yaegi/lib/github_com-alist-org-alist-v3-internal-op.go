// Code generated by 'yaegi extract github.com/alist-org/alist/v3/internal/op'. DO NOT EDIT.

package lib

import (
	"github.com/alist-org/alist/v3/internal/op"
	"go/constant"
	"go/token"
	"reflect"
)

func init() {
	Symbols["github.com/alist-org/alist/v3/internal/op/op"] = map[string]reflect.Value{
		// function, constant and variable definitions
		"Cancel2FAById":                reflect.ValueOf(op.Cancel2FAById),
		"Cancel2FAByUser":              reflect.ValueOf(op.Cancel2FAByUser),
		"ClearCache":                   reflect.ValueOf(op.ClearCache),
		"Copy":                         reflect.ValueOf(op.Copy),
		"CreateMeta":                   reflect.ValueOf(op.CreateMeta),
		"CreateStorage":                reflect.ValueOf(op.CreateStorage),
		"CreateUser":                   reflect.ValueOf(op.CreateUser),
		"DISABLED":                     reflect.ValueOf(constant.MakeFromLiteral("\"disabled\"", token.STRING, 0)),
		"DeleteMetaById":               reflect.ValueOf(op.DeleteMetaById),
		"DeleteSettingItemByKey":       reflect.ValueOf(op.DeleteSettingItemByKey),
		"DeleteStorageById":            reflect.ValueOf(op.DeleteStorageById),
		"DeleteUserById":               reflect.ValueOf(op.DeleteUserById),
		"DisableStorage":               reflect.ValueOf(op.DisableStorage),
		"DropStorage":                  reflect.ValueOf(op.DropStorage),
		"EnableStorage":                reflect.ValueOf(op.EnableStorage),
		"Get":                          reflect.ValueOf(op.Get),
		"GetAdmin":                     reflect.ValueOf(op.GetAdmin),
		"GetAllStorages":               reflect.ValueOf(op.GetAllStorages),
		"GetBalancedStorage":           reflect.ValueOf(op.GetBalancedStorage),
		"GetDriverInfoMap":             reflect.ValueOf(op.GetDriverInfoMap),
		"GetDriverNames":               reflect.ValueOf(op.GetDriverNames),
		"GetDriverNew":                 reflect.ValueOf(op.GetDriverNew),
		"GetDriverNewMap":              reflect.ValueOf(op.GetDriverNewMap),
		"GetGuest":                     reflect.ValueOf(op.GetGuest),
		"GetMainItems":                 reflect.ValueOf(op.GetMainItems),
		"GetMetaById":                  reflect.ValueOf(op.GetMetaById),
		"GetMetaByPath":                reflect.ValueOf(op.GetMetaByPath),
		"GetMetas":                     reflect.ValueOf(op.GetMetas),
		"GetNearestMeta":               reflect.ValueOf(op.GetNearestMeta),
		"GetPublicSettingItems":        reflect.ValueOf(op.GetPublicSettingItems),
		"GetPublicSettingsMap":         reflect.ValueOf(op.GetPublicSettingsMap),
		"GetSettingItemByKey":          reflect.ValueOf(op.GetSettingItemByKey),
		"GetSettingItemInKeys":         reflect.ValueOf(op.GetSettingItemInKeys),
		"GetSettingItems":              reflect.ValueOf(op.GetSettingItems),
		"GetSettingItemsByGroup":       reflect.ValueOf(op.GetSettingItemsByGroup),
		"GetSettingItemsInGroups":      reflect.ValueOf(op.GetSettingItemsInGroups),
		"GetSettingsMap":               reflect.ValueOf(op.GetSettingsMap),
		"GetStorageAndActualPath":      reflect.ValueOf(op.GetStorageAndActualPath),
		"GetStorageByMountPath":        reflect.ValueOf(op.GetStorageByMountPath),
		"GetStorageVirtualFilesByPath": reflect.ValueOf(op.GetStorageVirtualFilesByPath),
		"GetUnwrap":                    reflect.ValueOf(op.GetUnwrap),
		"GetUserById":                  reflect.ValueOf(op.GetUserById),
		"GetUserByName":                reflect.ValueOf(op.GetUserByName),
		"GetUserByRole":                reflect.ValueOf(op.GetUserByRole),
		"GetUsers":                     reflect.ValueOf(op.GetUsers),
		"HandleObjsUpdateHook":         reflect.ValueOf(op.HandleObjsUpdateHook),
		"HandleSettingItemHook":        reflect.ValueOf(op.HandleSettingItemHook),
		"HasStorage":                   reflect.ValueOf(op.HasStorage),
		"Key":                          reflect.ValueOf(op.Key),
		"Link":                         reflect.ValueOf(op.Link),
		"List":                         reflect.ValueOf(op.List),
		"LoadStorage":                  reflect.ValueOf(op.LoadStorage),
		"MakeDir":                      reflect.ValueOf(op.MakeDir),
		"Move":                         reflect.ValueOf(op.Move),
		"MustSaveDriverStorage":        reflect.ValueOf(op.MustSaveDriverStorage),
		"Other":                        reflect.ValueOf(op.Other),
		"Put":                          reflect.ValueOf(op.Put),
		"RegisterDriver":               reflect.ValueOf(op.RegisterDriver),
		"RegisterObjsUpdateHook":       reflect.ValueOf(op.RegisterObjsUpdateHook),
		"RegisterSettingItemHook":      reflect.ValueOf(op.RegisterSettingItemHook),
		"RegisterStorageHook":          reflect.ValueOf(op.RegisterStorageHook),
		"Remove":                       reflect.ValueOf(op.Remove),
		"Rename":                       reflect.ValueOf(op.Rename),
		"RootName":                     reflect.ValueOf(constant.MakeFromLiteral("\"root\"", token.STRING, 0)),
		"SaveSettingItem":              reflect.ValueOf(op.SaveSettingItem),
		"SaveSettingItems":             reflect.ValueOf(op.SaveSettingItems),
		"UnRegisterDriver":             reflect.ValueOf(op.UnRegisterDriver),
		"UpdateMeta":                   reflect.ValueOf(op.UpdateMeta),
		"UpdateStorage":                reflect.ValueOf(op.UpdateStorage),
		"UpdateUser":                   reflect.ValueOf(op.UpdateUser),
		"WORK":                         reflect.ValueOf(constant.MakeFromLiteral("\"work\"", token.STRING, 0)),

		// type definitions
		"New":             reflect.ValueOf((*op.New)(nil)),
		"ObjsUpdateHook":  reflect.ValueOf((*op.ObjsUpdateHook)(nil)),
		"SettingItemHook": reflect.ValueOf((*op.SettingItemHook)(nil)),
		"StorageHook":     reflect.ValueOf((*op.StorageHook)(nil)),
	}
}
