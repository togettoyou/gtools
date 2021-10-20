package gtools

import "testing"

func TestJWK2PEMTool(t *testing.T) {
	js := `{
    "e": "AQAB",
    "n": "wVKQLBUqOBiay2dkn9TlbfuaF40_edIKUmdLq6OlvzEMrP4IDzdOk50TMO0nfjJ6v5830_5x0vRg5bzZQeKpHniR0sw7qyoSI6n2eSkSnFt7P-N8gv2KWnwzVs_h9FDdeLOeVOU8k_qzkph3_tmBV7ZZG-4_DEvgvat6ifEC-WzzYqofsIrTiTT7ZFxTqid1q6zrrsmyU2DQH3WdgFiOJVVlN2D0BuZu5X7pGZup_RcWzt_9T6tQsGeU1juSuuUk_9_FVDXNNCTObfKCTKXqjW95ZgAI_xVrMeQC5nXlMh6VEaXfO83oy1j36wUoVUrUnkANhp-dnjTdvJgwN82dGQ",
    "kty": "RSA",
    "alg": "RS256",
    "use": "sig",
    "kid": "1ziyHTmy3_804C9McPCGTDfabB58A4CedoVkupywUyM"
}`
	pem, err := JWK2PEM(js)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(pem)
}
