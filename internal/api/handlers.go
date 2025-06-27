package api

import (
	"ethFees/internal/services"
	"net/http"

	"github.com/labstack/echo/v4"
)

func handleGetFees(c echo.Context) error {
	address := c.Param("address")

	if !services.IsValidEthAddress(address) {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid Ethereum address",
		})
	}

	resp, err := services.GetTransactions(address)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	outTxs := services.GetOutTransactions(resp.Result, address)
	totalFeeWei := services.CalculateTotalFeeWei(outTxs)
	totalFeeEth := services.WeiToEthString(totalFeeWei)
	totalFeeGwei := services.WeiToGweiString(totalFeeWei)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"address":          address,
		"out_transactions": outTxs,
		"total_fees_eth":   totalFeeEth,
		"total_fees_gwei":  totalFeeGwei,
		"total_out_txs":    len(outTxs),
	})
}
