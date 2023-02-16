package config

import (
	"github.com/stretchr/testify/require"
	"regexp"
	"testing"
)

func TestMailRegexp(t *testing.T) {
	re := regexp.MustCompile(mailRegexp)

	require.True(t, re.MatchString("a@b"))
	require.True(t, re.MatchString("jsquirrel_github_9a6d@packetloss.de"))
	require.True(t, re.MatchString("Super1+Super0-9Unusual324oi9e73289472347@Hihihi.there.com"))
	require.False(t, re.MatchString(" space@start"))
	require.False(t, re.MatchString("space@end "))
	require.False(t, re.MatchString("space@in side"))
	require.False(t, re.MatchString("more space@inside.de"))
}
