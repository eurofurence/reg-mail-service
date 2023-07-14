package config

import (
	"github.com/stretchr/testify/require"
	"regexp"
	"testing"
)

func TestMailRegexp(t *testing.T) {
	re := regexp.MustCompile(mailRegexp)

	require.True(t, re.MatchString("A <a@b>"))
	require.True(t, re.MatchString("jsquirrel_github_9a6d@packetloss.de"))
	require.True(t, re.MatchString("Super1+Super0-9Unusual324oi9e73289472347 <Super1+Super0-9Unusual324oi9e73289472347@Hihihi.there.com>"))
	require.False(t, re.MatchString("Space < space@start>"))
	require.False(t, re.MatchString("Space <space@end >"))
	require.False(t, re.MatchString("Space <space@in side>"))
	require.False(t, re.MatchString("Space <more space@inside.de>"))
	require.True(t, re.MatchString("no_tags_for_name@defined.de"))
	require.False(t, re.MatchString("Empty Tags <>"))
}
