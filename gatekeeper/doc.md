Learn more about [Contraints](https://github.com/open-policy-agent/frameworks/tree/master/constraint)

Constraints [spec.match](https://open-policy-agent.github.io/gatekeeper/website/docs/howto#the-match-field) spec.parameters field. 
[Input.review](https://open-policy-agent.github.io/gatekeeper/website/docs/howto#input-review) field from rego file
3 types of audit : prom metrics, logs of audit pod<[configure it](https://open-policy-agent.github.io/gatekeeper/website/docs/audit#configuring-audit)>, constraint’s status
Constraints spec.enforcementAction can have value dryRun, warn, deny
Syncing data to the cache is possible via Config. Cached data are accessible in rego file under data.inventory . ex : data.inventory.namespace[<namespace>][groupVersion][<kind>][<name>]

Exempting ns using spec.match[].excludeNamespace [field in Config](https://open-policy-agent.github.io/gatekeeper/website/docs/exempt-namespaces#exempting-namespaces-from-gatekeeper-using-config-resource). Or [by flag](https://open-policy-agent.github.io/gatekeeper/website/docs/exempt-namespaces#exempting-namespaces-from-the-gatekeeper-admission-webhook-using---exempt-namespace-flag)
Customizing [startup](https://open-policy-agent.github.io/gatekeeper/website/docs/customize-startup) & [admission behavior](https://open-policy-agent.github.io/gatekeeper/website/docs/customize-startup)

More features: Metrics, Debugging, recovery, cloud, failing



[Mutation CRDs](https://github.com/open-policy-agent/gatekeeper/tree/master/apis/mutations/v1)
Assign
	applyTo
	Match
	Location string
	AssignParameters

AssignMetadata
	Match
	Location string
	AssignMetadataParameters

ModifySet
	applyTo
	Match
	Location string
	ModifysetParameters

AssignParameters
	pathTests []pathtest
	Assign assignField
AssignMetadataParameters
	pathTests []pathtest
AssignMetadataParameters
	pathTests []pathtest
	Operation -> merge/prune
Values struct {FromList []interface{} `json:"fromList,omitempty"`}
Match & applyTo  are described [here](https://github.com/open-policy-agent/gatekeeper/tree/master/pkg/mutation/match)

