# Ethereum Gas Analyzer

A web application to analyze and track how much you've spent on Ethereum transaction fees. It fetches your last 10,000 transactions, calculates total gas spent, and displays detailed information about outgoing transactions.

## Features

- 🔍 Analyze Ethereum addresses for gas usage
- 📊 Shows total fees in ETH and GWEI
- 🧾 Lists outgoing transactions with details (hash, recipient, gas used, gas price, timestamp)
- ⚡ Fast, modern UI (Tailwind CSS)
- 🌐 Simple REST API backend in Go

### 🚀 Live Demo

https://ethfees.up.railway.app/

---

## Getting Started

### Prerequisites

- Go 1.24+
- Etherscan API Key (free at [etherscan.io/apis](https://etherscan.io/apis))

### Clone the repository

```bash
git clone https://github.com/SalvadorBertazzo/ethFees.git
cd ethFees
```

### Environment Variables

Create a `.env` file in the root directory:

```env
ETHERSCAN_API_KEY=your_etherscan_api_key
PORT=8080 # Optional, defaults to 8080
```

### Run Locally

```bash
go run cmd/server/main.go
```

Visit [http://localhost:8080](http://localhost:8080) in your browser.

---

## Project Structure

```
ethFees/
  ├── cmd/server/main.go         # Entry point for the Go server
  ├── internal/
  │   ├── api/                  # API handlers, routes, server setup
  │   ├── models/               # Data types
  │   └── services/             # Etherscan integration, business logic
  ├── web/                      # Static frontend 
  ├── go.mod, go.sum            # Go dependencies
  └── .env                      # Environment variables
```

---

## API

### `GET /api/fees/:address`

Returns:
```json
{
  "address": "0x...",
  "out_transactions": [ ... ],
  "total_fees_eth": "0.123456789",
  "total_fees_gwei": "123456.78",
  "total_out_txs": 42
}
```

---

## License

MIT

---

**Made by Salvador Bertazzo** 