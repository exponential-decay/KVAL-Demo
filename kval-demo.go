package main

import (
   "fmt"
   "github.com/exponential-decay/KVAL-Parse"
) 

var insert_value = "INS Prime Bucket >> Secondary Bucket >> Tertiary Bucket >>>> Key :: Value"
var get_bucket_contents = "GET Prime Bucket >> Secondary Bucket >> Tertiary Bucket"   
var get_value = "GET Prime Bucket >> Secondary Bucket >> Tertiary Bucket >>>> Key"
var get_value_from_key_pattern = "GET Prime Bucket >> Secondary Bucket >> Tertiary Bucket >>>> {PAT}"
var get_key_from_value = "GET Prime Bucket >> Secondary Bucket >> Tertiary Bucket >>>> _ :: Value"
var get_key_from_value_pattern = "GET Prime Bucket >> Secondary Bucket >> Tertiary Bucket >>>> _ :: {PAT}"
var does_key_exist = "LIS Prime Bucket >> Secondary Bucket >> Tertiary Bucket >>>> Key"
var does_bucket_exist = "LIS Prime Bucket >> Secondary Bucket >> Tertiary Bucket"
var delete_bucket = "DEL Prime Bucket >> Secondary Bucket >> Tertiary Bucket"
var delete_key = "DEL Prime Bucket >> Secondary Bucket >> Tertiary Bucket >>>> Key"
var delete_value = "DEL Prime Bucket >> Secondary Bucket >> Tertiary Bucket >>>> Key :: _ "
var rename_key = "REN Prime Bucket >> Secondary Bucket >> Tertiary Bucket >>>> Key => New Key"
var rename_bucket = "REN Prime Bucket >> Secondary Bucket >> Tertiary Bucket => Third Bucket"


func main() {
   
   fmt.Println("KVAL demo implementation.")

   //kvalparse.Parse(insert_value)
   //kvalparse.Parse(get_value)
   kq, err := kvalparse.Parse(get_value_from_key_pattern)
   //kq, err := kvalparse.Parse(get_key_from_value)
   //kvalparse.Parse(get_key_from_value_pattern)
   //kvalparse.Parse(does_key_exist)
   //kvalparse.Parse(does_bucket_exist) 
   //kvalparse.Parse(delete_bucket)
   //kvalparse.Parse(delete_key)
   //kq, err := kvalparse.Parse(delete_value)
   //kq, err := kvalparse.Parse(rename_key)
   //kq, err := kvalparse.Parse(rename_bucket)

   if err != nil {
      fmt.Println(err)
   }
   fmt.Println(kq)
}
