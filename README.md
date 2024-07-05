# fssou-api


## Infrastructure as Code

### How to validate the Terraform code

Initialize the Terraform configuration
```bash
terraform init -backend-config="bucket=iac.francl.in" -backend-config="key=terraform/state/702404391" -backend-config="region=us-east-1"
```

Plan the Terraform configuration
```bash
terraform plan
```

Apply the Terraform configuration to create the resources
```bash
terraform apply
```
