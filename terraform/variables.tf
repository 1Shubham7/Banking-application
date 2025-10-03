variable "aws_region" {
  description = "AWS region"
  type        = string
  default     = "ap-south-1"
}

variable "subnet_id" {
  description = "Subnet ID for EC2 instances"
  type        = string
  default     = "subnet-0a213b04359448ab4"
}

variable "vpc_id" {
  description = "VPC ID"
  type        = string
  default     = "vpc-00b7aed46eeaad5f9"
}

variable "key_pair_name" {
  description = "EC2 Key Pair name"
  type        = string
  default     = "smyik-keypair"
}

variable "ami_id" {
  description = "AMI ID for EC2 instances"
  type        = string
  default     = "ami-02d26659fd82cf299"
}

variable "instance_type" {
  description = "EC2 instance type"
  type        = string
  default     = "t3.medium"
}
