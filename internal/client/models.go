// Copyright (c) HashiCorp, Inc.

package client

type ApiKey struct {
	ID           int    `json:"id"`
	AllowedPath  string `json:"allowedPath"`
	ApiKey       string `json:"apiKey"`
	CreatedAt    string `json:"createdAt"`
	Creator      string `json:"creator"`
	CreatorEmail string `json:"creatorEmail"`
	ExpiredAt    string `json:"expiredAt"`
	Extra        string `json:"extra"`
	Name         string `json:"name"`
	Type         string `json:"type"`
	UpdatedAt    string `json:"updatedAt"`
	Updater      string `json:"updater"`
	UpdaterEmail string `json:"updaterEmail"`
}

type ApiKeyCreate struct {
	AllowedPath string `json:"allowedPath"`
	ExpiredAt   string `json:"expiredAt"`
	Name        string `json:"name"`
	Type        string `json:"type"`
}

type JsonBody struct {
	Code    int      `json:"code"`
	Success bool     `json:"success"`
	Message string   `json:"message"`
	Causes  []string `json:"causes"`
	Data    string   `json:"data"`
}
