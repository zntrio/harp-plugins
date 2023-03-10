// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

package container

import (
	"context"
	"fmt"
	"io"
	"net/http"
)

type httpLoader struct {
	scheme string
	host   string
}

// Reader returns the file Reader
func (d *httpLoader) Reader(ctx context.Context, key string) (io.ReadCloser, error) {
	// Prepare query
	q, err := http.NewRequestWithContext(ctx, http.MethodGet, fmt.Sprintf("%s://%s/%s", d.scheme, d.host, key), http.NoBody)
	if err != nil {
		return nil, fmt.Errorf("http: unable to prepare http query: %w", err)
	}

	// Query
	resp, err := http.DefaultClient.Do(q)
	if err != nil {
		return nil, fmt.Errorf("http: unable to query remote bundle URL: %w", err)
	}

	// No error
	return resp.Body, nil
}
