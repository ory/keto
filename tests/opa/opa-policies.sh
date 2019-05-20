#!/usr/bin/env bash
cd ../..
(cd engine/ladon/rego/core; curl -X PUT --data-binary @effect.rego localhost:8181/v1/policies/core/effect)
(cd engine/ladon/rego/core; curl -X PUT --data-binary @role.rego localhost:8181/v1/policies/core/role)
(cd engine/ladon/rego/condition; curl -X PUT --data-binary @boolean.rego localhost:8181/v1/policies/condition/boolean)
(cd engine/ladon/rego/condition; curl -X PUT --data-binary @cidr.rego localhost:8181/v1/policies/condition/cidr)
(cd engine/ladon/rego/condition; curl -X PUT --data-binary @condition.rego localhost:8181/v1/policies/condition/condition)
(cd engine/ladon/rego/condition; curl -X PUT --data-binary @helpers.rego localhost:8181/v1/policies/condition/helpers)
(cd engine/ladon/rego/condition; curl -X PUT --data-binary @resource_contains.rego localhost:8181/v1/policies/condition/resource_contains)
(cd engine/ladon/rego/condition; curl -X PUT --data-binary @string_equal.rego localhost:8181/v1/policies/condition/string_equal)
(cd engine/ladon/rego/condition; curl -X PUT --data-binary @string_match.rego localhost:8181/v1/policies/condition/string_match)
(cd engine/ladon/rego/condition; curl -X PUT --data-binary @string_pairs_equal.rego localhost:8181/v1/policies/condition/string_pairs_equal)
(cd engine/ladon/rego/condition; curl -X PUT --data-binary @string_subject_equal.rego localhost:8181/v1/policies/condition/string_subject_equal)

(cd engine/ladon/rego/exact; curl -X PUT --data-binary @main.rego localhost:8181/v1/policies/exact/main)
(cd engine/ladon/rego/regex; curl -X PUT --data-binary @main.rego localhost:8181/v1/policies/regex/main)
(cd engine/ladon/rego/glob; curl -X PUT --data-binary @main.rego localhost:8181/v1/policies/glob/main)
