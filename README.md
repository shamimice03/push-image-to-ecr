Steps:


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

### ECR

`Registry:` An ECR registry is provided to each AWS account; we can create image repositories in our registry and store images in them.

`Repository:` An ECR image repository contains our Docker images.

`Repository policy:` We can control access to our repositories and the images within them with repository policies.

`Authorization token:` Our Docker client must authenticate to Amazon ECR registries as an AWS user before it can push and pull images. The AWS CLI get-login command provides us with authentication credentials to pass to Docker.

`Image:` We can push and pull container images to our repositories.


```
 aws ecr describe-repositories
 
 aws ecr get-login-password --region ap-northeast-1 | sudo docker login --username AWS --password-stdin 391178969547.dkr.ecr.ap-northeast-1.amazonaws.com/cloudterms-prod
 
 sudo docker pull 391178969547.dkr.ecr.ap-northeast-1.amazonaws.com/cloudterms-prod:6591809cfd585396a8769dcfda95a8ed9624f798

 sudo docker images

 docker run -d -p 8080:8080 --network=host 391178969547.dkr.ecr.ap-northeast-1.amazonaws.com/cloudterms-prod:6591809cfd585396a8769dcfda95a8ed9624f798

```

```
          git config user.email ${{ secrets.USER_EMAIL }}
          git config user.name ${{ secrets.USERNAME }}
```