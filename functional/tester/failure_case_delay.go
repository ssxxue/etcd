// Copyright 2018 The etcd Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package tester

import (
	"time"

	"go.uber.org/zap"
)

type failureDelay struct {
	Failure
	delayDuration time.Duration
}

func (f *failureDelay) Inject(clus *Cluster) error {
	if err := f.Failure.Inject(clus); err != nil {
		return err
	}
	if f.delayDuration > 0 {
		clus.lg.Info(
			"wait after inject",
			zap.Duration("delay", f.delayDuration),
			zap.String("desc", f.Failure.Desc()),
		)
		time.Sleep(f.delayDuration)
	}
	return nil
}
