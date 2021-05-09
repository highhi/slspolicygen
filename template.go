package slspolicygen

var policyTemplate = `
{
	"Version": "2012-10-17",
	"Statement": [
		{
			"Sid": "cloudformation",
			"Effect": "Allow",
			"Action": [
				"cloudformation:DescribeStackEvents",
				"cloudformation:CreateStack",
				"cloudformation:DeleteStack",
				"cloudformation:UpdateStack",
				"cloudformation:DescribeStackResources",
				"cloudformation:DescribeStackResource",
				"cloudformation:DescribeStacks",
				"cloudformation:ListStackResources"
			],
			"Resource": "arn:aws:cloudformation:ap-northeast-1:{{ .account }}:stack/{{ .servicename }}*"
		},
		{
			"Sid": "iam0",
			"Effect": "Allow",
			"Action": [
				"cloudformation:ValidateTemplate"
			],
			"Resource": "*"
		},
		{
			"Sid": "iam1",
			"Effect": "Allow",
			"Action": [
				"iam:PassRole"
			],
			"Resource": "arn:aws:iam::{{ .account }}:role/AppDynamoDBAccess"
		},
		{{- if .s3Required -}}
		{
			"Sid": "s3",
			"Effect": "Allow",
			"Action": [
				"s3:Get*",
				"s3:List*",
				"s3:Put*",
				"s3:CreateBucket",
				"s3:DeleteBucket",
				"s3:DeleteBucketPolicy",
				"s3:DeleteObject"
			],
			"Resource": "arn:aws:s3:::{{ .servicename }}*"
		},
		{{ end }}
		{
			"Sid": "cloudwatch",
			"Effect": "Allow",
			"Action": [
				"logs:DescribeLogGroups",
				"logs:CreateLogGroup",
				"logs:DeleteLogGroup",
				"logs:PutRetentionPolicy"
			],
			"Resource": "arn:aws:logs:ap-northeast-1:{{ .account }}:*"
		},
		{
			"Sid": "cloudwatchevents",
			"Effect": "Allow",
			"Action": [
				"events:PutRule",
				"events:DescribeRule",
				"events:DeleteRule",
				"events:PutTargets",
				"events:RemoveTargets"
			],
			"Resource": "arn:aws:events:ap-northeast-1:{{ .account }}:rule/{{ .servicename }}*"
		},
		{{- if .dynamoDBRequired -}}
		{
			"Sid": "dynamodb",
			"Effect": "Allow",
			"Action": [
				"dynamodb:DescribeTable",
				"dynamodb:CreateTable",
				"dynamodb:DeleteTable"
			],
			"Resource": [
				"arn:aws:events:ap-northeast-1:{{ .account }}:rule/{{ .servicename }}*",
				"arn:aws:dynamodb:ap-northeast-1:{{ .account }}:table*"
			]
		},
		{{ end }}
		{
			"Sid": "apigateway",
			"Effect": "Allow",
			"Action": [
				"apigateway:*"
			],
			"Resource": "arn:aws:apigateway:ap-northeast-1::/restapis*"
		},
		{
			"Sid": "lambda",
			"Effect": "Allow",
			"Action": [
				"lambda:GetFunction",
				"lambda:DeleteFunction",
				"lambda:CreateFunction",
				"lambda:GetFunctionConfiguration",
				"lambda:ListVersionsByFunction",
				"lambda:AddPermission",
				"lambda:RemovePermission",
				"lambda:PublishVersion",
				"lambda:UpdateFunctionCode",
				"lambda:ListAliases",
				"lambda:UpdateFunctionConfiguration"
			],
			"Resource": "arn:aws:lambda:ap-northeast-1:{{ .account }}:function:{{ .servicename }}*"
		}
	]
}
`
