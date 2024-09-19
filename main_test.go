package myzap_test

import (
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/verbeux-ai/myzap"
)

var client *myzap.Client

func TestMain(m *testing.M) {
	_ = godotenv.Load(".env")

	client = myzap.NewClient(
		myzap.WithToken(os.Getenv("TOKEN")),
		myzap.WithSessionKey(os.Getenv("SESSION_KEY")),
		myzap.WithBaseUrl(os.Getenv("BASE_URL")),
	)

	os.Exit(m.Run())
}
