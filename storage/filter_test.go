package storage

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestListRequest_Filter(t *testing.T) {
	for i := range paramsReq {
		t.Run(fmt.Sprintf("Filter Policies: case=%s", paramsReq[i].target), func(t *testing.T) {
			l := ListRequest{
				Collection: "filter_test",
				Value:      &polReq,
				FilterFunc: ListByQuery,
			}
			assert.Equal(t, &polRes[i], l.Filter(paramsReq[i].target, 0, 100).Value)
		})

		t.Run(fmt.Sprintf("Filter Roles: case=%s", paramsReq[i].target), func(t *testing.T) {
			l := ListRequest{
				Collection: "filter_test",
				Value:      &rolReq,
				FilterFunc: ListByQuery,
			}
			assert.Equal(t, &rolRes[i], l.Filter(paramsReq[i].target,0 , 100).Value)
		})
	}
}
