Steps:

1. OIDC -> AWS
2. Secrets ->
3. TEST ECR Connection

- Two IAM Role 
- Two different policy for dev and prod repo 
- create two branch (dev and main)
- validation test on connection with aws 
- Two repo (dev and prod) *
- use github environment options to push image to ECR

### How GitHub OIDC works with AWS IAM 
- Image Source [Link](https://www.codecentric.de/wissens-hub/blog/secretless-connections-from-github-actions-to-aws-using-oidc)

![image](https://github.com/shamimice03/github-actions-lab/assets/19708705/3cb418a0-20d5-4973-b350-ef5977735217)

> GitHubActionsRoleToPushOnECR-cloudterms-dev


> Trust:
```
{
	"Version": "2012-10-17",
	"Statement": [
		{
			"Effect": "Allow",
			"Principal": {
				"Federated": "arn:aws:iam::391178969547:oidc-provider/token.actions.githubusercontent.com"
			},
			"Action": "sts:AssumeRoleWithWebIdentity",
			"Condition": {
				"StringEquals": {
					"token.actions.githubusercontent.com:aud": "sts.amazonaws.com"
				},
				"StringLike": {
					"token.actions.githubusercontent.com:sub": "repo:shamimice03/push-image-to-ecr:ref:refs/heads/dev"
				}
			}
		}
	]
}
```

> Policy: ECR-cloudterms-prod-repo-policy and ECR-cloudterms-dev-repo-policy
```
{
    "Version": "2012-10-17",
    "Statement": [
        {
            "Sid": "GetAuthorizationToken",
            "Effect": "Allow",
            "Action": [
                "ecr:GetAuthorizationToken"
            ],
            "Resource": "*"
        },
        {
            "Sid": "AllowPushPull",
            "Effect": "Allow",
            "Action": [
                "ecr:BatchGetImage",
                "ecr:BatchCheckLayerAvailability",
                "ecr:CompleteLayerUpload",
                "ecr:GetDownloadUrlForLayer",
                "ecr:InitiateLayerUpload",
                "ecr:PutImage",
                "ecr:UploadLayerPart"
            ],
            "Resource": "arn:aws:ecr:ap-northeast-1:391178969547:repository/cloudterms-prod"
        }
    ]
}
```


```
 if: github.event.pull_request.merged == true || github.event_name == 'push'
```


```
sequenceDiagram
    autonumber
    participant A as Github Actions
    participant B as Oidc Provider
    participant C as AWS IAM

    A->>+B: Request JWT
    B->>-A: Issue signed JWT
    A->>+C: Request Access Token
    C-->>+B: Verify Token
    B-->>-C: Valid
    C->>-A: Issue Role Access Session Token
```