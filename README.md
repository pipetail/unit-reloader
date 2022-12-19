# SQS Systemd unit reloader

This utility performs start|stop|restart on the given Systemd unit when it receives SQS message. See more details in the [official AWS documentation](https://aws.amazon.com/premiumsupport/knowledge-center/sqs-s3-event-notification-sse/).

## AWS configuration

### SQS queue configuration

Queue needs to be configured with the access policy that allows perform `sqs:SendMessage`
to S3 bucket.

```json
{
  "Version": "2008-10-17",
  "Id": "__default_policy_ID",
  "Statement": [
    {
      "Sid": "__owner_statement",
      "Effect": "Allow",
      "Principal": {
        "AWS": "arn:aws:iam::123456789123:root"
      },
      "Action": "SQS:*",
      "Resource": "arn:aws:sqs:eu-west-1:123456789123:events"
    },
    {
      "Sid": "s3_notification",
      "Effect": "Allow",
      "Principal": {
        "Service": "s3.amazonaws.com"
      },
      "Action": "SQS:SendMessage",
      "Resource": "arn:aws:sqs:eu-west-1:123456789123:events",
      "Condition": {
        "StringEquals": {
          "aws:SourceAccount": "123456789123"
        },
        "ArnLike": {
          "aws:SourceArn": "arn:aws:s3:::unit-reloader-test"
        }
      }
    }
  ]
}
```

### S3 bucket configuration

Create an event notification for the `s3:ObjectCreated:Put` and `s3:ObjectCreated:Post` events and set SQS qeue as the destination.

## Example

```bash
AWS_PROFILE=personal ./main -queue-url='https://sqs.eu-west-1.amazonaws.com/123456789123/events' -unit=nginx.service -action=restart
```
