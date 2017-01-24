backend "etcd" {
  address = "http://localhost:2379"
  path = "vault"
}

listener "tcp" {
 address = "127.0.0.1:8200"
 tls_disable = 1
}
