package main

import (
	"bytes"
	"fmt"
	"os"

	"github.com/hashicorp/terraform/helper/hashcode"
)

func volumeAttachmentID(name, volumeID, instanceID string) string {
	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintf("%s-", name))
	buf.WriteString(fmt.Sprintf("%s-", instanceID))
	buf.WriteString(fmt.Sprintf("%s-", volumeID))

	return fmt.Sprintf("vai-%d", hashcode.String(buf.String()))
}

func main() {
	cmd := os.Args[0]
	if len(os.Args) > 3 {
		mount := os.Args[1]
		volID := os.Args[2]
		instanceID := os.Args[3]
		fmt.Println(volumeAttachmentID(mount, volID, instanceID))
	} else {
		fmt.Printf("Usage: %s <mount point> <volume id> <instance id>\n", cmd)
	}
}
