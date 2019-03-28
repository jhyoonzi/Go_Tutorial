provider "aws" {
  access_key = "<AWS_ACCESS_KEY>"
  secret_key = "<AWS_SECRET_KEY>"
  region = "ap-northeast-2"
}

variable "ami" {
  description = "Amazone Linux AMI"
  default = "ami-02bcbb802e03574ba"
}

variable "Instance_type" {
  description = "instance type"
  default = "t2.micro"
}

variable "key_path" {
  description = "SSH public key path"
  default = "/home/core/.ssh/id_rsa.pub"
}

variable "bootstrap_path" {
  description = "Script to install docker engine"
  default = "install-docker.sh"
}




#test 