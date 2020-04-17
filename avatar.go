package main

import (
	"crypto/md5"
	"errors"
	"fmt"
	"io"
	"strings"
)

// GravatarURL : Gravatar URL
const GravatarURL = "//www.gravatar.com/avatar"

// ErrNoAvatarURL : AvatarインスタンスがアバターのURLを返すことができない場合に発生するエラー
var ErrNoAvatarURL = errors.New("chat: アバターのURLを取得できません。")

type (
	// Avatar : ユーザーのプロフィール画像を表す型
	Avatar interface {
		GetAvatarURL(c *client) (string, error)
	}

	// AuthAvatar :
	AuthAvatar struct{}

	// GravatarAvatar :
	GravatarAvatar struct{}
)

// UseAuthAvatar : アバターを使用する
var UseAuthAvatar AuthAvatar

// UseGravatar : Gravatarを使用する
var UseGravatar GravatarAvatar

// GetAvatarURL : アバターURLを取得する
func (AuthAvatar) GetAvatarURL(c *client) (string, error) {
	if url, ok := c.userData["avatar_url"]; ok {
		if urlStr, ok := url.(string); ok {
			return urlStr, nil
		}
	}

	return "", ErrNoAvatarURL
}

func (GravatarAvatar) GetAvatarURL(c *client) (string, error) {
	if email, ok := c.userData["email"]; ok {
		if emailStr, ok := email.(string); ok {
			m := md5.New()
			io.WriteString(m, strings.ToLower(emailStr))
			return fmt.Sprintf("%s/%x", GravatarURL, m.Sum(nil)), nil
		}
	}

	return "", ErrNoAvatarURL
}
