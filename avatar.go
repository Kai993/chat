package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"path/filepath"
)

// GravatarURL : Gravatar URL
const GravatarURL = "//www.gravatar.com/avatar"

// ErrNoAvatarURL : AvatarインスタンスがアバターのURLを返すことができない場合に発生するエラー
var ErrNoAvatarURL = errors.New("chat: アバターのURLを取得できません。")

type (
	// Avatar : ユーザーのプロフィール画像を表す型
	Avatar interface {
		GetAvatarURL(ChatUser) (string, error)
	}

	// AuthAvatar : OAuth認証で取得したアバター
	AuthAvatar struct{}

	// GravatarAvatar : Gravatarで取得したアバター
	GravatarAvatar struct{}

	// FileSystemAvatar : アップロードしたアバター
	FileSystemAvatar struct{}

	// TryAvatars : すべての実装を切り替えながらURLを取得する
	TryAvatars []Avatar
)

// UseAuthAvatar : Oauth認証で取得したアバターを使用する
var UseAuthAvatar AuthAvatar

// UseGravatar : Gravatarで取得したアバターを使用する
var UseGravatar GravatarAvatar

// UseFileSystemAvatar : アップロードしたアバターを使用する
var UseFileSystemAvatar FileSystemAvatar

// GetAvatarURL : アバターURLを取得する
func (AuthAvatar) GetAvatarURL(u ChatUser) (string, error) {
	url := u.AvatarURL()
	if url != "" {
		return url, nil
	}
	return "", ErrNoAvatarURL
}

// GetAvatarURL : GravatarURLを取得する
func (GravatarAvatar) GetAvatarURL(u ChatUser) (string, error) {
	return fmt.Sprintf("%s/%s", GravatarURL, u.UniqueID()), nil
}

// GetAvatarURL : ./avatars/から画像のURLを取得する
func (FileSystemAvatar) GetAvatarURL(u ChatUser) (string, error) {
	if files, err := ioutil.ReadDir("avatars"); err == nil {
		for _, file := range files {
			if file.IsDir() {
				continue
			}

			if match, _ := filepath.Match(u.UniqueID()+"*", file.Name()); match {
				return fmt.Sprintf("/avatars/%s", file.Name()), nil
			}
		}
	}
	return "", ErrNoAvatarURL
}

func (a TryAvatars) GetAvatarURL(u ChatUser) (string, error) {
	for _, avatar := range a {
		if url, err := avatar.GetAvatarURL(u); err == nil {
			return url, nil
		}
	}
	return "", ErrNoAvatarURL
}
