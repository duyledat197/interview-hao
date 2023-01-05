{{define "delete"}}Feature: delete {{.CamelCase}}

    Background:
        Given a signed in "admin"
        And a background

	# authenticate
    Scenario Outline: authenticate when delete {{.CamelCase}}
        Given a signed in "<role>"
        When user delete {{.CamelCase}}
        Then returns "<status code>" status code

        Examples:
            | role           | status code |
            | admin          |             |

	# delete {{.CamelCase}}
    Scenario: delete {{.CamelCase}}
        When user delete {{.CamelCase}}
        Then <%[2]s>returns "OK" status code
        And {{.CamelCase}} have been deleted correctly

	# delete {{.CamelCase}} again
    Scenario: delete {{.CamelCase}} again
        Given user delete {{.CamelCase}}
        When user delete {{.CamelCase}} again
        Then <%[2]s>returns "NotFound" status code
{{end}}