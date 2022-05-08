// Copyright (c) 2022 Trim21 <trim21.me@gmail.com>
//
// SPDX-License-Identifier: AGPL-3.0-only
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published
// by the Free Software Foundation, version 3.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.
// See the GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program. If not, see <https://www.gnu.org/licenses/>

package domain

import (
	"context"
	"time"

	"github.com/bangumi/server/model"
)

// AuthRepo presents an authorization.
type AuthRepo interface {
	// GetByToken return an authorized user by a valid access token.
	GetByToken(ctx context.Context, token string) (Auth, error)
	GetPermission(ctx context.Context, groupID uint8) (Permission, error)

	// GetByEmail return (Auth, HashedPassword, error)
	GetByEmail(ctx context.Context, email string) (Auth, []byte, error)
}

type GroupID = uint8

// Auth is the basic authorization represent a user.
type Auth struct {
	RegTime    time.Time
	ID         model.UIDType // user id
	GroupID    GroupID
	Permission Permission `json:"-"` // disable cache for this field.
}

const nsfwThreshold = -time.Hour * 24 * 60

// AllowNSFW return if current user is allowed to see NSFW resource.
func (u Auth) AllowNSFW() bool {
	if u.ID == 0 {
		return false
	}

	return u.RegTime.Add(nsfwThreshold).Before(time.Now())
}

type AuthService interface {
	GetByToken(ctx context.Context, token string) (Auth, error)
	ComparePassword(hashed []byte, password string) (bool, error)
	GetByTokenWithCache(ctx context.Context, token string) (Auth, error)
	Login(ctx context.Context, email, password string) (Auth, bool, error)
	GetPermission(ctx context.Context, id GroupID) (Permission, error)
}

type Permission struct {
	UserList           bool
	ManageUserGroup    bool
	ManageUserPhoto    bool
	ManageTopicState   bool
	ManageReport       bool
	UserBan            bool
	ManageUser         bool
	UserGroup          bool
	UserWikiApply      bool `doc:"申请 wiki 人"`
	UserWikiApprove    bool
	DoujinSubjectErase bool
	DoujinSubjectLock  bool
	SubjectEdit        bool
	SubjectLock        bool
	SubjectRefresh     bool
	SubjectRelated     bool
	SubjectMerge       bool
	SubjectErase       bool
	SubjectCoverLock   bool
	SubjectCoverErase  bool
	MonoEdit           bool
	MonoLock           bool
	MonoMerge          bool
	MonoErase          bool
	EpEdit             bool
	EpMove             bool
	EpMerge            bool
	EpLock             bool
	EpErase            bool
	Report             bool
	ManageApp          bool
	AppErase           bool
}
