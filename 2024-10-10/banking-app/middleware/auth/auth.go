package auth

import (
	"bankingApp/constants"
	"bankingApp/helper"
	user "bankingApp/user/service"
	"context"
	"errors"
	"fmt"
	"net/http"
)

func AuthenticationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Authentication Middleware Called")
		token, err := getAuthTokenFromHeader(r)
		if err != nil {
			helper.SendInvalidAuthError(w)
			return
		}
		claims, err1 := helper.ValidateJwtToken(token)
		fmt.Println("Validation Completed")
		if err1 != nil {
			fmt.Println(err1)
			helper.SendErrorWithCustomMessage(w, err1.Error())
			return
		}
		fmt.Println(claims)

		ctx := context.WithValue(r.Context(), "claims", claims)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func ValidateAdminPermissionsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Admin Validation Middleware Called")
		fmt.Print(r.Context())
		claims := r.Context().Value("claims").(*helper.Claims)
		fmt.Print(claims.UserId)
		admin, err := user.GetAdminInterfaceWithPassById(claims.UserId)
		if err != nil {
			helper.SendErrorWithCustomMessage(w, err.Error())
			return
		}
		ctx := context.WithValue(r.Context(), constants.ClaimsAdminKey, admin)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
func ValidateCustomerPermissionMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Customer Validation Middleware Called")
		fmt.Print(r.Context())
		claims := r.Context().Value("claims").(*helper.Claims)
		fmt.Print(claims.UserId)
		customer, err := user.GetStaffInterfaceWithPassById(claims.UserId)
		if err != nil {
			helper.SendErrorWithCustomMessage(w, err.Error())
			return
		}
		ctx := context.WithValue(r.Context(), constants.ClaimsCustomerKey, customer)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func getAuthTokenFromHeader(r *http.Request) (string, error) {
	headers := r.Header
	fmt.Println(headers)
	tempTokenHeader, ok := headers["Authorization"]
	if !ok || len(tempTokenHeader) == 0 {
		return "", errors.New("Token Not found")
	}
	token := tempTokenHeader[0]

	return token, nil
}
