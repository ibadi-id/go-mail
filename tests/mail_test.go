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
	"github.com/ainsleyclark/go-mail"
	"github.com/stretchr/testify/suite"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

// MailTestSuite defines the helper used for mail
// testing.
type MailTestSuite struct {
	suite.Suite
}

// TestMail
//
// Assert testing has begun.
func TestMail(t *testing.T) {
	suite.Run(t, new(MailTestSuite))
}

const (
	// DataPath defines where the test data resides.
	DataPath = "testdata"
	// PNGName defines the PNG name for testing.
	PNGName = "gopher.png"
)

// GetTransmission
//
// Returns a dummy transition for testing with an
// attachment.
func (t *MailTestSuite) GetTransmission() *mail.Transmission {
	wd, err := os.Getwd()
	t.NoError(err)
	path := filepath.Dir(wd) + string(os.PathSeparator) + DataPath + string(os.PathSeparator) + PNGName

	file, err := ioutil.ReadFile(path)
	if err != nil {
		t.Fail("error getting attachment with the path: "+path, err)
	}

	return &mail.Transmission{
		Recipients: []string{"ainsley@reddico.co.uk"},
		Subject:    "Test - Go Mail",
		HTML:       "<h1>Hello from Go Mail!</h1>",
		PlainText:  "",
		Attachments: mail.Attachments{
			mail.Attachment{
				Filename: "gopher.png",
				Bytes:    file,
			},
		},
	}
}