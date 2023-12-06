// Code generated by smithy-go-codegen DO NOT EDIT.

// Package paymentcryptography provides the API client, operations, and parameter
// types for Payment Cryptography Control Plane.
//
// Amazon Web Services Payment Cryptography Control Plane APIs manage encryption
// keys for use during payment-related cryptographic operations. You can create,
// import, export, share, manage, and delete keys. You can also manage Identity and
// Access Management (IAM) policies for keys. For more information, see Identity
// and access management (https://docs.aws.amazon.com/payment-cryptography/latest/userguide/security-iam.html)
// in the Amazon Web Services Payment Cryptography User Guide. To use encryption
// keys for payment-related transaction processing and associated cryptographic
// operations, you use the Amazon Web Services Payment Cryptography Data Plane (https://docs.aws.amazon.com/payment-cryptography/latest/DataAPIReference/Welcome.html)
// . You can perform actions like encrypt, decrypt, generate, and verify
// payment-related data. All Amazon Web Services Payment Cryptography API calls
// must be signed and transmitted using Transport Layer Security (TLS). We
// recommend you always use the latest supported TLS version for logging API
// requests. Amazon Web Services Payment Cryptography supports CloudTrail for
// control plane operations, a service that logs Amazon Web Services API calls and
// related events for your Amazon Web Services account and delivers them to an
// Amazon S3 bucket you specify. By using the information collected by CloudTrail,
// you can determine what requests were made to Amazon Web Services Payment
// Cryptography, who made the request, when it was made, and so on. If you don't
// conﬁgure a trail, you can still view the most recent events in the CloudTrail
// console. For more information, see the CloudTrail User Guide (https://docs.aws.amazon.com/awscloudtrail/latest/userguide/)
// .
package paymentcryptography
