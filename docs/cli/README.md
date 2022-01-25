# Connect to DekaBox using AWS CLI

## Prequistes

1. Have installed aws cli. Download aws cli from [here](https://aws.amazon.com/cli/)
2. Have DekaBox account

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


## 2. Create a bucket.

```
aws s3api --endpoint-url https://bucket.cloud.lintasarta.co.id:8082 create-bucket --bucket bucket-for-testing
```

If the bucket is created successfully, the location of the bucket is returned, as seen in the following example:
"Location": "/bucket-for-testing"

## 3. Upload an object.

```
aws s3api --endpoint-url https://bucket.cloud.lintasarta.co.id:8082 put-object --bucket bucket-for-testing --key file-test.pdf --body C:\test-s3-cli\files\file-test.pdf
```

If the object is uploaded successfully, an Etag is returned which is a hash of the object data.

## 4. List the contents of the bucket to verify that the object was uploaded.

```
aws s3api --endpoint-url https://bucket.cloud.lintasarta.co.id:8082 list-objects --bucket bucket-for-testing
```

## 5. Delete the object.

```
aws s3api --endpoint-url https://bucket.cloud.lintasarta.co.id:8082  delete-object --bucket bucket-for-testing --key file-test.pdf
```

## 6. Delete the bucket.

```
aws s3api --endpoint-url https://bucket.cloud.lintasarta.co.id:8082 delete-bucket --bucket bucket-for-testing
```