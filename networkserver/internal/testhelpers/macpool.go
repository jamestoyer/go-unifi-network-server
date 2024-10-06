// Copyright 2024 James Toyer
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

package testhelpers

import (
	"net"
	"sync"
)

type MACPool struct {
	lock sync.Mutex
	pool []*net.HardwareAddr
}

func NewMACPool() *MACPool {
	pool := &MACPool{}

	// As per https://datatracker.ietf.org/doc/html/rfc7042#section-2.1.1, 00-00-5E-00-53-XX has been reserved for
	// documentation. We're going to use that here as a bit of a hack as there appears to be no specific "testing" block.
	for i := 0; i < 256; i++ {
		mac := net.HardwareAddr{0x00, 0x00, 0x5e, 0x00, 0x53, byte(i)}
		pool.pool = append(pool.pool, &mac)
	}

	return pool
}

func (p *MACPool) MAC() (*net.HardwareAddr, func()) {
	p.lock.Lock()
	defer p.lock.Unlock()

	mac := p.pool[0]
	p.pool = p.pool[1:]

	release := func() {
		p.lock.Lock()
		defer p.lock.Unlock()
		p.pool = append(p.pool, mac)
	}

	return mac, release
}
