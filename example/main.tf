provider "custom" {
  name               = "custom files"
  location_files = "C:\\Users\\juan\\dev\\provider-test-files\\"
}

resource "custom_jjoc_file" "file1" {
  name = "file1"
  content   = "hola este es el contenido de file 1"
}


resource "custom_jjoc_file" "file3" {
  name = "file3"
  content   = "hola este es el contenido de file 3"
}