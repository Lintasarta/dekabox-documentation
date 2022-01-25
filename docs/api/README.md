# Connect to DekaBox using AWS S3 API

## 1. Configure the DekaBox S3 credential by using AWS S3 cli

Try to run following command

```
aws configure
```
It will asked about

Enter configuration mode: 
Enter the AWS Access Key ID for the account you created.
Enter the AWS Secret Access key for the account you created.
Enter the default region to use, for example, id-teknopark, id-tbs.
Enter the default output format to use, or press Enter to select JSON.

## 2. (optional) using environment variables in golang

## 3. Using following library to connect with 

```
github.com/aws/aws-sdk-go/aws
github.com/aws/aws-sdk-go/aws/session
github.com/aws/aws-sdk-go/service/s3
```