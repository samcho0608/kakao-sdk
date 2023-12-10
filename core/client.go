package core

import (
	"errors"
	"github.com/go-resty/resty/v2"
	"strings"
)

var (
	errUnsetRESTAPIKey = errors.New("unset REST API key")
)

type KakaoSDKClient struct {
	// TODO(@samcho0608):: 기본적으로 resty의 Client를 사용하되, 차차 다양한 Client들을 지원할 수 있도록 interface로 변경해야합니다.
	*resty.Client
}

type ClientOptions struct {
	baseURL    string
	restAPIKey string
}

type ClientOption func(*ClientOptions)

func DefaultOptions() *ClientOptions {
	return &ClientOptions{
		baseURL:    defaultKakaoApiHost,
		restAPIKey: globalOptions.restAPIKey,
	}
}

func WithBaseURL(baseURL string) ClientOption {
	return func(o *ClientOptions) {
		o.baseURL = baseURL
	}
}

func WithAPICollectionPath(path string) ClientOption {
	return func(o *ClientOptions) {
		if strings.HasPrefix(path, "/") {
			o.baseURL += path
			return
		}

		o.baseURL += "/" + path
	}
}

func Default(options ...ClientOption) (*KakaoSDKClient, error) {
	o := DefaultOptions()
	for _, option := range options {
		option(o)
	}

	if err := validateOptions(o); err != nil {
		return nil, err
	}

	client := resty.New().
		SetAuthScheme(kakaoAuthScheme).
		SetAuthToken(o.restAPIKey).
		SetBaseURL(o.baseURL)

	return &KakaoSDKClient{Client: client}, nil
}

func NewClient(client *resty.Client, options ...ClientOption) (*KakaoSDKClient, error) {
	o := &ClientOptions{
		baseURL: defaultKakaoApiHost,
	}
	for _, option := range options {
		option(o)
	}

	if err := validateOptions(o); err != nil {
		return nil, err
	}

	client.
		SetAuthScheme(kakaoAuthScheme).
		SetAuthToken(globalOptions.restAPIKey).
		SetBaseURL(o.baseURL)

	return &KakaoSDKClient{Client: client}, nil
}

func validateOptions(options *ClientOptions) error {
	if options.restAPIKey == "" {
		return errUnsetRESTAPIKey
	}

	if options.baseURL == "" {
		return errors.New("empty base URL")
	}

	return nil
}
