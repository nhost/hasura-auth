package controller_test

import (
	"encoding/json"
	"errors"
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/nhost/hasura-auth/go/api"
	"github.com/nhost/hasura-auth/go/controller"
	"github.com/nhost/hasura-auth/go/controller/mock"
	"github.com/nhost/hasura-auth/go/sql"
	"go.uber.org/mock/gomock"
)

func unmarshalUserWebauthnVerifyRequest(
	t *testing.T,
	b []byte,
) *api.UserAddSecurityKeyVerifyRequest {
	t.Helper()

	var v *api.UserAddSecurityKeyVerifyRequest
	if err := json.Unmarshal(b, &v); err != nil {
		t.Fatal(err)
	}

	return v
}

func TestPostUserWebauthnVerify(t *testing.T) { //nolint:maintidx
	t.Parallel()

	userID := uuid.MustParse("DB477732-48FA-4289-B694-2886A646B6EB")
	securityKeyID := uuid.MustParse("307b758d-c0b0-4ce3-894b-f8ddec753c29")

	jwtTokenFn := func() *jwt.Token {
		return &jwt.Token{
			Raw:    "",
			Method: jwt.SigningMethodHS256,
			Header: map[string]any{
				"alg": "HS256",
				"typ": "JWT",
			},
			Claims: jwt.MapClaims{
				"exp": float64(time.Now().Add(900 * time.Second).Unix()),
				"https://hasura.io/jwt/claims": map[string]any{
					"x-hasura-allowed-roles": []any{"user"},
					"x-hasura-default-role":  "user",
					"x-hasura-user-id":       userID.String(),
				},
				"iat": float64(time.Now().Unix()),
				"iss": "hasura-auth",
				"sub": userID.String(),
			},
			Signature: []byte{},
			Valid:     true,
		}
	}

	_, webauthnTouchIDChallenge := webAuthnTouchID(t)
	_, webauthnWindowsHelloChallenge := webAuthnWindowsHello(t)

	cases := []testRequest[api.PostUserWebauthnVerifyRequestObject, api.PostUserWebauthnVerifyResponseObject]{
		{ //nolint:dupl
			name:   "success - TouchID with nickname",
			config: getConfig,
			db: func(ctrl *gomock.Controller) controller.DBClient {
				mock := mock.NewMockDBClient(ctrl)

				mock.EXPECT().GetUser(
					gomock.Any(),
					userID,
				).Return(sql.AuthUser{ //nolint:exhaustruct
					ID:          userID,
					Email:       sql.Text("jane@acme.com"),
					DisplayName: "Jane Doe",
					Disabled:    false,
				}, nil)

				mock.EXPECT().InsertSecurityKey(
					gomock.Any(),
					gomock.Any(),
				).Return(securityKeyID, nil)

				return mock
			},
			request: api.PostUserWebauthnVerifyRequestObject{
				Body: unmarshalUserWebauthnVerifyRequest(
					t,
					//nolint:lll
					[]byte(
						`{"credential":{"id":"LychOomEPgZu4XNwiDvzlP5hd1U","rawId":"LychOomEPgZu4XNwiDvzlP5hd1U","response":{"attestationObject":"o2NmbXRkbm9uZWdhdHRTdG10oGhhdXRoRGF0YViY0RE6Bmg2J-FxNrC8136ZQSeTWKWtdni_Lpfv5XR4bDtdAAAAAPv8MAcVTk7MjAtuAgVX170AFC8nITqJhD4GbuFzcIg785T-YXdVpQECAyYgASFYIFfhtYKglcQa82Wd3cJoz2Y1JTKlhiL798bGCG2pyWR_Ilggo1CUEbgnUq5G7FY6OzptcSQQZq6yV3XVu5iM0MWRH2U","clientDataJSON":"eyJjaGFsbGVuZ2UiOiJ6em56dGp2RlZVTTBFMnA4WlY2c2hYRWN3MmY0dGJ6NVJyZlpXazRWUFhJIiwib3JpZ2luIjoiaHR0cHM6Ly9yZWFjdC1hcG9sbG8uZXhhbXBsZS5uaG9zdC5pbyIsInR5cGUiOiJ3ZWJhdXRobi5jcmVhdGUifQ","transports":["internal"],"publicKeyAlgorithm":-7,"publicKey":"MFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEV-G1gqCVxBrzZZ3dwmjPZjUlMqWGIvv3xsYIbanJZH-jUJQRuCdSrkbsVjo7Om1xJBBmrrJXddW7mIzQxZEfZQ","authenticatorData":"0RE6Bmg2J-FxNrC8136ZQSeTWKWtdni_Lpfv5XR4bDtdAAAAAPv8MAcVTk7MjAtuAgVX170AFC8nITqJhD4GbuFzcIg785T-YXdVpQECAyYgASFYIFfhtYKglcQa82Wd3cJoz2Y1JTKlhiL798bGCG2pyWR_Ilggo1CUEbgnUq5G7FY6OzptcSQQZq6yV3XVu5iM0MWRH2U"},"type":"public-key","clientExtensionResults":{},"authenticatorAttachment":"platform"},"nickname":"TouchID"}`,
					),
				),
			},
			expectedResponse: api.PostUserWebauthnVerify200JSONResponse{
				Id:       securityKeyID.String(),
				Nickname: ptr("TouchID"),
			},
			expectedJWT:       nil,
			jwtTokenFn:        jwtTokenFn,
			getControllerOpts: []getControllerOptsFunc{},
		},

		{ //nolint:dupl
			name:   "success - Windows Hello without nickname",
			config: getConfig,
			db: func(ctrl *gomock.Controller) controller.DBClient {
				mock := mock.NewMockDBClient(ctrl)

				mock.EXPECT().GetUser(
					gomock.Any(),
					userID,
				).Return(sql.AuthUser{ //nolint:exhaustruct
					ID:                  userID,
					Email:               sql.Text("jane@acme.com"),
					DisplayName:         "Jane Doe",
					AvatarUrl:           "",
					Locale:              "en",
					Disabled:            false,
					EmailVerified:       true,
					PhoneNumber:         sql.Text(""),
					PhoneNumberVerified: false,
					DefaultRole:         "user",
					IsAnonymous:         false,
					Ticket:              sql.Text(""),
					TicketExpiresAt:     sql.TimestampTz(time.Now().Add(24 * time.Hour)),
					ActiveMfaType:       sql.Text(""),
					NewEmail:            sql.Text(""),
					Metadata:            []byte("{}"),
					CreatedAt:           sql.TimestampTz(time.Now()),
					UpdatedAt:           sql.TimestampTz(time.Now()),
					LastSeen:            sql.TimestampTz(time.Now()),
				}, nil)

				mock.EXPECT().InsertSecurityKey(
					gomock.Any(),
					gomock.Any(),
				).Return(securityKeyID, nil)

				return mock
			},
			request: api.PostUserWebauthnVerifyRequestObject{
				Body: unmarshalUserWebauthnVerifyRequest(
					t,
					//nolint:lll
					[]byte(
						`{"credential":{"id":"t4r2_E24k3bp-LwQUz5M2xazSsWfZpATRPtaelkfqfc","rawId":"t4r2_E24k3bp-LwQUz5M2xazSsWfZpATRPtaelkfqfc","response":{"attestationObject":"o2NmbXRjdHBtZ2F0dFN0bXSmY2FsZzn__mNzaWdZAQDD6T1Xbcklo2ZbVD93TxxUh4LIlQgJopKlIEiqFGsYcvrzzR4D6IdDN0uQbNRcoS1ZKmzQ_v2gXmj8yorBt9LJ8zN4jSzUjoq4Yp_yZrZtVFwnNTTvPdvMxMUQoMS-lbzTZz_-w1nrkfzkGs_r_Wks-i-wKo5gVi45t1mjjuYijdNBPNNBD9MFXLjQXgfIR8u1KxckxqdaxTSl2E4jzRuC5W7IY0a6XUrgz_Z6fI1C780XdvrkXdWeni-9l4Nj3e5cKtCjHvwx-01mcEU2Kk1t3s9xIegMGJ0rQvySIzkiL7PhiMbLp0eCjczUaFtI9FLvU1h69waTiOaUi-myUunZY3ZlcmMyLjBjeDVjglkFvTCCBbkwggOhoAMCAQICEDYL-Azg3EijplSpYoe60dswDQYJKoZIhvcNAQELBQAwQjFAMD4GA1UEAxM3TkNVLUlOVEMtS0VZSUQtRUE5NTBEOTg3QkNGRjBERjlBQUMxRkVGRkI4QUI0ODAzRkVGMkNBMjAeFw0yMzEyMDMwODIzMDVaFw0yODA1MTIyMDQ3NTdaMAAwggEiMA0GCSqGSIb3DQEBAQUAA4IBDwAwggEKAoIBAQDFEDvH0Q1QuWKBNbm69EkzR9ybpL0_29R0vRIj2MahV9ZYK6FN_LjF12Ai7G7YTvGiIS5skMiFwzBhanIAcHpLZaaMb9Rywnp5cjp6414crt648nzvbq_oOQG_acO-LqfPTj2I-zhie3nAQz6r9KV7jK7I3p_-2DxRrHzyAlo71gOC8MsX0RARfbLslnsLGra_CWSATp3cEuTKh0PRsERNK85mSQ85pIiAUS5AAEFhk19sT3CJdjBr6sOz5cg0JO6hQ-upnrRK6yWJqF599OdN1Fm9uIX68mDiMl3fA3vwEItBe2lC4EJ_jkVs1KkLmvTQkPTRmgw6RAUFqS7HcCPRAgMBAAGjggHrMIIB5zAOBgNVHQ8BAf8EBAMCB4AwDAYDVR0TAQH_BAIwADBtBgNVHSABAf8EYzBhMF8GCSsGAQQBgjcVHzBSMFAGCCsGAQUFBwICMEQeQgBUAEMAUABBACAAIABUAHIAdQBzAHQAZQBkACAAIABQAGwAYQB0AGYAbwByAG0AIAAgAEkAZABlAG4AdABpAHQAeTAQBgNVHSUECTAHBgVngQUIAzBQBgNVHREBAf8ERjBEpEIwQDEWMBQGBWeBBQIBDAtpZDo0OTRFNTQ0MzEOMAwGBWeBBQICDANBREwxFjAUBgVngQUCAwwLaWQ6MDI1ODAwMTIwHwYDVR0jBBgwFoAUH5exkeyU6vAkYZshYb-lI1ji_rkwHQYDVR0OBBYEFPb2n64JiqOUAfDaHCOS6pBB74tOMIGzBggrBgEFBQcBAQSBpjCBozCBoAYIKwYBBQUHMAKGgZNodHRwOi8vYXpjc3Byb2RuY3VhaWtwdWJsaXNoLmJsb2IuY29yZS53aW5kb3dzLm5ldC9uY3UtaW50Yy1rZXlpZC1lYTk1MGQ5ODdiY2ZmMGRmOWFhYzFmZWZmYjhhYjQ4MDNmZWYyY2EyL2JmZWMyZGUwLWJjNzYtNGNlNi04ZWUwLTk3MjJhYWI1MjBlNS5jZXIwDQYJKoZIhvcNAQELBQADggIBADx3TcZ-t0jsuIOCv5Qa7QvfikgtiKV2BnwAfg_Sy_GfX4r3Wf5YL6B-GlGzdUv3aXWN8wucdVzoK-0MqVYCyIvUHwOZB-vxsTQl9vDIe144aHSHqUH8MkgxKixnMsi6-ODw-hDebWQxfSOXqqnG1s91qDiEXRcPFDtOTKlH4GS13nNNa8qfZf2NriuALFpqmmpooWbOJG857xgtLCIvMTiiW_oF55l-d6oqfK2rkfGWvIMXbKgaXFIC5mDEkIESSZCr4Fmdal51r3-lsk0-SBPu8tsapBlm1yHhWmjqHR8fsAtREj2Q4qegqozC38QBJgPJh047P3khL65W-iHc9_nI5b-QW3S3F4RYe8iGCvccz9PCHjjIgpgmPW0aaO4pVmDDTKokZAgWsc-t4h-xVSuEMAixjqvsalP1oSfzP8bAZwewOPSHLdtTFQnT_3JcusD31PThAyU5lC-VaI24I6wOk2F1Yrow3ze1jZXn4QVczhuA61ThbN-0fAfMh1KbrXjcJG5mOjV68yA0BwWxuOKYeT1u7uj6zt-yBBbNAvX_gSwrtQ0kSzMq2AvvlkVbyI25QnDYbKFHtbSdog4ghb018-q7rUepsqaDcolNWocqrNOiA785WvmSMi7b_szRr8fDUTcZCFJpxL4BMmmxbkhUc6Wkfs6q5cCsgk72PcBnWQbwMIIG7DCCBNSgAwIBAgITMwAAB4cFQab9pgXr9wAAAAAHhzANBgkqhkiG9w0BAQsFADCBjDELMAkGA1UEBhMCVVMxEzARBgNVBAgTCldhc2hpbmd0b24xEDAOBgNVBAcTB1JlZG1vbmQxHjAcBgNVBAoTFU1pY3Jvc29mdCBDb3Jwb3JhdGlvbjE2MDQGA1UEAxMtTWljcm9zb2Z0IFRQTSBSb290IENlcnRpZmljYXRlIEF1dGhvcml0eSAyMDE0MB4XDTIyMDUxMjIwNDc1N1oXDTI4MDUxMjIwNDc1N1owQjFAMD4GA1UEAxM3TkNVLUlOVEMtS0VZSUQtRUE5NTBEOTg3QkNGRjBERjlBQUMxRkVGRkI4QUI0ODAzRkVGMkNBMjCCAiIwDQYJKoZIhvcNAQEBBQADggIPADCCAgoCggIBAM0r9233MxhyPRqWif286clCz-LmWAWLPcr6faoLCnJ300lt_XMArj0xZ9iSI5q1cPWTTa5J9RPS8Bs1rIzsdoJtJGz8m11UGfKmpQIjVc7-eXm8j47Zzjvcw99XjHwYnsCpEGw4-rEkh_yUMpWbqqUzS5nIb9Ce8g1b2cvh9bKv7qpqVspd0YmX819faMFbSBMI1ZT-fIpEm2PWVjivXnFdkKOv5hp_4aWlO9ivneZg685QN3TsAiooQ3n1hEgLESkK_1YsEWiI_iEGpXNuZSSokoLA9h-Mox-QUTrEEmc7A-g2mDoVYX4kCQrSEVGTEqAwuhUAtUrTzjGo6M2Re7ZS2Sr7JlVhpeXhCS3jqRkp6MJc5zSHP-0YyfsQ7fs4pA7DBF-5I_BX9oN-m64kNtr7BQHbnzeT-i-oqAn6KVbjbOqoAcZCWuVsoU2tnRrnTADH0rRREsQq-YOCcJU5aQ0GtEfVqM2arbeR2nmvqxEWp21mFTWwQfMQUnjhjiA1qzCqk8JbvuNfbU8jDy90L0RbUc75VQhFGRmu8BAyHjmN-nFR-cbwbBO40eV5Wz0282-tBRMX_TQAZjC_tK7idY2-M6MmABtor3bLR_ylJQg3r5Z3zQqY44u0zktB5WdVbSykr23oki4yT1DfZmzrxloUaLm6o0WfqIJdIZ3ATtwbAgMBAAGjggGOMIIBijAOBgNVHQ8BAf8EBAMCAoQwGwYDVR0lBBQwEgYJKwYBBAGCNxUkBgVngQUIAzAWBgNVHSAEDzANMAsGCSsGAQQBgjcVHzASBgNVHRMBAf8ECDAGAQH_AgEAMB0GA1UdDgQWBBQfl7GR7JTq8CRhmyFhv6UjWOL-uTAfBgNVHSMEGDAWgBR6jArOL0hiF-KU0a5VwVLscXSkVjBwBgNVHR8EaTBnMGWgY6Bhhl9odHRwOi8vd3d3Lm1pY3Jvc29mdC5jb20vcGtpb3BzL2NybC9NaWNyb3NvZnQlMjBUUE0lMjBSb290JTIwQ2VydGlmaWNhdGUlMjBBdXRob3JpdHklMjAyMDE0LmNybDB9BggrBgEFBQcBAQRxMG8wbQYIKwYBBQUHMAKGYWh0dHA6Ly93d3cubWljcm9zb2Z0LmNvbS9wa2lvcHMvY2VydHMvTWljcm9zb2Z0JTIwVFBNJTIwUm9vdCUyMENlcnRpZmljYXRlJTIwQXV0aG9yaXR5JTIwMjAxNC5jcnQwDQYJKoZIhvcNAQELBQADggIBAHm8r8YHw3ozhenIyk4I5zVuDLol2FJiABcnWPk4DtzO63J7trqM2Cr_bsBUd2BDVxS8R9TQDR8iSXbPH_BbHHC7XYK2NjaGTHpV7f5GBloLDal-FXP0glD8bX7aOUp1ks4RJtv9scF8TnQgUtaWsweVi8sn1km4QLaE4tutdBkyvkIyR26ZihVm67Wmpg-JbQkt4ksB840YtmUgXbnmbV8byQQAvpYC5dl1aJBGSyz_sMgivHw11pXJAfgKurpwfG9IC5-k5-9vPa3XjyT484wT24R1gc4Zj_jEfh7z6w1ppxd9XbxYv0fHg5xCPoWt-tFndKuZOxDVWBnNzJ6zCw5Tezbax1PUpYljElwP2mylJkeK8EzbVwbMJUzW4uKKRe5kfTCxDT0gArVUWdqHhEY34rzkx9wI22f-mQl6NgcGW290AuEQ0L_Ni0Qqj_P2lC1YlTrAr90QxWEwouVZ7BLD9eHa_TBqelrE1kdd6NAzorU3m4aAVwW3BfEGxE54y5kSi8QTqC9CTsuAAGyPBuw0N_cr16KNus2F3pgNbQKHv7fblVaQNY4c9q0zL3nU1T4aVJGz8N0-hWc4H5j3hf3xhRL-jiNljBT2l11sOOmo4SjYCBRwdtBnVSJUrx4T2OrJ5klpUuZUrrbF7IO8IXxQTnsvI8g5MgDAKwtltBnHZ3B1YkFyZWFYdgAjAAsABAByACCd_8vzbDg65pn7mGjcbcuJ1xU4hL4oA5IsEkFYv60irgAQABAAAwAQACCc5JpkK9fmO9nCNd1rYQ7jd7GOro71OAkhaN4GxP2DdQAgsPo5B-oUPuIa2KCvefgsmxzDZdJDWjoRDa3v9zmTnrVoY2VydEluZm9Yof9UQ0eAFwAiAAt3Gqx7OUvHQid6kxmYviEHjoZ6nPhunCfjGfluPrvZWAAUaaePbyq16tk9DUi5NaYPaBbFKNYAAAACLwgJ6vX85sF7no1_ATCwK8mFs9NtACIAC9NjnLHa29t1i8vfJtZQytQSxB3v9YkNaBpktQ3nXitaACIAC7FXrVKJyNWcApoiMfNCaSiweyShCjhneC4Fx2mZoLKlaGF1dGhEYXRhWKTREToGaDYn4XE2sLzXfplBJ5NYpa12eL8ul-_ldHhsO0UAAAAACJhwWMrcS4G24TDeUNy-lgAgt4r2_E24k3bp-LwQUz5M2xazSsWfZpATRPtaelkfqfelAQIDJiABIVggnOSaZCvX5jvZwjXda2EO43exjq6O9TgJIWjeBsT9g3UiWCCw-jkH6hQ-4hrYoK95-CybHMNl0kNaOhENre_3OZOetQ","clientDataJSON":"eyJ0eXBlIjoid2ViYXV0aG4uY3JlYXRlIiwiY2hhbGxlbmdlIjoienY5bFBUSnBPbGd4emxyS1dsLXRHN0FkeGVVSWJDd3hxVjhNRlpaTlJkQSIsIm9yaWdpbiI6Imh0dHBzOi8vcmVhY3QtYXBvbGxvLmV4YW1wbGUubmhvc3QuaW8iLCJjcm9zc09yaWdpbiI6ZmFsc2V9","transports":["internal"],"publicKeyAlgorithm":-7},"type":"public-key","clientExtensionResults":{},"authenticatorAttachment":"platform"}}`,
					),
				),
			},
			expectedResponse: api.PostUserWebauthnVerify200JSONResponse{
				Id:       securityKeyID.String(),
				Nickname: ptr(""),
			},
			expectedJWT:       nil,
			jwtTokenFn:        jwtTokenFn,
			getControllerOpts: []getControllerOptsFunc{},
		},

		{
			name: "webauthn disabled",
			config: func() *controller.Config {
				config := getConfig()
				config.WebauthnEnabled = false

				return config
			},
			db: func(ctrl *gomock.Controller) controller.DBClient {
				mock := mock.NewMockDBClient(ctrl)

				return mock
			},
			request: api.PostUserWebauthnVerifyRequestObject{
				Body: unmarshalUserWebauthnVerifyRequest(
					t,
					//nolint:lll
					[]byte(
						`{"credential":{"id":"LychOomEPgZu4XNwiDvzlP5hd1U","rawId":"LychOomEPgZu4XNwiDvzlP5hd1U","response":{"attestationObject":"o2NmbXRkbm9uZWdhdHRTdG10oGhhdXRoRGF0YViY0RE6Bmg2J-FxNrC8136ZQSeTWKWtdni_Lpfv5XR4bDtdAAAAAPv8MAcVTk7MjAtuAgVX170AFC8nITqJhD4GbuFzcIg785T-YXdVpQECAyYgASFYIFfhtYKglcQa82Wd3cJoz2Y1JTKlhiL798bGCG2pyWR_Ilggo1CUEbgnUq5G7FY6OzptcSQQZq6yV3XVu5iM0MWRH2U","clientDataJSON":"eyJjaGFsbGVuZ2UiOiJ6em56dGp2RlZVTTBFMnA4WlY2c2hYRWN3MmY0dGJ6NVJyZlpXazRWUFhJIiwib3JpZ2luIjoiaHR0cHM6Ly9yZWFjdC1hcG9sbG8uZXhhbXBsZS5uaG9zdC5pbyIsInR5cGUiOiJ3ZWJhdXRobi5jcmVhdGUifQ","transports":["internal"],"publicKeyAlgorithm":-7,"publicKey":"MFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEV-G1gqCVxBrzZZ3dwmjPZjUlMqWGIvv3xsYIbanJZH-jUJQRuCdSrkbsVjo7Om1xJBBmrrJXddW7mIzQxZEfZQ","authenticatorData":"0RE6Bmg2J-FxNrC8136ZQSeTWKWtdni_Lpfv5XR4bDtdAAAAAPv8MAcVTk7MjAtuAgVX170AFC8nITqJhD4GbuFzcIg785T-YXdVpQECAyYgASFYIFfhtYKglcQa82Wd3cJoz2Y1JTKlhiL798bGCG2pyWR_Ilggo1CUEbgnUq5G7FY6OzptcSQQZq6yV3XVu5iM0MWRH2U"},"type":"public-key","clientExtensionResults":{},"authenticatorAttachment":"platform"},"nickname":"TouchID"}`,
					),
				),
			},
			expectedResponse: controller.ErrorResponse{
				Error:   "disabled-endpoint",
				Message: "This endpoint is disabled",
				Status:  409,
			},
			expectedJWT:       nil,
			jwtTokenFn:        jwtTokenFn,
			getControllerOpts: []getControllerOptsFunc{},
		},

		{
			name:   "no jwt token",
			config: getConfig,
			db: func(ctrl *gomock.Controller) controller.DBClient {
				mock := mock.NewMockDBClient(ctrl)

				return mock
			},
			request: api.PostUserWebauthnVerifyRequestObject{
				Body: unmarshalUserWebauthnVerifyRequest(
					t,
					//nolint:lll
					[]byte(
						`{"credential":{"id":"LychOomEPgZu4XNwiDvzlP5hd1U","rawId":"LychOomEPgZu4XNwiDvzlP5hd1U","response":{"attestationObject":"o2NmbXRkbm9uZWdhdHRTdG10oGhhdXRoRGF0YViY0RE6Bmg2J-FxNrC8136ZQSeTWKWtdni_Lpfv5XR4bDtdAAAAAPv8MAcVTk7MjAtuAgVX170AFC8nITqJhD4GbuFzcIg785T-YXdVpQECAyYgASFYIFfhtYKglcQa82Wd3cJoz2Y1JTKlhiL798bGCG2pyWR_Ilggo1CUEbgnUq5G7FY6OzptcSQQZq6yV3XVu5iM0MWRH2U","clientDataJSON":"eyJjaGFsbGVuZ2UiOiJ6em56dGp2RlZVTTBFMnA4WlY2c2hYRWN3MmY0dGJ6NVJyZlpXazRWUFhJIiwib3JpZ2luIjoiaHR0cHM6Ly9yZWFjdC1hcG9sbG8uZXhhbXBsZS5uaG9zdC5pbyIsInR5cGUiOiJ3ZWJhdXRobi5jcmVhdGUifQ","transports":["internal"],"publicKeyAlgorithm":-7,"publicKey":"MFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEV-G1gqCVxBrzZZ3dwmjPZjUlMqWGIvv3xsYIbanJZH-jUJQRuCdSrkbsVjo7Om1xJBBmrrJXddW7mIzQxZEfZQ","authenticatorData":"0RE6Bmg2J-FxNrC8136ZQSeTWKWtdni_Lpfv5XR4bDtdAAAAAPv8MAcVTk7MjAtuAgVX170AFC8nITqJhD4GbuFzcIg785T-YXdVpQECAyYgASFYIFfhtYKglcQa82Wd3cJoz2Y1JTKlhiL798bGCG2pyWR_Ilggo1CUEbgnUq5G7FY6OzptcSQQZq6yV3XVu5iM0MWRH2U"},"type":"public-key","clientExtensionResults":{},"authenticatorAttachment":"platform"},"nickname":"TouchID"}`,
					),
				),
			},
			expectedResponse: controller.ErrorResponse{
				Error:   "invalid-request",
				Message: "The request payload is incorrect",
				Status:  400,
			},
			expectedJWT:       nil,
			jwtTokenFn:        nil,
			getControllerOpts: []getControllerOptsFunc{},
		},

		{
			name:   "user not found",
			config: getConfig,
			db: func(ctrl *gomock.Controller) controller.DBClient {
				mock := mock.NewMockDBClient(ctrl)

				mock.EXPECT().GetUser(
					gomock.Any(),
					userID,
				).Return(sql.AuthUser{}, pgx.ErrNoRows) //nolint:exhaustruct

				return mock
			},
			request: api.PostUserWebauthnVerifyRequestObject{
				Body: unmarshalUserWebauthnVerifyRequest(
					t,
					//nolint:lll
					[]byte(
						`{"credential":{"id":"LychOomEPgZu4XNwiDvzlP5hd1U","rawId":"LychOomEPgZu4XNwiDvzlP5hd1U","response":{"attestationObject":"o2NmbXRkbm9uZWdhdHRTdG10oGhhdXRoRGF0YViY0RE6Bmg2J-FxNrC8136ZQSeTWKWtdni_Lpfv5XR4bDtdAAAAAPv8MAcVTk7MjAtuAgVX170AFC8nITqJhD4GbuFzcIg785T-YXdVpQECAyYgASFYIFfhtYKglcQa82Wd3cJoz2Y1JTKlhiL798bGCG2pyWR_Ilggo1CUEbgnUq5G7FY6OzptcSQQZq6yV3XVu5iM0MWRH2U","clientDataJSON":"eyJjaGFsbGVuZ2UiOiJ6em56dGp2RlZVTTBFMnA4WlY2c2hYRWN3MmY0dGJ6NVJyZlpXazRWUFhJIiwib3JpZ2luIjoiaHR0cHM6Ly9yZWFjdC1hcG9sbG8uZXhhbXBsZS5uaG9zdC5pbyIsInR5cGUiOiJ3ZWJhdXRobi5jcmVhdGUifQ","transports":["internal"],"publicKeyAlgorithm":-7,"publicKey":"MFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEV-G1gqCVxBrzZZ3dwmjPZjUlMqWGIvv3xsYIbanJZH-jUJQRuCdSrkbsVjo7Om1xJBBmrrJXddW7mIzQxZEfZQ","authenticatorData":"0RE6Bmg2J-FxNrC8136ZQSeTWKWtdni_Lpfv5XR4bDtdAAAAAPv8MAcVTk7MjAtuAgVX170AFC8nITqJhD4GbuFzcIg785T-YXdVpQECAyYgASFYIFfhtYKglcQa82Wd3cJoz2Y1JTKlhiL798bGCG2pyWR_Ilggo1CUEbgnUq5G7FY6OzptcSQQZq6yV3XVu5iM0MWRH2U"},"type":"public-key","clientExtensionResults":{},"authenticatorAttachment":"platform"},"nickname":"TouchID"}`,
					),
				),
			},
			expectedResponse: controller.ErrorResponse{
				Error:   "invalid-email-password",
				Message: "Incorrect email or password",
				Status:  401,
			},
			expectedJWT:       nil,
			jwtTokenFn:        jwtTokenFn,
			getControllerOpts: []getControllerOptsFunc{},
		},

		{ //nolint:dupl
			name:   "user disabled",
			config: getConfig,
			db: func(ctrl *gomock.Controller) controller.DBClient {
				mock := mock.NewMockDBClient(ctrl)

				mock.EXPECT().GetUser(
					gomock.Any(),
					userID,
				).Return(sql.AuthUser{ //nolint:exhaustruct
					ID:          userID,
					Email:       sql.Text("jane@acme.com"),
					DisplayName: "Jane Doe",
					Disabled:    true,
				}, nil)

				return mock
			},
			request: api.PostUserWebauthnVerifyRequestObject{
				Body: unmarshalUserWebauthnVerifyRequest(
					t,
					//nolint:lll
					[]byte(
						`{"credential":{"id":"LychOomEPgZu4XNwiDvzlP5hd1U","rawId":"LychOomEPgZu4XNwiDvzlP5hd1U","response":{"attestationObject":"o2NmbXRkbm9uZWdhdHRTdG10oGhhdXRoRGF0YViY0RE6Bmg2J-FxNrC8136ZQSeTWKWtdni_Lpfv5XR4bDtdAAAAAPv8MAcVTk7MjAtuAgVX170AFC8nITqJhD4GbuFzcIg785T-YXdVpQECAyYgASFYIFfhtYKglcQa82Wd3cJoz2Y1JTKlhiL798bGCG2pyWR_Ilggo1CUEbgnUq5G7FY6OzptcSQQZq6yV3XVu5iM0MWRH2U","clientDataJSON":"eyJjaGFsbGVuZ2UiOiJ6em56dGp2RlZVTTBFMnA4WlY2c2hYRWN3MmY0dGJ6NVJyZlpXazRWUFhJIiwib3JpZ2luIjoiaHR0cHM6Ly9yZWFjdC1hcG9sbG8uZXhhbXBsZS5uaG9zdC5pbyIsInR5cGUiOiJ3ZWJhdXRobi5jcmVhdGUifQ","transports":["internal"],"publicKeyAlgorithm":-7,"publicKey":"MFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEV-G1gqCVxBrzZZ3dwmjPZjUlMqWGIvv3xsYIbanJZH-jUJQRuCdSrkbsVjo7Om1xJBBmrrJXddW7mIzQxZEfZQ","authenticatorData":"0RE6Bmg2J-FxNrC8136ZQSeTWKWtdni_Lpfv5XR4bDtdAAAAAPv8MAcVTk7MjAtuAgVX170AFC8nITqJhD4GbuFzcIg785T-YXdVpQECAyYgASFYIFfhtYKglcQa82Wd3cJoz2Y1JTKlhiL798bGCG2pyWR_Ilggo1CUEbgnUq5G7FY6OzptcSQQZq6yV3XVu5iM0MWRH2U"},"type":"public-key","clientExtensionResults":{},"authenticatorAttachment":"platform"},"nickname":"TouchID"}`,
					),
				),
			},
			expectedResponse: controller.ErrorResponse{
				Error:   "disabled-user",
				Message: "User is disabled",
				Status:  401,
			},
			expectedJWT:       nil,
			jwtTokenFn:        jwtTokenFn,
			getControllerOpts: []getControllerOptsFunc{},
		},

		{ //nolint:dupl
			name: "anonymous user",
			config: func() *controller.Config {
				config := getConfig()
				config.RequireEmailVerification = false

				return config
			},
			db: func(ctrl *gomock.Controller) controller.DBClient {
				mock := mock.NewMockDBClient(ctrl)

				mock.EXPECT().GetUser(
					gomock.Any(),
					userID,
				).Return(sql.AuthUser{ //nolint:exhaustruct
					ID:            userID,
					Email:         sql.Text(""),
					DisplayName:   "Anonymous",
					EmailVerified: false,
					Disabled:      false,
				}, nil)

				return mock
			},
			request: api.PostUserWebauthnVerifyRequestObject{
				Body: unmarshalUserWebauthnVerifyRequest(
					t,
					//nolint:lll
					[]byte(
						`{"credential":{"id":"LychOomEPgZu4XNwiDvzlP5hd1U","rawId":"LychOomEPgZu4XNwiDvzlP5hd1U","response":{"attestationObject":"o2NmbXRkbm9uZWdhdHRTdG10oGhhdXRoRGF0YViY0RE6Bmg2J-FxNrC8136ZQSeTWKWtdni_Lpfv5XR4bDtdAAAAAPv8MAcVTk7MjAtuAgVX170AFC8nITqJhD4GbuFzcIg785T-YXdVpQECAyYgASFYIFfhtYKglcQa82Wd3cJoz2Y1JTKlhiL798bGCG2pyWR_Ilggo1CUEbgnUq5G7FY6OzptcSQQZq6yV3XVu5iM0MWRH2U","clientDataJSON":"eyJjaGFsbGVuZ2UiOiJ6em56dGp2RlZVTTBFMnA4WlY2c2hYRWN3MmY0dGJ6NVJyZlpXazRWUFhJIiwib3JpZ2luIjoiaHR0cHM6Ly9yZWFjdC1hcG9sbG8uZXhhbXBsZS5uaG9zdC5pbyIsInR5cGUiOiJ3ZWJhdXRobi5jcmVhdGUifQ","transports":["internal"],"publicKeyAlgorithm":-7,"publicKey":"MFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEV-G1gqCVxBrzZZ3dwmjPZjUlMqWGIvv3xsYIbanJZH-jUJQRuCdSrkbsVjo7Om1xJBBmrrJXddW7mIzQxZEfZQ","authenticatorData":"0RE6Bmg2J-FxNrC8136ZQSeTWKWtdni_Lpfv5XR4bDtdAAAAAPv8MAcVTk7MjAtuAgVX170AFC8nITqJhD4GbuFzcIg785T-YXdVpQECAyYgASFYIFfhtYKglcQa82Wd3cJoz2Y1JTKlhiL798bGCG2pyWR_Ilggo1CUEbgnUq5G7FY6OzptcSQQZq6yV3XVu5iM0MWRH2U"},"type":"public-key","clientExtensionResults":{},"authenticatorAttachment":"platform"},"nickname":"TouchID"}`,
					),
				),
			},
			expectedResponse: controller.ErrorResponse{
				Error:   "forbidden-anonymous",
				Message: "Forbidden, user is anonymous.",
				Status:  403,
			},
			expectedJWT:       nil,
			jwtTokenFn:        jwtTokenFn,
			getControllerOpts: []getControllerOptsFunc{},
		},

		{
			name:   "email not verified",
			config: getConfig,
			db: func(ctrl *gomock.Controller) controller.DBClient {
				mock := mock.NewMockDBClient(ctrl)

				mock.EXPECT().GetUser(
					gomock.Any(),
					userID,
				).Return(sql.AuthUser{ //nolint:exhaustruct
					ID:          userID,
					Email:       sql.Text("jane@acme.com"),
					DisplayName: "Jane Doe",
					Disabled:    false,
				}, nil)

				return mock
			},
			request: api.PostUserWebauthnVerifyRequestObject{
				Body: unmarshalUserWebauthnVerifyRequest(
					t,
					//nolint:lll
					[]byte(
						`{"credential":{"id":"LychOomEPgZu4XNwiDvzlP5hd1U","rawId":"LychOomEPgZu4XNwiDvzlP5hd1U","response":{"attestationObject":"o2NmbXRkbm9uZWdhdHRTdG10oGhhdXRoRGF0YViY0RE6Bmg2J-FxNrC8136ZQSeTWKWtdni_Lpfv5XR4bDtdAAAAAPv8MAcVTk7MjAtuAgVX170AFC8nITqJhD4GbuFzcIg785T-YXdVpQECAyYgASFYIFfhtYKglcQa82Wd3cJoz2Y1JTKlhiL798bGCG2pyWR_Ilggo1CUEbgnUq5G7FY6OzptcSQQZq6yV3XVu5iM0MWRH2U","clientDataJSON":"eyJjaGFsbGVuZ2UiOiJ6em56dGp2RlZVTTBFMnA4WlY2c2hYRWN3MmY0dGJ6NVJyZlpXazRWUFhJIiwib3JpZ2luIjoiaHR0cHM6Ly9yZWFjdC1hcG9sbG8uZXhhbXBsZS5uaG9zdC5pbyIsInR5cGUiOiJ3ZWJhdXRobi5jcmVhdGUifQ","transports":["internal"],"publicKeyAlgorithm":-7,"publicKey":"MFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEV-G1gqCVxBrzZZ3dwmjPZjUlMqWGIvv3xsYIbanJZH-jUJQRuCdSrkbsVjo7Om1xJBBmrrJXddW7mIzQxZEfZQ","authenticatorData":"0RE6Bmg2J-FxNrC8136ZQSeTWKWtdni_Lpfv5XR4bDtdAAAAAPv8MAcVTk7MjAtuAgVX170AFC8nITqJhD4GbuFzcIg785T-YXdVpQECAyYgASFYIFfhtYKglcQa82Wd3cJoz2Y1JTKlhiL798bGCG2pyWR_Ilggo1CUEbgnUq5G7FY6OzptcSQQZq6yV3XVu5iM0MWRH2U"},"type":"public-key","clientExtensionResults":{},"authenticatorAttachment":"platform"},"nickname":"TouchID"}`,
					),
				),
			},
			expectedResponse: controller.ErrorResponse{
				Error:   "unverified-user",
				Message: "User is not verified.",
				Status:  401,
			},
			expectedJWT:       nil,
			jwtTokenFn:        jwtTokenFn,
			getControllerOpts: []getControllerOptsFunc{},
		},

		{
			name:   "invalid credential data",
			config: getConfig,
			db: func(ctrl *gomock.Controller) controller.DBClient {
				mock := mock.NewMockDBClient(ctrl)

				mock.EXPECT().GetUser(
					gomock.Any(),
					userID,
				).Return(sql.AuthUser{ //nolint:exhaustruct
					ID:                  userID,
					Email:               sql.Text("jane@acme.com"),
					DisplayName:         "Jane Doe",
					AvatarUrl:           "",
					Locale:              "en",
					Disabled:            false,
					EmailVerified:       true,
					PhoneNumber:         sql.Text(""),
					PhoneNumberVerified: false,
					DefaultRole:         "user",
					IsAnonymous:         false,
					Ticket:              sql.Text(""),
					TicketExpiresAt:     sql.TimestampTz(time.Now().Add(24 * time.Hour)),
					ActiveMfaType:       sql.Text(""),
					NewEmail:            sql.Text(""),
					Metadata:            []byte("{}"),
					CreatedAt:           sql.TimestampTz(time.Now()),
					UpdatedAt:           sql.TimestampTz(time.Now()),
					LastSeen:            sql.TimestampTz(time.Now()),
				}, nil)

				return mock
			},
			request: api.PostUserWebauthnVerifyRequestObject{
				Body: unmarshalUserWebauthnVerifyRequest(
					t,
					//nolint:lll
					[]byte(
						`{"credential":{"id":"invalid-credential","rawId":"invalid","response":{"attestationObject":"invalid","clientDataJSON":"invalid","transports":[],"publicKeyAlgorithm":-7},"type":"public-key","clientExtensionResults":{},"authenticatorAttachment":"platform"},"nickname":"TouchID"}`,
					),
				),
			},
			expectedResponse: controller.ErrorResponse{
				Error:   "invalid-request",
				Message: "The request payload is incorrect",
				Status:  400,
			},
			expectedJWT:       nil,
			jwtTokenFn:        jwtTokenFn,
			getControllerOpts: []getControllerOptsFunc{},
		},

		{
			name:   "database insert error",
			config: getConfig,
			db: func(ctrl *gomock.Controller) controller.DBClient {
				mock := mock.NewMockDBClient(ctrl)

				mock.EXPECT().GetUser(
					gomock.Any(),
					userID,
				).Return(sql.AuthUser{ //nolint:exhaustruct
					ID:          userID,
					Email:       sql.Text("jane@acme.com"),
					DisplayName: "Jane Doe",
					Disabled:    false,
				}, nil)

				mock.EXPECT().InsertSecurityKey(
					gomock.Any(),
					gomock.Any(),
				).Return(uuid.Nil, errors.New("database error")) //nolint:goerr113

				return mock
			},
			request: api.PostUserWebauthnVerifyRequestObject{
				Body: unmarshalUserWebauthnVerifyRequest(
					t,
					//nolint:lll
					[]byte(
						`{"credential":{"id":"LychOomEPgZu4XNwiDvzlP5hd1U","rawId":"LychOomEPgZu4XNwiDvzlP5hd1U","response":{"attestationObject":"o2NmbXRkbm9uZWdhdHRTdG10oGhhdXRoRGF0YViY0RE6Bmg2J-FxNrC8136ZQSeTWKWtdni_Lpfv5XR4bDtdAAAAAPv8MAcVTk7MjAtuAgVX170AFC8nITqJhD4GbuFzcIg785T-YXdVpQECAyYgASFYIFfhtYKglcQa82Wd3cJoz2Y1JTKlhiL798bGCG2pyWR_Ilggo1CUEbgnUq5G7FY6OzptcSQQZq6yV3XVu5iM0MWRH2U","clientDataJSON":"eyJjaGFsbGVuZ2UiOiJ6em56dGp2RlZVTTBFMnA4WlY2c2hYRWN3MmY0dGJ6NVJyZlpXazRWUFhJIiwib3JpZ2luIjoiaHR0cHM6Ly9yZWFjdC1hcG9sbG8uZXhhbXBsZS5uaG9zdC5pbyIsInR5cGUiOiJ3ZWJhdXRobi5jcmVhdGUifQ","transports":["internal"],"publicKeyAlgorithm":-7,"publicKey":"MFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEV-G1gqCVxBrzZZ3dwmjPZjUlMqWGIvv3xsYIbanJZH-jUJQRuCdSrkbsVjo7Om1xJBBmrrJXddW7mIzQxZEfZQ","authenticatorData":"0RE6Bmg2J-FxNrC8136ZQSeTWKWtdni_Lpfv5XR4bDtdAAAAAPv8MAcVTk7MjAtuAgVX170AFC8nITqJhD4GbuFzcIg785T-YXdVpQECAyYgASFYIFfhtYKglcQa82Wd3cJoz2Y1JTKlhiL798bGCG2pyWR_Ilggo1CUEbgnUq5G7FY6OzptcSQQZq6yV3XVu5iM0MWRH2U"},"type":"public-key","clientExtensionResults":{},"authenticatorAttachment":"platform"},"nickname":"TouchID"}`,
					),
				),
			},
			expectedResponse: controller.ErrorResponse{
				Error:   "internal-server-error",
				Message: "Internal server error",
				Status:  500,
			},
			expectedJWT:       nil,
			jwtTokenFn:        jwtTokenFn,
			getControllerOpts: []getControllerOptsFunc{},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			ctrl := gomock.NewController(t)

			c, jwtGetter := getController(t, ctrl, tc.config, tc.db, tc.getControllerOpts...)

			ctx := t.Context()
			if tc.jwtTokenFn != nil {
				ctx = jwtGetter.ToContext(ctx, tc.jwtTokenFn())
			}

			// Mock webauthn session storage for successful test cases
			if c.Webauthn != nil && tc.name == "success - TouchID with nickname" {
				sessionData := controller.WebauthnChallenge{
					Session: webauthnTouchIDChallenge.Session,
					User:    webauthnTouchIDChallenge.User,
					Options: webauthnTouchIDChallenge.Options,
				}
				c.Webauthn.Storage["zznztjvFVUM0E2p8ZV6shXEcw2f4tbz5RrfZWk4VPXI"] = sessionData
			}

			if c.Webauthn != nil && tc.name == "success - Windows Hello without nickname" {
				sessionData := controller.WebauthnChallenge{
					Session: webauthnWindowsHelloChallenge.Session,
					User:    webauthnWindowsHelloChallenge.User,
					Options: webauthnWindowsHelloChallenge.Options,
				}
				c.Webauthn.Storage["zv9lPTJpOlgxzlrKWl-tG7AdxeUIbCwxqV8MFZZNRdA"] = sessionData
			}

			if c.Webauthn != nil && tc.name == "database insert error" {
				sessionData := controller.WebauthnChallenge{
					Session: webauthnTouchIDChallenge.Session,
					User:    webauthnTouchIDChallenge.User,
					Options: webauthnTouchIDChallenge.Options,
				}
				c.Webauthn.Storage["zznztjvFVUM0E2p8ZV6shXEcw2f4tbz5RrfZWk4VPXI"] = sessionData
			}

			assertRequest(
				ctx,
				t,
				c.PostUserWebauthnVerify,
				tc.request,
				tc.expectedResponse,
			)
		})
	}
}
