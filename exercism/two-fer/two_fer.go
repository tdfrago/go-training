/*Package twofer takes in a var name string
if name is not empty:
returns One for "name", one for me.
else:
returns One for you, one for me.
*/

package twofer

import "fmt"

//Sharwith appends a name to the string One for "name", one for me.
func ShareWith(name string) string {
	fmt.Scan(&name)
	if name == "" {
		name = "you"
	}
	output := "One for " + name + ", one for me."
	return output
}
