package main

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/mariomang/hitokoto-go"
	"github.com/mariomang/hitokoto-go/constants"
	"github.com/mariomang/hitokoto-go/op"
)

func main() {

	github := os.Getenv("GITHUB_TOKEN")
	if github == "" {
		panic(errors.New("the environment variable GITHUB_TOKEN not found"))
	}

	executor := hitokoto.NewExecutor()

	req := &op.HitokotoRequest{
		Type: []constants.HitokotoType{
			constants.Animation,
			constants.Cartoon,
			constants.Game,
			constants.Internet,
			constants.Literature,
			constants.MoviesAndTv,
			constants.Netease,
			constants.Original,
			constants.Other,
			constants.Poetry,
			constants.Philosophy,
			constants.PettyTrick,
		},
		Encode:    constants.EncodeJson,
		Charset:   constants.CharsetUTF8,
		Callback:  "",
		Select:    "",
		MinLength: 0,
		MaxLength: 0,
	}

	resp := &op.HitokotoResponse{}
	if err := executor.Do(&constants.APIHitokoto, req, resp); err != nil {
		panic(err)
	}

	fmt.Printf("================================\n")
	fmt.Printf("Hitokoto: \n")
	fmt.Printf("-- ID: %d\n", resp.ID)
	fmt.Printf("-- Hitokoto: %s\n", resp.Hitokoto)
	fmt.Printf("-- Type: %s\n", resp.Type)
	fmt.Printf("-- From: %s\n", resp.From)
	fmt.Printf("-- FromWho: %s\n", resp.FromWho)
	fmt.Printf("-- Creator: %s\n", resp.Creator)
	fmt.Printf("-- CreatorUID: %d\n", resp.CreatorUID)
	fmt.Printf("-- Reviewer: %d\n", resp.Reviewer)
	fmt.Printf("-- UUID: %s\n", resp.UUID)
	fmt.Printf("-- CommitFrom: %s\n", resp.CommitFrom)
	fmt.Printf("-- Length: %d\n", resp.Length)
	fmt.Printf("-- CreatedAt: %s\n", resp.CreatedAt.Format(time.RFC3339))

	if err := UpdateGithubUserBio(github, resp.Hitokoto); err != nil {
		panic(err)
	}
}

func UpdateGithubUserBio(token string, bio string) error {

	url := "https://api.github.com/user"
	payload := []byte(`{"bio":"` + bio + `"}`)

	req, err := http.NewRequest(http.MethodPatch, url, bytes.NewReader(payload))
	if err != nil {
		return err
	}

	req.Header.Set("Accept", "application/vnd.github+json")
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("X-GitHub-Api-Version", "2022-11-28")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	fmt.Println(string(b))

	return nil
}
