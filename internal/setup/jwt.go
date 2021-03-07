package setup

import (
    "context"
    "github.com/coreos/go-oidc/v3/oidc"
)

var IdTokenVerifier *oidc.IDTokenVerifier

func SetupIdTokenVerifier() {
    ctx := context.Background()
    provider, err := oidc.NewProvider(ctx, "http://localhost:8080/auth/realms/hazel")
    if err != nil {
        panic(err)
    }

    IdTokenVerifier = provider.Verifier(&oidc.Config{ClientID: "account"})
}
