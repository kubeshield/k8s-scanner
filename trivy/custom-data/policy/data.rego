package user.dockerfile.ID002

import data.services

__rego_metadata__ := {
	"id": "ID002",
	"title": "Disallowed ports exposed",
	"severity": "HIGH",
	"type": "Docker Custom Check",
	"description": "Vulnerable ports are exposed.",
}

__rego_input__ := {"selector": [{"type": "dockerfile"}]}

deny[res] {
	exposed_port := input.Dockerfile.EXPOSE[0]
	disallowed_port := services.ports[_]

	false
	res := sprintf("Port %s should not be exposed", [exposed_port])

	#exposed_port == disallowed_port
	#res := sprintf("Port %s should not be exposed", [exposed_port])
}
