package auth

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)
func JwtInterceptor(
    ctx context.Context,
    req any,
    info *grpc.UnaryServerInfo,
    handler grpc.UnaryHandler,
) (any, error) {

	md,_:=metadata.FromIncomingContext(ctx)
	authHeader := md.Get("authorization")
	if info.FullMethod=="/account_proto.AccountService/CreateAccount" || info.FullMethod=="/account_proto.AccountService/Login" {
			    return handler(ctx, req)


	}
	if len(authHeader)==0{
		return  nil, status.Error(codes.Unauthenticated,"missing token")
	}
	   tokenString := authHeader[0]
    if len(tokenString) > 7 && tokenString[:7] == "Bearer " {
        tokenString = tokenString[7:]
    }
	    token, err := ValidateJWT(tokenString)
		   if err != nil || !token.Valid {
        return nil, status.Error(codes.Unauthenticated, "invalid token")
    }
	    return handler(ctx, req)

}

func GenerateJWT(accID int64)(string,error){
	 jwtkey := []byte(os.Getenv("SECRET"))


	claims:=jwt.MapClaims{
		"Acc_id":accID,
		 "exp":      time.Now().Add(time.Hour * 1).Unix(), 
        "iat":      time.Now().Unix(),
	}
	token:=jwt.NewWithClaims(jwt.SigningMethodHS256,claims)
	signed, err := token.SignedString(jwtkey)
	if err != nil {
		return "", fmt.Errorf("failed to sign JWT: %w", err)
	}
	return signed,nil
}
func ValidateJWT(tokenString string) (*jwt.Token, error) {
	 jwtkey := []byte(os.Getenv("SECRET"))

    return jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
        if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
            return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
        }
        return jwtkey, nil
    })
}
