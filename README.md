# ETH Fees

Web app that calculates how much an Ethereum address has spent on gas fees. It pulls the last 10,000 transactions from Etherscan, filters outgoing ones, and shows the total cost in ETH, GWEI, and USD.

Live demo: https://ethfees.up.railway.app/

<div align="center">
  <img src="assets/demo.gif" width="350" />
</div>

## How it works

You enter an Ethereum address. The backend queries the Etherscan API to get the transaction history, filters only outgoing transactions (where the address is the sender), and calculates the total gas fees. It also fetches the current ETH/USD price to show the equivalent in dollars, both as a total and per transaction.

## Stack

- **Backend:** Go (Echo framework)
- **Frontend:** HTML + Tailwind CSS (CDN)
- **API:** Etherscan v2

## Setup

Requirements: Go 1.24+ and an [Etherscan API key](https://etherscan.io/apis) (free).

```bash
git clone https://github.com/SalvadorBertazzo/ethFees.git
cd ethFees
cp .env.example .env
```

Edit `.env` and add your API key:

```
ETHERSCAN_API_KEY=your_key_here
```

Run:

```bash
go run cmd/server/main.go
```

Open http://localhost:8080. Port is configurable via the `PORT` environment variable.

## API

### GET /api/fees/:address

Returns the fee breakdown for the given address:

```json
{
  "address": "0x...",
  "total_out_txs": 42,
  "total_fees_eth": "0.045231000000000000",
  "total_fees_gwei": "45231000.000000",
  "total_fees_usd": "125.40",
  "eth_price_usd": "2772.50",
  "out_transactions": [...]
}
```

## Project structure

```
cmd/server/main.go          Entry point
internal/api/               Routes, handlers, server config
internal/models/            Data types
internal/services/          Etherscan client, fee calculation, validation
web/                        Static frontend (HTML, favicon)
```

## License

MIT

---

Made by [Salvador Bertazzo](https://github.com/SalvadorBertazzo)
