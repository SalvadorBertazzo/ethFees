package services

import (
	"encoding/json"
	"ethFees/internal/models"
	"fmt"
	"io"
	"math/big"
	"net/http"
	"os"
	"regexp"
	"strings"
)

func GetTransactions(address string) (*models.EtherscanResponse, error) {
	apiKey := os.Getenv("ETHERSCAN_API_KEY")

	url := fmt.Sprintf(
		"https://api.etherscan.io/api?module=account&action=txlist&address=%s&startblock=0&endblock=99999999&sort=desc&apikey=%s",
		address, apiKey,
	)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var etherscanResp models.EtherscanResponse
	err = json.Unmarshal(body, &etherscanResp)
	if err != nil {
		return nil, err
	}

	return &etherscanResp, nil
}

func GetOutTransactions(transactions []models.Transaction, walletAddress string) []models.Transaction {
	outTransactions := make([]models.Transaction, 0)
	walletAddress = strings.ToLower(walletAddress)

	for _, tx := range transactions {
		if strings.ToLower(tx.From) == walletAddress {
			outTransactions = append(outTransactions, tx)
		}
	}

	return outTransactions
}

func CalculateTotalFeeWei(transactions []models.Transaction) *big.Int {
	totalFee := big.NewInt(0)

	for _, tx := range transactions {
		// Convert string from JSON to big.Int
		gasUsed := new(big.Int)
		gasPrice := new(big.Int)

		gasUsed.SetString(tx.GasUsed, 10)
		gasPrice.SetString(tx.GasPrice, 10)

		// Calculate fee
		fee := new(big.Int)
		fee.Mul(gasUsed, gasPrice)

		totalFee.Add(totalFee, fee)
	}
	return totalFee
}

func WeiToEthString(wei *big.Int) string {
	fWei := new(big.Float).SetInt(wei)
	ethValue := new(big.Float).Quo(fWei, big.NewFloat(1e18))
	return ethValue.Text('f', 18) // 18 decimals
}

func WeiToGweiString(wei *big.Int) string {
	fWei := new(big.Float).SetInt(wei)
	gweiValue := new(big.Float).Quo(fWei, big.NewFloat(1e9))
	return gweiValue.Text('f', 6) // 6 decimales
}

func IsValidEthAddress(address string) bool {
	if len(address) != 42 {
		return false
	}
	if !strings.HasPrefix(address, "0x") && !strings.HasPrefix(address, "0X") {
		return false
	}
	// Regex: after 0x, 40 hexadecimal characters
	match, _ := regexp.MatchString("^0x[0-9a-fA-F]{40}$", address)
	return match
}
