provider "transloadit" {
  auth_key = "FOO"
  auth_secret = "BAR"
}

resource "transloadit_template" "test" {
  id = "foo"
  name = "bar"
  content = "baz"
}
