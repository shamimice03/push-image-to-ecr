Steps:

1. OIDC -> AWS
2. Secrets ->
3. TEST ECR Connection

- Two IAM Role 
- Two different policy for dev and prod repo
- create two branch (dev and main)
- validation test on connection with aws 
- Two repo (dev and prod)
- use github environment options to push image to ECR

### How GitHub OIDC works with AWS IAM 
- Image Source [Link](https://www.codecentric.de/wissens-hub/blog/secretless-connections-from-github-actions-to-aws-using-oidc)

![image](https://github.com/shamimice03/github-actions-lab/assets/19708705/3cb418a0-20d5-4973-b350-ef5977735217)

