terraform {
  required_version = ">= 1.0"
  
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 5.0"
    }
  }
}

provider "aws" {
  region = var.aws_region
}

# Data source to get default VPC
data "aws_vpc" "default" {
  default = true
}

# Data source to get default subnet
data "aws_subnet" "default" {
  id = var.subnet_id
}

# Security Group for K8s cluster
resource "aws_security_group" "k8s_sg" {
  name        = "smyik-security-group"
  description = "Security group for my smiyk. this uses the default VPC"
  vpc_id      = var.vpc_id

  # SSH access from your IP
  ingress {
    from_port   = 22
    to_port     = 22
    protocol    = "tcp"
    cidr_blocks = ["49.36.217.103/32"]
  }

  # HTTP access
  ingress {
    description = "Public HTTP access"
    from_port   = 80
    to_port     = 80
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }

  # HTTPS access
  ingress {
    from_port   = 443
    to_port     = 443
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }

  # Frontend access
  ingress {
    description = "Frontend access"
    from_port   = 3000
    to_port     = 3000
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }

  # Backend access
  ingress {
    description = "Backend access"
    from_port   = 8080
    to_port     = 8080
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }

  # K8s API server
  ingress {
    description = "K8s API server"
    from_port   = 6443
    to_port     = 6443
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }

  # etcd server client API
  ingress {
    description = "etcd server client API"
    from_port   = 2379
    to_port     = 2380
    protocol    = "tcp"
    cidr_blocks = ["172.31.0.0/16"]
  }

  # Kubelet API
  ingress {
    description = "Kubelet API"
    from_port   = 10250
    to_port     = 10250
    protocol    = "tcp"
    cidr_blocks = ["172.31.0.0/16"]
  }

  # kube-controller-manager
  ingress {
    description = "kube-controller-manager"
    from_port   = 10257
    to_port     = 10257
    protocol    = "tcp"
    cidr_blocks = ["172.31.0.0/16"]
  }

  # kube-scheduler
  ingress {
    description = "kube-scheduler"
    from_port   = 10259
    to_port     = 10259
    protocol    = "tcp"
    cidr_blocks = ["172.31.0.0/16"]
  }

  # NodePort services TCP
  ingress {
    description = "NodePort services"
    from_port   = 30000
    to_port     = 32767
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }

  # NodePort services UDP
  ingress {
    description = "NodePort services UDP"
    from_port   = 30000
    to_port     = 32767
    protocol    = "udp"
    cidr_blocks = ["0.0.0.0/0"]
  }

  # Pod-to-pod communication - separate rules to match AWS state
  ingress {
    description = "Pod-to-pod communication"
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["10.0.0.0/8"]
  }

  ingress {
    description = "Internal cluster communication"
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["172.31.0.0/16"]
  }

  # Self-referencing rule for pod to pod communication
  ingress {
    description = "pod to pod communication"
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    self        = true
  }

  # Egress rule - allow all outbound traffic
  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }
}

# ECR Repository
resource "aws_ecr_repository" "smyik" {
  name                 = "smyik"
  image_tag_mutability = "MUTABLE"

  image_scanning_configuration {
    scan_on_push = false
  }
}

# K8s Master Node
resource "aws_instance" "k8s_master" {
  ami                    = var.ami_id
  instance_type          = var.instance_type
  key_name              = var.key_pair_name
  subnet_id             = var.subnet_id
  vpc_security_group_ids = [aws_security_group.k8s_sg.id]
  source_dest_check      = false  # Required for Kubernetes networking

  tags = {
    Name = "k8s-master"
  }

  lifecycle {
    prevent_destroy = false
  }
}

# K8s Worker Node
resource "aws_instance" "k8s_worker" {
  ami                    = var.ami_id
  instance_type          = var.instance_type
  key_name              = var.key_pair_name
  subnet_id             = var.subnet_id
  vpc_security_group_ids = [aws_security_group.k8s_sg.id]
  source_dest_check      = false  # Required for Kubernetes networking

  tags = {
    Name = "k8s-worker"
  }

  lifecycle {
    prevent_destroy = false
  }
}
