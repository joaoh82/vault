backend "etcd" {
  address = "http://localhost:2379"
  path = "vault"
}

listener "tcp" {
 tls_disable = 1
}
