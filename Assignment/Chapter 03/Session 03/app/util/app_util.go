package util

import (
	"fmt"
	"ojam-test/c3/s2/app/constant"
	"strings"
)

func CheckSuitabilityProductTypeAndUserRole(productTypeName, userRoleName string) bool {
	if userRoleName == string(constant.UserRole.HEAD) {
		return true
	}

	if productTypeName == string(constant.Product.CAR) && userRoleName == string(constant.UserRole.ADMIN_CAR) {
		return true
	} else if productTypeName == string(constant.Product.MOTOR) && userRoleName == string(constant.UserRole.ADMIN_MOTOR) {
		return true
	} else if productTypeName == string(constant.Product.CYCLE) && userRoleName == string(constant.UserRole.ADMIN_CYCLE) {
		return true
	} else {
		return false
	}
}

func CheckSuitabilityProductTypeIdAndUserRoleId(productTypeId, userRoleId int) bool {
	if userRoleId == int(constant.UserRoleId.HEAD) {
		return true
	}

	if productTypeId == int(constant.ProductId.CAR) && userRoleId == int(constant.UserRoleId.ADMIN_CAR) {
		return true
	} else if productTypeId == int(constant.ProductId.MOTOR) && userRoleId == int(constant.UserRoleId.ADMIN_MOTOR) {
		return true
	} else if productTypeId == int(constant.ProductId.CYCLE) && userRoleId == int(constant.UserRoleId.ADMIN_CYCLE) {
		return true
	} else {
		return false
	}
}

func CheckSuitabilityProduct(productName, productTypeName string) bool {
	return strings.Contains(
		strings.ToLower(productName),
		strings.ToLower(productTypeName),
	)
}

func GetUserRoleId(userRoleName string) (int, error) {
	switch userRoleName {

	case string(constant.UserRole.HEAD):
		return 1, nil
	case string(constant.UserRole.ADMIN_CAR):
		return 2, nil
	case string(constant.UserRole.ADMIN_MOTOR):
		return 3, nil
	case string(constant.UserRole.ADMIN_CYCLE):
		return 4, nil
	default:
		return 0, fmt.Errorf(constant.USER_ROLE_NAME_NOT_FOUND)
	}
}

func GetProductTypeId(productTypeName string) (int, error) {
	switch productTypeName {

	case string(constant.Product.CAR):
		return 1, nil
	case string(constant.Product.MOTOR):
		return 2, nil
	case string(constant.Product.CYCLE):
		return 3, nil
	default:
		return 0, fmt.Errorf(constant.PRODUCT_TYPE_NAME_NOT_FOUND)
	}
}

func CheckHeadRole(userRoleId int) bool {
	return userRoleId == int(constant.UserRoleId.HEAD)
}
