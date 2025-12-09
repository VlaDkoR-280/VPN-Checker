package telegram

type Answer struct {
	Status      bool   `json:"ok"`
	ErrorCode   int    `json:"error_code"`
	Description string `json:"description"`
}

//{"ok":false,"error_code":429,"description":"Too Many Requests: retry after 77241","parameters":{"retry_after":77241}}
