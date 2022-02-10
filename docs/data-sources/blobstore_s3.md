---
page_title: "Data Source nexus_blobstore_s3"
subcategory: "Blobstore"
description: |-
  Use this data source to get details of an existing Nexus S3 blobstore.
---
# Data Source nexus_blobstore_s3
Use this data source to get details of an existing Nexus S3 blobstore.
## Example Usage
```terraform
data "nexus_blobstore_s3" "aws" {
	name = "blobstore-s3"
}
```
<!-- schema generated by tfplugindocs -->
## Schema

### Required

- **name** (String) Blobstore name

### Read-Only

- **blob_count** (Number) Count of blobs
- **bucket_configuration** (List of Object) The S3 bucket configuration. (see [below for nested schema](#nestedatt--bucket_configuration))
- **id** (String) Used to identify data source at nexus
- **soft_quota** (List of Object) Soft quota of the blobstore (see [below for nested schema](#nestedatt--soft_quota))
- **total_size_in_bytes** (Number) The total size of the blobstore in Bytes

<a id="nestedatt--bucket_configuration"></a>
### Nested Schema for `bucket_configuration`

Read-Only:

- **advanced_bucket_connection** (List of Object) (see [below for nested schema](#nestedobjatt--bucket_configuration--advanced_bucket_connection))
- **bucket** (List of Object) (see [below for nested schema](#nestedobjatt--bucket_configuration--bucket))
- **bucket_security** (List of Object) (see [below for nested schema](#nestedobjatt--bucket_configuration--bucket_security))
- **encryption** (List of Object) (see [below for nested schema](#nestedobjatt--bucket_configuration--encryption))

<a id="nestedobjatt--bucket_configuration--advanced_bucket_connection"></a>
### Nested Schema for `bucket_configuration.advanced_bucket_connection`

Read-Only:

- **endpoint** (String)
- **force_path_style** (Boolean)
- **signer_type** (String)


<a id="nestedobjatt--bucket_configuration--bucket"></a>
### Nested Schema for `bucket_configuration.bucket`

Read-Only:

- **expiration** (Number)
- **name** (String)
- **prefix** (String)
- **region** (String)


<a id="nestedobjatt--bucket_configuration--bucket_security"></a>
### Nested Schema for `bucket_configuration.bucket_security`

Read-Only:

- **access_key_id** (String)
- **role** (String)
- **secret_access_key** (String)
- **session_token** (String)


<a id="nestedobjatt--bucket_configuration--encryption"></a>
### Nested Schema for `bucket_configuration.encryption`

Read-Only:

- **encryption_key** (String)
- **encryption_type** (String)



<a id="nestedatt--soft_quota"></a>
### Nested Schema for `soft_quota`

Read-Only:

- **limit** (Number)
- **type** (String)
