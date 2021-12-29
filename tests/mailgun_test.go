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
	"fmt"
	"github.com/ainsleyclark/go-mail/drivers"
	"github.com/ainsleyclark/go-mail/mail"
	"os"
)

func (t *MailTestSuite) Test_MailGun() {
	tx := t.GetTransmission()

	mailer, err := drivers.NewMailgun(mail.Config{
		URL:         os.Getenv("MAILGUN_URL"),
		APIKey:      os.Getenv("MAILGUN_API_KEY"),
		FromAddress: os.Getenv("MAILGUN_FROM_ADDRESS"),
		FromName:    os.Getenv("MAILGUN_FROM_NAME"),
		Domain:      os.Getenv("MAILGUN_DOMAIN"),
	})
	if err != nil {
		t.FailNow("Error creating client: " + err.Error())
	}

	result, err := mailer.Send(tx)
	if err != nil {
		t.FailNow("Error sending MailGun email: " + err.Error())
	}

	fmt.Printf("%+v\n", result)
	t.Equal(200, result.StatusCode)
	t.NotNil(result.Body)
	t.NotEmpty(result.Message)
	t.NotEmpty(result.ID)
}
