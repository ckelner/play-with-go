// For https://github.com/hashicorp/terraform/issues/3208

package main

import "fmt"
import "github.com/aws/aws-sdk-go/aws"

func main() {
  var cookie_expire int
  fmt.Println(int64(cookie_expire))
  fmt.Println(aws.Int64(int64(cookie_expire)))
}
