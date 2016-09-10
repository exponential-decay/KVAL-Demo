package main

import (
   "fmt"
   "github.com/exponential-decay/KVAL-Parse"
) 

func main() {

   //var insert_value = "INS Prime Bucket >> Secondary Bucket >> Tertiary Bucket >>>> Key :: Value"
   //var get_value = "GET Prime Bucket >> Secondary Bucket >> Tertiary Bucket"   
   //var get_value = "GET Prime Bucket >> Secondary Bucket >> Tertiary Bucket >>>> Key"
   //var get_value_from_key_pattern = "GET Prime Bucket >> Secondary Bucket >> Tertiary Bucket >>>> {PAT}"
   //var get_key_from_value = "GET Prime Bucket >> Secondary Bucket >> Tertiary Bucket >>>> _ :: Value"
   //var get_key_from_value_pattern = "GET Prime Bucket >> Secondary Bucket >> Tertiary Bucket >>>> _ :: {PAT}"
   //var does_key_exist = "LIS Prime Bucket >> Secondary Bucket >> Tertiary Bucket >>>> Key"
   //var does_bucket_exist = "LIS Prime Bucket >> Secondary Bucket >> Tertiary Bucket"
   //var delete_bucket = "DEL Prime Bucket >> Secondary Bucket >> Tertiary Bucket"
   //var delete_key = "DEL Prime Bucket >> Secondary Bucket >> Tertiary Bucket >>>> Key"
   var delete_value = "DEL Prime Bucket >> Secondary Bucket >> Tertiary Bucket >>>> Key :: _ "
   //var rename_key = "REN Prime Bucket >> Secondary Bucket >> Tertiary Bucket >>>> Key => Key"
   //var rename_bucket = "REN Prime Bucket >> Secondary Bucket >> Tertiary Bucket => Third Bucket"
   
   fmt.Println("KVAL demo implementation.")

   //kvalparse.Query(insert_value)
   //kvalparse.Query(get_value)
   //kvalparse.Query(get_value_from_key_pattern)
   //kvalparse.Query(get_key_from_value)
   //kvalparse.Query(get_key_from_value_pattern)
   //kvalparse.Query(does_key_exist)
   //kvalparse.Query(does_bucket_exist) 
   //kvalparse.Query(delete_bucket)
   //kvalparse.Query(delete_key)
   kvalparse.Query(delete_value)
   //kvalparse.Query(rename_key)
   //kvalparse.Query(rename_bucket)

}
