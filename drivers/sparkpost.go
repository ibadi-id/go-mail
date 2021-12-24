// Copyright 2020 The Go Mail Authors. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package mail

import (
	sp "github.com/SparkPost/gosparkpost"
	"github.com/ainsleyclark/go-mail"
	"strings"
)

// sparkPost represents the data for sending mail via the
// SparkPost API. Configuration, the client and the
// main send function are parsed for sending
// data.
type sparkPost struct {
	cfg    mail.Config
	client sp.Client
	send   sparkSendFunc
}

// sparkSendFunc defines the function for ending SparkPost
// transmissions.
type sparkSendFunc func(t *sp.Transmission) (id string, res *sp.Response, err error)

const (
	// SparkAPIVersion defines the default API version for
	// SparkPost.
	SparkAPIVersion = 1
)

// Creates a new SparkPost client. Configuration is
// validated before initialisation.
func newSparkPost(cfg mail.Config) (*sparkPost, error) {
	err := cfg.Validate()
	if err != nil {
		return nil, err
	}

	config := &sp.Config{
		BaseUrl:    cfg.URL,
		ApiKey:     cfg.APIKey,
		ApiVersion: SparkAPIVersion,
		Verbose:    true,
	}

	var client sp.Client
	err = client.Init(config)
	if err != nil {
		return nil, err
	}

	return &sparkPost{
		cfg:    cfg,
		client: client,
		send:   client.Send,
	}, nil
}

// Send posts the go mail Transmission to the SparkPost
// API. Transmissions are validated before sending
// and attachments are added. Returns an error
// upon failure.
func (s *sparkPost) Send(t *mail.Transmission) (mail.Response, error) {
	err := t.Validate()
	if err != nil {
		return mail.Response{}, err
	}

	headerTo := strings.Join(t.Recipients, ",")

	content := sp.Content{
		HTML: t.HTML,
		Text: t.PlainText,
		From: sp.From{
			Email: s.cfg.FromAddress,
			Name:  s.cfg.FromName,
		},
		Subject: t.Subject,
		Headers: make(map[string]string),
	}

	tx := &sp.Transmission{
		Recipients: []sp.Recipient{},
	}

	for _, r := range t.Recipients {
		tx.Recipients = append(tx.Recipients.([]sp.Recipient), sp.Recipient{
			Address: sp.Address{Email: r, HeaderTo: headerTo},
		})
	}

	if t.HasCC() {
		for _, c := range t.CC {
			tx.Recipients = append(tx.Recipients.([]sp.Recipient), sp.Recipient{
				Address: sp.Address{Email: c, HeaderTo: headerTo},
			})
			content.Headers["cc"] = strings.Join(t.CC, ",")
		}
	}

	if t.HasBCC() {
		for _, b := range t.BCC {
			tx.Recipients = append(tx.Recipients.([]sp.Recipient), sp.Recipient{
				Address: sp.Address{Email: b, HeaderTo: headerTo},
			})
		}
	}

	if t.Attachments.Exists() {
		content.Attachments = s.addAttachments(t.Attachments)
	}

	tx.Content = content

	id, response, err := s.send(tx)
	if err != nil {
		return mail.Response{}, err
	}

	if len(response.Errors) > 0 {
		return mail.Response{}, response.Errors
	}

	return mail.Response{
		StatusCode: response.HTTP.StatusCode,
		Body:       string(response.Body),
		Headers:    response.HTTP.Header,
		ID:         id,
		Message:    response.Verbose,
	}, nil
}

// addAttachments transforms a go mail attachments to
// SparkPost attachments.
func (s *sparkPost) addAttachments(a mail.Attachments) []sp.Attachment {
	var att []sp.Attachment
	for _, v := range a {
		att = append(att, sp.Attachment{
			MIMEType: v.Mime(),
			Filename: v.Filename,
			B64Data:  v.B64(),
		})
	}
	return att
}
