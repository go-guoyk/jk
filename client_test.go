package jk

import "os"

func ClientFromTestENV() *Client {
	return New(ClientOptions{
		URL:      os.Getenv("JK_TEST_URL"),
		Username: os.Getenv("JK_TEST_USERNAME"),
		Password: os.Getenv("JK_TEST_PASSWORD"),
	})
}
