// +build e2e

package e2e

import (
	"context"
	"encoding/base64"
	"fmt"
	"net/http"
	"os"
	"testing"

	"github.com/consensys/quorum-key-manager/pkg/client"
	"github.com/consensys/quorum-key-manager/pkg/common"
	"github.com/consensys/quorum-key-manager/src/stores/api/types"
	"github.com/consensys/quorum-key-manager/tests"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

var ecdsaPrivKey, _ = base64.StdEncoding.DecodeString("2zN8oyleQFBYZ5PyUuZB87OoNzkBj6TM4BqBypIOfhw=")
var eddsaPrivKey, _ = base64.StdEncoding.DecodeString("X9Yz/5+O42+eOodHCUBhA4VMD2ZQy5CMAQ6lXqvDUZGGbioek5qYuzJzTNZpTHrVjjFk7iFe3FYwfpxZyNPxtIaFB5gb9VP9IcHZewwNZly821re7RkmB8pGdjywygPH")

type keysTestSuite struct {
	suite.Suite
	err              error
	ctx              context.Context
	keyManagerClient *client.HTTPClient
	cfg              *tests.Config
}

func (s *keysTestSuite) SetupSuite() {
	if s.err != nil {
		s.T().Error(s.err)
	}

	s.keyManagerClient = client.NewHTTPClient(&http.Client{}, &client.Config{
		URL: s.cfg.KeyManagerURL,
	})
}

func (s *keysTestSuite) TearDownSuite() {
	if s.err != nil {
		s.T().Error(s.err)
	}
}

func TestKeyManagerKeys(t *testing.T) {
	s := new(keysTestSuite)

	s.ctx = context.Background()
	sig := common.NewSignalListener(func(signal os.Signal) {
		s.err = fmt.Errorf("interrupt signal was caught")
		t.FailNow()
	})
	defer sig.Close()

	s.cfg, s.err = tests.NewConfig()
	suite.Run(t, s)
}

func (s *keysTestSuite) TestCreate() {
	s.Run("should create a new key successfully: Secp256k1/ECDSA", func() {
		keyID := "my-key-ecdsa"
		request := &types.CreateKeyRequest{
			Curve:            "secp256k1",
			SigningAlgorithm: "ecdsa",
			Tags: map[string]string{
				"myTag0": "tag0",
				"myTag1": "tag1",
			},
		}

		key, err := s.keyManagerClient.CreateKey(s.ctx, s.cfg.HashicorpKeyStore, keyID, request)
		require.NoError(s.T(), err)

		assert.NotEmpty(s.T(), key.PublicKey)
		assert.Equal(s.T(), request.SigningAlgorithm, key.SigningAlgorithm)
		assert.Equal(s.T(), request.Curve, key.Curve)
		assert.Equal(s.T(), keyID, key.ID)
		assert.Equal(s.T(), request.Tags, key.Tags)
		assert.False(s.T(), key.Disabled)
		assert.NotEmpty(s.T(), key.CreatedAt)
		assert.NotEmpty(s.T(), key.UpdatedAt)
		assert.True(s.T(), key.ExpireAt.IsZero())
		assert.True(s.T(), key.DeletedAt.IsZero())
		assert.True(s.T(), key.DestroyedAt.IsZero())
	})

	s.Run("should create a new key successfully: BN254/EDDSA", func() {
		keyID := "my-key-eddsa"
		request := &types.CreateKeyRequest{
			Curve:            "bn254",
			SigningAlgorithm: "eddsa",
			Tags: map[string]string{
				"myTag0": "tag0",
				"myTag1": "tag1",
			},
		}

		key, err := s.keyManagerClient.CreateKey(s.ctx, s.cfg.HashicorpKeyStore, keyID, request)
		require.NoError(s.T(), err)

		assert.NotEmpty(s.T(), key.PublicKey)
		assert.Equal(s.T(), request.SigningAlgorithm, key.SigningAlgorithm)
		assert.Equal(s.T(), request.Curve, key.Curve)
		assert.Equal(s.T(), keyID, key.ID)
		assert.Equal(s.T(), request.Tags, key.Tags)
		assert.False(s.T(), key.Disabled)
		assert.NotEmpty(s.T(), key.CreatedAt)
		assert.NotEmpty(s.T(), key.UpdatedAt)
		assert.True(s.T(), key.ExpireAt.IsZero())
		assert.True(s.T(), key.DeletedAt.IsZero())
		assert.True(s.T(), key.DestroyedAt.IsZero())
	})

	s.Run("should parse errors successfully", func() {
		keyID := "my-key"
		request := &types.CreateKeyRequest{
			Curve:            "bn254",
			SigningAlgorithm: "eddsa",
			Tags: map[string]string{
				"myTag0": "tag0",
				"myTag1": "tag1",
			},
		}

		key, err := s.keyManagerClient.CreateKey(s.ctx, "inexistentStoreName", keyID, request)
		require.Nil(s.T(), key)

		httpError := err.(*client.ResponseError)
		assert.Equal(s.T(), 404, httpError.StatusCode)
	})

	s.Run("should fail with bad request if curve is not supported", func() {
		keyID := "my-key"
		request := &types.CreateKeyRequest{
			Curve:            "invalidCurve",
			SigningAlgorithm: "eddsa",
			Tags: map[string]string{
				"myTag0": "tag0",
				"myTag1": "tag1",
			},
		}

		key, err := s.keyManagerClient.CreateKey(s.ctx, s.cfg.HashicorpKeyStore, keyID, request)
		require.Nil(s.T(), key)

		httpError := err.(*client.ResponseError)
		assert.Equal(s.T(), 400, httpError.StatusCode)
	})

	s.Run("should fail with bad request if signing algorithm is not supported", func() {
		keyID := "my-key"
		request := &types.CreateKeyRequest{
			Curve:            "secp256k1",
			SigningAlgorithm: "invalidSigningAlgorithm",
			Tags: map[string]string{
				"myTag0": "tag0",
				"myTag1": "tag1",
			},
		}

		key, err := s.keyManagerClient.CreateKey(s.ctx, s.cfg.HashicorpKeyStore, keyID, request)
		require.Nil(s.T(), key)

		httpError := err.(*client.ResponseError)
		assert.Equal(s.T(), 400, httpError.StatusCode)
	})
}

func (s *keysTestSuite) TestImport() {
	s.Run("should create a new key successfully: Secp256k1/ECDSA", func() {
		keyID := "my-key-import-ecdsa"
		request := &types.ImportKeyRequest{
			Curve:            "secp256k1",
			PrivateKey:       ecdsaPrivKey,
			SigningAlgorithm: "ecdsa",
			Tags: map[string]string{
				"myTag0": "tag0",
				"myTag1": "tag1",
			},
		}

		key, err := s.keyManagerClient.ImportKey(s.ctx, s.cfg.HashicorpKeyStore, keyID, request)
		require.NoError(s.T(), err)

		assert.Equal(s.T(), "BFVSFJhqUh9DQJwcayNtsWdDMvqq8R/EKnBHqwd4Hr5vCXTyJlqKfYIgj4jCGixVZjsz5a+S2RklJRFjjoLf+LI=", key.PublicKey)
		assert.Equal(s.T(), request.SigningAlgorithm, key.SigningAlgorithm)
		assert.Equal(s.T(), request.Curve, key.Curve)
		assert.Equal(s.T(), keyID, key.ID)
		assert.Equal(s.T(), request.Tags, key.Tags)
		assert.False(s.T(), key.Disabled)
		assert.NotEmpty(s.T(), key.CreatedAt)
		assert.NotEmpty(s.T(), key.UpdatedAt)
		assert.True(s.T(), key.ExpireAt.IsZero())
		assert.True(s.T(), key.DeletedAt.IsZero())
		assert.True(s.T(), key.DestroyedAt.IsZero())
	})

	s.Run("should create a new key successfully: BN254/EDDSA", func() {
		keyID := "my-key-import-eddsa"
		request := &types.ImportKeyRequest{
			Curve:            "bn254",
			SigningAlgorithm: "eddsa",
			PrivateKey:       eddsaPrivKey,
			Tags: map[string]string{
				"myTag0": "tag0",
				"myTag1": "tag1",
			},
		}

		key, err := s.keyManagerClient.ImportKey(s.ctx, s.cfg.HashicorpKeyStore, keyID, request)
		require.NoError(s.T(), err)

		assert.Equal(s.T(), "X9Yz/5+O42+eOodHCUBhA4VMD2ZQy5CMAQ6lXqvDUZE=", key.PublicKey)
		assert.Equal(s.T(), request.SigningAlgorithm, key.SigningAlgorithm)
		assert.Equal(s.T(), request.Curve, key.Curve)
		assert.Equal(s.T(), keyID, key.ID)
		assert.Equal(s.T(), request.Tags, key.Tags)
		assert.False(s.T(), key.Disabled)
		assert.NotEmpty(s.T(), key.CreatedAt)
		assert.NotEmpty(s.T(), key.UpdatedAt)
		assert.True(s.T(), key.ExpireAt.IsZero())
		assert.True(s.T(), key.DeletedAt.IsZero())
		assert.True(s.T(), key.DestroyedAt.IsZero())
	})

	s.Run("should fail with bad request if curve is not supported", func() {
		keyID := "my-key-import"
		request := &types.ImportKeyRequest{
			Curve:            "invalidCurve",
			SigningAlgorithm: "eddsa",
			PrivateKey:       ecdsaPrivKey,
			Tags: map[string]string{
				"myTag0": "tag0",
				"myTag1": "tag1",
			},
		}

		key, err := s.keyManagerClient.ImportKey(s.ctx, s.cfg.HashicorpKeyStore, keyID, request)
		require.Nil(s.T(), key)

		httpError := err.(*client.ResponseError)
		assert.Equal(s.T(), 400, httpError.StatusCode)
	})

	s.Run("should fail with bad request if signing algorithm is not supported", func() {
		keyID := "my-key-import"
		request := &types.ImportKeyRequest{
			Curve:            "secp256k1",
			SigningAlgorithm: "invalidSigningAlgorithm",
			PrivateKey:       ecdsaPrivKey,
			Tags: map[string]string{
				"myTag0": "tag0",
				"myTag1": "tag1",
			},
		}

		key, err := s.keyManagerClient.ImportKey(s.ctx, s.cfg.HashicorpKeyStore, keyID, request)
		require.Nil(s.T(), key)

		httpError := err.(*client.ResponseError)
		assert.Equal(s.T(), 400, httpError.StatusCode)
	})
}

func (s *keysTestSuite) TestGet() {
	keyID := "my-key-get"
	request := &types.ImportKeyRequest{
		Curve:            "secp256k1",
		SigningAlgorithm: "ecdsa",
		PrivateKey:       ecdsaPrivKey,
		Tags: map[string]string{
			"myTag0": "tag0",
			"myTag1": "tag1",
		},
	}

	key, err := s.keyManagerClient.ImportKey(s.ctx, s.cfg.HashicorpKeyStore, keyID, request)
	require.NoError(s.T(), err)

	s.Run("should get a key successfully", func() {
		keyRetrieved, err := s.keyManagerClient.GetKey(s.ctx, s.cfg.HashicorpKeyStore, key.ID)
		require.NoError(s.T(), err)

		assert.Equal(s.T(), "BFVSFJhqUh9DQJwcayNtsWdDMvqq8R/EKnBHqwd4Hr5vCXTyJlqKfYIgj4jCGixVZjsz5a+S2RklJRFjjoLf+LI=", keyRetrieved.PublicKey)
		assert.Equal(s.T(), keyID, keyRetrieved.ID)
		assert.Equal(s.T(), request.Tags, keyRetrieved.Tags)
		assert.False(s.T(), keyRetrieved.Disabled)
		assert.NotEmpty(s.T(), keyRetrieved.CreatedAt)
		assert.NotEmpty(s.T(), keyRetrieved.UpdatedAt)
		assert.True(s.T(), keyRetrieved.ExpireAt.IsZero())
		assert.True(s.T(), keyRetrieved.DeletedAt.IsZero())
		assert.True(s.T(), keyRetrieved.DestroyedAt.IsZero())
	})
}

func (s *keysTestSuite) TestList() {
	keyID := "my-key-list"
	request := &types.ImportKeyRequest{
		Curve:            "secp256k1",
		SigningAlgorithm: "ecdsa",
		PrivateKey:       ecdsaPrivKey,
		Tags: map[string]string{
			"myTag0": "tag0",
			"myTag1": "tag1",
		},
	}

	key, err := s.keyManagerClient.ImportKey(s.ctx, s.cfg.HashicorpKeyStore, keyID, request)
	require.NoError(s.T(), err)

	s.Run("should get all key ids successfully", func() {
		ids, err := s.keyManagerClient.ListKeys(s.ctx, s.cfg.HashicorpKeyStore)
		require.NoError(s.T(), err)

		assert.GreaterOrEqual(s.T(), len(ids), 1)
		assert.Contains(s.T(), ids, key.ID)
	})

	s.Run("should parse errors successfully", func() {
		ids, err := s.keyManagerClient.ListKeys(s.ctx, "inexistentStoreName")
		require.Empty(s.T(), ids)

		httpError := err.(*client.ResponseError)
		assert.Equal(s.T(), 404, httpError.StatusCode)
	})
}

func (s *keysTestSuite) TestSignVerify() {
	data := []byte("my data to sign")
	hashedPayload := crypto.Keccak256(data)

	s.Run("should sign a new payload successfully: Secp256k1/ECDSA", func() {
		keyID := "my-key-sign-ecdsa"
		request := &types.ImportKeyRequest{
			Curve:            "secp256k1",
			PrivateKey:       ecdsaPrivKey,
			SigningAlgorithm: "ecdsa",
		}

		key, err := s.keyManagerClient.ImportKey(s.ctx, s.cfg.HashicorpKeyStore, keyID, request)
		require.NoError(s.T(), err)

		requestSign := &types.SignBase64PayloadRequest{
			Data: hashedPayload,
		}
		signature, err := s.keyManagerClient.SignKey(s.ctx, s.cfg.HashicorpKeyStore, key.ID, requestSign)
		require.NoError(s.T(), err)

		assert.Equal(s.T(), "YzQeLIN0Sd43Nbb0QCsVSqChGNAuRaKzEfujnERAJd0523aZyz2KXK93KKh+d4ws3MxAhc8qNG43wYI97Fzi7Q==", signature)

		sigB, err := base64.StdEncoding.DecodeString(signature)
		require.NoError(s.T(), err)
		pubKeyB, err := base64.StdEncoding.DecodeString(key.PublicKey)
		require.NoError(s.T(), err)

		verifyRequest := &types.VerifyKeySignatureRequest{
			Data:             hashedPayload,
			Signature:        sigB,
			Curve:            key.Curve,
			SigningAlgorithm: key.SigningAlgorithm,
			PublicKey:        pubKeyB,
		}
		err = s.keyManagerClient.VerifyKeySignature(s.ctx, s.cfg.HashicorpKeyStore, verifyRequest)
		require.NoError(s.T(), err)
	})

	s.Run("should sign and verify a new payload successfully: BN254/EDDSA", func() {
		keyID := "my-key-sign-eddsa"
		request := &types.ImportKeyRequest{
			Curve:            "bn254",
			SigningAlgorithm: "eddsa",
			PrivateKey:       eddsaPrivKey,
		}
		key, err := s.keyManagerClient.ImportKey(s.ctx, s.cfg.HashicorpKeyStore, keyID, request)
		require.NoError(s.T(), err)

		requestSign := &types.SignBase64PayloadRequest{
			Data: data,
		}
		signature, err := s.keyManagerClient.SignKey(s.ctx, s.cfg.HashicorpKeyStore, key.ID, requestSign)
		require.NoError(s.T(), err)

		assert.Equal(s.T(), "tdpR9JkX7lKSugSvYJX2icf6/uQnCAmXG9v/FG26vS0AcBqg6eVakZQNYwfic/Ec3LWqzSbXg54TBteQq6grdw==", signature)

		sigB, _ := base64.StdEncoding.DecodeString(signature)
		pubKeyB, _ := base64.StdEncoding.DecodeString(key.PublicKey)
		verifyRequest := &types.VerifyKeySignatureRequest{
			Data:             data,
			Signature:        sigB,
			Curve:            key.Curve,
			SigningAlgorithm: key.SigningAlgorithm,
			PublicKey:        pubKeyB,
		}
		err = s.keyManagerClient.VerifyKeySignature(s.ctx, s.cfg.HashicorpKeyStore, verifyRequest)
		require.NoError(s.T(), err)
	})
}
