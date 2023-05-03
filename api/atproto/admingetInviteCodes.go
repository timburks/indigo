// Code generated by cmd/lexgen (see Makefile's lexgen); DO NOT EDIT.

package atproto

// schema: com.atproto.admin.getInviteCodes

import (
	"context"

	"github.com/bluesky-social/indigo/xrpc"
)

// AdminGetInviteCodes_Output is the output of a com.atproto.admin.getInviteCodes call.
type AdminGetInviteCodes_Output struct {
	Codes  []*ServerDefs_InviteCode `json:"codes" cborgen:"codes"`
	Cursor *string                  `json:"cursor,omitempty" cborgen:"cursor,omitempty"`
}

// AdminGetInviteCodes calls the XRPC method "com.atproto.admin.getInviteCodes".
func AdminGetInviteCodes(ctx context.Context, c *xrpc.Client, cursor string, limit int64, sort string) (*AdminGetInviteCodes_Output, error) {
	var out AdminGetInviteCodes_Output

	params := map[string]interface{}{
		"cursor": cursor,
		"limit":  limit,
		"sort":   sort,
	}
	if err := c.Do(ctx, xrpc.Query, "", "com.atproto.admin.getInviteCodes", params, nil, &out); err != nil {
		return nil, err
	}

	return &out, nil
}