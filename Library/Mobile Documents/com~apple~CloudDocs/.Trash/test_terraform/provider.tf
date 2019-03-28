provider "aws" {
  access_key = "jhyoon/Desktop/aws/terraform_key/credentials.csv"
  secret_key = "jhyoon/Desktop/aws/terraform_key/credentials.csv"
  region     = "ap-northeast-1"
}

resource "aws_key_pair" "terraform" {
  key_name   = "terraform"
  public_key = "AKIAJOO3JZQ7ZKB7LILQ"
} 