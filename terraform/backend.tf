terraform {
  backend "s3" {
    bucket         = "banking-app-terraform-state"
    key            = "banking-app/terraform.tfstate"
    region         = "ap-south-1"
    encrypt        = true
  }
}
