package actions

import "testing"

func TestGenerate(t *testing.T){
	lines := []string{"# Table of contents", "Generate table of contents", "## Inject them to file", "Inject Table of contents to file"}
	result := []string{"Table of contents:\n", "- [Table of contents](#table-of-contents)", "    - [Inject them to file](#inject-them-to-file)"}

	generate := GenerateToc(lines)

	equality := testEq(generate, result)

	if equality == false {
		t.Errorf("Value was incorrect, got: %v, want: %v.", generate, result)

	}
}


func testEq(a, b []string) bool {

    // If one is nil, the other must also be nil.
    if (a == nil) != (b == nil) {
        return false;
    }

    if len(a) != len(b) {
        return false
    }

    for i := range a {
        if a[i] != b[i] {
            return false
        }
    }

    return true
}
