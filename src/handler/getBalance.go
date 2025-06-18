package handler

import (
	"fetch-saldo/src/helper"
	"fetch-saldo/src/models"
	"sync"

	"github.com/gofiber/fiber/v2"
	"github.com/panjf2000/ants/v2"
	"resty.dev/v3"
)

func GetBalance(c *fiber.Ctx) error {
	apiKey := c.Get("X-API-Key")
	if apiKey == "" {
		return c.Status(401).JSON(fiber.Map{"error": "Missing X-API-Key"})
	}

	if !helper.GetAPIKeyCache(apiKey) {
		exists := models.ApiExist(apiKey)
		helper.SetAPIKeyCache(apiKey, exists)

		if !exists {
			return c.Status(403).JSON(fiber.Map{
				"error": "Invalid or unauthorized API key",
			})
		}
	}

	type request struct {
		Wallets []string `json:"wallets"`
	}
	var req request
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "invalid request body",
		})
	}

	if len(req.Wallets) == 0 {
		return c.Status(400).JSON(fiber.Map{
			"error": "Missing required fields: wallets: []",
		})
	}

	type walletBalance struct {
		Wallet  string `json:"wallet"`
		Balance int    `json:"balance"`
		Error   string `json:"error,omitempty"`
	}
	results := make([]walletBalance, len(req.Wallets))
	var wg sync.WaitGroup
	var mutex sync.Mutex

	pool, _ := ants.NewPool(5)
	defer pool.Release()

	for i, wallet := range req.Wallets {
		wg.Add(1)
		i, wallet := i, wallet

		_ = pool.Submit(func() {
			defer wg.Done()

			if cachedBalance, ok := helper.GetCacheWallet(wallet); ok {
				mutex.Lock()
				results[i] = walletBalance{
					Wallet:  wallet,
					Balance: cachedBalance,
				}
				mutex.Unlock()
				return
			}

			type rpcRequest struct {
				Jsonrpc string `json:"jsonrpc"`
				ID      int    `json:"id"`
				Method  string `json:"method"`
				Params  []any  `json:"params"`
			}
			payload := rpcRequest{
				Jsonrpc: "2.0",
				ID:      1,
				Method:  "getBalance",
				Params:  []any{wallet},
			}

			type responseGetBalance struct {
				Jsonrpc string `json:"jsonrpc"`
				ID      int    `json:"id"`
				Result  struct {
					Context struct {
						Slot int `json:"slot"`
					} `json:"context"`
					Value int `json:"value"`
				} `json:"result"`
				Error any `json:"error,omitempty"`
			}
			var res responseGetBalance

			client := resty.New()

			_, err := client.R().
				SetBody(payload).
				SetResult(&res).
				Post(helper.RPC_URI + apiKey)

			mutex.Lock()
			defer mutex.Unlock()

			if err != nil {
				results[i] = walletBalance{
					Wallet: wallet,
					Error:  err.Error(),
				}
				return
			}

			balance := res.Result.Value
			helper.SetCacheWallet(wallet, balance)

			results[i] = walletBalance{
				Wallet:  wallet,
				Balance: balance,
			}
		})
	}

	wg.Wait()
	return c.JSON(results)
}
