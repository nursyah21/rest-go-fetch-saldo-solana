package main

import (
	"github.com/go-resty/resty/v2"
	"github.com/gofiber/fiber/v2"
)

type BalanceRequest struct {
	Wallets []string `json:"wallets"`
}

type RpcRequest struct {
	Jsonrpc string `json:"jsonrpc"`
	ID      int    `json:"id"`
	Method  string `json:"method"`
	Params  []any  `json:"params"`
}

type RpcBalanceResult struct {
	Context struct {
		Slot int `json:"slot"`
	} `json:"context"`
	Value int `json:"value"`
}

type RpcResponse struct {
	Jsonrpc string           `json:"jsonrpc"`
	ID      int              `json:"id"`
	Result  RpcBalanceResult `json:"result"`
	Error   any              `json:"error,omitempty"`
}

func GetBalance(c *fiber.Ctx) error {
	var req BalanceRequest
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

	payload := RpcRequest{
		Jsonrpc: "2.0",
		ID:      1,
		Method:  "getBalance",
		Params:  []any{req.Wallets[0]},
	}

	var res RpcResponse
	client := resty.New()
	_, err := client.R().
		SetBody(payload).
		SetResult(&res).
		Post(RPC_HELIUS)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"wallet":  req.Wallets[0],
		"balance": res.Result.Value,
	})
}
