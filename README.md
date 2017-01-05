# Terraform Volume Attachment ID (tfvolattid)
Regenerate the "vai" ID for Terraform state file volume attachment

If you haven't already, you will need to pull down the hashcode package and put it in your GOROOT directory:
https://raw.githubusercontent.com/hashicorp/terraform/master/helper/hashcode/hashcode.go

Once this is done you can do "go run tfvolattid.go <mount point> <volume id> <instance id>"

Your volume attachment in your state file will look something like this:
```json
"aws_volume_attachment.myebs_ebs_att.1": {
    "type": "aws_volume_attachment",
    "depends_on": [
        "aws_ebs_volume.myebs_ebs",
    ],
    "primary": {
        "id": "vai-3314652726",
        "attributes": {
            "device_name": "/dev/xvde",
            "id": "vai-3314652726",
            "instance_id": "i-0bbced25e",
            "volume_id": "vol-0f43568a8"
        },
        "meta": {},
        "tainted": false
    },
    "deposed": [],
    "provider": ""
},
```
