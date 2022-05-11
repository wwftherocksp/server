// Copyright (c) 2022 Sociosarbis <136657577@qq.com>
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

package handler

import (
	"errors"
	"reflect"
	"strings"

	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	"github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zh_translations "github.com/go-playground/validator/v10/translations/zh" //nolint:importas
	"go.uber.org/zap"

	"github.com/bangumi/server/cache"
	"github.com/bangumi/server/config"
	"github.com/bangumi/server/domain"
	"github.com/bangumi/server/internal/errgo"
	"github.com/bangumi/server/web/captcha"
	"github.com/bangumi/server/web/rate"
	"github.com/bangumi/server/web/session"
)

var errTranslationNotFound = errors.New("failed to find translation for zh")

func New(
	cfg config.AppConfig,
	s domain.SubjectService,
	c domain.CharacterService,
	p domain.PersonService,
	a domain.AuthService,
	e domain.EpisodeRepo,
	r domain.RevisionRepo,
	index domain.IndexRepo,
	user domain.UserRepo,
	cache cache.Generic,
	captcha captcha.Manager,
	session session.Manager,
	rateLimit rate.Manager,
	log *zap.Logger,
) (Handler, error) {

	validate, trans, err := getValidator()
	if err != nil {
		return Handler{}, err
	}

	return Handler{
		cfg:                  cfg,
		cache:                cache,
		log:                  log,
		rateLimit:            rateLimit,
		session:              session,
		p:                    p,
		s:                    s,
		a:                    a,
		u:                    user,
		e:                    e,
		c:                    c,
		i:                    index,
		r:                    r,
		captcha:              captcha,
		v:                    validate,
		validatorTranslation: trans,
	}, nil
}

type Handler struct {
	rateLimit            rate.Manager
	s                    domain.SubjectService
	p                    domain.PersonService
	a                    domain.AuthService
	session              session.Manager
	captcha              captcha.Manager
	e                    domain.EpisodeRepo
	c                    domain.CharacterService
	u                    domain.UserRepo
	cache                cache.Generic
	i                    domain.IndexRepo
	r                    domain.RevisionRepo
	validatorTranslation ut.Translator
	log                  *zap.Logger
	v                    *validator.Validate
	cfg                  config.AppConfig
}

func getValidator() (*validator.Validate, ut.Translator, error) {
	validate := validator.New()
	uni := ut.New(en.New(), zh.New())

	// this is usually know or extracted from http 'Accept-Language' header
	// also see uni.FindTranslator(...)
	trans, found := uni.GetTranslator(zh.New().Locale())
	if !found {
		return nil, nil, errTranslationNotFound
	}

	err := zh_translations.RegisterDefaultTranslations(validate, trans)
	if err != nil {
		return nil, nil, errgo.Wrap(err, "failed to register translation")
	}

	validate.RegisterTagNameFunc(func(field reflect.StructField) string {
		tag := field.Tag.Get("json")
		if tag == "" {
			return field.Name
		}

		name := strings.SplitN(tag, ",", 2)[0]
		if name == "-" {
			return ""
		}

		return name
	})

	return validate, trans, nil
}
