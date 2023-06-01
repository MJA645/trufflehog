	"fmt"
		parser.FromReader(context.Background(), r, commitChan)
		NewParser().FromReader(context.Background(), r, commitChan)
		parser.FromReader(context.Background(), bigReader, commitChan)
func TestMessageParsing(t *testing.T) {
	r := bytes.NewReader([]byte(fmt.Sprintf(singleCommitTabbedContent, `     `)))
	commitChan := make(chan Commit)
	parser := NewParser()
	expectedMessage := singleCommitSingleDiffMessage

	go func() {
		parser.FromReader(context.Background(), r, commitChan)
	}()
	for commit := range commitChan {
		if commit.Message.String() != expectedMessage {
			t.Errorf("Message does not match. Got:\n%s\nexpected:\n%s", commit.Message.String(), expectedMessage)
		}
	}
}

const singleCommitTabbedContent = `commit 70001020fab32b1fcf2f1f0e5c66424eae649826 (HEAD -> master, origin/master, origin/HEAD)
Author: Dustin Decker <humanatcomputer@gmail.com>
Date:   Mon Mar 15 23:27:16 2021 -0700

    Update aws

diff --git a/aws b/aws
index 2ee133b..12b4843 100644
--- a/aws
+++ b/aws
@@ -1,7 +1,5 @@
+[default]
%saws_access_key_id = AKIAXYZDQCEN4B6JSJQI
+aws_secret_access_key = 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
`
