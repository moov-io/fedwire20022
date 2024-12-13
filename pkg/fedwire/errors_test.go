package fedwire_test

import (
	"fmt"
	"testing"

	"github.com/moov-io/fedwire20022/pkg/fedwire"

	"github.com/stretchr/testify/require"
)

func TestErrorCodes(t *testing.T) {
	ec := fedwire.IsError("9910")
	require.NotNil(t, ec)
	require.Equal(t, "9910", ec.Code)
	require.Equal(t, fedwire.ErrTemporary, ec.Level)

	ec = fedwire.IsError("be07")
	require.NotNil(t, ec)
	require.Equal(t, "BE07", ec.Code)
	require.Equal(t, fedwire.ErrFatal, ec.Level)

	ec = fedwire.IsError("ac03")
	require.NotNil(t, ec)
	require.Equal(t, "AC03", ec.Code)
	require.Equal(t, fedwire.ErrAccountFatal, ec.Level)

	ec = fedwire.IsError("AM13")
	require.NotNil(t, ec)
	require.Equal(t, "AM13", ec.Code)
	require.Equal(t, fedwire.ErrTemporary, ec.Level)

	ec = fedwire.IsError("AM09")
	require.NotNil(t, ec)
	require.Equal(t, "AM09", ec.Code)
	require.Equal(t, fedwire.ErrLogic, ec.Level)

	ec = fedwire.IsError("")
	require.Nil(t, ec)

	ec = fedwire.IsError("9999")
	require.Nil(t, ec)
}

func TestErrorLevel(t *testing.T) {
	// just make sure .Error() doesn't recusively panic
	level := fedwire.ErrNetwork
	require.Equal(t, "Network issue", level.Error())
	require.Equal(t, "Network issue", fmt.Sprintf("%v", level))
	require.Equal(t, "Network issue", fmt.Errorf("%w", level).Error())
}
