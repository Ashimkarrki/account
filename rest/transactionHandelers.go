package rest

import (
	pb "account/transaction_proto"
	"context"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)
type Transactionhandeler struct{
	Client pb.TransactionServiceClient
}

func (h Transactionhandeler) CreateTransaction(c echo.Context)  error{
	var req pb.TransactionRequest
	if err:=c.Bind(&req);err!=nil{
		return c.JSON(http.StatusBadRequest,map[string]string{"error":"invalid request"})
	}
	res,err:=h.Client.CreateTransaction(context.Background(),&req)
	if err!=nil{
		return c.JSON(http.StatusBadRequest,map[string]string{"error":"invalid request"})
	}
	return  c.JSON(http.StatusOK,res)
}

func (h Transactionhandeler) ReadTransaction(c echo.Context)  error{
	id:=c.Param("id")
	idX,err:=strconv.Atoi(id)
	if err != nil {
				return c.JSON(http.StatusBadRequest,map[string]string{"error":"invalid request"})
	}
	res,err:=h.Client.ReadTransaction(context.Background(),&pb.ReadTransactionRequest{TransId: int64(idX)})
	if err!=nil{
		return c.JSON(http.StatusBadRequest,map[string]string{"error":"invalid request"})
	}
	return  c.JSON(http.StatusOK,res)
}