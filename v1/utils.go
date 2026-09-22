package v1

type Tags map[string]string

// CIRunIDLabel labels a CI run's instances and their sub-resources with a per-run value so a
// post-run sweep can delete them. Production never sets it.
const CIRunIDLabel = "ci-run-id"
