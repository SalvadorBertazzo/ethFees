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
	"strconv"
	"strings"
)

func GetTransactions(address string) (*models.EtherscanResponse, error) {
	apiKey := os.Getenv("ETHERSCAN_API_KEY")
	if apiKey == "" {
		return nil, fmt.Errorf("ETHERSCAN_API_KEY is not defined: configure the variable in .env or in the system")
	}

	url := fmt.Sprintf(
		"https://api.etherscan.io/v2/api?chainid=1&module=account&action=txlist&address=%s&startblock=0&endblock=99999999&sort=desc&apikey=%s",
		address, apiKey,
	)

	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error connecting to Etherscan: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var rawResp models.EtherscanRawResponse
	err = json.Unmarshal(body, &rawResp)
	if err != nil {
		return nil, err
	}

	if rawResp.Status != "1" {
		return nil, fmt.Errorf("etherscan API error: %s", string(rawResp.Result))
	}

	var transactions []models.Transaction
	err = json.Unmarshal(rawResp.Result, &transactions)
	if err != nil {
		return nil, err
	}

	return &models.EtherscanResponse{
		Status:  rawResp.Status,
		Message: rawResp.Message,
		Result:  transactions,
	}, nil
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

func GetEthPriceUSD() (float64, error) {
	apiKey := os.Getenv("ETHERSCAN_API_KEY")
	if apiKey == "" {
		return 0, fmt.Errorf("ETHERSCAN_API_KEY is not defined: configure the variable in .env or in the system")
	}

	url := fmt.Sprintf(
		"https://api.etherscan.io/v2/api?chainid=1&module=stats&action=ethprice&apikey=%s",
		apiKey,
	)

	resp, err := http.Get(url)
	if err != nil {
		return 0, fmt.Errorf("error connecting to Etherscan: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}

	var rawResp models.EtherscanRawResponse
	err = json.Unmarshal(body, &rawResp)
	if err != nil {
		return 0, err
	}

	if rawResp.Status != "1" {
		return 0, fmt.Errorf("etherscan API error: %s", string(rawResp.Result))
	}

	var priceResult models.EthPriceResult
	err = json.Unmarshal(rawResp.Result, &priceResult)
	if err != nil {
		return 0, err
	}

	price, err := strconv.ParseFloat(priceResult.EthUSD, 64)
	if err != nil {
		return 0, fmt.Errorf("error parsing ETH price: %w", err)
	}

	return price, nil
}
