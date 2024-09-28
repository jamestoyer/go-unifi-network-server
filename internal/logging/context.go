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

package logging

import (
	"context"
	"log/slog"
)

type contextKey string

const (
	slogAttrs contextKey = "network_server_slog_attrs"
)

type Handler struct {
	slog.Handler
}

// Handle adds any slog attributes in the context to the slog.Record before calling the underlying handler.
func (h Handler) Handle(ctx context.Context, r slog.Record) error {
	if attrs, ok := ctx.Value(slogAttrs).([]slog.Attr); ok {
		for _, v := range attrs {
			r.AddAttrs(v)
		}
	}

	return h.Handler.Handle(ctx, r)
}

// CtxWithValues adds slog attributes to the given context so that they can be added to a slog.Record when an
// appropriate slog.Handler is set.
func CtxWithValues(parent context.Context, attrs ...slog.Attr) context.Context {
	var allAttrs []slog.Attr
	if v, ok := parent.Value(slogAttrs).([]slog.Attr); ok {
		allAttrs = v
	}

	allAttrs = append(allAttrs, attrs...)
	return context.WithValue(parent, slogAttrs, allAttrs)
}
