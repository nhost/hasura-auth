package controller

/*
func (ctrl *Controller) GetOauthTwitter( //nolint:ireturn
	_ context.Context,
	_ api.GetOauthTwitterRequestObject,
) (api.GetOauthTwitterResponseObject, error) {
	verifier := oauth2.GenerateVerifier()

	url := ctrl.Oauth.Twitter.AuthCodeURL(
		"state",
		oauth2.S256ChallengeOption(verifier),
	)

	return api.GetOauthTwitter302Response{
		Headers: api.GetOauthTwitter302ResponseHeaders{
			Location:  url,
			SetCookie: fmt.Sprintf("verifier=%s; HttpOnly; Max-Age=3600; Secure: true;", verifier),
		},
	}, nil
}

func (ctrl *Controller) GetOauthTwitterCallback( //nolint:ireturn
	ctx context.Context,
	request api.GetOauthTwitterCallbackRequestObject,
) (api.GetOauthTwitterCallbackResponseObject, error) {
	// state := request.Params.State
	code := request.Params.Code

	tok, err := ctrl.Oauth.Twitter.Exchange(
		ctx,
		code,
		oauth2.VerifierOption(request.Params.Verifier),
	)
	if err != nil {
		return nil, fmt.Errorf("error exchanging code: %w", err)
	}

	b, err := json.Marshal(tok)
	if err != nil {
		return nil, fmt.Errorf("error marshalling token: %w", err)
	}
	fmt.Println("tok:", string(b))

	client := ctrl.Oauth.Twitter.Client(ctx, tok)
	// resp, err := client.Get("https://api.twitter.com/2/users/me?user.fields=profile_image_url")
	resp, err := client.Get("https://api.twitter.com/1.1/account/verify_credentials.json?include_email=true")
	if err != nil {
		return nil, fmt.Errorf("error getting user info: %w", err)
	}
	defer resp.Body.Close()

	b, err = io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}

	fmt.Println("user info:", string(b))

	fmt.Println("response:", resp.Status)

	return api.GetOauthTwitterCallback302Response{
		Headers: api.GetOauthTwitterCallback302ResponseHeaders{
			Location: "http://localhost:3000",
		},
	}, nil
}
*/
