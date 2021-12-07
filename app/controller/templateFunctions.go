package controller

import (
	"forgottennw/app/model"
)

func Add(x, y int) int {
	return x + y
}

func Minus(x, y int) int {
	return x - y
}

func ConvertPermissionLevelMap(u map[string]interface{}) string {
	user := model.Map2User(u)
	return user.ConvertPermissionLevel()
}

func GetRoleName(roleID string) string {
	role, _ := model.GetRoleById(roleID)
	return role.Name
}

func GetWeaponName(weaponID string) string {
	weapon, _ := model.GetWeaponById(weaponID)
	return weapon.Name
}
