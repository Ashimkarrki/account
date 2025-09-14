package rest

import (
	pb_acc "account/account_proto"
	"context"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)
type AccountHandelers struct{
	Client pb_acc.AccountServiceClient
}
type updateAcc struct{
	UserName string
	Amt int64
}

func (ah AccountHandelers) AdddAccount(c echo.Context) error{
	var req pb_acc.AccountRequest
	if err:=c.Bind(&req);err!=nil{
		return c.JSON(http.StatusBadRequest,map[string]string{"error":"invalid request"})
	}
	if req.Password=="" || req.UserName=="" {
				return c.JSON(http.StatusBadRequest,map[string]string{"error":"invalid request"})

	}


	resp,err:=ah.Client.CreateAccount(context.Background(),&req)
if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
		return c.JSON(http.StatusOK, resp)
}

func (ah AccountHandelers) GetAccount(c echo.Context) error{
	id:=c.Param("id")
		idX, err := strconv.Atoi(id)
		if err!=nil{
		return c.JSON(http.StatusBadRequest,map[string]string{"error":"invalid request"})
		}
	req:=&pb_acc.ReadAccountRequest{AccId: int64(idX)}
	resp,err:=ah.Client.ReadAccount(context.Background(),req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return  c.JSON(http.StatusOK,resp)

}

func (ah AccountHandelers) UpdateAccount(c echo.Context)error{
	var r updateAcc
	id:=c.Param("id")
		idX, err := strconv.Atoi(id)
		
		if err!=nil{
		return c.JSON(http.StatusBadRequest,map[string]string{"error":"invalid request"})
		}
		if err:=c.Bind(&r); err!=nil{
			return c.JSON(http.StatusBadRequest,map[string]string{"error":"invalid request"})
		}
		req:=&pb_acc.UpdateAccountRequest{AccId: int64(idX),Amt: r.Amt,UserName: r.UserName}
		
		resp,err:=ah.Client.UpdateAccount(context.Background(),req)
		if err != nil {
		return c.JSON(http.StatusBadRequest,map[string]string{"error":"invalid request"})
		}
		return c.JSON(http.StatusOK,resp)
}

func (ah AccountHandelers) DeleteAccount(c echo.Context)error{
		id:=c.Param("id")
		idX, err := strconv.Atoi(id)
		
		if err!=nil{
		return c.JSON(http.StatusBadRequest,map[string]string{"error":"invalid request"})
		}
		req:=&pb_acc.DeleteAccountRequest{AccId: int64(idX)}
		res,err:=ah.Client.DeleteAccount(context.Background(),req)
			if err != nil {
		return c.JSON(http.StatusBadRequest,map[string]string{"error":"invalid request"})
		}
				return c.JSON(http.StatusOK,res)


	
}