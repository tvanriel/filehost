http {
  address = ":8080"
  debug = false
  rate_limit = 20
}

s3 {
  ssl = false

  access_key = "minio"
  secret_key = "miniosecret"
  endpoint = "localhost:9000"
}

logging {
  development = true
}

files {
  bucket = "public"
  directory = "dir"
}

web {
  s3_public_url_prefix = "http://localhost:9000/public/dir"
}
