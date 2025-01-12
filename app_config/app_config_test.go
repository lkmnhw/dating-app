package app_config

import (
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGet(t *testing.T) {
	type env struct {
		envKey   string
		envValue string
	}

	type arg struct {
		isProd bool
		env    []env
	}

	tests := []struct {
		name string
		inp  arg
		exp  *AppConfig
	}{
		{
			name: "test get test env first initialize",
			inp: arg{
				isProd: false,
				env: []env{
					{
						envKey:   "TESTING_DATABASE_SOURCE",
						envValue: "TESTING_DATABASE_SOURCE",
					},
					{
						envKey:   "TESTING_DATABASE_NAME",
						envValue: "TESTING_DATABASE_NAME",
					},
				},
			},
			exp: &AppConfig{
				ENV: ENV{
					TESTING_DATABASE_SOURCE: "TESTING_DATABASE_SOURCE",
					TESTING_DATABASE_NAME:   "TESTING_DATABASE_NAME",
					COLLECTION_USERS:        "test_users",
					COLLECTION_PROFILES:     "test_profiles",
					COLLECTION_MATCHES:      "test_matches",
				},
			},
		},
		{
			name: "test get test second first initialize",
			inp: arg{
				isProd: false,
				env: []env{
					{
						envKey:   "TESTING_DATABASE_SOURCE",
						envValue: "TESTING_DATABASE_SOURCE",
					},
					{
						envKey:   "TESTING_DATABASE_NAME",
						envValue: "TESTING_DATABASE_NAME",
					},
				},
			},
			exp: &AppConfig{
				ENV: ENV{
					TESTING_DATABASE_SOURCE: "TESTING_DATABASE_SOURCE",
					TESTING_DATABASE_NAME:   "TESTING_DATABASE_NAME",
					COLLECTION_USERS:        "test_users",
					COLLECTION_PROFILES:     "test_profiles",
					COLLECTION_MATCHES:      "test_matches",
				},
			},
		},
		{
			name: "test get prod env first initialize",
			inp: arg{
				isProd: true,
				env: []env{
					{
						envKey:   "PORT",
						envValue: "PORT",
					},
					{
						envKey:   "DATABASE_SOURCE",
						envValue: "DATABASE_SOURCE",
					},
					{
						envKey:   "DATABASE_NAME",
						envValue: "DATABASE_NAME",
					},
					{
						envKey:   "COLLECTION_USERS",
						envValue: "COLLECTION_USERS",
					},
					{
						envKey:   "COLLECTION_PROFILES",
						envValue: "COLLECTION_PROFILES",
					},
					{
						envKey:   "COLLECTION_MATCHES",
						envValue: "COLLECTION_MATCHES",
					},
				},
			},
			exp: &AppConfig{
				ENV: ENV{
					PORT:                "PORT",
					DATABASE_SOURCE:     "DATABASE_SOURCE",
					DATABASE_NAME:       "DATABASE_NAME",
					COLLECTION_USERS:    "COLLECTION_USERS",
					COLLECTION_PROFILES: "COLLECTION_PROFILES",
					COLLECTION_MATCHES:  "COLLECTION_MATCHES",
				},
			},
		},
		{
			name: "test get prod env second initialize",
			inp: arg{
				isProd: true,
				env: []env{
					{
						envKey:   "PORT",
						envValue: "PORT",
					},
					{
						envKey:   "DATABASE_SOURCE",
						envValue: "DATABASE_SOURCE",
					},
					{
						envKey:   "DATABASE_NAME",
						envValue: "DATABASE_NAME",
					},
					{
						envKey:   "COLLECTION_USERS",
						envValue: "COLLECTION_USERS",
					},
					{
						envKey:   "COLLECTION_PROFILES",
						envValue: "COLLECTION_PROFILES",
					},
					{
						envKey:   "COLLECTION_MATCHES",
						envValue: "COLLECTION_MATCHES",
					},
				},
			},
			exp: &AppConfig{
				ENV: ENV{
					PORT:                "PORT",
					DATABASE_SOURCE:     "DATABASE_SOURCE",
					DATABASE_NAME:       "DATABASE_NAME",
					COLLECTION_USERS:    "COLLECTION_USERS",
					COLLECTION_PROFILES: "COLLECTION_PROFILES",
					COLLECTION_MATCHES:  "COLLECTION_MATCHES",
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for _, e := range tt.inp.env {
				os.Setenv(e.envKey, e.envValue)
			}
			act := Get(tt.inp.isProd)
			assert.Equal(t, tt.exp, act, "Test act and tt.exp is not equal")
		})
	}
	os.Clearenv()
}

func TestCheckENV(t *testing.T) {
	type arg struct {
		envKey   string
		envValue string
	}

	tests := []struct {
		name     string
		inp      arg
		expError bool
	}{
		{
			name: "test check env with error env PORT not set",
			inp: arg{
				envKey:   "KEY",
				envValue: "VAL",
			},
			expError: true,
		},
		{
			name: "test check env with error env DATABASE_SOURCE not set",
			inp: arg{
				envKey:   "PORT",
				envValue: "PORT",
			},
			expError: true,
		},
		{
			name: "test check env with error env COLLECTION_USERS not set",
			inp: arg{
				envKey:   "DATABASE_SOURCE",
				envValue: "DATABASE_SOURCE",
			},
			expError: true,
		},
		{
			name: "test check env with error env COLLECTION_PROFILES not set",
			inp: arg{
				envKey:   "DATABASE_NAME",
				envValue: "DATABASE_NAME",
			},
			expError: true,
		},
		{
			name: "test check env with error env COLLECTION_PROFILES not set",
			inp: arg{
				envKey:   "COLLECTION_USERS",
				envValue: "COLLECTION_USERS",
			},
			expError: true,
		},
		{
			name: "test check env with error env COLLECTION_MATCHES not set",
			inp: arg{
				envKey:   "COLLECTION_PROFILES",
				envValue: "COLLECTION_PROFILES",
			},
			expError: true,
		},
		{
			name: "test check env without error",
			inp: arg{
				envKey:   "COLLECTION_MATCHES",
				envValue: "COLLECTION_MATCHES",
			},
			expError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			os.Setenv(tt.inp.envKey, tt.inp.envValue)
			actErr := checkENV()
			if (actErr != nil) != tt.expError {
				t.Errorf("TestCheckENV() error = %v, expError %v", actErr, tt.expError)
				return
			}
		})
	}
	os.Clearenv()
}

func TestInitConfigPanic(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			log.Printf("error test init config not recoverable, %v\n", err)
		}
	}()
	initConfig()
}
