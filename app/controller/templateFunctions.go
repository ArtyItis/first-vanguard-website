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

func ContainsWeapon(w map[string]interface{}, weapons []model.Weapon) bool {
	weapon := model.Map2Weapon(w)
	for _, w := range weapons {
		if w.Name == weapon.Name {
			return true
		}
	}
	return false
}

func ContainsRole(r map[string]interface{}, roles []model.Role) bool {
	role := model.Map2Role(r)
	for _, w := range roles {
		if w.Name == role.Name {
			return true
		}
	}
	return false
}

func ConvertPermissionLevelMap(u map[string]interface{}) string {
	user := model.Map2User(u)
	return ConvertPermissionLevelUser(user)
}

func ConvertPermissionLevelUser(user model.User) string {
	switch user.Permission_level {
	case 0:
		return "Siedler"
	case 200:
		return "Offizier"
	case 400:
		return "Konsul"
	case 600:
		return "Gouverneur"
	case 1000:
		return "Website-Admin"
	}
	return "nicht definiert"
}
