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

func GetWeaponByType(weaponID string, weaponType string) bool {
	weapon, _ := model.GetWeaponById(weaponID)
	return weapon.Type == weaponType
}

//used in taxes.html
func GetUsersByCompany(users []map[string]interface{}, companyName string) (companyMembers []model.User) {
	for _, u := range users {
		user := model.Map2User(u)
		if user.Company == companyName {
			companyMembers = append(companyMembers, user)
		}
	}
	return companyMembers
}
