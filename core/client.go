package core

import (
	"errors"
	"github.com/go-resty/resty/v2"
)

var (
	errUnsetRESTAPIKey = errors.New("unset REST API key")
)

type KakaoSDKClient struct {
	restyCli *resty.Client
}

type ClientOptions struct {
	baseURL string
}

type ClientOption func(*ClientOptions)

func WithBaseURL(baseURL string) ClientOption {
	return func(o *ClientOptions) {
		o.baseURL = baseURL
	}
}

func Default(options ...ClientOption) (*KakaoSDKClient, error) {
	o := &ClientOptions{
		baseURL: defaultKakaoApiHost,
	}
	for _, option := range options {
		option(o)
	}

	if globalOptions.restAPIKey == "" {
		return nil, errUnsetRESTAPIKey
	}

	restyCli := resty.New().
		SetAuthScheme(kakaoAuthScheme).
		SetAuthToken(globalOptions.restAPIKey).
		SetBaseURL(o.baseURL)

	return &KakaoSDKClient{restyCli: restyCli}, nil
}

func NewClient(restyCli *resty.Client, options ...ClientOption) (*KakaoSDKClient, error) {
	o := &ClientOptions{
		baseURL: defaultKakaoApiHost,
	}
	for _, option := range options {
		option(o)
	}

	if globalOptions.restAPIKey == "" {
		return nil, errUnsetRESTAPIKey
	}

	restyCli.
		SetAuthScheme(kakaoAuthScheme).
		SetAuthToken(globalOptions.restAPIKey).
		SetBaseURL(o.baseURL)

	return &KakaoSDKClient{restyCli: restyCli}, nil
}
