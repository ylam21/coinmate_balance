# coinmate_balance

`coinmate_balance` is a small CLI tool that prints the balance of your BTC portfolio for your CoinMate account without requiring you to log in manually or provide 2FA each time.

It exists to make checking your BTC portfolio quick and automated using your CoinMate API key.
---

## Prerequisites

* Go 1.24 or higher installed
* A CoinMate account
* A CoinMate API key with the following credentials:

  * `COINMATE_CLIENT_ID`
  * `COINMATE_PUBLIC_KEY`
  * `COINMATE_PRIVATE_KEY`

---

## Setup

1. Create a CoinMate API key from your account.
2. Export the credentials as environment variables in your shell:

```bash
export COINMATE_CLIENT_ID="your_client_id"
export COINMATE_PUBLIC_KEY="your_public_key"
export COINMATE_PRIVATE_KEY="your_private_key"
```

3. Verify that the environment variables are set:

```bash
echo $COINMATE_CLIENT_ID
echo $COINMATE_PUBLIC_KEY
echo $COINMATE_PRIVATE_KEY
```

---

## Build

To compile the executable binary:

```bash
cd coinmate_balance
go build -o balance ./main.go
```

---

## Usage

Run the program:

```bash
./balance
```

It will fetch your BTC balance from CoinMate and print:

* Current BTC price in USD
* BTC balance in portfolio
* Equivalent value of BTC balance in CZK (via EUR conversion)

---

## Notes

* Currently, only BTC balances are supported.
* CoinMate API supports only `BTC_EUR` and `BTC_CZK` tickers.
---

## Example Output

```text
BTC Price: 34500 USD
BTC Balance: 0.001 BTC
BTC Balance in CZK: 722 CZK
```
