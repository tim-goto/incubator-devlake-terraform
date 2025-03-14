---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "devlake_apikeys Data Source - devlake"
subcategory: ""
description: |-
  
---

# devlake_apikeys (Data Source)





<!-- schema generated by tfplugindocs -->
## Schema

### Read-Only

- `apikeys` (Attributes List) (see [below for nested schema](#nestedatt--apikeys))

<a id="nestedatt--apikeys"></a>
### Nested Schema for `apikeys`

Read-Only:

- `allowed_path` (String) The API URL or endpoint that the API key is permitted to access. It defines the specific resources that the key can interact with.
- `api_key` (String) The returned apikey, will always be empty for security reasons. If you want to use the apikey, you need to create a new resource.
- `created_at` (String) When the apikey was created.
- `creator` (String) Who created the apikey, there is no user management yet though.
- `creator_email` (String) Email of the person who created the apikey, there is no user management yet though.
- `expired_at` (String) When the apikey expires.
- `extra` (String) Currently not used.
- `id` (Number) Numeric identifier for the apikey.
- `name` (String) The name of the apikey.
- `type` (String) The apikey type. Currently only 'devlake' is a valid value.
- `updated_at` (String) When the apikey was last updated. Serves no purpose as the UPDATE endpoint doesn't do anything.
- `updater` (String) Who updated the apikey, there is no user management yet though.
- `updater_email` (String) Email of the person who updated the apikey, there is no user management yet though.
