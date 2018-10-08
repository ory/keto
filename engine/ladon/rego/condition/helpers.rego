package ory.condition

cast_string_empty(r, key) = value {
  not r[key]
  value := ""
}{
  cast_string(r[key], value)
}
